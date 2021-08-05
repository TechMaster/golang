package controller

import (
	"context"
	"crypto/sha1"
	"encoding/json"
	"errors"
	"html/template"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"
	"techmaster-anonymous/constant"
	"techmaster-anonymous/helper"
	"techmaster-anonymous/model"
	package_model "techmaster-anonymous/package-model"

	"github.com/asaskevich/govalidator"
	"github.com/go-pg/pg/orm"
	"github.com/jinzhu/copier"
	"github.com/kataras/iris/v12"
	"github.com/rs/xid"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"

	p "github.com/wuriyanto48/go-pbkdf2"

	"github.com/microcosm-cc/bluemonday"

	uuid "github.com/satori/go.uuid"
)

func (c *Controller) Login(ctx iris.Context) {
	req := model.LoginRequest{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Request sai định dạng"))
		return
	}

	if !govalidator.IsEmail(req.Email) {
		package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Định dạng email không hợp lệ"))
		return
	}
	if req.Password == "" {
		package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Vui lòng nhập mật khẩu"))
		return
	}

	// Lấy thông tin tài khoản từ CSDL
	var user model.UserInfo
	if len(req.Email) > 0 {
		// Kiểm tra tài khoản theo Email
		rs, err := c.DB.Query(&user, `
			SELECT id, email, avatar, full_name, roles, phone, 
			email_confirmed, password, access_failed_count, lockout_end, 
			user_status, salt, verify_email_token_end
			FROM  auth.users
			WHERE email = ? `, req.Email)

		if err != nil {
			package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Tài khoản hoặc mật khẩu không chính xác."))
			return
		}

		if rs.RowsReturned() == 0 {
			package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Tài khoản hoặc mật khẩu không chính xác."))
			return
		}

	} else if len(req.Email) < 1 && len(req.Phone) > 0 {
		package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Hiện chưa hỗ trợ đăng nhập bằng số điện thoại."))
		return
	}

	// Trường hợp tài khoản chưa kích hoạt
	if !user.EmailConfirmed {
		// Kiểm tra TOKEN còn thời gian hiệu lực không
		timeLeft := user.VerifyEmailTokenEnd.Sub(time.Now())

		if timeLeft > 0 {
			// Nếu TOKEN còn hiệu trong 1 ngày thì thông báo
			package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Một email kích hoạt đã được gửi tới hòm thư của bạn. Vui lòng kích hoạt tài khoản"))
			return
		} else {
			// Nếu TOKEN không còn hiệu lực thì gửi lại email
			uuidToken := uuid.NewV4()

			randomStr, errRanStr := package_model.GenerateRandomString(20)
			if errRanStr != nil {
				package_model.NewError(ctx, iris.StatusInternalServerError, errors.New("Không thể gửi lại email kích hoạt"))
				return
			}
			token := uuidToken.String() + randomStr

			var userResend package_model.User
			// Gửi mail kích hoạt
			userResend.Id = user.Id
			userResend.Email = user.Email
			userResend.Phone = user.Phone
			userResend.FullName = user.FullName
			userResend.VerifyEmailToken = token
			userResend.VerifyEmailTokenEnd = time.Now().Add(1440 * time.Minute)
			_, err = c.DB.Model(&userResend).Column("verify_email_token", "verify_email_token_end").WherePK().Update()
			if err != nil {
				package_model.NewError(ctx, iris.StatusInternalServerError, errors.New("Không thể gửi lại email kích hoạt"))
				return
			}

			key := userResend.VerifyEmailToken + "/" + userResend.Id

			var emailReq package_model.CreateAccountUser
			emailReq.Key = key
			emailReq.FullName = user.FullName
			emailReq.Email = user.Email

			url := c.Config.SignupMail
			c.SendMailHelper(url, emailReq)

			package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Tài khoản chưa kích hoạt. Một email kích hoạt đã được gửi tới hòm thư của bạn."))
			return
		}
	}

	// Trường hợp tài khoản bị khóa
	if !user.UserStatus {
		package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Tài khoản đã bị khóa."))
		return
	}

	// Trường hợp tài khoản đang bị lock
	var t time.Time
	remainingTime := time.Now().Sub(user.LockoutEnd)
	if user.LockoutEnd != t && remainingTime < 0 {
		remainingMinutes := strconv.FormatFloat(-math.Round(remainingTime.Minutes()), 'f', 0, 64)
		package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Đăng nhập sai quá nhiều. Vui lòng thử lại sau "+remainingMinutes+" phút"))
		return
	}

	// Trường hợp nhập sai quá số lần quy định
	if user.AccessFailedCount == c.Config.Authenticate.MaxAccessFailed {
		lockoutTime := time.Duration(c.Config.Authenticate.LockoutInMinutes)
		lockoutTimeMinutes := strconv.Itoa(int(c.Config.Authenticate.LockoutInMinutes))
		_, err = c.DB.Exec(`
			UPDATE auth.users 
			SET access_failed_count = 0, lockout_end = ?
			WHERE id = ? and email = ?`, time.Now().Add(time.Minute*lockoutTime), user.Id, user.Email)
		if err != nil {
			package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Không thể cập nhật thời gian tạm khoá tài khoản do đăng nhập sai quá nhiều"))
			return
		}
		package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Đăng nhập sai quá nhiều. Vui lòng thử lại sau "+lockoutTimeMinutes+" phút"))
		return
	}

	// Kiểm tra mật khẩu
	pass := p.NewPassword(sha1.New, 50, 64, 10000)
	if pass.VerifyPassword(req.Password, user.Password, user.Salt) {
		if len(user.Roles) == 0 {
			user.Roles = []int32{}
		}

		// Mật khẩu đúng reset số lần nhập sai
		_, err = c.DB.Exec(`UPDATE auth.users SET access_failed_count = 0, lockout_end = null WHERE id = ? and email = ?`, user.Id, user.Email)
		if err != nil {
			package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Không thể kiểm tra số lần nhập sai"))
			return
		}

		tx, err := c.DB.Begin()
		if err != nil {
			package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Không thể kết nối với cơ sở dữ liệu"))
			return
		}

		// Đồng bộ hóa giỏ hàng
		err = SynchronizeShoppingCart(user, tx, ctx)
		if err != nil {
			package_model.NewError(ctx, iris.StatusInternalServerError, errors.New("Không thể đồng bộ hoá giỏ hàng"))
			return
		}

		tx.Commit()

		// Lấy ra các key mà user đã sở hữu
		var keys []string
		_, err = c.DB.Query(&keys, `
			SELECT distinct unnest(auth_role.keys) roles 
			FROM auth.users users 
			INNER JOIN auth.role as auth_role 
			ON auth_role.id = ANY(users.roles)
			WHERE users.email = ? AND auth_role.status = true`, req.Email)
		if err != nil {
			package_model.NewError(ctx, iris.StatusInternalServerError, errors.New("Không thể kiểm tra quyền hạn user hiện tại"))
			return
		}

		// Set session
		session := constant.Sess.Start(ctx)

		var auth = package_model.Authenticate{
			Authenticated: true,
			UserId:        user.Id,
			UserFullName:  user.FullName,
			Keys:          keys,
			UserEmail:     user.Email,
			UserAvatar:    user.Avatar,
			UserPhone:     user.Phone,
			ExpiredTime:   time.Now().Add(8760 * time.Hour),
		}
		// set session
		session.Set("authenticate", auth)

		if len(user.Phone) == 0 {
			ctx.JSON(auth)
		} else {
			ctx.JSON(auth)
		}
	} else {
		// Mật khẩu sai tăng số lần nhập sai
		_, err = c.DB.Exec(`
			UPDATE auth.users
			SET access_failed_count = ?
			WHERE id = ? and email = ?`, user.AccessFailedCount+1, user.Id, user.Email)
		if err != nil {
			package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Không thể cập nhật số lần đăng nhập sai"))
			return
		}
		package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Tài khoản hoặc mật khẩu không chính xác"))
		return
	}
}

func (c *Controller) UpdateRequiredInfo(ctx iris.Context) {
	var req model.UpdateRequiredInfo
	err := ctx.ReadJSON(&req)
	if err != nil {
		package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Định dạng request sai"))
		return
	}
	p := bluemonday.StrictPolicy()

	if !govalidator.IsNumeric(p.Sanitize(req.Phone)) || len(p.Sanitize(req.Phone)) > 20 || len(p.Sanitize(req.Phone)) <= 8 {
		package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Định dạng số điện thoại không hợp lệ"))
		return
	}

	var userID string
	userAuthen, isLogin := package_model.GetAuthentication(constant.Sess, ctx)
	if isLogin {
		userID = userAuthen.UserId
	} else {
		package_model.NewError(ctx, http.StatusBadRequest, errors.New("Bạn không có quyền"))
		return
	}

	var user package_model.User
	user.Id = userID
	user.Phone = p.Sanitize(req.Phone)
	_, err = c.DB.Model(&user).Column("phone").WherePK().Update()
	if err != nil {
		package_model.NewError(ctx, http.StatusInternalServerError, errors.New("Cập nhật thất bại"))
		return
	}

	// Set session
	session := constant.Sess.Start(ctx)
	userAuthen.UserPhone = req.Phone
	session.Set("authenticate", userAuthen)

	ctx.JSON("Cập nhập thành công")
	ctx.StatusCode(iris.StatusOK)
}

func (c *Controller) Logout(ctx iris.Context) {
	session := constant.Sess.Start(ctx)
	// Revoke users authentication
	session.Delete("authenticate")

	// Xoá cookie giỏ hàng
	ctx.SetCookie(
		&http.Cookie{
			Name:     package_model.COOKIE_NAME,
			HttpOnly: true,
			MaxAge:   -1,
			Expires:  time.Now().Add(-100 * time.Hour),
			Secure:   true,
		},
	)

	ctx.SetCookie(
		&http.Cookie{
			Name:     constant.CookieNameForSessionID,
			HttpOnly: true,
			MaxAge:   -1,
			Expires:  time.Now().Add(-100 * time.Hour),
			Secure:   true,
		},
	)

	ctx.StatusCode(200)
	ctx.JSON("Đăng xuất thành công")
}

func (c *Controller) VerifyEmail(ctx iris.Context) {
	token := ctx.Params().GetStringDefault("token", "")
	userID := ctx.Params().GetStringDefault("userId", "")

	if len(token) == 0 || len(userID) == 0 {
		ctx.View("/error/error.html")
		return
	}

	var user package_model.User
	err := c.DB.Model(&user).Where("id = ?", userID).Select()
	if err != nil {
		ctx.View("/error/error.html")
		return
	}

	if user.EmailConfirmed {
		// Tài khoản đã kích hoạt
		ctx.ViewData("is_used", true)
	} else {
		if user.VerifyEmailToken != token {
			// Mã kích hoạt không hợp lệ
			ctx.View("/error/error.html")
			return
		}

		if user.VerifyEmailTokenEnd.Unix() < time.Now().Unix() {
			ctx.ViewData("is_expired", true)
		} else {
			if len(user.Password) > 0 {
				// Tài khoản tự đăng ký
				user.EmailConfirmed = true
				_, err = c.DB.Model(&user).Column("email_confirmed").Where("id = ?", userID).Update()

				if err != nil {
					ctx.View("/error/error.html")
					return
				}

				// Lấy ra các key mà user đã sở hữu
				var keys []string
				_, err = c.DB.Query(&keys, `
					SELECT distinct unnest(auth_role.keys) roles FROM auth.users users 
					INNER JOIN auth.role as auth_role ON auth_role.id = ANY(users.roles)
					where users.email = ? and auth_role.status = true`, user.Email)
				if err == nil {
					// Set session
					session := constant.Sess.Start(ctx)
					var auth = package_model.Authenticate{
						Authenticated: true,
						UserId:        user.Id,
						UserFullName:  user.FullName,
						Keys:          keys,
						UserEmail:     user.Email,
						UserAvatar:    user.Avatar,
						UserPhone:     user.Phone,
						ExpiredTime:   time.Now().Add(48 * time.Hour),
					}
					// set session
					session.Set("authenticate", auth)
					var userInfo = UserInfoData{
						Authenticated:    true,
						UserId:           user.Id,
						UserFullName:     user.FullName,
						UserEmail:        user.Email,
						UserAvatar:       user.Avatar,
						UserPhone:        user.Phone,
						CountCartItem:    0,
						GithubClientId:   "",
						GoogleClientId:   "",
						FacebookClientId: "",
					}
					ctx.ViewData("user_info", userInfo)
				}

				ctx.ViewData("success", true)
			} else {
				// Tài khoản tạo tự động
				ctx.ViewData("userId", userID)
				ctx.ViewData("token", token)
				ctx.ViewData("set_password", true)
			}
		}
	}
	ctx.View("xac-thuc.html")
}

func (c *Controller) SetPasswordAndConfirm(ctx iris.Context) {
	var req model.SetPassword
	err := ctx.ReadJSON(&req)
	if err != nil {
		package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Định dạng request sai"))
		return
	}

	if len(req.Password) == 0 {
		package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Vui lòng nhập mật khẩu"))
		return
	}

	var user package_model.User
	err = c.DB.Model(&user).Where("id = ?", req.UserId).Select()
	if err != nil {
		package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Tài khoản chưa tồn tại trong hệ thống"))
		return
	}
	if user.EmailConfirmed {
		package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Tài khoản đã được kích hoạt trước đó"))
		return
	}
	if user.VerifyEmailToken != req.Token {
		package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Token không hợp lệ"))
		return
	}
	if user.VerifyEmailTokenEnd.Unix() < time.Now().Unix() {
		package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Token hết hạn"))
		return
	}

	// Hash password
	hashedPassword, salt := helper.HashPassword(req.Password)

	// Update thông tin
	user.Password = hashedPassword
	user.Salt = salt
	user.EmailConfirmed = true
	upadtedColumns := []string{"password", "salt", "email_confirmed"}
	_, err = c.DB.Model(&user).Column(upadtedColumns...).WherePK().Update()
	if err != nil {
		package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Không thể lưu mật khẩu. Vui lòng thử lại"))
		return
	}

	// Lấy ra các key mà user đã sở hữu
	var keys []string
	_, err = c.DB.Query(&keys, `
		SELECT distinct unnest(auth_role.keys) roles FROM auth.users users 
		INNER JOIN auth.role as auth_role ON auth_role.id = ANY(users.roles)
		where users.email = ? and auth_role.status = true`, user.Email)
	if err == nil {
		// Set session
		session := constant.Sess.Start(ctx)
		var auth = package_model.Authenticate{
			Authenticated: true,
			UserId:        user.Id,
			UserFullName:  user.FullName,
			Keys:          keys,
			UserEmail:     user.Email,
			UserAvatar:    user.Avatar,
			UserPhone:     user.Phone,
			ExpiredTime:   time.Now().Add(48 * time.Hour),
		}
		// set session
		session.Set("authenticate", auth)
		var userInfo = UserInfoData{
			Authenticated:    true,
			UserId:           user.Id,
			UserFullName:     user.FullName,
			UserEmail:        user.Email,
			UserAvatar:       user.Avatar,
			UserPhone:        user.Phone,
			CountCartItem:    0,
			GithubClientId:   "",
			GoogleClientId:   "",
			FacebookClientId: "",
		}
		ctx.ViewData("user_info", userInfo)
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON("Đặt mật khẩu thành công")
}

func (c *Controller) SendConfirmRequest(ctx iris.Context) {
	var req model.SendConfirmRequest
	err := ctx.ReadJSON(&req)
	if err != nil {
		package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Định dạng request sai"))
		return
	}

	// Validate thông tin
	if !govalidator.IsEmail(req.Email) {
		package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Định dạng email không hợp lệ"))
		return
	}

	var user package_model.User
	err = c.DB.Model(&user).Where("email = ?", req.Email).Select()
	if err != nil {
		package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Email chưa tồn tại trong hệ thống"))
		return
	}
	if user.EmailConfirmed {
		package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Tài khoản đã được kích hoạt trước đó"))
		return
	}

	// Gửi mail kích hoạt
	uuidToken := uuid.NewV4()
	randomStr, errRanStr := helper.GenerateRandomString(20)
	if errRanStr != nil {
		package_model.NewError(ctx, iris.StatusInternalServerError, errors.New("Lỗi khi gửi email kích hoạt"))
		return
	}
	token := uuidToken.String() + randomStr
	user.VerifyEmailToken = token
	user.VerifyEmailTokenEnd = time.Now().Add(30 * time.Minute)
	err = user.UpdateUser(c.DB, []string{"verify_email_token", "verify_email_token_end"})
	if err != nil {
		package_model.NewError(ctx, iris.StatusInternalServerError, errors.New("Không thể gửi email kích hoạt"))
		return
	}
	key := user.VerifyEmailToken + "/" + user.Id

	var emailReq package_model.CreateAccountUser
	emailReq.Key = key
	emailReq.FullName = user.FullName
	emailReq.Email = user.Email

	url := c.Config.SignupMail
	c.SendMailHelper(url, emailReq)

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON("Gửi email kích hoạt thành công")
}

func (c *Controller) SendForgotPasswordRequest(ctx iris.Context) {
	var req model.SendConfirmRequest
	err := ctx.ReadJSON(&req)
	if err != nil {
		package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Định dạng request sai"))
		return
	}

	// Validate thông tin
	if !govalidator.IsEmail(req.Email) {
		package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Định dạng email không hợp lệ"))
		return
	}

	var user package_model.User
	err = c.DB.Model(&user).Where("email = ?", req.Email).Select()
	if err != nil {
		package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Email chưa tồn tại trong hệ thống"))
		return
	}
	if !user.EmailConfirmed {
		package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Tài khoản của bạn chưa được kích hoạt. Vui lòng kích hoạt tài khoản để tiếp tục."))
		return
	}
	if !user.UserStatus {
		package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Tài khoản của bạn đã bị khoá. Vui lòng liên hệ với chúng tôi để được giải quyết."))
		return
	}

	// Lưu token
	uuidToken := uuid.NewV4()
	randomStr, errRanStr := helper.GenerateRandomString(20)
	if errRanStr != nil {
		package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Token không hợp lệ"))
		return
	}

	token := uuidToken.String() + randomStr
	tokenEnd := time.Now().Add(10 * time.Minute)
	_, err = c.DB.Exec(`UPDATE auth.users SET verify_email_token = ?, verify_email_token_end = ? WHERE id = ?`, token, tokenEnd, user.Id)
	if err != nil {
		return
	}

	var emailReq ForgotPasswordAccount
	emailReq.Id = user.Id
	emailReq.FullName = user.FullName
	emailReq.Email = user.Email
	emailReq.Token = token

	url := c.Config.ForgotMail
	c.SendMailHelper(url, emailReq)

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON("Gửi email hỗ trợ thành công")
}

func (c *Controller) SetPasswordForgot(ctx iris.Context) {
	var req model.SetPassword
	err := ctx.ReadJSON(&req)
	if err != nil {
		package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Định dạng request sai"))
		return
	}

	if len(req.Password) == 0 {
		package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Vui lòng nhập mật khẩu"))
		return
	}

	var user package_model.User
	err = c.DB.Model(&user).Where("id = ?", req.UserId).Select()
	if err != nil {
		package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Tài khoản chưa tồn tại trong hệ thống"))
		return
	}
	if !user.EmailConfirmed {
		package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Tài khoản chưa kích hoạt. Vui lòng kích hoạt để tiếp tục"))
		return
	}
	if !user.UserStatus {
		package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Tài khoản đang bị khoá. Hãy liên lạc với chúng tôi để được giải quyết"))
		return
	}
	if user.VerifyEmailToken != req.Token {
		package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Token không hợp lệ"))
		return
	}
	if user.VerifyEmailTokenEnd.Unix() < time.Now().Unix() {
		package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Token hết hạn"))
		return
	}

	// Hash password
	hashedPassword, salt := helper.HashPassword(req.Password)

	// Update thông tin
	user.Password = hashedPassword
	user.Salt = salt
	updatedColumns := []string{"password", "salt"}
	_, err = c.DB.Model(&user).Column(updatedColumns...).WherePK().Update()
	if err != nil {
		package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Không thể lưu mật khẩu. Vui lòng thử lại"))
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON("Đổi mật khẩu thành công")
}

func (c *Controller) SignUp(ctx iris.Context) {
	req := model.SignUpRequest{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Request sai định dạng"))
		return
	}
	p := bluemonday.StrictPolicy()

	// Validate thông tin
	if p.Sanitize(req.Email) == "" || req.Password == "" || p.Sanitize(req.FullName) == "" || p.Sanitize(req.Phone) == "" {
		package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Vui lòng nhập đủ thông tin"))
		return
	}
	if !govalidator.IsEmail(p.Sanitize(req.Email)) {
		package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Định dạng email không hợp lệ"))
		return
	}

	// Kiểm tra tài khoản theo Email
	var userID string
	rs, err := c.DB.Query(&userID, `
		SELECT id FROM auth.users WHERE email = ? `, req.Email)
	if err != nil {
		package_model.NewError(ctx, iris.StatusInternalServerError, errors.New("Không thể kiểm tra thông tin đăng ký"))
		return
	}
	if rs.RowsReturned() > 0 {
		package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Email đã tồn tại trong hệ thống. Vui lòng sử dụng email khác."))
		return
	}

	// Hash password
	hashedPassword, salt := helper.HashPassword(req.Password)

	tx, err := c.DB.Begin()
	if err != nil {
		package_model.NewError(ctx, iris.StatusInternalServerError, errors.New("Không thể kết nối cơ sở dữ liệu"))
		return
	}

	var user package_model.User
	user.Id = xid.New().String()
	user.FullName = p.Sanitize(req.FullName)
	user.Password = hashedPassword
	user.Salt = salt
	user.Phone = p.Sanitize(req.Phone)
	user.Email = p.Sanitize(strings.ToLower(req.Email))
	user.EmailConfirmed = false
	user.UserStatus = true
	user.Roles = []int32{2}
	err = user.InsertUser(tx, c.Config.SignupMail)
	if err != nil {
		tx.Rollback()
		package_model.NewError(ctx, iris.StatusInternalServerError, err)
		return
	}

	tx.Commit()

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON("Đăng ký thành công. Một email kích hoạt đã được gửi tới hòm thư của bạn")
}

type ForgotPasswordAccount struct {
	Id       string
	FullName string
	Email    string
	Token    string
}

func (c *Controller) ForgotPassword(ctx iris.Context) {
	req := model.ForgotPasswordRequest{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Request sai định dạng"))
		return
	}

	// Validate thông tin
	if !govalidator.IsEmail(req.Email) {
		package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Định dạng email không hợp lệ"))
		return
	}

	var user package_model.User
	err = c.DB.Model(&user).Where("email = ?", req.Email).Select()
	if err != nil {
		package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Email chưa tồn tại trong hệ thống"))
		return
	}
	if !user.EmailConfirmed {
		package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Tài khoản của bạn chưa được kích hoạt. Vui lòng kích hoạt tài khoản để tiếp tục."))
		return
	}
	if !user.UserStatus {
		package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Tài khoản của bạn đã bị khoá. Vui lòng liên hệ với chúng tôi để được giải quyết."))
		return
	}

	// Lưu token
	uuidToken := uuid.NewV4()
	randomStr, errRanStr := helper.GenerateRandomString(20)
	if errRanStr != nil {
		package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Token không hợp lệ"))
		return
	}

	token := uuidToken.String() + randomStr
	tokenEnd := time.Now().Add(10 * time.Minute)
	_, err = c.DB.Exec(`UPDATE auth.users SET verify_email_token = ?, verify_email_token_end = ? WHERE id = ?`, token, tokenEnd, user.Id)
	if err != nil {
		return
	}

	var emailReq ForgotPasswordAccount
	emailReq.Id = user.Id
	emailReq.FullName = user.FullName
	emailReq.Email = user.Email
	emailReq.Token = token

	url := c.Config.ForgotMail
	c.SendMailHelper(url, emailReq)

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON("Thành công")
}

func (c *Controller) GetSetPasswordPage(ctx iris.Context) {
	token := ctx.Params().GetStringDefault("token", "")
	userID := ctx.Params().GetStringDefault("userId", "")
	if len(token) == 0 || len(userID) == 0 {
		ctx.View("/error/error.html")
		return
	}

	var user package_model.User
	err := c.DB.Model(&user).Where("id = ?", userID).Select()
	if err != nil {
		ctx.View("/error/error.html")
		return
	}
	if !user.EmailConfirmed {
		ctx.View("/error/error.html")
		return
	}
	if !user.UserStatus {
		ctx.View("/error/error.html")
		return
	}
	if user.VerifyEmailToken != token {
		ctx.View("/error/error.html")
		return
	}
	if user.VerifyEmailTokenEnd.Unix() < time.Now().Unix() {
		// Hết hạn
		ctx.ViewData("is_expired", true)
	} else {
		ctx.ViewData("set_password", true)
		ctx.ViewData("token", token)
		ctx.ViewData("userId", userID)
	}

	ctx.View("change-password.html")
}

func (c *Controller) LoginGithub(ctx iris.Context) {
	code := ctx.URLParamDefault("code", "")
	if len(code) == 0 {
		package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Request sai định dạng"))
		return
	}

	config := &oauth2.Config{
		ClientID:     package_model.GITHUB_CLIENT_ID,
		ClientSecret: package_model.GITHUB_CLIENT_SECRET,
		Endpoint:     github.Endpoint,
		Scopes:       []string{"user", "user:email"},
	}
	accessToken, err := config.Exchange(context.Background(), code)
	if err != nil {
		HandlerError(c.DB, ctx, "error", errors.New("Token không hợp lệ."))
		return
	}

	client := config.Client(oauth2.NoContext, accessToken)
	resp, err := client.Get(package_model.GITHUB_OAUTH_URL)
	if err != nil {
		HandlerError(c.DB, ctx, "error", errors.New("Không lấy được thông tin người dùng từ Github."))
		return
	}
	defer resp.Body.Close()
	response, _ := ioutil.ReadAll(resp.Body)
	if err != nil {
		HandlerError(c.DB, ctx, "error", errors.New("Không lấy được thông tin người dùng từ Github."))
		return
	}

	var data model.GithubProfile
	err = json.Unmarshal(response, &data)
	if err != nil {
		HandlerError(c.DB, ctx, "error", errors.New("Không lấy được thông tin người dùng từ Github."))
		return
	}

	// Lấy thông tin private email
	if data.Email == "" {
		resp2, err := client.Get(package_model.GITHUB_OAUTH_EMAIL_URL)
		if err != nil {
			HandlerError(c.DB, ctx, "error", errors.New("Không lấy được thông tin người dùng từ Github."))
			return
		}
		defer resp2.Body.Close()

		emailInfo, _ := ioutil.ReadAll(resp2.Body)
		if err != nil {
			HandlerError(c.DB, ctx, "error", errors.New("Không lấy được thông tin người dùng từ Github."))
			return
		}

		var emails []model.EmailGithubInfo
		err = json.Unmarshal(emailInfo, &emails)
		if err != nil {
			HandlerError(c.DB, ctx, "error", errors.New("Không lấy được thông tin người dùng từ Github."))
			return
		}

		for _, email := range emails {
			if email.Primary == true {
				data.Email = email.Email
				data.EmailVerified = email.Verified
				break
			}
		}
	}

	tx, err := c.DB.Begin()
	if err != nil {
		HandlerError(c.DB, ctx, "error", errors.New("Lỗi kết nối cơ sở dữ liệu."))
		return
	}

	var user model.UserInfo
	// Kiểm tra tài khoản theo Email
	rs, err := tx.Query(&user, `
		SELECT id, email, avatar, full_name, roles, phone,
		email, email_confirmed, access_failed_count, lockout_end, user_status
		FROM  auth.users
		WHERE email = ? `, data.Email)
	if err != nil {
		tx.Rollback()
		HandlerError(c.DB, ctx, "error", errors.New("Không thể kiểm tra thông tin"))
		return
	}
	if rs.RowsReturned() == 0 {
		// Tài khoản chưa tồn tại trong hệ thống
		var userInsert package_model.User
		userInsert.Id = xid.New().String()
		if len(data.Name) == 0 {
			userInsert.FullName = data.Login
		} else {
			userInsert.FullName = data.Name
		}
		userInsert.Avatar = data.AvatarUrl
		userInsert.Email = data.Email
		userInsert.UserStatus = true
		userInsert.EmailConfirmed = true
		userInsert.Roles = []int32{2}
		userInsert.CreatedAt = time.Now()
		err = tx.Insert(&userInsert)
		if err != nil {
			tx.Rollback()
			HandlerError(c.DB, ctx, "error", errors.New("Không thể đăng nhập bằng Github"))
			return
		}

		copier.Copy(&user, &userInsert)
	} else {
		// Trường hợp tài khoản bị khóa
		if !user.UserStatus {
			tx.Rollback()
			HandlerError(c.DB, ctx, "error", errors.New("Không thể đăng nhập bằng Github"))
			return
		}
	}

	// Đồng bộ hoá giỏ hàng
	err = SynchronizeShoppingCart(user, tx, ctx)
	if err != nil {
		tx.Rollback()
		HandlerError(c.DB, ctx, "error", errors.New("Không thể đồng bộ hoá giỏ hàng"))
		return
	}

	tx.Commit()

	// Lấy ra các key mà user đã sở hữu
	var keys []string
	_, err = c.DB.Query(&keys, `
		SELECT distinct unnest(auth_role.keys) roles FROM auth.users users 
		INNER JOIN auth.role as auth_role ON auth_role.id = ANY(users.roles)
		where users.email = ? and auth_role.status = true`, user.Email)
	if err != nil {
		package_model.NewError(ctx, iris.StatusInternalServerError, errors.New("Không thể kiểm tra quyền hạn user hiện tại"))
		return
	}

	// Set session
	session := constant.Sess.Start(ctx)
	var auth = package_model.Authenticate{
		Authenticated: true,
		UserId:        user.Id,
		UserFullName:  user.FullName,
		Keys:          keys,
		UserEmail:     user.Email,
		UserAvatar:    user.Avatar,
		UserPhone:     user.Phone,
		ExpiredTime:   time.Now().Add(48 * time.Hour),
	}
	// set session
	session.Set("authenticate", auth)
	ctx.Redirect("/")
}

func (c *Controller) LoginFacebook(ctx iris.Context) {
	req := model.LoginFacebookRequest{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Request sai định dạng"))
		return
	}

	accessToken := req.AccessToken

	token := &oauth2.Token{
		AccessToken: accessToken,
	}
	config := &oauth2.Config{
		ClientID:     package_model.FACEBOOK_CLIENT_ID,
		ClientSecret: package_model.FACEBOOK_CLIENT_SECRET,
		Endpoint:     facebook.Endpoint,
		Scopes:       []string{"email", "name", "picture"},
	}
	client := config.Client(oauth2.NoContext, token)
	resp, err := client.Get(package_model.FACEBOOK_OAUTH_URL)
	if err != nil {
		package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Không thể đăng nhập bằng Facebook"))
		return
	}
	defer resp.Body.Close()

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Không thể đăng nhập bằng Facebook"))
		return
	}
	var data model.ProfileFacebook
	err = json.Unmarshal(response, &data)
	if err != nil {
		package_model.NewError(ctx, http.StatusBadRequest, errors.New("Không thể nhận thông tin từ Facebook."))
		return
	}

	if data.Email == "" {
		package_model.NewError(ctx, http.StatusBadRequest, errors.New("Không thể sử dụng Facebook để đăng nhập vì bạn chưa liên kết email cho Facebook."))
		return
	}

	tx, err := c.DB.Begin()
	if err != nil {
		HandlerError(c.DB, ctx, "error", errors.New("Lỗi kết nối cơ sở dữ liệu."))
		return
	}

	var user model.UserInfo
	// Kiểm tra tài khoản theo Email
	rs, err := tx.Query(&user, `
		SELECT id, email, avatar, full_name, roles, phone,
		email, email_confirmed, access_failed_count, lockout_end, user_status
		FROM  auth.users
		WHERE email = ? `, data.Email)
	if err != nil {
		tx.Rollback()
		package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Không thể kiểm tra thông tin"))
		return
	}

	if rs.RowsReturned() == 0 {
		// Tài khoản chưa tồn tại trong hệ thống
		var userInsert package_model.User
		userInsert.Id = xid.New().String()
		userInsert.FullName = data.Name
		userInsert.Avatar = data.Avatar.Data.Url
		userInsert.Email = data.Email
		userInsert.UserStatus = true
		userInsert.EmailConfirmed = true
		userInsert.Roles = []int32{2}
		userInsert.CreatedAt = time.Now()
		err = tx.Insert(&userInsert)
		if err != nil {
			tx.Rollback()
			package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Không thể đăng nhập bằng Facebook"))
			return
		}

		copier.Copy(&user, &userInsert)
	} else {
		// Trường hợp tài khoản bị khóa
		if !user.UserStatus {
			package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Tài khoản đã bị khóa."))
			return
		}
	}

	// Đồng bộ hoá giỏ hàng
	err = SynchronizeShoppingCart(user, tx, ctx)
	if err != nil {
		tx.Rollback()
		package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Không thể đồng bộ hoá giỏ hàng"))
		return
	}

	tx.Commit()

	// Lấy ra các key mà user đã sở hữu
	var keys []string
	_, err = c.DB.Query(&keys, `
		SELECT distinct unnest(auth_role.keys) roles FROM auth.users users 
		INNER JOIN auth.role as auth_role ON auth_role.id = ANY(users.roles)
		where users.email = ? and auth_role.status = true`, user.Email)
	if err != nil {
		package_model.NewError(ctx, iris.StatusInternalServerError, errors.New("Không thể kiểm tra quyền hạn user hiện tại"))
		return
	}

	// Set session
	session := constant.Sess.Start(ctx)
	var auth = package_model.Authenticate{
		Authenticated: true,
		UserId:        user.Id,
		UserFullName:  user.FullName,
		Keys:          keys,
		UserEmail:     user.Email,
		UserAvatar:    user.Avatar,
		UserPhone:     user.Phone,
		ExpiredTime:   time.Now().Add(48 * time.Hour),
	}
	// set session
	session.Set("authenticate", auth)

	if len(user.Phone) == 0 {
		ctx.JSON(false)
	} else {
		ctx.JSON(true)
	}

	ctx.StatusCode(iris.StatusOK)
}

func (c *Controller) LoginGoogle(ctx iris.Context) {
	req := model.LoginGoogleRequest{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString(err.Error())
		return
	}

	// Lấy AccessToken từ Code frontend gửi lên
	token := &oauth2.Token{
		AccessToken: req.AccessToken,
	}
	config := &oauth2.Config{
		ClientID:     package_model.GOOGLE_CLIENT_ID,
		ClientSecret: package_model.GOOGLE_CLIENT_SECRET,
		Endpoint:     google.Endpoint,
		RedirectURL:  "https://techmaster.vn",
		Scopes:       []string{"email", "profile"},
	}

	// Đọc thông tin user từ google apis
	client := config.Client(oauth2.NoContext, token)
	resp, err := client.Get(package_model.GOOGLE_OAUTH_URL)
	if err != nil {
		package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Không thể đăng nhập bằng Google"))
		return
	}
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	profile := model.GoogleProfile{}
	err = json.Unmarshal(data, &profile)
	if err != nil {
		package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Không thể đăng nhập bằng Google"))
		return
	}
	tx, err := c.DB.Begin()
	if err != nil {
		HandlerError(c.DB, ctx, "error", errors.New("Lỗi kết nối cơ sở dữ liệu."))
		return
	}

	var user model.UserInfo
	// Kiểm tra tài khoản theo Email
	rs, err := tx.Query(&user, `
		SELECT id, email, avatar, full_name, roles, phone,
		email, email_confirmed, access_failed_count, lockout_end, user_status
		FROM  auth.users
		WHERE email = ? `, profile.Email)
	if err != nil {
		tx.Rollback()
		package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Không thể kiểm tra thông tin"))
		return
	}

	if rs.RowsReturned() == 0 {
		// Tài khoản chưa tồn tại trong hệ thống
		var userInsert package_model.User
		userInsert.Id = xid.New().String()
		userInsert.FullName = profile.Name
		userInsert.Avatar = profile.Picture
		userInsert.Email = profile.Email
		userInsert.UserStatus = true
		userInsert.EmailConfirmed = true
		userInsert.Roles = []int32{2}
		userInsert.CreatedAt = time.Now()
		err = tx.Insert(&userInsert)
		if err != nil {
			tx.Rollback()
			package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Không thể đăng nhập bằng Google"))
			return
		}

		copier.Copy(&user, &userInsert)
	} else {
		// Trường hợp tài khoản bị khóa
		if !user.UserStatus {
			package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Tài khoản đã bị khóa."))
			return
		}
	}

	// Đồng bộ hoá giỏ hàng
	err = SynchronizeShoppingCart(user, tx, ctx)
	if err != nil {
		tx.Rollback()
		package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Không thể đồng bộ hoá giỏ hàng"))
		return
	}

	tx.Commit()

	// Lấy ra các key mà user đã sở hữu
	var keys []string
	_, err = c.DB.Query(&keys, `
		SELECT distinct unnest(auth_role.keys) roles FROM auth.users users 
		INNER JOIN auth.role as auth_role ON auth_role.id = ANY(users.roles)
		where users.email = ? and auth_role.status = true`, user.Email)
	if err != nil {
		package_model.NewError(ctx, iris.StatusInternalServerError, errors.New("Không thể kiểm tra quyền hạn user hiện tại"))
		return
	}
	// Set session
	session := constant.Sess.Start(ctx)
	var auth = package_model.Authenticate{
		Authenticated: true,
		UserId:        user.Id,
		UserFullName:  user.FullName,
		Keys:          keys,
		UserEmail:     user.Email,
		UserAvatar:    user.Avatar,
		UserPhone:     user.Phone,
		ExpiredTime:   time.Now().Add(48 * time.Hour),
	}
	// set session
	session.Set("authenticate", auth)

	if len(user.Phone) == 0 {
		ctx.JSON(false)
	} else {
		ctx.JSON(true)
	}

	ctx.StatusCode(iris.StatusOK)
}

func SynchronizeShoppingCart(user model.UserInfo, DB orm.DB, ctx iris.Context) error {
	cartID, err := GetCartIDFromCookie(ctx)
	if err == nil && cartID > 0 {
		// Có giỏ hàng, kiểm tra trong giỏ có sản phẩm không
		var count int32
		_, err = DB.Query(&count, `SELECT count(item_id) FROM sale.cart_item WHERE cart_id = ?`, cartID)
		if err != nil {
			return err
		}
		if count > 0 {
			// Có sản phẩm, mapping tài khoản của user với giỏ hàng trong cookie
			_, err = DB.Exec(`UPDATE sale.cart SET user_id = null WHERE user_id = ?`, user.Id)
			if err != nil {
				return err
			}
			_, err = DB.Exec(`UPDATE sale.cart SET user_id = ?, time = now() WHERE id = ?`, user.Id, cartID)
			if err != nil {
				return err
			}
		} else {
			// Không có sản phẩm, kiểm tra trong database có giỏ hàng của user không
			var oldCartId int64
			count, err := DB.Query(&oldCartId, `SELECT id FROM sale.cart WHERE user_id = ?`, user.Id)
			if err != nil {
				return err
			}
			if count.RowsReturned() == 0 {
				// Không có giỏ hàng trong database, mapping tài khoản của user với giỏ hàng trong cookie
				_, err = DB.Exec(`UPDATE sale.cart SET user_id = ?, time = now() WHERE id = ?`, user.Id, cartID)
				if err != nil {
					return err
				}
			} else {
				// Có giỏ hàng trong database, lấy giỏ hàng đó trả về
				_, err = DB.Exec(`UPDATE sale.cart SET time = now() WHERE id = ?`, oldCartId)
				if err != nil {
					return err
				}
				cartID = oldCartId
			}
		}
	} else {
		// Không có giỏ hảng, kiểm tra trong database có giỏ hàng của user không
		var oldCartId int64
		count, err := DB.Query(&oldCartId, `SELECT id FROM sale.cart WHERE user_id = ?`, user.Id)
		if err != nil {
			return err
		}
		if count.RowsReturned() == 0 {
			// Không có giỏ hàng trong database, tạo giỏ hàng mới
			var cart package_model.Cart
			cart.UserId = user.Id
			cart.Time = time.Now()
			err = DB.Insert(&cart)
			if err != nil {
				return err
			}
			cartID = cart.Id
		} else {
			// Có giỏ hàng trong database
			cartID = oldCartId
		}
	}

	SetCookieCartID(cartID, ctx)

	return nil
}

func (c *Controller) GetSiteCanAccess(ctx iris.Context) {
	var userID string
	userAuthen, isLogin := package_model.GetAuthentication(constant.Sess, ctx)
	if isLogin {
		userID = userAuthen.UserId
	} else {
		package_model.NewError(ctx, http.StatusBadRequest, errors.New("Bạn không có quyền"))
		return
	}

	var user package_model.User
	err := c.DB.Model(&user).Where("id = ?", userID).Select()
	if err != nil {
		package_model.NewError(ctx, http.StatusBadRequest, errors.New("Bạn không có quyền"))
		return
	}

	type CanAccessSite struct {
		Admin   bool
		Teacher bool
	}
	var rsp CanAccessSite

	for _, r := range user.Roles {
		if r != 2 {
			rsp.Admin = true
		}
		if r == 3 {
			rsp.Teacher = true
		}
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(rsp)
}

type InfomationOfStudentAdvisory struct {
	FullName     string
	Email        string
	Phone        string
	Note         string
	Link         string
	TypeItem     string
	NameItem     string
	UserReceiver []string
}

func (c *Controller) SubmitAdvisory(ctx iris.Context) {
	req := model.SubmitAdvisoryRequest{}
	err := ctx.ReadJSON(&req)
	if err != nil {
		package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Định dạng request sai"))
		return
	}
	p := bluemonday.StrictPolicy()

	if len(p.Sanitize(req.Link)) == 0 || len(p.Sanitize(req.FullName)) == 0 || len(p.Sanitize(req.Email)) == 0 || len(p.Sanitize(req.Phone)) == 0 {
		package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Vui lòng nhập đủ thông tin"))
		return
	}
	if !govalidator.IsEmail(req.Email) {
		package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Định dạng email không hợp lệ"))
		return
	}

	var advisory model.Advisory
	advisory.Id = xid.New().String()
	advisory.FullName = p.Sanitize(req.FullName)
	advisory.Link = p.Sanitize(req.Link)
	advisory.Email = p.Sanitize(req.Email)
	advisory.Phone = p.Sanitize(req.Phone)
	advisory.Info = p.Sanitize(req.Info)
	advisory.ItemId = req.ItemId
	advisory.CreatedAt = time.Now()
	advisory.IsProcessed = false
	advisory.IsEdit = false
	err = c.DB.Insert(&advisory)
	if err != nil {
		package_model.NewError(ctx, iris.StatusBadRequest, errors.New("Lỗi. Không thể gửi yêu cầu tư vấn."))
		return
	}

	// ======== Lấy thông tin gửi Email ========
	var nameItem string
	var supporters []model.SupporterInfo
	if req.Type == constant.COURSE_TYPE {
		_, err = c.DB.Query(&nameItem, `
		SELECT title
		FROM courses.course_package
		WHERE id = ?`, req.ItemId)

		if err != nil {
			ReturnErrorPage(c.DB, ctx, err.Error())
			return
		}

		_, err = c.DB.Query(&supporters, `
		SELECT users.id, users.email, users.full_name, users.phone , users.avatar
		FROM courses.course_package AS course, auth.users AS users
		WHERE users.id = ANY(course.supporters)
		AND course.id = ?`, req.ItemId)

		if err != nil {
			ReturnErrorPage(c.DB, ctx, err.Error())
			return
		}
		if len(supporters) == 0 {
			supporters = []model.SupporterInfo{}
		}
	}

	if req.Type == constant.TRACK_TYPE {
		_, err = c.DB.Query(&nameItem, `
		SELECT title
		FROM track.track
		WHERE id = ?`, req.ItemId)

		if err != nil {
			ReturnErrorPage(c.DB, ctx, err.Error())
			return
		}
		_, err = c.DB.Query(&supporters, `
		SELECT users.id, users.email, users.full_name, users.phone , users.avatar
		FROM track.track AS track, auth.users AS users
		WHERE users.id = ANY(track.supporters)
		AND track.id = ?`, req.ItemId)

		if err != nil {
			ReturnErrorPage(c.DB, ctx, err.Error())
			return
		}
		if len(supporters) == 0 {
			supporters = []model.SupporterInfo{}
		}
	}

	var supportersEmail []string
	for i := 0; i < len(supporters); i++ {
		supportersEmail = append(supportersEmail, supporters[i].Email)
	}

	// ======== Gửi Email ========
	var emailReq InfomationOfStudentAdvisory
	emailReq.FullName = advisory.FullName
	emailReq.Email = advisory.Email
	emailReq.Phone = advisory.Phone
	emailReq.Note = advisory.Info
	emailReq.Link = advisory.Link
	emailReq.NameItem = nameItem
	emailReq.UserReceiver = supportersEmail

	url := c.Config.AdvisoryMail
	c.SendMailHelper(url, emailReq)

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON("Gửi tư vấn thành công")
}

func (c *Controller) GetFormReserve(ctx iris.Context){
	
	_, isLogin := package_model.GetAuthentication(constant.Sess, ctx)
	if isLogin {
	} else {
		package_model.NewError(ctx, http.StatusBadRequest, errors.New("Bạn không có quyền"))
		return
	}
	classId := ctx.Params().Get("classId")
	userId := ctx.Params().Get("userID")
	type UserClassReserve struct{
		StartDate time.Time
		Reason template.HTML
		ClassName string
		NumberAttendance int
		Email string
		FullName string
		Session int
		TrackName string
		CourseName string
		Phone string
	}

	var userclass UserClassReserve
	_,err := c.DB.Query(&userclass,`SELECT reserve.start_date, reserve.reason, reserve.class_name, reserve.number_attendance,reserve.session,users.full_name,users.email,users.phone,courses.title AS course_name,track.title AS track_name 
	FROM class.bao_luu AS reserve 
	JOIN auth.users AS users ON users.id = reserve.user_id 
	JOIN class.class AS classes ON classes.id=reserve.class_id 
	LEFT JOIN courses.course_package AS courses ON courses.id = classes.course_id 
	LEFT JOIN track.track AS track ON courses.id = ANY(track.course_ids) 
	WHERE reserve.class_id = ? and users.id = ?`, classId,userId)

	if err != nil{
		ReturnErrorPage(c.DB, ctx, err.Error())
	}
	ctx.ViewData("info_reserve",userclass)
	ctx.View("/form_reserve.html")
}
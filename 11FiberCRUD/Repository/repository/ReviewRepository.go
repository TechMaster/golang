package repository

import (
	"errors"

	"github.com/TechMaster/golang/11FiberCRUD/Repository/model"
)

type ReviewRepository struct {
	reviews map[int64]*model.Review
	autoID  int64 //đây là biến đếm tự tăng gán giá trị cho id của Book

	//Danh sách các đối tượng quan sát được trừu tượng hoá qua interface Observer
	observerList []Observer
}

var ReviewRepo ReviewRepository //Khai báo biến toàn cục, global variable

func init() { //func init luôn chạy đầu tiên khi chúng ta import package
	ReviewRepo = ReviewRepository{autoID: 0}
	ReviewRepo.reviews = make(map[int64]*model.Review)
}

//Pointer receiver ~ method trong Java. Đối tượng chủ thể là *BookRepo
func (r *ReviewRepository) getAutoID() int64 {
	r.autoID += 1
	return r.autoID
}

func (r *ReviewRepository) CreateNewReview(review *model.Review) int64 {
	nextID := r.getAutoID() //giống trong CSDL quan hệ sequence.NETX_VAL()
	review.Id = nextID
	r.reviews[nextID] = review //tạo mới một phần tử trong map, gán key bằng nextID
	return nextID
}

func (r *ReviewRepository) GetAllReviews() map[int64]*model.Review {
	return r.reviews
}

func (r *ReviewRepository) FindReviewById(Id int64) (*model.Review, error) {
	if review, ok := r.reviews[Id]; ok {
		return review, nil //tìm được
	} else {
		return nil, errors.New("review not found")
	}
}

func (r *ReviewRepository) DeleteReviewById(Id int64) error {
	if _, ok := r.reviews[Id]; ok {
		delete(r.reviews, Id)
		return nil
	} else {
		return errors.New("book not found")
	}
}

func (r *ReviewRepository) UpdateReview(review *model.Review) error {
	if _, ok := r.reviews[review.Id]; ok {
		r.reviews[review.Id] = review
		return nil //tìm được
	} else {
		return errors.New("review not found")
	}
}

func (r *ReviewRepository) UpsertReview(review *model.Review) int64 {
	if _, ok := r.reviews[review.Id]; ok {
		r.reviews[review.Id] = review //tìm thấy thì update
		return review.Id
	} else { //không tìm thấy thì tạo mới
		return r.CreateNewReview(review)
	}
}

//--------------
/*
Khi một rewview thay đổi thì chạy làm hàm này để tính lại average book rating
*/
func (r *ReviewRepository) ComputeBookRatingWhenAReviewChage(review *model.Review) (bookId int64, averageRating float32) {
	bookId = review.BookId
	return bookId, r.ComputeBookRating(bookId)
}

func (r *ReviewRepository) ComputeBookRating(bookId int64) float32 {
	count := 0
	sum := 0
	for _, review := range r.reviews {
		if review.BookId == bookId {
			count++
			sum += review.Rating
		}
	}
	return float32(sum / count)
}

//-------------- Impement interface Publisher ----
func (r *ReviewRepository) RegisterObserver(o Observer) {
	r.observerList = append(r.observerList, o)
}

func (r *ReviewRepository) RemoveObserver(o Observer) {
	found := false
	i := 0
	for ; i < len(r.observerList); i++ {
		if r.observerList[i] == o {
			found = true
			break
		}
	}
	if found {
		r.observerList = append(r.observerList[:i], r.observerList[i+1:]...)
	}
}

func (r *ReviewRepository) NotifyObserver(id int64) {
	review := r.reviews[id]
	bookId, averageRating := r.ComputeBookRatingWhenAReviewChage(review)

	for _, observer := range r.observerList {
		observer.Update(bookId, averageRating)
	}
}

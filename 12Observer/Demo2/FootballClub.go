package main

// FootballClub đối tượng sẽ phát đi thông báo thay đổi: Publisher
type FootballClub struct {
	name string // tên câu lạc bộ bóng đá
	// Khi point thay đổi giá trị cần phát đi thông báo cho tất cả các đối tượng tuân thủ Observer
	point int
	//Danh sách các đối tượng quan sát được trừu tượng hoá qua interface Observer
	observerList []Observer
}

// Hàm constructor của FootballClub
func NewFootballClub(name string) *FootballClub {
	return &FootballClub{
		name:         name,
		point:        0,
		observerList: make([]Observer, 0),
	}
}

// Đăng ký một đối tượng sẽ nhận thông báo thay đổi
func (fbc *FootballClub) RegisterObserver(o Observer) {
	fbc.observerList = append(fbc.observerList, o)
}

// Loại bỏ đối tượng ra khỏi danh sách thông báo dạng slice
func (fbc *FootballClub) RemoveObserver(o Observer) {
	found := false
	i := 0
	for ; i < len(fbc.observerList); i++ {
		if fbc.observerList[i] == o {
			found = true
			break
		}
	}
	if found {
		fbc.observerList = append(fbc.observerList[:i], fbc.observerList[i+1:]...)
	}
}

// NotifyObserver thông báo đến tất cả các Observer trong danh sách
func (fbc *FootballClub) NotifyObserver() {
	for _, observer := range fbc.observerList {
		observer.Update(fbc.point)
	}
}

// SetValue khi thay đổi giá trị thì thông báo đến tất cả các Observer
func (fbc *FootballClub) SetPoint(point int) {
	fbc.point = point
	fbc.NotifyObserver()
}

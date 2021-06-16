package repository

type Publisher interface {
	RegisterObserver(o Observer)
	RemoveObserver(o Observer)
	NotifyObserver(id int64) //Thông báo khi có thay đổi ở bản ghi id
}

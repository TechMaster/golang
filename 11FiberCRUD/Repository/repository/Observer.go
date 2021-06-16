package repository

// Interface đối tượng nào muốn nhận được thông báo sẽ phải tuân thủ
type Observer interface {
	//Thông tin cần cập nhật ở đây là giá trị int
	Update(bookId int64, averageRating float32) error
}

package entity

//Person object for REST(CRUD)
type Users struct {
	User_id  int    //`gorm:"primary_key;not null"`
	F_name   string //`gorm:"not null"`
	L_name   string //`gorm:"not null"`
	Password string //`gorm:"not null"`
	Email    string //`gorm:"unique;not null"`
	Mo_no    string //`gorm:"not null"`
	Address  string //`gorm:"not null"`
}

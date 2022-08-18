package domain

type Role struct {
	Id   uint64 `json:"id" gorm:"primary_key:auto_increment"`
	Name string `json:"name" gorm:"type:varchar(20);not null, unique"`
}
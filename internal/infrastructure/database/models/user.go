package models

type User struct {
	BaseModel
	Email    string `gorm:"uniqueIndex;not null;size:255" json:"email"`
	Username string `gorm:"uniqueIndex;not null;size:100" json:"username"`
	Password string `gorm:"not null;size:255" json:"-"`
	FullName string `gorm:"size:255" json:"full_name"`
	Phone    string `gorm:"size:20" json:"phone"`
	Address  string `gorm:"size:500" json:"address"`
	City     string `gorm:"size:100" json:"city"`
	Province string `gorm:"size:100" json:"province"`
	ZipCode  string `gorm:"size:20" json:"zip_code"`
}

func (User) TableName() string {
	return "users"
}

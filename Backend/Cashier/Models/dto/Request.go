package dto

type RequestCashierRegister struct {
	Name     string `json:"name" gorm:"type:varchar(255)" validate:"required"`
	Phone    string `json:"phone" gorm:"type:varchar(255)" validate:"required"`
	Username string `json:"username" gorm:"type:varchar(255)" validate:"required,min=5"`
	Password string `json:"password" gorm:"type:varchar(255)" validate:"required,min=5"`
}

type RequestCashierLogin struct {
	Username string `json:"username" gorm:"type:varchar(255)" validate:"required,min=5"`
	Password string `json:"password" gorm:"type:varchar(255)" validate:"required,min=5"`
}

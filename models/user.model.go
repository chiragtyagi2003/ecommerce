package models

type User struct {
	ID        uint    `json:"id" gorm:"primary_key"`
	FirstName string  `json:"firstName" gorm:"type:varchar(75)"`
	LastName  string  `json:"lastName" gorm:"type:varchar(75)"`
	Email     string  `json:"email" gorm:"type:varchar(100)"`
	Password  *string `json:"password,omitempty" gorm:"type:varchar(200)"`
	Address   string  `json:"address" gorm:"type:varchar(200)"`
	Telephone string  `json:"telephone" gorm:"type:varchar(20)"`
	CartTotal float32 `json:"total" gorm:"type:float"`
	// Items     *[]CartItem `json:"items,omitempty" gorm:"foreignKey:UserId"`
}

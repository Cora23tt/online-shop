package onlinedilerv3

type User struct {
	ID          int     `json:"-" db:"id"`
	FirstName   string  `json:"first_name"`
	LastName    string  `json:"last_name"`
	Email       string  `json:"email" binding:"required"`
	Password    string  `json:"password" binding:"required"`
	PhoneNumber string  `json:"phone_number"`
	Latitude    float64 `json:"latitude"`
	Longtitude  float64 `json:"longtitude"`
}

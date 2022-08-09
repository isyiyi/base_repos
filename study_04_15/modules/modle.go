package modules


type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Patient struct {
	Name string `json:"name"`
	Gender string `json:"gender"`
	Age int `json:"age"`
	Address string `json:"address"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	Symptom string `json:"symptom"`
	Doctor string `json:"doctor"`
}
package models

type Student struct {
	Name        string  `json:"name"`
	Age         int     `json:"age"`
	GPA         float64 `json:"gpa"`
	Gender      string  `json:"gender"`
	YearJoined  int     `json:"year_joined"`
	CourseName  string  `json:"course_name"`
	Email       string  `json:"email"`
	PhoneNumber string  `json:"phone_number"`
	FeeBalance  float64 `json:"fee_balance"`
}

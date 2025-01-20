package domain

type User struct {
	ID       int    `json:"id" bson:"_id,omitempty"`
	Username string `json:"username" bson:"name"`
	Email    string `json:"email" bson:"email"`
}

type UserRepository interface {
	CreateUser(user *User) error
	GetUserByID(id int) (*User, error)
	UpdateUser(user *User) error
	DeleteUser(id int) error
}

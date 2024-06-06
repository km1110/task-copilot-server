package model

type User struct {
	ID          string
	FirebaseUID string
	Name        string
	Role        Role
	Status      bool
}

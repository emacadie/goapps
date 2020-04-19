package models

type User struct {
	ID int
	FirstName string
	LastName string
}

var (
	users []*User // slice holding pointers to User objects
	// by using pointers we can manipulate them without copying them around
	nextID = 1
	// could say nextID int32 = 1
)

func GetUsers() []* User {
	return users
}

func AddUser(u User) (User, error) {
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}



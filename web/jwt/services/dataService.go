package services

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

func GetUsers() map[string]string {
	return users
}

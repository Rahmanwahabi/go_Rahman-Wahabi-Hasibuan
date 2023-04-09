package main

import "fmt"

// Struct field untuk username dan password menggunakan tipe data string
type user struct {
	id       int
	username string
	password string
}

// Struct field userservice diganti namanya menjadi userService dan diberikan penjelasan
type userService struct {
	users []user // users adalah slice dari struct user
}

// Method getAllUsers tidak memerlukan parameter, variabel u diganti menjadi _ karena tidak digunakan
func (us userService) getAllUsers() []user {
	return us.users
}

// Method getUserByID mengembalikan error jika user dengan id tersebut tidak ditemukan
func (us userService) getUserByID(id int) (user, error) {
	for _, u := range us.users {
		if id == u.id {
			return u, nil
		}
	}

	// Mengembalikan error jika user tidak ditemukan
	return user{}, fmt.Errorf("user with ID %d not found", id)
}

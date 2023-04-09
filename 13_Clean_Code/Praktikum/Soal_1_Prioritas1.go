package main

// Struct field untuk username dan password menggunakan tipe data int, seharusnya menggunakan tipe data string
type user struct {
	id       int
	username int
	password int
}

// Struct field t tidak dijelaskan dengan baik, sebaiknya memberikan penjelasan
type userservice struct {
	t []user
}

// Method getallusers tidak memerlukan parameter, sebaiknya mengubah u menjadi _ karena variabel u tidak digunakan
func (u userservice) getallusers() []user {
	return u.t
}

// Method getuserbyid sebaiknya mengembalikan error jika user dengan id tersebut tidak ditemukan
func (u userservice) getuserbyid(id int) user {
	for _, r := range u.t {
		if id == r.id {
			return r
		}
	}

	// Mengembalikan struct user dengan nilai kosong jika user tidak ditemukan, sebaiknya mengembalikan error
	return user{}
}

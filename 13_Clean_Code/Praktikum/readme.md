SOAL 1 :
Terdapat beberapa kekurangan dalam penulisan kode tersebut:

Struct field untuk username dan password menggunakan tipe data int, seharusnya menggunakan tipe data string.
Struct field t pada userservice tidak dijelaskan dengan baik, sebaiknya memberikan penjelasan.
Method getallusers tidak memerlukan parameter, sehingga variabel u sebaiknya diganti menjadi _ karena tidak digunakan.
Method getuserbyid sebaiknya mengembalikan error jika user dengan id tersebut tidak ditemukan, bukan hanya mengembalikan struct user dengan nilai kosong.



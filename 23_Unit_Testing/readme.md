Unit testing adalah proses pengujian yang dilakukan pada level terkecil dalam pengembangan perangkat lunak, yaitu pada level unit. Dalam bahasa pemrograman Go (Golang), terdapat paket bawaan yang disebut "testing" yang menyediakan alat dan fungsi untuk melakukan unit testing secara efektif. Berikut adalah rangkuman tentang unit testing di Golang:

Membuat file pengujian: Untuk setiap file sumber kode Go yang ingin diuji, buat file terpisah dengan nama yang sama ditambahkan dengan "_test.go". Misalnya, jika ada file "calc.go" yang ingin diuji, buatlah file baru dengan nama "calc_test.go".

Import paket "testing": Di file pengujian, impor paket "testing" dengan menambahkan pernyataan import "testing" di awal file.

Membuat fungsi pengujian: Di file pengujian, buat fungsi-fungsi pengujian dengan menggunakan tipe data *testing.T sebagai parameter. Contoh: func TestAdd(t *testing.T) { ... }. Nama fungsi pengujian harus diawali dengan kata "Test" dan harus dieksekusi secara otomatis oleh framework pengujian.

Menggunakan fungsi-fungsi pengujian: Dalam fungsi pengujian, gunakan fungsi-fungsi pengujian yang disediakan oleh paket "testing" seperti t.Run, t.Error, t.Errorf, t.Fail, dan sebagainya untuk melakukan pengujian.

Menjalankan unit testing: Untuk menjalankan unit testing, Anda dapat menggunakan perintah go test di direktori yang berisi file-file pengujian. Go akan secara otomatis menemukan dan menjalankan semua fungsi pengujian yang sesuai.

Melihat hasil pengujian: Setelah unit testing selesai dijalankan, Anda akan melihat hasilnya di konsol. Go akan memberikan informasi tentang unit pengujian mana yang berhasil dan mana yang gagal.

Menulis assersi: Untuk memverifikasi bahwa hasil yang diperoleh sesuai dengan yang diharapkan, Anda dapat menggunakan fungsi assersi seperti t.Errorf untuk memeriksa kesalahan atau t.Fail untuk menggagalkan pengujian secara langsung.

Menangani pengujian yang kompleks: Go juga menyediakan fitur-fitur untuk menangani pengujian yang lebih kompleks, seperti tabel pengujian, subtes, setup dan teardown, mocking, dan banyak lagi.

Dengan menggunakan unit testing di Golang, Anda dapat meningkatkan keandalan dan kualitas perangkat lunak yang dikembangkan, serta memudahkan proses debugging dan perbaikan jika terjadi masalah.





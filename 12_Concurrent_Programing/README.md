Dalam Resume ini, saya akan membahas tentang 2 topik penting dalam pemrograman Go, yaitu Goroutine dan Channel.

Concurrent programming adalah gaya pemrograman yang memungkinkan beberapa tugas (tasks) berjalan secara bersamaan pada waktu yang sama. Pada Go (Golang), concurrent programming didukung secara langsung oleh bahasa tersebut dan sangat mudah untuk diterapkan.

Go memiliki fitur goroutine, yang merupakan unit eksekusi yang independen dari proses utama. Goroutine adalah fungsi yang berjalan secara asinkron, sehingga bisa berjalan secara bersamaan dengan goroutine lainnya. Go juga memiliki fitur channel, yaitu saluran yang digunakan untuk mengirim dan menerima data antar goroutine.

Dengan fitur goroutine dan channel, Go memungkinkan programmer untuk membangun program yang berjalan secara parallel dengan mudah. Sebagai contoh, jika terdapat sebuah program yang membutuhkan waktu cukup lama untuk mengeksekusi sebuah fungsi, maka fungsi tersebut dapat dijalankan pada sebuah goroutine yang terpisah. Dengan demikian, program tidak perlu menunggu fungsi tersebut selesai dieksekusi untuk melanjutkan eksekusi program yang lain.

Selain itu, Go juga mendukung pembuatan pool goroutine, yaitu sebuah kumpulan goroutine yang bisa digunakan secara bergantian untuk mengeksekusi beberapa tugas yang membutuhkan waktu eksekusi yang lama. Dengan pembuatan pool goroutine, penggunaan sumber daya CPU dapat dioptimalkan dan waktu eksekusi program dapat dipercepat.
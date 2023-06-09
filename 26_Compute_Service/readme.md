Compute Service di Golang adalah kemampuan untuk mengelola komputasi dan pemrosesan dalam bahasa pemrograman Go. Golang menyediakan berbagai fitur dan pustaka yang memungkinkan pengembang untuk membuat dan mengelola layanan komputasi yang efisien dan skalabel.

Beberapa poin penting tentang Compute Service di Golang:

Concurrency: Golang memiliki dukungan bawaan untuk konkurensi yang kuat. Dengan menggunakan goroutine dan channel, pengembang dapat membuat aplikasi yang dapat menjalankan banyak tugas secara paralel, meningkatkan kinerja dan efisiensi pemrosesan.

Pengelolaan Memori: Golang memiliki manajemen memori yang efisien dengan garbage collector (GC) yang canggih. GC secara otomatis mengelola alokasi dan pemulihan memori yang tidak lagi digunakan, mengurangi beban tugas pengembang.

Pustaka dan Framework: Golang memiliki berbagai pustaka dan framework yang membantu pengembang dalam membuat layanan komputasi. Beberapa pustaka populer termasuk net/http untuk pengembangan web, database/sql untuk interaksi dengan database, dan banyak lagi.

Skalabilitas: Golang dirancang untuk skalabilitas. Kemampuan konkurensi yang kuat dan manajemen memori yang efisien memungkinkan aplikasi untuk menangani lalu lintas yang tinggi dengan performa yang baik.

Deployment: Golang memungkinkan pengembang untuk dengan mudah mendeploy aplikasi mereka. Dalam Golang, aplikasi dikompilasi menjadi file biner yang dapat dieksekusi secara independen, membuat proses deployment menjadi lebih sederhana.

Secara keseluruhan, Compute Service di Golang memberikan pengembang alat yang kuat dan efisien untuk mengelola komputasi dan pemrosesan. Dengan dukungan konkurensi, manajemen memori yang baik, pustaka yang kaya, dan kemampuan deployment yang mudah, Golang menjadi pilihan yang baik untuk membangun layanan komputasi yang skalabel dan efisien.

disini saya menggunkan 2 platfrom yaitu :
AWS (Amazon Web Services):

AWS menyediakan layanan cloud computing yang luas, termasuk Amazon Elastic Compute Cloud (EC2) untuk menjalankan aplikasi.
Untuk melakukan deployment aplikasi Golang di AWS EC2, Anda dapat membuat instance EC2, mengonfigurasi lingkungan server, dan menginstal Golang pada instance tersebut.
Setelah itu, Anda dapat mengunggah kode Golang Anda ke instance EC2 dan menjalankannya dengan menggunakan perintah go run atau mengompilasi menjadi file biner dan menjalankannya secara langsung.
Google Cloud:

Google Cloud juga menyediakan layanan komputasi cloud, termasuk Google Compute Engine (GCE) untuk menjalankan aplikasi.
Untuk deployment aplikasi Golang di Google Cloud, Anda dapat membuat instance GCE, mengonfigurasi lingkungan server, dan menginstal Golang pada instance tersebut.
Selanjutnya, Anda dapat mengunggah kode Golang Anda ke instance GCE dan menjalankannya dengan menggunakan perintah go run atau mengompilasi menjadi file biner dan menjalankannya secara langsung.
Sekarang, baik di AWS maupun di Google Cloud, Anda juga dapat menggunakan layanan manajemen kontainer seperti AWS Elastic Container Service (ECS) atau Google Kubernetes Engine (GKE) untuk deployment yang lebih canggih. Dalam pendekatan ini, Anda akan membangun dan mengemas aplikasi Golang Anda menjadi container Docker yang dapat dijalankan di lingkungan yang dikelola secara otomatis oleh platform cloud.

Selain itu, AWS dan Google Cloud menyediakan beragam layanan lain yang dapat Anda manfaatkan untuk memperluas dan mengelola aplikasi Golang, seperti penyimpanan data, panggilan API, manajemen sumber daya, dan masih banyak lagi.

Penting untuk mencatat bahwa proses deployment dapat lebih kompleks tergantung pada persyaratan dan arsitektur aplikasi Anda. Jadi, pastikan Anda merujuk pada dokumentasi resmi dari AWS dan Google Cloud untuk panduan yang lebih mendalam dan detail mengenai deployment di platform mereka masing-masing.
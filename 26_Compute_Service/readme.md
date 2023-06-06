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
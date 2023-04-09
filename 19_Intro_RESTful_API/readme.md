RESTful API (Representational State Transfer API) adalah sebuah protokol untuk mengakses sumber daya (resource) pada suatu web service. RESTful API menggunakan protokol HTTP untuk melakukan pertukaran data dan komunikasi antara client dan server. Pada umumnya, RESTful API mendukung empat jenis HTTP method yaitu GET, POST, PUT, dan DELETE.

Berikut penjelasan tentang empat jenis HTTP method pada RESTful API:

GET
HTTP method GET digunakan untuk meminta informasi dari suatu resource yang tersedia pada server. Contoh penggunaan GET pada RESTful API adalah ketika seorang pengguna ingin mendapatkan informasi tentang suatu produk pada suatu e-commerce. HTTP request dengan method GET akan mengirimkan permintaan informasi produk yang diinginkan oleh pengguna, kemudian server akan merespon dengan mengirimkan data informasi produk tersebut pada body response.

POST
HTTP method POST digunakan untuk menambahkan suatu resource baru pada server. Contoh penggunaan POST pada RESTful API adalah ketika seorang pengguna ingin menambahkan suatu data baru pada suatu aplikasi. HTTP request dengan method POST akan mengirimkan data informasi yang ingin ditambahkan pada server, kemudian server akan merespon dengan mengirimkan status berhasil atau gagal pada body response.

PUT
HTTP method PUT digunakan untuk memperbarui suatu resource yang tersedia pada server. Contoh penggunaan PUT pada RESTful API adalah ketika seorang pengguna ingin memperbarui suatu data informasi pada suatu aplikasi. HTTP request dengan method PUT akan mengirimkan data informasi yang ingin diperbarui pada server, kemudian server akan merespon dengan mengirimkan status berhasil atau gagal pada body response.

DELETE
HTTP method DELETE digunakan untuk menghapus suatu resource yang tersedia pada server. Contoh penggunaan DELETE pada RESTful API adalah ketika seorang pengguna ingin menghapus suatu data pada suatu aplikasi. HTTP request dengan method DELETE akan mengirimkan permintaan untuk menghapus data pada server, kemudian server akan merespon dengan mengirimkan status berhasil atau gagal pada body response.

Dalam penggunaannya, setiap HTTP method pada RESTful API dapat memiliki parameter-parameter tertentu yang diperlukan sesuai dengan kebutuhan aplikasi. Selain itu, setiap HTTP method pada RESTful API juga memiliki response code yang digunakan untuk memberikan informasi mengenai status request yang dilakukan oleh client. Response code yang sering digunakan pada RESTful API antara lain adalah 200 (OK), 201 (Created), 204 (No Content), 400 (Bad Request), 401 (Unauthorized), 404 (Not Found), 500 (Internal Server Error), dan lain-lain.
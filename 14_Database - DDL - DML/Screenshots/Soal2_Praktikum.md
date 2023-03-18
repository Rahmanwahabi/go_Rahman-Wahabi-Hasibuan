1. Create database alta_online_shop.
<img src="https://i.ibb.co/yNxrXxc/Screenshot-2023-03-19-020948.png" width="512px"/>

2. Dari schema Olshop yang telah kamu kerjakan di, Implementasikanlah menjadi table pada MySQL.
    1. Create table user.
    2. Create table product, product type, operators, product description, payment_method.
    3. Create table transaction, transaction detail.
<img src="https://i.ibb.co/3mPzkdy/Screenshot-2023-03-19-021237.png" width="512px"/>
3. Create tabel kurir dengan field id, name, created_at, updated_at.
<img src="https://i.ibb.co/ScVMcxy/Screenshot-2023-03-19-021736.png" width="512px"/>
4. Tambahkan ongkos_dasar column di tabel kurir.
<img src="https://i.ibb.co/B6X5nxW/Screenshot-1860.png" width="512px"/>
5. Rename tabel kurir menjadi shipping.
<img src="https://i.ibb.co/SnFZ6tP/Screenshot-1861.png" width="512px"/>
6. Hapus / Drop tabel shipping karena ternyata tidak dibutuhkan.
<img src="https://i.ibb.co/k6K7kw2/Screenshot-1864.png" width="512px"/>
7. Silahkan menambahkan entity baru dengan relation 1-to-1, 1-to-many, many-to-many. Seperti:
    1. 1-to-1: payment method description.
    2. 1-to-many: user dengan alamat.
    3. many-to-many: user dengan payment method menjadi user_payment_method_detail.
    <img src="https://i.ibb.co/qyRx9mc/Screenshot-2023-03-19-022934.png" width="512px"/>
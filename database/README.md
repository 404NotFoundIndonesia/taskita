## DATABASE TASKITA

### EXPORT
Buka CLI, pastikan `mysql` sudah menyala, kemudian masukkan perintah berikut:
```shell
mysqldump -u NAMA_USER -p NAMA_DATABASE > taskita.sql
```

Apabila menggunakan mysql pada docker container silakan masukkan perintah:
```shell
docker exec NAMA_CONTAINER /usr/bin/mysqldump -u NAMA_USER --password=PASSWORD_DATABASE NAMA_DATABASE > taskita.sql
```

Pastikan direktori aktif ada di direktori ini dan silakan ganti `NAMA_USER` dan `NAMA_DATABASE` sesuai dengan yang ada di komputer Anda.


### IMPORT
Buka CLI, pastikan `mysql` sudah menyala, kemudian masukkan perintah berikut:
```shell
mysql -u NAMA_USER -p NAMA_DATABASE < taskita.sql
```

Apabila menggunakan mysql pada docker container silakan masukkan perintah:
```shell
cat taskita.sql | docker exec -i NAMA_CONTAINER /usr/bin/mysql -u NAMA_USER --password=PASSWORD_DATABASE NAMA_DATABASE
```

Pastikan direktori aktif ada di direktori ini dan silakan ganti `NAMA_USER` dan `NAMA_DATABASE` sesuai dengan yang ada di komputer Anda.

# Aplikasi RESTful API Untuk Mengelola Pembelian Dealer Mobil
Aplikasi Berbasis Restful API dengan bahasa pemrograman go yang berguna mengelola jual beli delaer mobil.
## Feature
- Aplikasi berbasis API yang siap digunakan pada sisi frontend
- Auth JWT
- Clean Code
## Tech
Aplikasi ini dibuat dengan :

- [GoLang](https://go.dev/) - Golang adalah singkatan dari Go Language (bahasa Go) yang merupakan bahasa pemrograman hasil garapan Google, dan dikelola langsung oleh Google.
- [GIN](https://gin-gonic.com/) - Gin adalah sebuah framework web berbahasa pemrograman Go (atau disebut juga Golang) yang digunakan untuk membangun aplikasi web dan RESTful API dengan cepat dan efisien. Gin dirancang untuk menjadi ringan, cepat, dan memiliki performa tinggi, sehingga cocok untuk membangun aplikasi web yang skala tinggi dan dapat menangani banyak permintaan dalam waktu singkat.
- [GORM](https://gorm.io/index.html) - GORM adalah sebuah library (pustaka) berbahasa pemrograman Go (Golang) yang digunakan sebagai Object-Relational Mapping (ORM) atau pustaka pemetaan objek ke basis data. Dengan GORM, Anda dapat dengan mudah berinteraksi dengan basis data relasional seperti MySQL, PostgreSQL, SQLite, dan sebagainya, tanpa perlu menulis perintah SQL secara langsung.
- [JWT GO](https://github.com/golang-jwt/jwt) - JWT adalah singkatan dari JSON Web Token, yang merupakan standar terbuka (open standard) untuk mentransmisikan informasi dalam bentuk token yang aman antara pihak-pihak tertentu. JWT biasanya digunakan untuk melakukan autentikasi dan pertukaran informasi antara sistem yang berbeda secara aman.
- [Viper](https://github.com/spf13/viper) - Viper adalah sebuah pustaka (library) berbahasa pemrograman Go (Golang) yang digunakan untuk mengelola konfigurasi aplikasi. Viper menyediakan cara yang mudah dan fleksibel untuk membaca, menulis, dan mengelola berbagai jenis konfigurasi aplikasi, seperti konfigurasi dalam bentuk file, variabel lingkungan (environment variables), atau sumber konfigurasi lainnya.

## Installation

Untuk menggunakan aplikasi ini dapat menggunakan clone atau donwload
```sh
git clone https://github.com/loethfii/Car-Dealer-Project-go.git
```

#### Instalasi dependesi / package
```sh
go mod tidy
```

#### Mengatur Environment
pastikan membuat file .env copy text yang berada di .env_example ke .env
##### penjelasan .env :
|
```sh
DB_USER= //berisi username database
DB_PASSWORD= //berisi password database kosongkan jika tidak ada
DB_HOST= // Host locak komputer bisa disisi localhost atau 127.0.0.1
DB_NAME= // Nama Database
JWT_KEY= // Jwt key yang berisi string acak untuk mengaman encrypsi jwt
```

## Jalankan Aplikasi
Jalankan aplikasi ini dengan pemerintah `go run main.go`
Disarankan menggunakan golang versi 20.xx karna saat membuat aplikasi ini penulis menggunakan versi tersebut.

## Daftar API
Untuk mengetahui daftar API dapat mengimport file bernama `Import PostMan.json` ke postman.

Link Download POST MAN [here](https://www.postman.com/)
## Kredit

Dibuat Oleh Arif Luthfi Romadhoni

[Klik Untuk Melihat Linkedin Saya](https://www.linkedin.com/in/arif-luthfi-romadhoni-a48149204/)




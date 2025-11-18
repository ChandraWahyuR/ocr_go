README SEMENTARA ALIAS SEINGAT CHANDRA BUAT GRPC

CONFIGURASI FILE PROTO DULU, lalu jalankan perintah buat file pb.

File pb tadi akan muncul method(di go kur method njay) panggil dan olah di folder mana saja bebas, saran dekat sama folder pb.
Kemudian karena struktur code saya yang independent dan akan di panggil kebutuhannya di app, tinggal init dial nya jika server ini buat fetch doang. kalau listen ya teruskan saja perintah code pb tadi di main, mirip init rest api di main.go.

SOAL 1
Struct Data:
Menyimpan dua array input (arr1 dan arr2), hasil operasi logika (hasilOR, hasilAND, hasilXOR), dan jumlah elemen (n).
Fungsi isiArray:
Mengisi array dengan input dari pengguna.
Memvalidasi bahwa hanya angka 0 atau 1 yang diterima.
Fungsi hitungOperasi:
Menghitung hasil operasi logika OR, AND, dan XOR untuk setiap elemen pada kedua array.
Fungsi cetakHasil:
Mencetak hasil operasi logika untuk setiap elemen.
Menampilkan total jumlah elemen yang menghasilkan nilai 1.

SOAL 2
Struct ArrayData:
Menyimpan elemen array dalam slice elements.
Fungsi Utama dalam Struct:
fillArray(size int): Mengisi array dengan elemen yang dimasukkan pengguna.
display(filter func(int) bool): Menampilkan elemen array berdasarkan kriteria tertentu menggunakan fungsi filter.
deleteAtIndex(index int): Menghapus elemen array pada indeks tertentu.
average(): Menghitung rata-rata nilai elemen array.
stdDev(): Menghitung standar deviasi elemen array.
frequency(target int): Menghitung frekuensi kemunculan elemen tertentu dalam array

SOAL 3
Struct Match:
Menyimpan informasi tentang pertandingan:
TeamA: Nama klub A.
TeamB: Nama klub B.
ScoreA: Skor klub A.
ScoreB: Skor klub B.
Winner: Pemenang pertandingan (bisa "Draw" jika seri).
Fungsi main:
Input Nama Klub: Program meminta nama klub A dan B.
Loop untuk Input Skor: Program terus meminta skor pertandingan antara kedua klub sampai pengguna memasukkan skor negatif (yang menghentikan input).
Skor valid (positif) dibandingkan, dan pemenang ditentukan berdasarkan skor yang lebih tinggi.
Hasil pertandingan disimpan dalam slice matches sebagai struct Match.
Output Hasil Pertandingan: Setelah input selesai, program menampilkan hasil dari setiap pertandingan (nama klub, skor, dan pemenang).
Output Daftar Klub Pemenang: Program menampilkan daftar klub yang memenangkan pertandingan (tidak termasuk hasil seri).

SOAL 4
Struct Table:
arr: Array dengan panjang MAX untuk menyimpan karakter-karakter (tipe data rune digunakan untuk mendukung karakter Unicode).
n: Variabel yang menyimpan jumlah elemen (panjang string) yang dimasukkan oleh pengguna.
Fungsi isiArray:
Fungsi ini meminta pengguna untuk memasukkan sebuah teks (string).
Teks tersebut diubah menjadi array karakter (rune) yang disimpan dalam field arr pada struct Table.
n diatur untuk mencatat panjang string yang dimasukkan.
Fungsi cetakArray:
Fungsi ini menampilkan karakter-karakter yang ada di dalam array arr pada struct Table.
Fungsi ini mengubah setiap rune menjadi string agar bisa dicetak dengan benar.
Fungsi balikkanArray:
Fungsi ini membalikkan urutan elemen-elemen array arr.
Digunakan teknik swapping dua elemen di array menggunakan dua indeks, i dan j, yang bergerak dari dua ujung array ke tengah.
Fungsi isPalindrome:
Fungsi ini memeriksa apakah array arr membentuk palindrome dengan membandingkan elemen pertama dengan elemen terakhir, elemen kedua dengan elemen kedua dari belakang, dan seterusnya.

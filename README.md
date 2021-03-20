# 3-17-2020

1. Buat fungsi finalGrade, yang menghitung nilai akhir siswa bergantung pada dua parameter: nilai untuk ujian dan jumlah proyek yang diselesaikan.
   Fungsi ini harus mengambil dua argumen: nilai ujian untuk ujian (dari 0 hingga 100); proyek - jumlah proyek yang diselesaikan (dari 0 ke atas);

   Fungsi ini harus mengembalikan angka (nilai akhir). Ada empat jenis nilai akhir:
   100, jika nilai ujian lebih dari 90 atau jika jumlah proyek yang diselesaikan lebih dari 10.
   90, jika nilai ujian lebih dari 75 dan jika jumlah proyek yang diselesaikan minimal 5.
   75, jika nilai ujian lebih dari 50 dan jika jumlah proyek yang diselesaikan minimal 2.
   0, dalam kasus lain

   Contoh:
   finalGrade(100, 12); // 100
   finalGrade(99, 0); // 100
   finalGrade(10, 15); // 100
   finalGrade(85, 5); // 90
   finalGrade(55, 3); // 75
   finalGrade(55, 0); // 0
   finalGrade(20, 2); // 0

2. Buat fungsi (atau tulis skrip di Shell) yang menggunakan bilangan bulat sebagai argumen dan mengembalikan "Genap" untuk bilangan genap atau "Ganjil" untuk bilangan ganjil.

3. Sangat sederhana, diberi angka, temukan kebalikannya.
   contoh :
   1: -1
   14: -14
   -34: 34

4. Sum of positive :
   Anda mendapatkan serangkaian angka, mengembalikan jumlah semua angka positif.
   Contoh [1, -4,7,12] => 1 + 7 + 12 = 2
   Catatan: jika tidak ada yang dijumlahkan, jumlah defaultnya adalah 0.

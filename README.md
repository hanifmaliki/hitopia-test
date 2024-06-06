# Hitopia Test

# Balanced Bracket

## Kompleksitas

Fungsi `balancedBracket` memiliki kompleksitas waktu O(n), di mana n adalah panjang string input. Hal ini karena setiap char dalam string diproses tepat satu kali.

Kompleksitas ruang juga O(n) dalam kasus terburuk, di mana semua char dalam string adalah bracket pembuka `(, {, [` yang perlu disimpan di dalam stack.

## Detail
- Kita mendeklarasikan sebuah `stack` untuk menyimpan char bracket pembuka.
- Kita menggunakan map `bracketPairs` untuk mencocokkan bracket penutup dengan bracket pembuka yang sesuai.
- Kita mengiterasi setiap char dalam input string.
Jika char adalah bracket pembuka `(, {, [`, kita menambahkannya ke stack.
- Jika char adalah bracket penutup `), }, ]`, kita periksa apakah stack kosong atau apakah char terakhir pada stack tidak cocok dengan bracket penutup saat ini menggunakan map `bracketPairs`. Jika tidak cocok atau stack kosong, kita kembalikan "NO".
- Jika cocok, kita hapus char terakhir dari stack.
- Setelah iterasi selesai, jika stack kosong, berarti semua bracket seimbang dan kita kembalikan "YES", jika tidak, kita kembalikan "NO".
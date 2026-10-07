# LeetCode Go

Solusi soal LeetCode menggunakan Go.

## Struktur

```
arrays/
  p0001_twosum/
    solution.go       # solusi
    solution_test.go  # test
```

Setiap soal ada di folder `p<nomor>_<nama_soal>`.

## Menjalankan test

```bash
# semua soal
go test ./...

# satu soal
go test -v ./arrays/p0001_twosum/
```

## Daftar soal

| No | Soal | Tingkat |
|---|---|---|
| 1 | [Two Sum](https://leetcode.com/problems/two-sum/) | Easy |
| 146 | [LRU Cache](https://leetcode.com/problems/lru-cache/) | Medium |
| 217 | [Contains Duplicate](https://leetcode.com/problems/contains-duplicate/) | Easy |
| 242 | [Valid Anagram](https://leetcode.com/problems/valid-anagram/) | Easy |
| 303 | [Range Sum Query - Immutable](https://leetcode.com/problems/range-sum-query-immutable/) | Easy |
| 1929 | [Concatenation of Array](https://leetcode.com/problems/concatenation-of-array/) | Easy |

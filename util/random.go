package util

import (
    "math/rand"
    "strings"
    "time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

var rng *rand.Rand

func init() {
    // Inisialisasi generator angka acak dengan sumber yang Anda tentukan sendiri
    seed := time.Now().UnixNano()
    rng = rand.New(rand.NewSource(seed))
}

// RandomInt generate a random integer between min and max
func RandomInt(min, max int64) int64 {
    return min + rng.Int63n(max-min+1)
}

func RandomString(n int) string {
    var sb strings.Builder
    k := len(alphabet)

    for i := 0; i < n; i++ {
        c := alphabet[rng.Intn(k)]
        sb.WriteByte(c)
    }

    return sb.String()
}

func RandomOwner() string {
    return RandomString(6)
}

func RandomMoney() int64 {
    return RandomInt(0, 1000)
}

func RandomCurrency() string {
    currencies := []string{"EUR", "USD", "IDR"}
    n := len(currencies)
    return currencies[rng.Intn(n)]
}

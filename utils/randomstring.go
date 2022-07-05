package utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func Randomstring() string {
	rand.Seed(time.Now().UTC().UnixNano())
	random := rand.Int()
	angka := []rune(fmt.Sprintf("%d", random))
	b := make([]rune, 8)
	for i := range b {
		b[i] = angka[rand.Intn(len(angka))]
	}
	hasil := string(b)
	inthasil, _ := strconv.Atoi(hasil)
	strhasil := strconv.Itoa(inthasil)
	return strhasil
}

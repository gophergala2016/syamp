package random

import (
	"time"
	"math/rand"
)

// Generator
func random(chn chan string) {
	alphnrk := []string {
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n",
		"o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "1", "2", "3",
		"4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F", "G", "H", "I",
		"J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "V", "W", "X", "Y",
		"Z",
	}
	max := len(alphnrk)
	 go func() {
		 for {
		 	rand.Seed(time.Now().UnixNano())
		 	num := rand.Intn(max)
			chn <- alphnrk[num]
		}
	 }()
}

func RandStr(digits int) string {
	c := make(chan string)
	var str string
	go random(c)
	for i := 0; i < digits; i++ {
		str += <-c
	}

	return str
}

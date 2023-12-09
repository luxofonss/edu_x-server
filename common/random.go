package common

import (
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func randSequence(n int) string {
	b := make([]rune, n)

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	for i := range b {
		b[i] = letters[r1.Intn(999999)%len(letters)]
	}

	return string(b)
}

func GetSalt(length int) string {
	if length < 0 {
		length = 50
	}

	return randSequence(length)
}

func GenCourseCode(length int) string {
	if length < 0 {
		length = 6
	}

	return randSequence(length)
}

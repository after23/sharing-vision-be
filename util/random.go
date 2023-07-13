package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"
var status = []string{"published", "draft", "thrashed"}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomString(n int) string {
	var sb strings.Builder
	l := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(l)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomTitle() string {
	return randomString(10)
}

func RandomContent() string {
	return randomString(50)
}

func RandomStatus() string {
	l := len(status)
	return status[rand.Intn(l)]
}

func RandomCategory() string {
	return randomString(5)
}
package util

import "math/rand"

func RandString(n int) string {
	var leffters = []byte("abcdefghijklmnopqrstuvwxyz")
	result := make([]byte, n)
	for i := range result {
		result[i] = leffters[rand.Intn(len(leffters))]
	}
	return string(result)
}

package utils

import "math/rand"

const letterByters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStrihngBytesRmndr(n int) []byte {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterByters[rand.Int63()%int64(len(letterByters))]
	}
	return b
}

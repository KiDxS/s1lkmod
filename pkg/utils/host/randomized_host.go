package host

import (
	"math/rand"
	"time"
)

func RandomizedId() (hostID string) {
	characters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	rand.Seed(time.Now().Unix())
	result := make([]byte, 24)
	for i := range result {
		result[i] = characters[rand.Intn(len(characters))]
	}
	hostID = string(result)
	return
}

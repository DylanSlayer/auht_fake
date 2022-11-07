package utils

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func GetRandomItem(list interface{}) string {
	rand.Seed(time.Now().Unix())
	index := rand.Intn(len(list.([]string)))
	return list.([]string)[index]
}
func GetNumberBetween(a, b int) int {
	rand.Seed(time.Now().Unix())
	return a + rand.Intn(b-a)
}
func GetRandomTel() string {
	rand.Seed(time.Now().Unix())
	telFirst := strings.Split("134,135,136,137,138,139,150,151,152,157,158,159,130,131,132,155,156,133,153", ",")
	first := GetRandomItem(telFirst)
	second := strconv.Itoa(GetNumberBetween(1, 888) + 10000)
	third := strconv.Itoa(GetNumberBetween(1, 9100) + 10000)
	return first + second + third
}

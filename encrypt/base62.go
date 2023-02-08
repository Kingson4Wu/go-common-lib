package encrypt

import (
	"math"
	"math/rand"
	"strings"
	"time"
)

var chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func Random62Factor() string {
	rand.Seed(time.Now().UnixNano())

	charArr := []rune(chars)
	size := 62

	for i := range charArr {
		r := rand.Intn(size)
		charArr[i], charArr[r+i] = charArr[r+i], charArr[i]
		size--
	}

	return string(charArr)
}

func Encode62(num int64) string {
	return Encode62WithFactor(num, chars)
}

func Decode62(str string) int64 {
	return Decode62WithFactor(str, chars)
}

func Encode62WithFactor(num int64, factor string) string {
	if len(factor) != 62 {
		return ""
	}

	var bytes []byte
	for num > 0 {
		bytes = append(bytes, factor[num%62])
		num = num / 62
	}
	reverse(bytes)
	return string(bytes)
}

func Decode62WithFactor(str string, factor string) int64 {
	if len(factor) != 62 {
		return 0
	}
	var num int64
	n := len(str)
	for i := 0; i < n; i++ {
		pos := strings.IndexByte(factor, str[i])
		num += int64(math.Pow(62, float64(n-i-1)) * float64(pos))
	}
	return num
}

func reverse(a []byte) {
	for left, right := 0, len(a)-1; left < right; left, right = left+1, right-1 {
		a[left], a[right] = a[right], a[left]
	}
}

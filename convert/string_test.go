package convert

import (
	"fmt"
	"testing"
)

func TestBytesToString(t *testing.T) {
	b := []byte("xx")
	s := BytesToString(b)
	fmt.Println(s)
}

func TestStringToBytes(t *testing.T) {

}

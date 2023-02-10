package encrypt_test

import (
	"fmt"
	"github.com/kingson4wu/go-common-lib/encrypt"
	"testing"
)

func TestEncode62(t *testing.T) {
	fmt.Println(encrypt.Encode62(78653623482744824))
	fmt.Println(encrypt.Decode62("5oEX9EWxw0"))
	fmt.Println(encrypt.Encode62(78653623482744825))
	fmt.Println(encrypt.Decode62("5oEX9EWxw1"))
	fmt.Println(encrypt.Encode62(88653623482744825))
	fmt.Println(encrypt.Decode62("6Y28kuzEnB"))
}

func TestEncode62Random(t *testing.T) {

	factor := encrypt.Random62Factor()
	fmt.Println(factor)
	result := encrypt.Encode62WithFactor(78653623482744824, factor)
	fmt.Println(result)
	fmt.Println(encrypt.Decode62WithFactor(result, factor))
}

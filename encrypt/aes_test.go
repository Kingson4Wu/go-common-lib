package encrypt_test

import (
	"fmt"
	"github.com/kingson4wu/go-common-lib/encrypt"
	"testing"
	"unicode/utf8"
)

func ExampleEncryptByAesWithKey() {

	in := "kxw"
	actual, _ := encrypt.EncryptByAesWithKey(in, "ABCDABCDABCDABCD")

	fmt.Println(actual)
}

func TestEncryptByAesWithKey(t *testing.T) {
	var (
		in       = "kxw"
		expected = "aDBA5g4yEryqeDty8fuW3A=="
	)
	actual, _ := encrypt.EncryptByAesWithKey(in, "ABCDABCDABCDABCD")
	if actual != expected {
		t.Errorf("EncryptByAesWithKey(%s) = %s; expected %s", in, actual, expected)
	}
}

//go test ./... Run the test cases in the current directory and all subdirectories

func TestDecryptByAesWithKey(t *testing.T) {
	var (
		in       = "aDBA5g4yEryqeDty8fuW3A=="
		expected = "kxw"
	)
	actual, _ := encrypt.DecryptByAesWithKey(in, "ABCDABCDABCDABCD")
	if actual != expected {
		t.Errorf("DecryptByAesWithKey(%s) = %s; expected %s", in, actual, expected)
	}
}

func BenchmarkEncryptByAesWithKey(b *testing.B) {
	var (
		in = "kxw"
	)
	b.ResetTimer()

	encrypt.EncryptByAesWithKey(in, "ABCDABCDABCDABCD")
}

//go test -bench=. -run=^$

func FuzzEncryptByAesWithKey(f *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testcases {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}
	f.Fuzz(func(t *testing.T, orig string) {
		rev, _ := encrypt.EncryptByAesWithKey(orig, "ABCDABCDABCDABCD")
		doubleRev, _ := encrypt.DecryptByAesWithKey(rev, "ABCDABCDABCDABCD")
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}

//go test -v -run=Coverage
//go test -run=Coverage -coverprofile=c.out

//go test -cover

//go test -cpuprofile=cpu.out

func TestEncrypt(t *testing.T) {

	in := "iTC\n\n\n"
	actual, _ := encrypt.EncryptByAesWithKey(in, "ABCDABCDABCDABCD")

	fmt.Println(actual)
}

func TestDecrypt(t *testing.T) {

	in := "MuXeSxz2yXnyRc="
	actual, _ := encrypt.DecryptByAesWithKey(in, "ABCDABCDABCDABCD")

	fmt.Println(actual)
}

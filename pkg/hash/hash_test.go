package hash

import (
	"fmt"
	"testing"
)

func TestGenerateHash(t *testing.T) {
	hashPass, err := PasswordHash("123456")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(hashPass)
}

func TestVerifyHash(t *testing.T) {
	hashPass, err := PasswordHash("123456")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(hashPass)
	very := PasswordVerify("123456", hashPass)
	if !very {
		fmt.Println(true)
		return
	}
	fmt.Println(false)
}

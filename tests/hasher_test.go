package tests

import (
	"gingorm/utils"
	"testing"
)

func TestHasher(t *testing.T) {
	password := "123456"
	hashed, _ := utils.HashPassword(password)
	println("password is: ", password)
	println("hashed is: ", hashed)

	result := utils.CheckPasswordHash(password, hashed)
	if !result {
		t.Errorf("sifre ve hash uyusmuyor")
	}

	t.Log("result is: ", result)

	falseResult := utils.CheckPasswordHash("wrong_password", hashed)
	if falseResult {
		t.Errorf("yanlis sifre ve hash uyusuyor")

	}

	println("result is: ", result)
	println("wrong result is: ", falseResult)
}

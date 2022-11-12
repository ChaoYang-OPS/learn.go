package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

var val = validator.New()

type RegisterRequest struct {
	UserName   string `validate:"gt=0"`
	Password   string `validate:"min=6,max=12"`
	PassRepeat string `validate:"eqfield=Password"`
	Email      string `validate:"email"`
}

type InfoRequest struct {
	UserName string `validate:"gt=0"`
	Password string `validate:"eqcsfield=Reg.Password"`
	Reg      RegisterRequest
}

func main() {
	req := RegisterRequest{
		UserName:   "TF",
		Password:   "12345678",
		PassRepeat: "12345678",
		Email:      "123qq.com",
	}
	if err := val.Struct(req); err != nil {
		fmt.Println(err) // Key: 'RegisterRequest.Email' Error:Field validation for 'Email' failed on the 'email' tag
	}
	info := InfoRequest{
		UserName: "123",
		Password: "1234",
	}
	// cross struct need check and impl
	info2 := InfoRequest{
		UserName: "123",
		Password: "12345678",
		Reg:      req,
	}
	if err := val.Struct(info); err != nil {
		fmt.Println(err) // Key: 'InfoRequest.Password' Error:Field validation for 'Password' failed on the 'eqcsfield' tag

	}
	if err := val.Struct(info2); err != nil {
		fmt.Println(err) // Key: 'InfoRequest.Password' Error:Field validation for 'Password' failed on the 'eqcsfield' tag
	}
}

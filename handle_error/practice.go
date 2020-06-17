package main

import (
	"fmt"
)

//======Define Error Variable=====//
type Error interface {
	full_message()
	valid() bool
}

type Standard struct {
	code    int
	message string
}

func (s Standard) full_message() {
	fmt.Printf("{ code: %v, message: %v}", s.code, s.message)
}

func (s Standard) valid() bool {
	if s.code == 200 {
		return true
	}
	return false
}

func New(code int, message string) Error {
	return Standard{code, message}
}

var (
	NotFound = New(404, "Not Found")
	Success  = New(200, "Success")
	Conflic  = New(409, "Conflic")
)

//======End Define Error Variable=====//

type ActiveRecord interface {
	update(params map[string]string) Error
}

type User struct {
	name  string
	phone string
}


func (u *User) update(params map[string]string) Error {
	u.name = params["name"]
	u.phone = params["phone"]

	return Success
}

func main() {
	user := User{"Nguyen thanh cong", "0338529345"}

	var model ActiveRecord = &user
	var params = map[string]string{"name": "Tran thi lien", "phone": "1234535"}

	fmt.Println(user)

	updated := model.update(params)
	
	if updated.valid() {
		fmt.Println("Cập nhật thành công")
	} else {
		updated.full_message()		
	}
}

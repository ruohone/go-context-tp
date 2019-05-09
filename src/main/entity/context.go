package entity

import "time"

type MyContext struct {
	Id string `json:"id"`
}

type LoginContext struct {
	MyContext
	Phone string `json:"phone"`
	Code  string `json:"code"`
}

type SaveLoginContext struct {
	MyContext
	Phone      string    `json:"phone"`
	Code       string    `json:"code"`
	CreateDate time.Time `json:"createDate"`
}

package main

type MyResponse struct {
	ID    string      `json:"id"`
	Stuff interface{} `json:"stuff"`
}

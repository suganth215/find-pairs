package main

// Request struct
type Request struct {
	Numbers []int `json:"numbers"`
	Target  int   `json:"target"`
}

// Response struct
type Response struct {
	Solutions [][]int `json:"solutions"`
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

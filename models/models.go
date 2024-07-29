package models

type Student struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int32  `json:"age"`
}

type Test struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

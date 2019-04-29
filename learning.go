package main

import "fmt"

type (
	User struct {
		Id   int
		Name string
	}
	Employee struct {
		User
		Company string
	}
)

func (e Employee) showDetail() {
	fmt.Printf("%d, %s, %s\n", e.Id, e.Name, e.Company)
}

func main() {
	pat := Employee{}
	pat.Id = 620245
	pat.Name = "Thawatchai Singngam"
	pat.Company = "KTB"

	pat.showDetail()
}

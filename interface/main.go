package main

import "fmt"

type printer interface {
	PrintOutInfo()
}

type book struct {
	Name  string
	Price float32
}

type drink struct {
	Name  string
	Price float32
}

func (b book) PrintOutInfo() {
	fmt.Printf("\nBook: Name(%s) Price(%.2f)\n", b.Name, b.Price)
}

func (d drink) PrintOutInfo() {
	fmt.Printf("Book: Name(%s) Price(%.2f)\n", d.Name, d.Price)
}

func main() {
	goWithWeb := book{
		Name:  "GoWithWeb",
		Price: 12.2,
	}

	coffee := drink{
		Name:  "coffee",
		Price: 4.2,
	}

	goWithWeb.PrintOutInfo()
	coffee.PrintOutInfo()

	info := []printer{goWithWeb, coffee}

	info[0].PrintOutInfo()
	info[1].PrintOutInfo()
}

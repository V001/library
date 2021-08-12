package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	fmt.Println("Hello, playground")
	type Inventory struct {
		Material int32
		Count    uint
		Array    []int32
	}
	sweaters := Inventory{25, 17, []int32{1, 2, 3}}
	tmpl, err := template.New("test").Parse(`
		{{ if and (eq .Age 33) (eq .Name “Ruslan”)


		
	 `)
	if err != nil {
		print(err)
	}
	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil {
		panic(err)
	}

}

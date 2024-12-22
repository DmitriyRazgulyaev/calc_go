package main

import "github.com/DmitriyRazgulyaev/calc_go/internal/application"

func main() {
	app := application.New()
	app.RunServer()
}

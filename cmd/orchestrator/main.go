package main

import "github.com/DmitriyRazgulyaev/calc_go/internal/orchestrator"

func main() {
	app := orchestrator.New()
	app.RunServer()
}

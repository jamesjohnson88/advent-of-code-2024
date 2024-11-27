package main

import (
	"github.com/jamesjohnson88/advent-of-code-2024/day-1/handlers"
	"log"
)

func main() {
	log.Print("Go module running...")

	taskOneResult := handlers.PerformTaskOne("data-1.txt")
	log.Printf("Task One Result: %v", taskOneResult)

	log.Print("Exiting module...")
}

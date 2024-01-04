package main

import "fmt"

func main() {
	lab := Dog{"Lab", "Yellow", 10}
	retriever := Dog{"Retriever", "Golden", 4}

	fmt.Printf("Breed: %v\nWeight: %v\n", lab.Breed, lab.Age)
	lab.Sleep()
	fmt.Printf("Breed: %v\nWeight: %v\n", retriever.Breed, retriever.Age)
	retriever.Awake()
}

type Dog struct {
	Breed string
	Color string
	Age   int
}

func (d Dog) Sleep() {
	fmt.Printf("%s is sleeping\n", d.Breed)
}

func (d Dog) Awake() {
	fmt.Printf("%s is awake\n", d.Breed)
}

func (d Dog) Bark() {
	fmt.Printf("%s is barking\n", d.Breed)
}

func (d Dog) Eating() {
	fmt.Printf("%s is eating\n", d.Breed)
}

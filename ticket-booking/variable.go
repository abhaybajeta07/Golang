package main

import "fmt"

func main() {
	var name string = "Abhay"
	surname := "Bajeta"
	var job string = "DevSecOps Engineer"
	fmt.Println(name, surname)
	fmt.Print("hi I am\n", name, "\n", surname, "is my surname\n")
	fmt.Printf("I am a %v and my name is %v", job, name)

}

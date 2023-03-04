package alert

import "fmt"

func Println(msg string) {
	fmt.Println(msg)
}

func Print(msg string) {
	fmt.Print(msg)
}

func Error(err error) {
	fmt.Println("Error:", err.Error())
	panic(err)
}

func Warn(msg string) {
	fmt.Println("Warn:", msg)
}

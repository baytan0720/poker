package alert

import "fmt"

func Print(msg string) {
	fmt.Println(msg)
}

func Error(err error) {
	fmt.Println("Error:", err.Error())
	panic(err)
}

func Warn(msg string) {
	fmt.Println("Warn:", msg)
}

package alert

import (
	"fmt"
	"os"
)

func Println(msg string) {
	fmt.Println(msg)
}

func Print(msg string) {
	fmt.Print(msg)
}

func Error(err error) {
	fmt.Println("Error:", err.Error())
	os.Exit(1)
}

func Warn(msg string) {
	fmt.Println("Warn:", msg)
}

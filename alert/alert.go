package alert

import (
	"fmt"
	"os"
)

func Error(msg string) {
	fmt.Println("Error", msg)
}

func Warm(msg string) {
	fmt.Println("Warning", msg)
}

func Fatal(msg string) {
	fmt.Println("Error:", msg)
	os.Exit(0)
}

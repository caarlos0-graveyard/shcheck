package status

import (
	"fmt"

	"github.com/fatih/color"
)

// Success status
func Success(file string) {
	print(" OK ", file, color.FgGreen)
}

// Fail status
func Fail(file string) {
	print("FAIL", file, color.FgRed)
}

// Info status
func Info(file string) {
	print(" ?? ", file, color.FgWhite)
}

func print(text, file string, col color.Attribute) {
	fmt.Printf(
		"[%s] %s\n",
		color.New(col).Sprint(text),
		color.New(color.Bold).Sprint(file),
	)
}

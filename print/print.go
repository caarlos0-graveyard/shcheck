package print

import "fmt"

func Success(file string) {
	fmt.Printf("\r  [ \033[00;32mOK\033[0m ] %s...\n", file)
}

func Fail(file string) {
	fmt.Printf("\r  [\033[0;31mFAIL\033[0m] %s...\n", file)
}

func Info(file string) {
	fmt.Printf("\r  [ \033[00;34m??\033[0m ] %s", file)
}
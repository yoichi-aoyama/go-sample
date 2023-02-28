package main

import "fmt"

func main() {
	fmt.Println(helloMessage("Hello Go"))
}

func helloMessage(msg string) string {
	return fmt.Sprintf("%s hello.", msg)
}

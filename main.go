package main

import (
	"fmt"
	str "strings"

	"example.com/numbers"
	"example.com/strings"
	"example.com/strings/greetings"
)

func main() {
	fmt.Println(numbers.IsPrime(19))

	fmt.Println(greetings.Welcome)

	fmt.Println(strings.Reverse("callicoder"))

	fmt.Println(str.Count("Go is Awesome. I love Go", "Go"))
}

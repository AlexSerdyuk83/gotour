package main

import (
	"fmt"
	"gotour/pkg"
)

func main() {
	// fmt.Println("My favorite number is", rand.Intn(10))
	// fmt.Printf("Now you have %g problems.\n", math.Round(math.Sqrt(7)))

	// fmt.Println(pkg.Sum(5, 7))

	// a, b := pkg.Swap("hello", "world")
	// fmt.Println(a, b)

	// fmt.Println(pkg.DomainForLocale("site.com", ""))
	// fmt.Println(pkg.DomainForLocale("site.com", "by"))

	// fmt.Println(pkg.ModifySpaces("site com", "dash"))
	// fmt.Println(pkg.ModifySpaces("site com", "www"))

	// fmt.Println(pkg.ErrorMessageToCode("OK"))
	fmt.Println(pkg.ErrorMessageToCode("CANCELLED"))
	fmt.Println(pkg.ErrorMessageToCode("err"))
}

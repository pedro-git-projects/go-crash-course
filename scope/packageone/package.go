package packageone

import "fmt"

var privateVar string = "I am private"
var PublicVar string = "I am public (or exported)"

func notExported() {

}

func Exported() {
	fmt.Println("Exported function")
}

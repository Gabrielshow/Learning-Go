// example of a program which shows the operating system on which it runs. 
// it has a local string variable getting its value by calling the Getnenv function (which is used to obtain environment-variables) from the os-package

package main
import (
    "fmt"
	"os"
)

func main() {
	var goos string = os.Getenv("GOOS")
	fmt.Printf("The operating system is %s\n", goos)
	path := os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)
}
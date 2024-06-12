package main 

import "fmt"

func main() {
    var username string = "mojo";
    var password string = "brilliancy";

    fmt.Println("Authorization Basic," username +":"+ password);
}
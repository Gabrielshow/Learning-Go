// if multiple packages are needed, they can be
// imported by a separate statement
// e.g
// import "fmt"
// import "os"
// or: import "fmt"; import "os"
// using the shorter way
// import ( "fmt"
// "os"
// )

// it can even be shorter:
// import ("fmt" ; "os") but gofmt enforces the distributed version

//visibility rule
// identifiers which start with a lowercase is not visible outside the package, but
// they are visible and usable in the whole package(like private)
// Suppose we have a thing(variable or function) called Thing(starts wiht T so it is exported)
// in a package pack1, then when pack1 is imported in the current package, Thing can be called using the usual
// dot-notation from OO-language: pack1.Thing (The pack1. is necessary!)

// packages also serve as namespaces and can help to avoid name-clashes (name-conflicts);
// variables with the same name in tow packages are differentiated by their package name, like:
// pack1.Thing and pack2.Thing

// A package can, if this is useful(for shortening, name conflicts), also be given another name(an alias), like
// import fm "fmt". the alias is then used in the following code:

package main

import fm "fmt" //alias

func main() {
	fm.Println("hello, world")
}

// importing a packagae which is not used in the rest of the code is a build error(for example: imported and not used: os).
// tis follows the Go-motto: "no unnecessary code!"

// the first { must be on the same line as the func-declaration}
// for small functions it is allowed that everything is written on one line, like for example:
// func Sum(a, b int) int { return a + b }
// schematically a general function looks like:
// func funtionName(parameter_list) (return_value_list) {
// -
// }
//  where parameter_list is of the form (param1 type1, param2 type2, ...)
// and return_value_list is of the form ( ret1 type1, ret2 type2, ..)

// Function names only start with a capital letter when the funciton has to be used outside the package
// then they follow PascalCasing, otherwise they follow camelCasing: every new word in the name stars with a capital letter.

// it is possible to define an alias for an existing type, like in:
// type IZ int
// var a IZ = 7
// having more types to define, you acn use the factored keyword form, as in:
// type (
// IZ int
// FZ float
// STR string
// )

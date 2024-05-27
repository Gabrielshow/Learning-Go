// Init functions
// Apart from global declarations with initialization, variables can also be initialized in init() function.
// <!-- This is a special funciton with the name init() which cannot be called, bu is executed automatically before the main() function in package main or at teh start of the import of the package that contain it -->

package trans
import “math”

var Pi float64
func init() {
 Pi = 4 * math.Atan(1) // init() function computes Pi
}
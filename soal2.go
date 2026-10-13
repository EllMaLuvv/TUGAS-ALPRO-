package main
import "fmt"
func main () {
var tb, bmi float64
var bb float64
fmt.Scan(&bmi, &tb)
bb = (bmi * (tb * tb))
fmt.Printf("%.0f", bb)
}
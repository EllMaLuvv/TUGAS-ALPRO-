package main
import "fmt"
 
func main() {
	var buy, disc int

	fmt.Scan(&buy, &disc)
	buy = buy - (buy * disc / 100)

	fmt.Println(buy)
}
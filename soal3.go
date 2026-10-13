package main
import (
	"fmt"
	"math"
)

func main() {
	var ax, ay, bx, by, cx, cy float64
	var ab, bc, ac, terpanjang float64

fmt.Scan(&ax, &ay, &bx, &by, &cx, &cy)

ab = math.Sqrt(math.Pow(bx-ax, 2) + math.Pow(by-ay, 2))
bc = math.Sqrt(math.Pow(cx-bx, 2) + math.Pow(cy-by, 2))
ac = math.Sqrt(math.Pow(cx-ax, 2) + math.Pow(cy-ay, 2))
terpanjang = math.Max(math.Max(ab, bc), ac)

fmt.Printf("%.0f", terpanjang)

}

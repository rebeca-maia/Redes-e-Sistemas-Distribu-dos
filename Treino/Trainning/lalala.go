package main
import "fmt"

func main() {
	var A, B, C float64

	fmt.Scanf("%.1f\n", &A)
	fmt.Scanf("%.1f\n", &B)
	fmt.Scanf("%.1f\n", &C)

	if A >= 0.0 && A <= 10.0 && B >= 0.0  && B <= 10.0 && C >= 0.0 && C<=10.0{
		MEDIA := ((2.0 * A) + (3.0 * B) + (5.0 * C)) / 10.0
		fmt.Printf("MEDIA = %.1f\n",MEDIA)
	}
}
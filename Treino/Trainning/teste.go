package main

import "fmt"

func main(){
	var A,B,C,MEDIA float64

	fmt.Scanf("%.1f\n", &A)
	fmt.Scanf("%.1f\n",&B)
	fmt.Scanf("%.1f\n",&C)

	if A >= 0 && A <= 10 && B >= 0 && C >= 0 && B <= 10 {
		MEDIA = ((2 * A) + (3 * B) + (5 * C)) / 10
	}
	fmt.Printf("MEDIA = %.1f\n",MEDIA)
}
package generics

import "fmt"

func ExampleMax() {
	ivals := []int{2, 4, 5, 1, 12, 11, 25, 29, 15, 14, 23, 23, 19}

	fmt.Println(Max(ivals))

	fvals := []float64{15, 14, 23, 23, 19}

	fmt.Println(Max(fvals))

	svals := []string{"z", "orange", "watermelon", "apple"}

	fmt.Println(Max(svals))

	// Output:
	// 29 <nil>
	// 23 <nil>
	// z <nil>
}

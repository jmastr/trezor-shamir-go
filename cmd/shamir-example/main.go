package main

import (
	"fmt"

	shamir "github.com/jmastr/trezor-shamir-go"
)

func main() {
	example := shamir.Shamir{
		ResultIndex:  1,
		ShareIndices: []uint8{5, 1, 10},
		ShareValues: [][]uint8{
			{129, 18, 104, 86, 236, 73, 176},
			{91, 188, 226, 91, 254, 197, 225},
			{69, 53, 151, 204, 224, 37, 19},
		},
		ShareCount: 3,
		Length:     7,
	}
	result, err := shamir.Interpolate(example.ResultIndex, example.ShareIndices, example.ShareValues, example.ShareCount, example.Length)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

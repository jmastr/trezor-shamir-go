package main

import (
	"fmt"

	shamir "github.com/jmastr/trezor-shamir-go"
)

type Shamir struct {
	resultIndex  uint8
	shareIndices []uint8
	shareValues  [][]uint8
	shareCount   uint8
	length       uint8
}

func main() {
	example := Shamir{
		resultIndex:  1,
		shareIndices: []uint8{5, 1, 10},
		shareValues: [][]uint8{
			{129, 18, 104, 86, 236, 73, 176},
			{91, 188, 226, 91, 254, 197, 225},
			{69, 53, 151, 204, 224, 37, 19},
		},
		shareCount: 3,
		length:     7,
	}
	result, err := shamir.Interpolate(example.resultIndex, example.shareIndices, example.shareValues, example.shareCount, example.length)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

package shamir

/*
#include "./trezor-firmware/crypto/memzero.c"
#include "./trezor-firmware/crypto/shamir.c"

bool interpolate(uint8_t *result,
			uint8_t result_index,
			const uint8_t *share_indices,
			const uint8_t *share_values_1d,
			uint8_t share_count,
			size_t length) {
	// convert 1D share_values_1d array to 2D array
	const uint8_t *share_values[share_count];
	uint8_t share_values_temp[share_count][length];
	int c = 0;
	for (int i = 0; i < share_count; i++) {
		for (int j = 0; j < length; j++) {
			share_values_temp[i][j] = share_values_1d[c];
			c++;
		}
		share_values[i] = share_values_temp[i];
	}
	// calculate
	return shamir_interpolate(result, result_index, share_indices, share_values, share_count, length);
}
*/
import "C"

import (
	"bytes"
	"errors"
)

var ErrInterpolateFailed = errors.New("interpolate failed")

// Interpolate computes f(x) given the Shamir shares (x_1, f(x_1)), ... , (x_m, f(x_m)).
//
// See `./trezor-firmware/crypto/shamir.h` for more info.
func Interpolate(resultIndex uint8, shareIndices []uint8, shareValues [][]uint8, shareCount uint8, length uint8) (result []uint8, err error) {
	// convert 2D shareValues array to 1D array, because cgo cannot handle 2D arrays well
	// https://stackoverflow.com/a/45299068
	shareValues1D := bytes.Join(shareValues, nil)
	// convert to C
	result = make([]uint8, length)
	resultC := (*C.uint8_t)(&result[0])
	resultIndexC := (C.uint8_t)(resultIndex)
	shareIndicesC := (*C.uint8_t)(&shareIndices[0])
	shareValues1DC := (*C.uint8_t)(&shareValues1D[0])
	shareCountC := (C.uint8_t)(shareCount)
	lengthC := (C.size_t)(length)
	// calculate
	r, err := C.interpolate(resultC, resultIndexC, shareIndicesC, shareValues1DC, shareCountC, lengthC)
	ret := bool(r)
	if !ret {
		result = []uint8{0} // to be compatible with trezor
		err = ErrInterpolateFailed
	}

	return
}

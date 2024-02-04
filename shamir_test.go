package shamir_test

import (
	"strconv"
	"testing"

	shamir "github.com/jmastr/trezor-shamir-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInterpolate(t *testing.T) {
	t.Parallel()
	// test vectors can be found in `./trezor-firmware/crypto/tests/test_check.c` line 6185
	testCases := []struct {
		result       []uint8
		resultIndex  uint8
		shareIndices []uint8
		shareValues  [][]uint8
		shareCount   uint8
		length       uint8
		ret          bool
	}{
		{
			result:       []uint8{7, 151, 168, 57, 186, 104, 218, 21, 209, 96, 106, 152, 252, 35, 210, 208, 43, 47, 13, 21, 142, 122, 24, 42, 149, 192, 95, 24, 240, 24, 148, 110},
			resultIndex:  0,
			shareIndices: []uint8{2},
			shareValues: [][]uint8{
				{7, 151, 168, 57, 186, 104, 218, 21, 209, 96, 106, 152, 252, 35, 210, 208, 43, 47, 13, 21, 142, 122, 24, 42, 149, 192, 95, 24, 240, 24, 148, 110},
			},
			shareCount: 1,
			length:     32,
			ret:        true,
		},
		{
			result:       []uint8{53},
			resultIndex:  255,
			shareIndices: []uint8{14, 10, 1, 13, 8, 7, 3, 11, 9, 4, 6, 0, 5, 12, 15, 2},
			shareValues: [][]uint8{
				{114},
				{41},
				{116},
				{67},
				{198},
				{109},
				{232},
				{39},
				{90},
				{241},
				{156},
				{75},
				{46},
				{181},
				{144},
				{175},
			},
			shareCount: 16,
			length:     1,
			ret:        true,
		},
		{
			result:       []uint8{91, 188, 226, 91, 254, 197, 225},
			resultIndex:  1,
			shareIndices: []uint8{5, 1, 10},
			shareValues: [][]uint8{
				{129, 18, 104, 86, 236, 73, 176},
				{91, 188, 226, 91, 254, 197, 225},
				{69, 53, 151, 204, 224, 37, 19},
			},
			shareCount: 3,
			length:     7,
			ret:        true,
		},
		{
			result:       []uint8{0},
			resultIndex:  1,
			shareIndices: []uint8{5, 1, 1},
			shareValues: [][]uint8{
				{129, 18, 104, 86, 236, 73, 176},
				{91, 188, 226, 91, 254, 197, 225},
				{69, 53, 151, 204, 224, 37, 19},
			},
			shareCount: 3,
			length:     7,
			ret:        false,
		},
		{
			result:       []uint8{0},
			resultIndex:  255,
			shareIndices: []uint8{3, 12, 3},
			shareValues: [][]uint8{
				{100, 176, 99, 142, 115, 192, 138},
				{54, 139, 99, 172, 29, 137, 58},
				{216, 119, 222, 40, 87, 25, 147},
			},
			shareCount: 3,
			length:     7,
			ret:        false,
		},
		{
			result:       []uint8{163, 120, 30, 243, 179, 172, 196, 137, 119, 17},
			resultIndex:  3,
			shareIndices: []uint8{1, 0, 12},
			shareValues: [][]uint8{
				{80, 180, 198, 131, 111, 251, 45, 181, 2, 242},
				{121, 9, 79, 98, 132, 164, 9, 165, 19, 230},
				{86, 52, 173, 138, 189, 223, 122, 102, 248, 157},
			},
			shareCount: 3,
			length:     10,
			ret:        true,
		},
	}

	for index, testCase := range testCases {
		index, testCase := index, testCase
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			t.Parallel()
			result, err := shamir.Interpolate(testCase.resultIndex, testCase.shareIndices, testCase.shareValues, testCase.shareCount, testCase.length)
			if testCase.ret {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
				assert.Equal(t, "interpolate failed", shamir.ErrInterpolateFailed.Error())
			}
			assert.Equal(t, testCase.result, result)
		})
	}
}

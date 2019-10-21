package encoder

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIsNumeric(t *testing.T) {
	require.Equal(t, isNumeric('a'), false, "a failed")
	require.Equal(t, isNumeric('0'), true, "0 failed")
	require.Equal(t, isNumeric('9'), true, "9 failed")
	require.Equal(t, isNumeric('5'), true, "5 failed")
	require.Equal(t, isNumeric('.'), false, ". failed")
	require.Equal(t, isNumeric('-'), false, "- failed")
	require.Equal(t, isNumeric(' '), false, "' ' failed")
}

func concat(str1 string, err error) string {
	if err != nil {
		return (str1 + err.Error())
	}
	return str1
}

func TestEncoder(t *testing.T) {
	require.Equal(t, concat(EncodeString("a4bc2d5e")), "aaaabccddddde", "a4bc2d5e failed")
	require.Equal(t, concat(EncodeString("abcd")), "abcd", "abcd failed")
	require.Equal(t, concat(EncodeString("45")), "first character can't be number", "error test failed")

	// (*)
	
	require.Equal(t, concat(EncodeString("qwe\\4\\5")), "qwe45", "qwe45 failed")
	require.Equal(t, concat(EncodeString("qwe\\45")), "qwe44444", "qwe44444 failed")
	require.Equal(t, concat(EncodeString("qwe\\\\5")), "qwe\\\\\\\\\\", "qwe\\\\\\\\\\ failed")
}

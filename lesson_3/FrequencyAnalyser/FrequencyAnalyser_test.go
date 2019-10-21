package FrequencyAnalyser

import (
	"github.com/stretchr/testify/require"
	"testing"
)
func TestGetRawString(t *testing.T) { 
	require.Equal(t, getRawString("Тест"), "тест", "Lower string test failed")
	require.Equal(t, getRawString("Тест."), "тест", "Punctuation mark test failed")
	require.Equal(t, getRawString("Тест.:;"), "тест", "Punctuation mark test2 failed")
	require.Equal(t, getRawString("Тест.:;а"), "тест.:;а", "Punctuation mark test3 failed") 

} 
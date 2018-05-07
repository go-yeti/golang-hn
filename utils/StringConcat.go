package utils

import (
	"bytes"
	"strconv"
)

// efficient way to concatenate strings
func StringConcat(words ...string) string {
	var buffer bytes.Buffer
	for _, word := range words {
		buffer.WriteString(word)
	}
	return buffer.String()
}

func IntArray2String(intArray []int, delim string) string {
	var buffer bytes.Buffer
	for i := 0; i < len(intArray); i++ {
		buffer.WriteString(strconv.Itoa(intArray[i]))
		if i != len(intArray)-1 {
			buffer.WriteString(delim)
		}
	}

	return buffer.String()
}

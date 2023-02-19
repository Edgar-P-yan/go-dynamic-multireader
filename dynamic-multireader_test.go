package dynamicmultireader

import (
	"io"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var data string = "Lorem ipsum dolor sit amet, consectetur adipiscing elit," +
	"sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. " +
	"Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. " +
	"Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. " +
	"Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."

func TestDynamicMultiReaderWithStringReaderAndLimitReader(t *testing.T) {
	originalReader := strings.NewReader(data)

	chunkSize := 444
	readSum := 0

	dmr := DynamicMultiReader(func(_ int) io.Reader {
		if readSum >= len(data) {
			return nil
		}

		readSum += chunkSize
		return io.LimitReader(originalReader, int64(chunkSize))
	})

	out, _ := ioutil.ReadAll(dmr)

	assert.Equal(t, data, string(out))
}

func TestReturnNil(t *testing.T) {
	dmr := DynamicMultiReader(func(_ int) io.Reader {
		return nil
	})

	out, _ := ioutil.ReadAll(dmr)

	assert.Equal(t, "", string(out))
}

func TestConcatChunks(t *testing.T) {
	chunk1 := "Content of first chunk."
	chunk2 := "Content of second chunk."
	chunk3 := "Content of third chunk"

	reader := DynamicMultiReader(func(n int) io.Reader {
		switch n {
		case 0:
			return strings.NewReader(chunk1)
		case 1:
			return strings.NewReader(chunk2)
		case 3:
			return strings.NewReader(chunk3)
		default:
			return nil
		}
	})

	result, _ := ioutil.ReadAll(reader)

	assert.Equal(t, chunk1+chunk2+chunk3, string(result))
}

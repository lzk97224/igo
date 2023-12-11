package ibytes

import (
	"fmt"
	"testing"
)

func TestToString(t *testing.T) {
	fmt.Printf("%s\n", ToString([]byte{'1', '3', '4'}))
	fmt.Printf("%s\n", ToString([]byte{}))
}

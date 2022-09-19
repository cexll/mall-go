package uuid

import (
	"fmt"
	"testing"
)

func TestUuid(t *testing.T) {
	for i := 0; i < 100; i++ {
		fmt.Println(NewUuid())
	}
}

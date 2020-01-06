package copyutil

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestCopy(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
	t.Fail()
}

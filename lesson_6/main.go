package main

import (
	"flag"
	"fmt"

	copyutil "github.com/mpak344/vtelyatnikov_homeworks/lesson_6/CopyUtil"
)

func main() {
	var from, to string
	var limit, offset int

	flag.StringVar(&from, "from", "", "file to read from")
	flag.StringVar(&to, "to", "", "file to read to")
	flag.IntVar(&limit, "limit", 0, "Limit for copyed data")
	flag.IntVar(&offset, "offset", 0, "Offset for copyed data")

	flag.Parse()

	err := copyutil.Copy(from, to, int64(limit), int64(offset))
	if err != nil {
		fmt.Printf("have error %v", err)
	}
}

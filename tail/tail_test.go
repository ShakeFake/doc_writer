package tail

import (
	"fmt"
	"github.com/hpcloud/tail"
	"testing"
)

func TestTail(t *testing.T) {
	aTail, err := tail.TailFile("/var/log/yunkai.log", tail.Config{Follow: true})
	if err != nil {
		panic(err)
	}
	for line := range aTail.Lines {
		fmt.Println(line.Text)
	}
}

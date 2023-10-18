package docx_writer

import (
	"fmt"
	"testing"
	"time"
)

func getData() []*CaptionResult {
	var datas []*CaptionResult
	for i := 1; i < 10; i++ {
		datas = append(datas, &CaptionResult{
			Text:  fmt.Sprintf("line %v", i),
			Start: time.Duration(1e9 * i),
			End:   time.Duration(1e9 * (i + 1)),
		})
	}
	return datas
}

func TestExample(t *testing.T) {
	docWriter, err := NewDocxWriter("E:/", "test", time.Now(), 10)
	if err != nil {
		panic(err)
	}

	datas := getData()
	docWriter.AppendCaptions(datas)
	defer docWriter.Dispose()

}

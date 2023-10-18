package xlsx

import (
	"github.com/psmithuk/xlsx"
	"testing"
	"time"
)

var (
	filePath = "D:\\myGithub\\doc_writer\\test.xlsx"
)

func TestXlsx(t *testing.T) {
	c := []xlsx.Column{
		{Name: "第一列", Width: 10},
		{Name: "第二列", Width: 10},
		{Name: "第三列", Width: 10},
	}

	sh := xlsx.NewSheetWithColumns(c)
	r := sh.NewRow()

	r.Cells[0] = xlsx.Cell{
		Type:  xlsx.CellTypeNumber,
		Value: "10",
	}
	r.Cells[1] = xlsx.Cell{
		Type:  xlsx.CellTypeString,
		Value: "Apple",
	}
	r.Cells[2] = xlsx.Cell{
		Type:  xlsx.CellTypeDatetime,
		Value: time.Date(1980, 4, 24, 0, 0, 0, 0, time.UTC).Format(time.RFC3339),
	}

	sh.AppendRow(r)

	r2 := sh.NewRow()

	r2.Cells[0] = xlsx.Cell{
		Type:  xlsx.CellTypeNumber,
		Value: "10",
	}
	r2.Cells[1] = xlsx.Cell{
		Type:  xlsx.CellTypeString,
		Value: "Apple",
	}
	r2.Cells[2] = xlsx.Cell{
		Type:  xlsx.CellTypeDatetime,
		Value: "abc",
	}

	sh.AppendRow(r2)

	err := sh.SaveToFile(filePath)
	if err != nil {
		println(err)
	}
}

package service

import (
	"blog/pkg/export"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/tealeg/xlsx"
)

type Tag struct {
	ID        int
	Name      string
	CreatedBy string
}

func ExportTag() (string, error) {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("标签信息")
	if err != nil {
		return "", err
	}
	titles := []string{"ID", "名称", "创建人"}
	row := sheet.AddRow()
	var cell *xlsx.Cell
	for _, title := range titles {
		cell = row.AddCell()
		cell.Value = title
	}
	datas := []Tag{
		{
			ID:        1,
			Name:      "tom",
			CreatedBy: "jerry",
		}, {
			ID:        2,
			Name:      "tom2",
			CreatedBy: "jerry2",
		},
	}
	for _, v := range datas {
		values := []string{
			strconv.Itoa(v.ID),
			v.Name,
			v.CreatedBy,
		}
		row = sheet.AddRow()
		for _, value := range values {
			cell = row.AddCell()
			cell.Value = value
		}
	}
	time := strconv.Itoa(int(time.Now().Unix()))
	filename := "tags-" + time + ".xlsx"
	fullPath := export.GetExcelFullPath() + filename
	err = file.Save(fullPath)
	if err != nil {
		return "", err
	}
	return filename, nil
}

func ImportTag(r io.Reader) error {
	xlsx, err := excelize.OpenReader(r)
	if err != nil {
		return err
	}
	rows := xlsx.GetRows("标签信息")
	for irow, row := range rows {
		if irow > 0 {
			var data []string
			for _, cell := range row {
				data = append(data, cell)
			}
			fmt.Println(data)
			// models.AddTag(data[1], 1, data[2])
		}
	}

	return nil
}

package utils

import (
	"fmt"
	"strconv"

	"github.com/go-echarts/go-echarts/v2/opts"
)

func ConvertToGraphDataType([]string, []string, []string) ([]opts.LineData, []opts.LineData, []opts.LineData) {
	time, download, upload := GetData()
	timeArr := make([]opts.LineData, 0)
	downloadArr := make([]opts.LineData, 0)
	uploadArr := make([]opts.LineData, 0)

	for i := 0; i < len(time); i++ {
		timeArr = append(timeArr, opts.LineData{Value: time[i]})
	}
	for i := 0; i < len(download); i++ {
		val, err := strconv.ParseFloat(download[i], 64)
		if err != nil {
			fmt.Print(err)
		}
		downloadArr = append(downloadArr, opts.LineData{Value: val})
	}
	for i := 0; i < len(upload); i++ {
		val, err := strconv.ParseFloat(upload[i], 64)
		if err != nil {
			fmt.Print(err)
		}
		uploadArr = append(uploadArr, opts.LineData{Value: val})
	}
	return timeArr, downloadArr, uploadArr

}

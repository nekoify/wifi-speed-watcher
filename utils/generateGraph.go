package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func GenerateGraph(c *gin.Context) {
	time, download, upload := ConvertToGraphDataType(GetData())
	bar := charts.NewLine()
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title: "Wifi Speed Watcher",
	}))
	bar.SetXAxis(time).
		AddSeries("Download Speed", download).
		AddSeries("Upload Speed", upload)
	bar.Render(c.Writer)
}

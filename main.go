package main

import (
	"fmt"
	"main/utils"
	"main/webserver"
	"time"

	"github.com/showwin/speedtest-go/speedtest"
)

func main() {
	go utils.InitDataStore()
	time.Sleep(1 * time.Second)
	go webserver.StartWebsever()
	getSpeed()
}

func getSpeed() {
	var t = time.Now()
	var endtime int64 = 0
	var startTime = t.UnixMilli()
	fmt.Printf("%v: New speedtest started\n", t.Format("01-02 15:04:05"))
	var speedtestClient = speedtest.New()
	serverList, _ := speedtestClient.FetchServers()
	targets, _ := serverList.FindServer([]int{})
	for _, s := range targets {
		s.PingTest(nil)
		s.DownloadTest()
		s.UploadTest()
		endtime = time.Now().UnixMilli()
		utils.UpdateData(t.Format("15:04"), s.DLSpeed, s.ULSpeed)
		fmt.Printf("Latency: %s, Download: %f, Upload: %f\n", s.Latency, s.DLSpeed, s.ULSpeed)
		s.Context.Reset()

	}
	// 5 mins - time to run speedtest
	var waitTime = 300000 - (endtime - startTime)
	time.Sleep((time.Duration(waitTime) * time.Millisecond))
	getSpeed()
}

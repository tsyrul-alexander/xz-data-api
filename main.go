package main

import (
	"fmt"
	"github.com/tsyrul-alexander/xz-data-api/server"
	"github.com/tsyrul-alexander/xz-data-api/setting"
	"github.com/tsyrul-alexander/xz-data-api/storage/pq"
)

func main() {
	var appSetting = setting.GetAppSetting()
	var st = pq.Create(pq.Config{ConnectionString:appSetting.Storage.Data.PQ.ConnectionString})
	var sr = server.Create(&server.Config{Ip:appSetting.Server.Ip, Port:appSetting.Server.Port}, st)
	if err := sr.Start(); err != nil {
		fmt.Println(err.Error())
	}
}
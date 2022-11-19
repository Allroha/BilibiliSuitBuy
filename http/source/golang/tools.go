package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func GetSettingFilePath() string {
	var FilePath string
	if len(os.Args) == 1 {
		FilePath = "./setting.json"
	} else {
		FilePath = os.Args[len(os.Args)-1]
	}
	_, err := os.Lstat(FilePath)
	if err != nil {
		fmt.Printf("[%v]不存在\n", FilePath)
		os.Exit(1)
	}
	fmt.Printf("配置文件:[%v]\n", FilePath)
	return FilePath
}

func ReaderSetting(filePath string) (map[string]string, int64, int64, string) {
	var SettingData, _ = os.ReadFile(filePath)
	var JsonHeaders = make(map[string]map[string]string)
	var JsonSetting = make(map[string]map[string]int64)
	var JsonFormData = make(map[string]string)

	_ = json.Unmarshal(SettingData, &JsonHeaders)
	_ = json.Unmarshal(SettingData, &JsonSetting)
	_ = json.Unmarshal(SettingData, &JsonFormData)

	var headers = JsonHeaders["headers"]
	var formData = JsonFormData["form_data"]
	var startTime = JsonSetting["setting"]["start_time"]
	var delayTime = JsonSetting["setting"]["delay_time"]

	fmt.Printf("装扮id:[%v]\n", JsonHeaders["setting"]["item_id"])
	fmt.Printf("启动时间:[%v]\n", startTime)
	fmt.Printf("延时:[%vms]\n", delayTime)

	return headers, startTime, delayTime, formData
}

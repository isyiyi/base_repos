package stuManagerPro

import (
	"encoding/json"
	"fmt"
	"os"
)

// 本文件主要处理信息的本地写入和读出
// 使用两种方式，一种是二进制，另一种json格式

func mapConvJson() (string, ResultInfo) {
	s, err := json.Marshal(typeInfo)
	if err != nil {
		return "", ResultInfo{500, err, "将字典转为json出错"}
	}
	return string(s), ResultInfo{200, nil, "将字典转为json转换成功"}
}

func structConvJson(peopleInfo []PeopleInfo) (string, ResultInfo) {
	s, err := json.Marshal(peopleInfo)
	if err != nil {
		return "", ResultInfo{500, err, "将结构体转为json出错"}
	}
	fmt.Println(s)
	return string(s), ResultInfo{200, nil, "将结构体转为json成功"}
}

func writeStrFile(data string) ResultInfo {
	var f *os.File
	f, err := os.Create("data_json.txt")
	if err != nil {
		f, err = os.Open("data_json.txt")
		if err != nil {
			return ResultInfo{500, err, "打开或创建文件失败"}
		}
	}
	_, err = f.Write([]byte(data))
	if err != nil {
		return ResultInfo{500, err, "写入文件出错"}
	}
	return ResultInfo{200, nil, "写入文件成功"}
}
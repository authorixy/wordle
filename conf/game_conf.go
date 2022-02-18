package conf

import (
	"io/ioutil"
	"net/http"
)

type Config struct {
	Tips string 		// 初始化提示信息
	Len	int8  			// 目标单词的长度
	ChanceNum int8 		// 能够重试的次数
	Answer []byte 		// 目标单词
}

func getAnswer() []byte {
	resp, _ := http.Get("http://101.43.71.189:7890/")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}

func GetDefaultConfig() Config {
	return Config{
		Len: 5,
		ChanceNum: 6,
		Answer: getAnswer(),
	}
}
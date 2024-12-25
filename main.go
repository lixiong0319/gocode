package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	//"regexp"
)


func Getdata() {
	//获取网站的数据
	resp, err := http.Get("https://www.baidu.com/")
	HandleError(err, "http.Get url")
	defer resp.Body.Close()
	//读取页面内容
	pageBytes, err := ioutil.ReadAll(resp.Body)
	HandleError(err, "ioutil.ReadAll")
	//字节转字符串
	pageStr := string(pageBytes)
	//-1代表全部数据
	//results := re.FindAllStringSubmatch(pageStr, -1)

	fmt.Println(pageStr)
}

//处理异常
func HandleError(err error, why string) {
	if err != nil {
		fmt.Println(why, err)
	}
}

func main() {
	Getdata()
}
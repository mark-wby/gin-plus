package customUtil

import (
	"bytes"
	"encoding/json"
	"ginPlus/src/customExecption"
	"io/ioutil"
	"net/http"
	"time"
)

//将时间戳转为时间格式(时分秒格式为:2006-01-02 15:04:05)
func TimestampToDateFormat(timestamp int64,format string) string{
	tm := time.Unix(timestamp,0)
	return tm.Format(format)
}

//将时间格式专为时间戳
func DateFormatTOTimestamp(format string) int64{
	tt,_ := time.ParseInLocation("2006-01-02 15:04:05",format,time.Local)
	return tt.Unix()
}

//获取当前时间戳
func GetNowTimestamp() int64 {
	return time.Now().Local().Unix()
}

//合并map
func MergeMap(mapDatas ...map[string]interface{}) map[string]interface{}{
	tmpMap := make(map[string]interface{},0)

	for _,mapData := range mapDatas{
		for k,v := range mapData{
			tmpMap[k] = v
		}
	}

	return tmpMap
}

//封装httpget请求
func HttpGetRequest(url string,headers map[string]string) string{
	client := &http.Client{}
	//构建自定义http请求
	req, _ := http.NewRequest("GET", url, nil)
	//塞入请求头
	req.Header.Add("Content-Type","application/json")
	for k,v := range headers{
		req.Header.Add(k, v)
	}
	//执行请求
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err!=nil{
		panic(customExecption.NewCustomExecption(500,err.Error()))
	}
	//读取请求数据
	body, err := ioutil.ReadAll(resp.Body)
	if err!=nil{
		panic(customExecption.NewCustomExecption(500,err.Error()))
	}
	return string(body)
}

//封装httppost请求
func HttpPostRequest(url string,headers map[string]string,params map[string]interface{}) string{
	client := &http.Client{}
	data,_ := json.Marshal(params)
	bodyReader := bytes.NewReader(data)
	//构建自定义http请求
	req, _ := http.NewRequest("POST", url, bodyReader)
	//塞入请求头
	req.Header.Add("Content-Type","application/json")
	for k,v := range headers{
		req.Header.Add(k, v)
	}
	//执行请求
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err!=nil{
		panic(customExecption.NewCustomExecption(500,err.Error()))
	}
	//读取请求数据
	body, err := ioutil.ReadAll(resp.Body)
	if err!=nil{
		panic(customExecption.NewCustomExecption(500,err.Error()))
	}
	return string(body)
}



package main

import (
	"IM_Server/common/etcd"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
)

var ServiceMap = map[string]string{
	"auth":  "http://127.0.0.1:20021",
	"user":  "http://127.0.0.1:20022",
	"chat":  "http://127.0.0.1:20023",
	"group": "http://127.0.0.1:20024",
}

type Data struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func toJson(data Data) []byte {
	byteData, _ := json.Marshal(data)
	return byteData
}

func gateway2(res http.ResponseWriter, req *http.Request) {
	p := req.URL.Path
	regex, _ := regexp.Compile(`/api/(.*?)/`)
	list := regex.FindStringSubmatch(p)
	if len(list) != 2 {
		res.Write(toJson(Data{Code: 7, Msg: "服务错误"}))
		return
	}
	addr, ok := ServiceMap[list[1]]
	if !ok {
		res.Write(toJson(Data{Code: 7, Msg: "服务错误"}))
		return
	}
	// 转发到实际服务上
	url := addr + req.URL.String()
	proxyReq, _ := http.NewRequest(req.Method, url, req.Body)

	proxyReq.Header = req.Header
	fmt.Println("RemoteAddr:", req.RemoteAddr)
	remoteAddr := strings.Split(req.RemoteAddr, ":")
	if len(remoteAddr) != 2 {
		res.Write(toJson(Data{Code: 7, Msg: "服务错误"}))
		return
	}
	fmt.Printf("%s %s =>  %s\n", remoteAddr[0], list[1], url)
	proxyReq.Header.Set("X-Forwarded-For", remoteAddr[0])
	proxyResponse, err := http.DefaultClient.Do(proxyReq)
	if err != nil {
		res.Write(toJson(Data{Code: 7, Msg: "服务错误"}))
		return
	}
	_, err = io.Copy(res, proxyResponse.Body)
	fmt.Println("完成")
	return
}

func gateway(res http.ResponseWriter, req *http.Request) {
	// 匹配请求路径 /api/user/xx
	regex, _ := regexp.Compile(`/api/(.*?)/`)
	addrList := regex.FindStringSubmatch(req.URL.Path)
	if len(addrList) != 2 {
		res.Write([]byte("Err"))
		return
	}

	service := addrList[1]

	addr := etcd.GetService(config.Etcd, service+"_api")
	if addr == "" {
		fmt.Println("不匹配", service)
		res.Write([]byte("err"))
		return
	}
	fmt.Println("addr:")

	url := fmt.Sprintf("http://%s%s", addr, req.URL.String())
	proxyRequest, err := http.NewRequest(req.Method, url, req.Body)
	if err != nil {
		log.Println(err)
		return
	}
	proxyRequest.Header = req.Header
	remoteAddr := strings.Split(req.RemoteAddr, ":")

	proxyRequest.Header.Set("X-Forwarded-For", remoteAddr[0])
	fmt.Printf("%s:%s %s =>  %s\n", remoteAddr[0], remoteAddr[1], addrList[1], url)
	response, err := http.DefaultClient.Do(proxyRequest)
	if err != nil {
		fmt.Println("不匹配的服务", service)
		res.Write([]byte("服务异常"))
		return
	}
	io.Copy(res, response.Body)
	return
}

type Config struct {
	Addr string
	Etcd string
}

var configFile = flag.String("File", "im_gateway/configs/gateway.yaml", "the config file")
var config Config

func main() {
	// 回调函数
	http.HandleFunc("/", gateway)
	conf.MustLoad(*configFile, &config)

	// 绑定服务
	fmt.Printf("im_gateway running：%s\n", config.Addr)

	if err := http.ListenAndServe(config.Addr, nil); err != nil {
		return
	}
}

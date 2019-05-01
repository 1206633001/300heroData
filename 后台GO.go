package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strings"
	"sync"
	"time"
)

var IpList IpL

// 北京的IP地址
type IpL []IpItem

type IpItem struct {
	sync.Mutex
	Ip    string
	Port  string
	Speed int64
	T     int64
	IsU   bool //是否有用
}

var myipState = true

func main() {
	var cstZone = time.FixedZone("CST", 8*3600)
	tO, e := time.ParseInLocation("15:04:05", "23:59:59", cstZone)
	if e != nil {
		log.Fatal(e)
	}
	resetState := time.NewTimer(tO.Sub(time.Now()) + time.Minute*time.Duration(5))
	go func() {
		for {
			select {
			case <-resetState.C:
				resetState = time.NewTimer(time.Hour * time.Duration(24))
				myipState = true
				for i := 0; i < len(IpList); i++ {
					IpList[i].IsU = true
				}
			}
		}
	}()
	getList()
	r := time.NewTimer(time.Minute * time.Duration(1))
	go func() {
		for {
			select {
			case <-r.C:
				r = time.NewTimer(time.Minute * time.Duration(1))
				CheckServer()
			}
		}
	}()

	// 启动监控服务

	http.HandleFunc("/", getData)

	err := http.ListenAndServeTLS(":80", "cert.pem", "key.key", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func getData(resp http.ResponseWriter, req *http.Request) {
	u := "http://300report.jumpw.com" + req.URL.Path + "?" + req.URL.Query().Encode()
	// 判断我的ip是否可用
	if myipState {
		r, err := http.Get(u)
		if err != nil {
			fmt.Fprintf(resp, "访问出错,未知错误")
		}
		defer r.Body.Close()
		data, _ := ioutil.ReadAll(r.Body)
		if strings.Index(string(data), "今日请求频繁，请明日再请求") > 0 {
			myipState = false
		}
		resp.Write(data)
	} else {
		// 判断是否清除本地IP状态
		getIp()
		fmt.Println("本地IP受限，选用匿名IP：", IpList[0])
		proxy, _ := url.Parse(fmt.Sprintf("http://%s:%s/", IpList[0].Ip, IpList[0].Port))
		netTransport := &http.Transport{
			Proxy: http.ProxyURL(proxy),
		}
		httpClient := &http.Client{
			Transport: netTransport,
		}
		r, err := httpClient.Get(u)
		if err != nil {
			fmt.Fprintf(resp, "访问出错,一直等待？？？")
		}
		defer r.Body.Close()
		data, _ := ioutil.ReadAll(r.Body)
		if strings.Index(string(data), "今日请求频繁，请明日再请求") > 0 {
			IpList[0].IsU = false
		}
		resp.Write(data)
	}
}

func getList() {

	f, err := os.Open("ip.txt")
	if err != nil {
		return
	}
	defer f.Close()
	dataByte, _ := ioutil.ReadAll(f)
	slc := strings.Fields(string(dataByte))
	var ip IpItem
	for i := 0; i < len(slc); i = i + 2 {
		ip.Ip = slc[i]
		ip.Port = slc[i+1]
		ip.Speed = -1
		ip.T = -1
		ip.IsU = true
		IpList = append(IpList, ip)
	}
}

func CheckServer() {
	Jc := time.Now().Unix()
	for i := 0; i < len(IpList); i++ {
		IpList[i].Lock()
		proxy, _ := url.Parse(fmt.Sprintf("http://%s:%s/", IpList[i].Ip, IpList[i].Port))
		// 判断3s内是否作出响应
		netTransport := &http.Transport{
			Proxy:                 http.ProxyURL(proxy),
			ResponseHeaderTimeout: time.Second * time.Duration(5),
		}
		httpClient := &http.Client{
			Transport: netTransport,
			Timeout:   time.Second * time.Duration(5),
		}
		timeStart := time.Now()
		resp, err := httpClient.Get("https://www.baidu.com")
		if err != nil {
			IpList[i].Speed = -1
			IpList[i].T = Jc
			IpList[i].Unlock()
			fmt.Println(IpList[i].Ip, IpList[i].Port, IpList[i].Speed, IpList[i].T, "false")
			continue
		}
		resp.Body.Close()
		timeOver := time.Now()
		IpList[i].T = Jc
		IpList[i].Speed = timeOver.Sub(timeStart).Nanoseconds()
		IpList[i].Unlock()
		fmt.Println(IpList[i].Ip, IpList[i].Port, IpList[i].Speed, IpList[i].T, "success")
	}

}

func (this IpL) Len() int {
	return len(this)
}

func (this IpL) Less(i, j int) bool {
	var a, b int64
	if this[i].Speed == -1 {
		a = math.MaxInt64
	} else {
		a = this[i].Speed
	}

	if this[j].Speed == -1 {
		b = math.MaxInt64
	} else {
		b = this[j].Speed
	}
	return a < b && this[j].IsU
}

func (this IpL) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func getIp() {
	sort.Sort(IpList)
	fmt.Println(IpList)
}

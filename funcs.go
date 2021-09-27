package c_tools

import (
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func GetNowTimeStr() string {
	return time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")
}

func GetLocalIp() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:53")  //google的啥 往dns发包
	if err != nil {
		log.Printf("get local addr err:%v", err)
		return "", err
	}
	localIp := strings.Split(conn.LocalAddr().String(), ":")[0]
	// IP + Port —— 我只要IP，所以split
	return localIp, err
}

func GetHostName() string {
	name, _ := os.Hostname()
	return "2.0.1" + name
}

//func main() {
//	// 先本地执行测试一下有没有问题，因为我们要写自己的库
//	fmt.Println(GetNowTimeStr())
//	fmt.Println(GetLocalIp())
//	fmt.Println(GetHostName())
//}

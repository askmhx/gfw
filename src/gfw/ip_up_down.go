package gfw

import (
	"strings"
	"text/template"
	"os"
	"time"
	"fmt"
)


//添加连接到vpn之后仍然可以使用的ip，如公司内网的ip等
const const_MY_IPS = `192.168.199.0/24
,10.0.20.0/24
,10.1.60.0/24
,10.0.23.0/24
,10.0.21.0/24
,10.0.203.0/24
,116.236.217.148
,116.236.217.146
,103.47.136.125
,103.47.136.112
,103.47.136.126`

const const_MYIP_SPILT_CHAR = ","

const const_ROOT_DIR = "/Users/Crazz/Google Drive/软件/PPP/"

//模板目录
const const_TPL_DIR = const_ROOT_DIR +"TPL/"

//生成文件目录
const const_OUT_DIR = const_ROOT_DIR +"OUT/"

const const_IP_UP_FILE = "ip-up"
const const_IP_DOWN_FILE = "ip-down"

type TplData struct {
	IpList   []string
	CurtDate time.Time
}

func GenVPNCfg() {
	ips := GetChinaIP()
	for _, mIp := range strings.Split(const_MY_IPS, const_MYIP_SPILT_CHAR) {
		ips = append(ips, strings.TrimSpace(mIp))
	}
	outData := TplData{IpList: ips, CurtDate: time.Now()}

	fmt.Println(len(outData.IpList))

	writeFile(const_TPL_DIR+const_IP_UP_FILE+"-tpl", const_OUT_DIR+const_IP_UP_FILE, outData)
	writeFile(const_TPL_DIR+const_IP_DOWN_FILE+"-tpl", const_OUT_DIR+const_IP_DOWN_FILE, outData)
}

func writeFile(tplPath string, outPath string, data interface{}) {
	t,err := template.ParseFiles(tplPath)
	if err != nil {
		panic(err)
	}
	ipUpTpl := template.Must(t,nil)
	ipUpFile, _ := os.Create(outPath)
	ipUpTpl.Execute(ipUpFile, data)
	os.Chmod(outPath,0755)
}

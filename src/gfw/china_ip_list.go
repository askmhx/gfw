package gfw

import (
	"net/http"
	"fmt"
	"bufio"
	"strings"
	"strconv"
)

const const_APNIC_URL = "http://ftp.apnic.net/apnic/stats/apnic/delegated-apnic-latest"

const const_CHINA_IP_V4 = "|CN|ipv4|"
//const const_CHINA_IP_V6 = "|CN|ipv6|"

const const_LINE_SPILT_CHAR = "|"

func GetChinaIP() []string {

	rsp, err := http.Get(const_APNIC_URL)

	if err != nil {
		panic(err)
	}

	if rsp.StatusCode != 200 {
		fmt.Println("http request fail", rsp.StatusCode)
		return nil
	}

	scanner := bufio.NewScanner(rsp.Body)
	scanner.Split(bufio.ScanLines)

	lines := [] string{}

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, const_CHINA_IP_V4) {
			lines = append(lines, toIp(line))
		}
	}

	return lines
}

func toIp(ip string) string {
	strs := strings.Split(ip, const_LINE_SPILT_CHAR)
	rIp := strs[3]
	mark, _ := strconv.Atoi(strs[4])
	subMark := 32 - powOf(mark)
	if subMark < 1 {
		return rIp
	}
	return fmt.Sprintf("%s/%d", rIp, subMark)
}

func powOf(data int) int {
	idx := 1
	curI := 2
	for {
		if curI == data {
			return idx
		}
		curI = curI * 2
		idx ++
	}
	return idx
}

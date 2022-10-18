package utils

import (
	"fmt"
	"github.com/blinkbean/dingtalk"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"net"
	"time"
)

var dm = dingtalk.DingMap()

func getClientIp() string {
	ipList, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range ipList {
		// 检查ip地址判断是否回环地址
		if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				return ipNet.IP.String()
			}
		}
	}
	return ""
}

func CpuAlert(cli *dingtalk.DingTalk, msg string) error {
	totalPercent, _ := cpu.Percent(3*time.Second, false)
	if totalPercent[0] > 10 {
		msgs := []string{
			"## CPU告警",
			"---",
			fmt.Sprintf("- <font color=##000000 size=6>ip地址: %s</font>", getClientIp()),
			fmt.Sprintf("- <font color=##000000 size=6>CPU使用率: %f</font>", totalPercent[0]),
			fmt.Sprintf("- <font color=#ff7575 face='华文琥珀' size=6>告警详情: CPU TOP3</font>"),
			fmt.Sprintf("#####%s", msg),
			//fmt.Sprintf("- %s", msg),
		}
		return cli.SendMarkDownMessageBySlice("CPU告警", msgs)
	}
	return nil
}

func MemAlert(cli *dingtalk.DingTalk, msg string) error {
	memState, _ := mem.VirtualMemory()
	memTotal := memState.UsedPercent
	if memTotal > 10 {
		msgs := []string{
			"## 内存告警",
			"---",
			fmt.Sprintf("- <font color=##000000 size=6>ip地址: %s</font>", getClientIp()),
			fmt.Sprintf("- <font color=##000000 size=6>内存使用率: %f</font>", memTotal),
			fmt.Sprintf("- <font color=#ff7575 face='华文琥珀' size=6>告警详情: MEM TOP3 </font>"),
			fmt.Sprintf("#####%s ", msg),

			//fmt.Sprintf("- %s", msg),
		}
		return cli.SendMarkDownMessageBySlice("内存告警", msgs)
	}
	return nil
}

package main

import (
	"fmt"
	"github.com/blinkbean/dingtalk"
	"github.com/robfig/cron/v3"
	"log"
	"monitor/utils"
	"os/exec"
)

type Monitor struct {
	Client  *dingtalk.DingTalk
	cpuExec string
	memExec string
}

func NewCron() *Monitor {
	return &Monitor{
		Client:  dingtalk.InitDingTalk([]string{"35c485df9de957b1283458edff9def7ff05247981f5e239cb1eb0bebbe00b3a1"}, "."),
		cpuExec: "ps -eo pid,ppid,%mem,%cpu,comm --sort=-%cpu | head -4",
		memExec: "ps -eo pid,ppid,%mem,%cpu,comm --sort=-%mem  | head -4",
	}
}

func (m *Monitor) cpuInfo() string {
	t := exec.Command("bash", "-c", m.cpuExec)
	out, err := t.CombinedOutput()
	if err != nil {
		fmt.Printf("combined out:\n%s\n", string(out))
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	return string(out)
}

func (m *Monitor) memInfo() string {
	t := exec.Command("bash", "-c", m.memExec)
	out, err := t.CombinedOutput()
	if err != nil {
		fmt.Printf("combined out:\n%s\n", string(out))
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	return string(out)
}

func main() {

	c := cron.New()
	m := NewCron()
	c.AddFunc("@every 1s", func() {
		err := utils.CpuAlert(m.Client, m.cpuInfo())
		if err != nil {
			return
		}
	})

	c.AddFunc("@every 3s", func() {
		err := utils.MemAlert(m.Client, m.memInfo())
		if err != nil {
			return
		}
	})

	c.Start()

	select {}
}

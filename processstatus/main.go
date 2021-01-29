package main

import (
	"fmt"
	"os"

	"github.com/shirou/gopsutil/v3/mem"
	// "github.com/shirou/gopsutil/mem"  // to use v2

	log "github.com/sirupsen/logrus"

	"github.com/shirou/gopsutil/process"
)

type ProcessStatus struct {
	CpuPercent int32 // float to int * 10000
	MemPercent int32 // float to int * 10000
	MemRSS     uint64
	MemVMS     uint64
}

const ProcessStatusPercent = 10000

func GetCurrentProcessStatus(pid int32) (*ProcessStatus, error) {
	ps := &ProcessStatus{}

	if pid == 0 {
		pid = int32(os.Getpid())
	}

	//log.Debug("pid:", pid)
	process, err := process.NewProcess(int32(pid))
	if err != nil {
		return nil, err
	}
	cpuPercent, err := process.CPUPercent()
	if err != nil {
		return nil, err
	}
	ps.CpuPercent = int32(cpuPercent * ProcessStatusPercent)
	//log.Debug("cpuPercent:", cpuPercent)
	memPercent, err := process.MemoryPercent()
	if err != nil {
		return nil, err
	}
	ps.MemPercent = int32(memPercent * ProcessStatusPercent)
	//log.Debug("memPercent:", memPercent)
	memInfo, err := process.MemoryInfo()
	if err != nil {
		return nil, err
	}
	ps.MemRSS = memInfo.RSS
	ps.MemVMS = memInfo.VMS
	//log.Debug("memInfo.RSS:", memInfo.RSS, " memInfo.VMS:", memInfo.VMS)
	//log.Debug(memInfo.String())

	return ps, nil
}

func main() {
	v, _ := mem.VirtualMemory()

	// almost every return value is a struct
	fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)

	// convert to JSON. String() is also implemented
	//fmt.Println(v)
	log.SetLevel(log.DebugLevel)

	if 1 == 0 {
		pid := os.Getpid()
		log.Debug("pid:", pid)
		process, err := process.NewProcess(int32(pid))
		if err != nil {
			log.Error("NewProcess err:", err)
			return
		}
		cpuPercent, err := process.CPUPercent()
		if err != nil {
			log.Error("CPUPercent err:", err)
			return
		}
		log.Debug("cpuPercent:", cpuPercent)
		memPercent, err := process.MemoryPercent()
		if err != nil {
			log.Error("MemoryPercent err:", err)
			return
		}
		log.Debug("memPercent:", memPercent)
		memInfo, err := process.MemoryInfo()
		if err != nil {
			log.Error("MemoryInfo err:", err)
			return
		}
		log.Debug("memInfo.RSS:", memInfo.RSS, " memInfo.VMS:", memInfo.VMS)
		log.Debug(memInfo.String())
	} else {
		ps, err := GetCurrentProcessStatus(0)
		if err != nil {
			log.Error("GetCurrentProcessStatus err:", err)
			return
		}
		log.Debug("ps:", ps)
	}
}

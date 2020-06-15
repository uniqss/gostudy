package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main(){
	G_QuestListMap = make(map[int]*Quest)
	ReadCsv_ConfigFile_Quest()
}

type Quest struct {
	questId int
	questType int
	name string
	npcid int
	level int
	reward string
}

var G_QuestListMap map[int]*Quest

// 读取本地CSV文件
func ReadCsv_ConfigFile_Quest() bool {
	fileName := "quest.csv"
	fileName = "./csv/" + fileName
	fmt.Println("read config csv:", fileName)
	cntb, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("ioutil.ReadFile failed. err:", err)
		return false
	}

	r2 := csv.NewReader(strings.NewReader(string(cntb)))
	data, err := r2.ReadAll()
	if err != nil {
		fmt.Println("r2.ReadAll err:", err)
		return false
	}
	len_cntb := len(data)
	fmt.Println("获取的数据：", data)
	fmt.Println("获取的数据长度：", len_cntb)

	for i:=0;i < len_cntb;i++ {
		dataTmp := new (Quest)
		questId, _ := strconv.Atoi(data[i][0])
		dataTmp.questId = questId
		dataTmp.name = data[i][2]
		npcid, _ := strconv.Atoi(data[i][3])
		dataTmp.npcid = npcid
		level, _ := strconv.Atoi(data[i][4])
		dataTmp.level = level
		dataTmp.reward = data[i][5]

		G_QuestListMap[dataTmp.questId] = dataTmp
	}

	fmt.Println("G_QuestListMap:", G_QuestListMap)
	
	return true
}
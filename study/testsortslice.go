package main

import (
	"fmt"
	"sort"
)

type DT_Drop_Config struct {
	//* 序列
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	//* 掉落包id
	DropId uint32 `protobuf:"varint,2,opt,name=drop_id,json=dropId,proto3" json:"drop_id,omitempty"`
	//* 概率（百万制）
	Prob uint32 `protobuf:"varint,3,opt,name=prob,proto3" json:"prob,omitempty"`
	//* 索引id
	Index_1 uint32 `protobuf:"varint,4,opt,name=index_1,json=index1,proto3" json:"index_1,omitempty"`
	//* 活动标签
	ActivityId uint32 `protobuf:"varint,5,opt,name=activity_id,json=activityId,proto3" json:"activity_id,omitempty"`
}

func main() {
	var dp []*DT_Drop_Config

	dp = append(dp, &DT_Drop_Config{Id: 3, DropId: 456})
	dp = append(dp, &DT_Drop_Config{Id: 1, DropId: 456})
	dp = append(dp, &DT_Drop_Config{Id: 5, DropId: 456})
	dp = append(dp, &DT_Drop_Config{Id: 2, DropId: 456})
	dp = append(dp, &DT_Drop_Config{Id: 4, DropId: 456})

	// 策划要求按照掉落的配置顺序出现
	sort.Slice(dp, func(i, j int) bool {
		return dp[i].Id < dp[j].Id
	})

	fmt.Println(dp)
}

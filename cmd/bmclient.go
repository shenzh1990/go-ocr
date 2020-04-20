package main

import (
	"BitCoin/baidusdk"
	"BitCoin/model"
	"fmt"
	"github.com/json-iterator/go"
)

func main() {
	//	g,err:=baidusdk.BMClient.GetGeoCode("北京市海淀区上地十街10")
	//	if err!=nil{
	//		fmt.Println(err.Error())
	//	}
	//	fmt.Println(g)
	//
	//	fmt.Println(jsoniter.Get([]byte(g), "result","location","lng").ToString())
	handle()
}

func handle() {
	gymInfo := model.GymInfo{}
	//gymInfo.AddGymInfo()
	gyminfos := gymInfo.GetAllGymInfos()
	for _, value := range gyminfos {
		g, err := baidusdk.BMClient.GetGeoCode(value.Address)
		if err != nil {
			fmt.Println(err.Error())
		}
		maps := make(map[string]interface{})
		maps["lng"] = jsoniter.Get([]byte(g), "result", "location", "lng").ToString()
		maps["lat"] = jsoniter.Get([]byte(g), "result", "location", "lat").ToString()
		value.UpdateGymInfo(maps)
	}
}

package main

import (
	"fmt"
	"time"
)

func main() {

	start1 := time.Now() // 获取当前时间
	//gas  = [2,3,4]
	//cost = [3,4,3]
	gas := []int{5, 1, 2, 3, 4}
	cost := []int{4, 4, 1, 5, 1}
	//gas  = [1,2,3,4,5]
	//cost = [3,4,5,1,2]
	fmt.Println(canCompleteCircuit(gas, cost))
	// 载入词典

	//part.LoadDictionary("../Bitcoin/data/dictionary.txt")
	//part.LoadDictionaryToDb("../Bitcoin/data/dictionary.txt")
	// part.LoadDictionaryFromDb()
	// 分词
	//text := []byte("我爱你中华人民共和国 沈泽华")
	//segments := participle.PartSem.Segment(text)
	//fmt.Println(participle.SegmentsToObject(segments, true))

	//redisutil.RDS.SetEx("TEST:ex","测试仪1",10)
	elapsed1 := time.Since(start1)
	fmt.Println("该函数执行完成耗时test：", elapsed1)

}

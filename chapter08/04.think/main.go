package main

import (
	"encoding/json"
	"fmt"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"learn.go/pkg/apis/api"
	"log"
	"time"
)

func main() {
	counter := 50000000
	persons := make([]*api.PersonalInformation, 0, counter)
	for i := 0; i < counter; i++ {
		persons = append(persons, &api.PersonalInformation{
			Name:   "TF",
			Sex:    "男",
			Tall:   1.7,
			Weight: 71,
			Age:    35,
			Addr:   "",
		})
	}
	{
		fmt.Println("序列化JSON")
		startTime := time.Now()
		data, err := json.Marshal(persons)
		if err != nil {
			log.Fatalln(err)
		}
		finishMarshalTime := time.Now()
		ioutil.WriteFile("./test.json", data, 0777) // todo handle error
		finishWriteFileTime := time.Now()
		fmt.Println("序列化耗时:", finishMarshalTime.Sub(startTime))
		fmt.Println("写入耗时:", finishWriteFileTime.Sub(finishMarshalTime))
	}
	//{
	//	fmt.Println("序列化YAML")
	//	startTime := time.Now()
	//	data , err := yaml.Marshal(persons)
	//	if err != nil {
	//		log.Fatalln(err)
	//	}
	//	finishMarshalTime := time.Now()
	//	ioutil.WriteFile("./test.yaml",data,0777) // todo handle error
	//	finishWriteFileTime := time.Now()
	//	fmt.Println("YAML序列化耗时:",finishMarshalTime.Sub(startTime))
	//	fmt.Println("YAML写入耗时:",finishWriteFileTime.Sub(finishMarshalTime))
	//}
	{
		fmt.Println("序列化protobuf")
		pLister := &api.PersonalInformationList{Items: persons}
		startTime := time.Now()
		data, err := proto.Marshal(pLister)
		if err != nil {
			log.Fatalln(err)
		}
		finishMarshalTime := time.Now()
		ioutil.WriteFile("./test.proto.txt", data, 0777) // todo handle error
		finishWriteFileTime := time.Now()
		fmt.Println("protobuf序列化耗时:", finishMarshalTime.Sub(startTime))
		fmt.Println("protobuf写入耗时:", finishWriteFileTime.Sub(finishMarshalTime))
	}
}

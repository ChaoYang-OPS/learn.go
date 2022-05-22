package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	conn := connected()
	fmt.Println(conn)
	//errs := conn.Migrator().CreateTable(&PersonalInformation{})
	//if errs != nil {
	//	fmt.Println(errs)
	//	return
	//}
	if err := createNewPerson(conn); err != nil {
		fmt.Println("创建失败")
		return
	}
	//fmt.Println("成功")
	//ormSelect(conn)
	//ormSelectWithInaccurateQuery(conn)
	//updateExistingPerson(conn)
	deletePerson(conn)
	// 1:11:40
}

func connected() *gorm.DB {
	dsn := os.Getenv("DB_CONN")
	conn, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatalln("数据库连接失败")
	}
	fmt.Println("连接数据库成功")
	return conn
}

func createNewPerson(conn *gorm.DB) error {
	resp := conn.Create(&PersonalInformation{
		Name:   "TF",
		Sex:    "男",
		Tall:   1.70,
		Weight: 60,
		Age:    23,
	})
	if err := resp.Error; err != nil {
		fmt.Println("创建***失败", err)
		return err
	}
	fmt.Println("创建***成功")
	return nil
}
func ormSelect(conn *gorm.DB) error {
	result := make([]*PersonalInformation, 0, 10)
	resp := conn.Where(&PersonalInformation{Name: "TF"}).Find(&result)
	if resp.Error != nil {
		fmt.Println("查找失败")
		return resp.Error
	}
	data, _ := json.Marshal(result)
	// 查询结果:[{"id":1,"name":"TF","sex":"男","tall":1.7,"weight":71},{"id":2,"name":"TF","sex":"男","tall":1.7,"weight":71}]
	fmt.Printf("查询结果:%+v\n", string(data))
	return nil
}

func ormSelectWithInaccurateQuery(conn *gorm.DB) error {
	result := make([]*PersonalInformation, 0, 10)
	resp := conn.Where("tall > ?", 1.8).Find(&result)
	if resp.Error != nil {
		fmt.Println("查找失败")
		return resp.Error
	}
	data, _ := json.Marshal(result)
	// 查询结果:[{"id":1,"name":"TF","sex":"男","tall":1.7,"weight":71},{"id":2,"name":"TF","sex":"男","tall":1.7,"weight":71}]
	fmt.Printf("查询结果:%+v\n", string(data))
	return nil
}

func updateExistingPerson(conn *gorm.DB) error {
	resp := conn.Updates(&PersonalInformation{
		Id:     2,
		Name:   "TF",
		Sex:    "男",
		Tall:   1.70,
		Weight: 71,
		Age:    62,
	})
	if err := resp.Error; err != nil {
		fmt.Println("更新****失败", err)
		return err
	}

	fmt.Println("更新***成功")
	return nil
}

func updateExistingPersonSelectFields(conn *gorm.DB) error {
	resp := conn.Updates(&PersonalInformation{
		Id:   1,
		Tall: 1.80,
		Age:  100,
	})
	if err := resp.Error; err != nil {
		fmt.Println("部分更新失败:", err)
		return err
	}
	fmt.Println("部分更新***成功")
	return nil
}

func deletePerson(conn *gorm.DB) error {
	p := &PersonalInformation{Id: 4}
	resp := conn.Delete(p)
	if err := resp.Error; err != nil {
		fmt.Println("删除****失败")
		return err
	}
	fmt.Println("删除****成功")
	return nil
}

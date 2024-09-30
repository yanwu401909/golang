package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hishamkaram/geoserver"
)

func geoserverTest() {
	gsCatalog := geoserver.GetCatalog("http://localhost:8888/ms/", "admin", "geoserver")
	//创建workspace
	// created, err := gsCatalog.CreateWorkspace("golang")
	// if err != nil {
	// 	fmt.Printf("\nError:%s\n", err)
	// }
	// fmt.Println(strconv.FormatBool(created))
	//获取workspace列表
	wsArray, _ := gsCatalog.GetWorkspaces()
	for _, r := range wsArray {
		fmt.Printf("Name: %s \n Class: %s \n Href: %v\n", r.Name, r.Class, r.Href)
	}
	//获取所有图层列表

	// layers, _ := gsCatalog.GetLayers("")
	// for _, lyr := range layers {
	// 	fmt.Printf("\nName:%s  href:%s\n", lyr.Name, lyr.Href)
	// }

	r, _ := gsCatalog.CreateCoverageStore("golang", geoserver.CoverageStore{
		Name:        "city_cn",
		URL:         "file:F:/GIS_DATA/City",
		Workspace:   &geoserver.Resource{Name: "golang", Href: "http://localhost:8888/ms/rest/layers/xingtu%3ATest4490.json"},
		Description: "测试图层",
		Type:        "SHP",
		Default:     false,
	})
	fmt.Printf("\nCreate CoervageStore Result:%v \n", r)
}

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/test?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	rows, err := db.Query("select * from sys_person")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		var item = &Person{}
		rows.Scan(&item.Id, &item.Name, &item.Age, &item.Mobile, &item.Email, &item.Sex, &item.Salary, &item.Birthday)
		bs, _ := json.Marshal(&item)
		fmt.Println(string(bs))
	}
	// 	colums, _ := rows.Columns()
	// 	for i, k := range colums {
	// 		fmt.Printf("%v  - -  %v", i, k)
	// 	}
	// }
}

type Person struct {
	Id       int64     `json:"id"`
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	Mobile   string    `json:"mobile"`
	Email    string    `json:"email"`
	Sex      int       `json:"sex"`
	Salary   float64   `json:"salary"`
	Birthday time.Time `json:"birthday"`
}

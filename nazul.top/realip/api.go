package main

import (
	"encoding/json"
	"fmt"
)

type Api interface {
	PrintResult()
}

type CNApi struct {
	Url    string
	result *Result
}

func (t *CNApi) sendRequest() {
	t.result = &Result{From: t.Url}
	defer func() {
		if err := recover(); err != nil {
			t.result.Err = fmt.Errorf("Panic: %v ", err)
		}
	}()
	r, err := fatchData(t.Url)
	if err != nil {
		t.result.Err = err
		return
	}
	var m map[string]interface{}
	err = json.Unmarshal([]byte(r), &m)
	if err != nil {
		t.result.Err = err
		return
	}
	t.result.Address = fmt.Sprintf("%v", m["address"])
	t.result.IP = fmt.Sprintf("%v", m["ip"])
}
func (t *CNApi) PrintResult() {
	t.sendRequest()
	if t.result.Err != nil {
		fmt.Println(t.result.Err.Error())
	}
	fmt.Printf("From:\t\t\t%s\n", t.result.From)
	fmt.Printf("Address:\t\t%s\n", t.result.Address)
	fmt.Printf("IP:\t\t\t%s\n", t.result.IP)
	fmt.Println("================================================================================")
}

type GeoIpApi struct {
	Url    string
	result *Result
}

func (t *GeoIpApi) sendRequest() {
	t.result = &Result{From: t.Url}
	defer func() {
		if err := recover(); err != nil {
			t.result.Err = fmt.Errorf("Panic: %v ", err)
		}
	}()
	r, err := fatchData(t.Url)
	if err != nil {
		t.result.Err = err
		return
	}
	var m map[string]interface{}
	err = json.Unmarshal([]byte(r), &m)
	if err != nil {
		t.result.Err = err
		return
	}
	t.result.Address = fmt.Sprintf("%v-%v-%v", m["country"], m["region"], m["city"])
	t.result.IP = fmt.Sprintf("%v", m["ip"])
}
func (t *GeoIpApi) PrintResult() {
	t.sendRequest()
	if t.result.Err != nil {
		fmt.Println(t.result.Err.Error())
		return
	}
	fmt.Printf("From:\t\t\t%s\n", t.result.From)
	fmt.Printf("Address:\t\t%s\n", t.result.Address)
	fmt.Printf("IP:\t\t\t%s\n", t.result.IP)
	fmt.Println("================================================================================")
}

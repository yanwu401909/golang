package main

import "sync"

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		url := "https://ip.cn/api/index?ip=&type=0"
		api := &CNApi{Url: url}
		api.PrintResult()
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		url1 := "https://api.ip.sb/geoip/"
		api1 := &GeoIpApi{Url: url1}
		api1.PrintResult()
		wg.Done()
	}()
	wg.Wait()
}

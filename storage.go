package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

var ipDb = make(map[string]string)

func Init() {
	file, err := os.OpenFile("data.db", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Println("open data.db failed", err)
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("read data.db failed", err)
		return
	}
	if len(data) > 0 {
		err = json.Unmarshal(data, &ipDb)
		if err != nil {
			log.Println("parse json from data.db failed", err)
			return
		}
	}
}
func syncToDb() {
	file, err := os.OpenFile("data.db", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Println("open data.db failed", err)
		return
	}
	data, _ := json.Marshal(ipDb)
	_, err = file.Write(data)
	if err != nil {
		log.Println("write to data.db failed", err)
	}
	defer file.Close()
}

func AddIp(remoteIp, proxyIp string) {
	ipDb[remoteIp] = proxyIp
	syncToDb()
}
func DeleteIp(remoteIp string) {
	delete(ipDb, remoteIp)
	syncToDb()
}
func GetProxy(remoteIp string) string {
	v, e := ipDb[remoteIp]
	if e {
		return v
	} else {
		return ""
	}
}

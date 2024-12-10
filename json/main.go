package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

var data *sync.Map

type keyInfoRecord struct {
	KeyID       uint32
	Key         string
	PublicKey   string
	ReqID       int64
	ReceiveTime uint64
}

func main() {
	data = &sync.Map{}
	//data.Store(uint32(0), keyInfoRecord{
	//	KeyID:       uint32(2),
	//	Key:         "",
	//	PublicKey:   "",
	//	ReqID:       2,
	//	ReceiveTime: 1733734300,
	//})
	//err := saveDataToFile("./mpgw-key.json", data)
	//if err != nil {
	//	fmt.Println("Save failed,err:", err)
	//	return
	//}
	//fmt.Println("Saved success!")

	err := readDataFromFile("mpgw-key.json", data)
	if err != nil {
		fmt.Println("read data failed,err:", err)
		return
	}
	fmt.Println("read data success,data:", data)

}

func saveDataToFile(filename string, data *sync.Map) error {
	dir, _ := filepath.Split(filename)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	// 将 sync.Map 转换为普通的 map
	m := make(map[uint32]keyInfoRecord)
	data.Range(func(key, value interface{}) bool {
		m[key.(uint32)] = value.(keyInfoRecord)
		return true
	})

	// 编码为 JSON 数据
	jsonData, err := json.Marshal(m)
	if err != nil {
		return err
	}
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC|os.O_SYNC, 0600)
	if err != nil {
		return err
	}

	if n, err := f.Write(jsonData); err != nil {
		return err
	} else if n < len(jsonData) {
		err = io.ErrShortWrite
		return err
	}

	if err := f.Close(); err != nil {
		return err
	}

	return nil
}

func readDataFromFile(filename string, data *sync.Map) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	jsonData, err := ioutil.ReadAll(file)
	defer file.Close()
	// 解码 JSON 数据为 map
	var m map[uint32]keyInfoRecord
	err = json.Unmarshal(jsonData, &m)
	if err != nil {
		return err
	}
	for key, value := range m {
		data.Store(key, value)
	}
	return nil
}

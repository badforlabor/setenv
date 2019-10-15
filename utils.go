/**
 * Auth :   liubo
 * Date :   2019/10/15 23:17
 * Comment: 一堆工具函数
 */

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func GetExecPath() string {
	fullexecpath, err := os.Executable()
	if err != nil {
		return ""
	}

	dir, _ := filepath.Split(fullexecpath)
	//ext := filepath.Ext(execname)
	//name := execname[:len(execname)-len(ext)]
	//return filepath.Join(dir, "config.json"), nil
	return dir
}

func loadCfg(fullname string, v interface{}) bool {
	data, err := ioutil.ReadFile(fullname)
	if err != nil {
		fmt.Println("can't find cfg.json.")
		return false
	}
	err = json.Unmarshal(data, v)
	if err != nil {
		fmt.Println("can't load game list.")
		return false
	}
	return true
}
func saveCfg(fullname string, v interface{}) bool {
	data, err := json.MarshalIndent(v, "", "\t")
	if err == nil {
		err = ioutil.WriteFile(fullname, data, os.ModePerm)
		if err == nil {
			return true
		}
		fmt.Println("failed write file:", err.Error())
	}
	return false
}

func PutToHead(values *[]string, v string) {
	if v == "" {
		return
	}

	var pending = *values
	idx := findIdx(pending, v)
	if idx >= 0 {
		pending = append(pending[:idx], pending[idx+1:]...)
	}
	ret := make([]string, 0)
	ret = append(ret, v)
	ret = append(ret, pending...)
	*values = ret
}
func PutToTail(values *[]string, v string) {
	if v == "" {
		return
	}

	*values = append(*values, v)
}
func contain(list []string, v string) bool {
	return findIdx(list, v) >= 0
}
func findIdx(list []string, v string) int {
	for i, it := range list {
		if it == v {
			return i
		}
	}
	return -1
}

func systemPause() {
	fmt.Println("按任意键继续...")
	b := make([]byte, 1)
	os.Stdin.Read(b)
	//bufio.NewReader(os.Stdin).ReadBytes('\n')
	//bufio.NewReader(os.Stdin).ReadByte()
}
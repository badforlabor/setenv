/**
 * Auth :   liubo
 * Date :   2019/10/15 22:07
 * Comment:
 */

package main

import (
	"fmt"
	"golang.org/x/sys/windows/registry"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const(
	OpRestore = "restore"
	OpSet = "set"
	OpHead = "head"
	OpTail = "tail"
)

func main() {
	// 读取配置文件
	workDir := GetExecPath()

	cfg := EnvConfig{}
	filePath := filepath.Join(workDir, "cfg.json")

	b := loadCfg(filePath, &cfg)
	if !b {
		panic("没有配置文件")
		return
	}

	// 处理
	if cfg.Op == OpSet {
		// 先备份所有涉及到的环境变量
		// 然后在设置
		restoreInfo := EnvConfig{}
		restoreInfo.Op = OpRestore
		for _, v := range cfg.Args {
			envValue := GetEnv(v.Key)
			if envValue != "" {
				one := v
				one.Op = OpSet
				one.Value = envValue
				restoreInfo.Args = append(restoreInfo.Args, one)
			}
		}

		// saveconfig.
		if len(restoreInfo.Args) > 0 {
			hostname, _ := os.Hostname()
			filename := fmt.Sprintf("%s-%s.json", hostname, time.Now().Format("20060102-150405"))
			fmt.Printf("filename:%s\n", filename)
			saveCfg(filename, restoreInfo)
		}

		// 处理环境变量
		for _, v := range cfg.Args {
			DoAction(v)
		}

	} else if cfg.Op == OpRestore {

	}

	pause()
}
func DoAction(one OneEnvOp) {
	if one.Op == OpSet {
		SetEnv(one.Key, one.Value)
	} else if one.Op == OpHead {
		oldValue := GetEnv(one.Key)
		valueList := strings.Split(oldValue, ";")
		newValueList := strings.Split(one.Value, ";")
		for i := len(newValueList)-1; i >= 0; i-- {
			PutToHead(&valueList, newValueList[i])
		}
		SetEnv(one.Key, strings.Join(valueList, ";"))
	} else if one.Op == OpTail {
		oldValue := GetEnv(one.Key)
		valueList := strings.Split(oldValue, ";")
		newValueList := strings.Split(one.Value, ";")
		for i := len(newValueList)-1; i >= 0; i-- {
			PutToTail(&valueList, newValueList[i])
		}
		SetEnv(one.Key, strings.Join(valueList, ";"))
	}
}
func GetEnv(key string) string {
	//machine, err := registry.OpenKey(registry.LOCAL_MACHINE, "SYSTEM\\CurrentControlSet\\Control\\Session Manager\\Environment", registry.ALL_ACCESS)
	user, err := registry.OpenKey(registry.CURRENT_USER, "Environment", registry.ALL_ACCESS)
	if err == nil {
		v, _, err :=  user.GetStringValue(key)
		if err == nil {
			return v
		}
		user.Close()
	}
	return ""
}
func SetEnv(key string, value string)  {
	user, err := registry.OpenKey(registry.CURRENT_USER, "Environment", registry.ALL_ACCESS)
	if err == nil {
		user.SetStringValue(key, value)
		user.Close()
	}
}



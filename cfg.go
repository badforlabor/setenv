/**
 * Auth :   liubo
 * Date :   2019/10/15 22:12
 * Comment: 由https://app.quicktype.io/自动生成
 */

package main

import "encoding/json"

func UnmarshalOneEnvOp(data []byte) (OneEnvOp, error) {
	var r OneEnvOp
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *OneEnvOp) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalEnvConfig(data []byte) (EnvConfig, error) {
	var r EnvConfig
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *EnvConfig) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type EnvConfig struct {
	Args []OneEnvOp `json:"args"`
	Op   string     `json:"op"`
}

type OneEnvOp struct {
	Key   string `json:"key"`
	Op    string `json:"op"`
	Value string `json:"value"`
}

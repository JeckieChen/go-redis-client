package service

import (
	"encoding/json"
	"errors"
	"go-redis-client/internal/define"
	"io/ioutil"
	"os"
)

func ConnectionList() ([]*define.Connection, error) {
	nowPath, _ := os.Getwd()
	data, err := ioutil.ReadFile(nowPath + string(os.PathSeparator) + define.ConfigName)
	if errors.Is(err, os.ErrNotExist) {
		return []*define.Connection{{
			Identity: "1",
			Name:     "1",
			Addr:     "1",
			Port:     "1",
			Username: "1",
			Password: "1",
		}}, nil
	}
	conf := new(define.Config)
	err = json.Unmarshal(data, err)
	if err != nil {
		return nil, err
	}
	return conf.Connection, nil
}

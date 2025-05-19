package service

import (
	"context"
	"go-redis-client/internal/define"
	"go-redis-client/internal/helper"
)

func SetValueDelete(req *define.SetValueRequest) error {
	rdb, err := helper.GetRedisClient(req.ConnIdentity, req.Db)
	err = rdb.SRem(context.Background(), req.Key, req.Value).Err()
	return err
}

func SetValueCreate(req *define.SetValueRequest) error {
	rdb, err := helper.GetRedisClient(req.ConnIdentity, req.Db)
	err = rdb.SAdd(context.Background(), req.Key, req.Value).Err()
	return err
}

package service

import (
	"context"
	"go-redis-client/internal/define"
	"go-redis-client/internal/helper"
)

func ListValueDelete(req *define.ListValueRequest) error {
	rdb, err := helper.GetRedisClient(req.ConnIdentity, req.Db)
	err = rdb.LRem(context.Background(), req.Key, 1, req.Value).Err()
	return err
}

func ListValueCreate(req *define.ListValueRequest) error {
	rdb, err := helper.GetRedisClient(req.ConnIdentity, req.Db)
	err = rdb.RPush(context.Background(), req.Key, req.Value).Err()
	return err
}

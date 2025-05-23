package test

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/go-redis/redis/v8"
)

var rdb = redis.NewClient(&redis.Options{
	Addr:     "10.10.16.39:6379",
	Password: "zg123456",
})

var ctx = context.Background()

func TestInfo(t *testing.T) {
	info, err := rdb.Info(context.Background()).Result()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(info)
	res, err := rdb.Info(context.Background(), "keyspace").Result()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
	dbs := strings.Split(res, "\n")
	m := make(map[string]int)
	fmt.Println(m)
	for i := 1; i < len(dbs)-1; i++ {
		v := strings.Split(dbs[i], ":")
		if len(v) < 2 {
			continue
		}
		vv := strings.Split(v[1], ",")
		if len(vv) < 3 {
			continue
		}
		keyNumber := strings.Split(vv[0], "=")
		if len(keyNumber) < 2 {
			continue
		}
		num, err := strconv.Atoi(keyNumber[1])
		if err != nil {
			continue
		}
		m[v[0]] = num
	}
	fmt.Println(m)
}

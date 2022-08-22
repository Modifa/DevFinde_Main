package services

import (
	"context"
	"encoding/json"
	"os"

	models "github.com/Modifa/DevFinde_Main/models"

	"github.com/go-redis/redis/v8"
)

var (
	ctx = context.Background()
)

func GetDeveloperProfile(Key string) (bool, *models.DeveloperProfile) {
	var cp *models.DeveloperProfile

	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDISSERVER_HOST") + ":" + os.Getenv("REDISSERVER_PORT"),
		Password: os.Getenv("REDISSERVER_PASSWORD"), // no password set
		DB:       3,                                 // use default DB
	})

	//err := rdb.Set(ctx, "TAXIMONEY:TAXIPROFILE:"+taxino, taxino, 0).Err()

	val, err := rdb.Get(ctx, Key).Result()

	defer rdb.Close()
	if err != nil {
		//panic(err)
		return false, cp
	}

	byt := []byte(val)

	if err := json.Unmarshal(byt, &cp); err != nil {
		panic(err)
	}

	//fmt.Println("key", val)

	return true, cp

}

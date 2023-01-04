package services

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	models "github.com/Modifa/DevFinde_Main/models"

	"github.com/go-redis/redis/v8"
)

var (
	ctx = context.Background()
)

func GetDeveloperProfile(Key string) (bool, *models.DeveloperProfile) {
	var cp *models.DeveloperProfile

	rdb := redis.NewClient(&redis.Options{
		Addr:       os.Getenv("REDISSERVER_HOST") + ":" + os.Getenv("REDISSERVER_PORT"),
		Password:   os.Getenv("REDISSERVER_PASSWORD"), // no password set
		DB:         0,
		MaxConnAge: 0, // use default DB
	})

	//err := rdb.Set(ctx, "TAXIMONEY:TAXIPROFILE:"+taxino, taxino, 0).Err()

	val, err := rdb.Get(ctx, "DEVELOPER:"+Key).Result()

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

//Set Redis Developer Profile
func SaveDeveloperprofile(User models.DeveloperProfile) error {
	rdb := redis.NewClient(&redis.Options{
		Addr:       os.Getenv("REDISSERVER_HOST") + ":" + os.Getenv("REDISSERVER_PORT"),
		Password:   os.Getenv("REDISSERVER_PASSWORD"), // no password set
		DB:         0,
		MaxConnAge: 0,
	})
	b, err := json.Marshal(User)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	/**/
	Newkey := strings.ToUpper(User.UserName)
	err = rdb.Set(ctx, "DEVELOPER:"+Newkey, b, 0).Err()
	return err
}

//Set Redis Developer Profile
func SaveDeveloperResume(User models.ResumeResponse) error {
	rdb := redis.NewClient(&redis.Options{
		Addr:       os.Getenv("REDISSERVER_HOST") + ":" + os.Getenv("REDISSERVER_PORT"),
		Password:   os.Getenv("REDISSERVER_PASSWORD"), // no password set
		DB:         0,
		MaxConnAge: 0,
	})
	b, err := json.Marshal(User)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	/**/

	Newkey := strings.ToUpper(User.Username)
	err = rdb.Set(ctx, "RESUME:"+Newkey, b, 0).Err()
	return err
}

//Save Portfolio Links
func SaveDeveloperLinks(DeveloperLinks models.LinksRequestReponse, Username string) error {
	rdb := redis.NewClient(&redis.Options{
		Addr:       os.Getenv("REDISSERVER_HOST") + ":" + os.Getenv("REDISSERVER_PORT"),
		Password:   os.Getenv("REDISSERVER_PASSWORD"), // no password set
		DB:         0,                                 /*LookUP*/
		MaxConnAge: 0,
	})

	b, err := json.Marshal(DeveloperLinks)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// err = rdb.Set(ctx, "PORTFOLIO:TRANSACTIONS:"+TransactionId, b, 0).Err()
	err = rdb.LPush(ctx, "LINKS:"+Username+":", b).Err()

	if err != nil {
		fmt.Println(err)
	}

	return err
}

//	err := rdb.Del(ctx, "TPAY:LOOKUPS:CITIES:"+ProvinceId).Err()
func ClearDeveloperLinks(Username string) error {
	rdb := redis.NewClient(&redis.Options{
		Addr:       os.Getenv("REDISSERVER_HOST") + ":" + os.Getenv("REDISSERVER_PORT"),
		Password:   os.Getenv("REDISSERVER_PASSWORD"), // no password set
		DB:         0,                                 /*LookUP*/
		MaxConnAge: 0,
	})

	Newkey := strings.ToUpper(Username)
	err := rdb.Del(ctx, "LINKS:"+Newkey+":").Err()

	if err != nil {
		fmt.Println(err)
	}

	return err
}

//
//Save Portfolio Links
func SaveDeveloperEnducation(DeveloperLinks models.Education, Username string) error {
	rdb := redis.NewClient(&redis.Options{
		Addr:       os.Getenv("REDISSERVER_HOST") + ":" + os.Getenv("REDISSERVER_PORT"),
		Password:   os.Getenv("REDISSERVER_PASSWORD"), // no password set
		DB:         0,                                 /*LookUP*/
		MaxConnAge: 0,
	})

	b, err := json.Marshal(DeveloperLinks)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// err = rdb.Set(ctx, "PORTFOLIO:TRANSACTIONS:"+TransactionId, b, 0).Err()
	err = rdb.LPush(ctx, "LINKS:"+Username+":", b).Err()

	if err != nil {
		fmt.Println(err)
	}

	return err
}

func GetAllDeveloperLinks(key string) []models.LinksRequestReponse {
	rdb := redis.NewClient(&redis.Options{
		Addr:       os.Getenv("REDISSERVER_HOST") + ":" + os.Getenv("REDISSERVER_PORT"),
		Password:   os.Getenv("REDISSERVER_PASSWORD"), // no password set
		DB:         0,                                 /*Transaction Redis DB*/
		MaxConnAge: 0,
	})

	Links := []models.LinksRequestReponse{}

	r, _ := rdb.LRange(ctx, key, 0, -1).Result()

	for _, val := range r {

		d := models.LinksRequestReponse{}

		json.Unmarshal(json.RawMessage(val), &d)

		Links = append(Links, d)

	}

	return Links
}

//
//Save Developer Experinces
func SaveDeveloperExperience(DeveloperLinks models.ExperienceResponseDB, Username string) error {
	rdb := redis.NewClient(&redis.Options{
		Addr:       os.Getenv("REDISSERVER_HOST") + ":" + os.Getenv("REDISSERVER_PORT"),
		Password:   os.Getenv("REDISSERVER_PASSWORD"), // no password set
		DB:         0,                                 /*LookUP*/
		MaxConnAge: 0,
	})

	b, err := json.Marshal(DeveloperLinks)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// err = rdb.Set(ctx, "PORTFOLIO:TRANSACTIONS:"+TransactionId, b, 0).Err()
	err = rdb.LPush(ctx, "EXPERIENCE:"+Username+":", b).Err()

	if err != nil {
		fmt.Println(err)
	}

	return err
}

//
func ClearDeveloperExperience(Username string) error {
	rdb := redis.NewClient(&redis.Options{
		Addr:       os.Getenv("REDISSERVER_HOST") + ":" + os.Getenv("REDISSERVER_PORT"),
		Password:   os.Getenv("REDISSERVER_PASSWORD"), // no password set
		DB:         0,                                 /*LookUP*/
		MaxConnAge: 0,
	})

	Newkey := strings.ToUpper(Username)
	err := rdb.Del(ctx, "EXPERIENCE:"+Newkey+":").Err()

	if err != nil {
		fmt.Println(err)
	}

	return err
}

//

func GetDeveloperResume(Key string) (bool, *models.ResumeResponse) {
	var cp *models.ResumeResponse

	rdb := redis.NewClient(&redis.Options{
		Addr:       os.Getenv("REDISSERVER_HOST") + ":" + os.Getenv("REDISSERVER_PORT"),
		Password:   os.Getenv("REDISSERVER_PASSWORD"), // no password set
		DB:         0,
		MaxConnAge: 0, // use default DB
	})

	//err := rdb.Set(ctx, "TAXIMONEY:TAXIPROFILE:"+taxino, taxino, 0).Err()

	val, err := rdb.Get(ctx, "RESUME:"+Key).Result()

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

//
func GetDeveloperEducation(key string) (bool, []models.Education) {
	rdb := redis.NewClient(&redis.Options{
		Addr:       os.Getenv("REDISSERVER_HOST") + ":" + os.Getenv("REDISSERVER_PORT"),
		Password:   os.Getenv("REDISSERVER_PASSWORD"), // no password set
		DB:         0,                                 /*Transaction Redis DB*/
		MaxConnAge: 0,
	})

	Education := []models.Education{}

	r, err := rdb.LRange(ctx, key, 0, -1).Result()
	defer rdb.Close()
	if err != nil {
		//panic(err)
		return false, Education
	}

	for _, val := range r {

		d := models.Education{}

		json.Unmarshal(json.RawMessage(val), &d)

		Education = append(Education, d)

	}

	return true, Education
}

//
func GetDeveloperExperienceRD(key string) (bool, []models.ExperienceResponseDB) {
	rdb := redis.NewClient(&redis.Options{
		Addr:       os.Getenv("REDISSERVER_HOST") + ":" + os.Getenv("REDISSERVER_PORT"),
		Password:   os.Getenv("REDISSERVER_PASSWORD"), // no password set
		DB:         0,                                 /*Transaction Redis DB*/
		MaxConnAge: 0,
	})

	Exp := []models.ExperienceResponseDB{}

	r, err := rdb.LRange(ctx, key, 0, -1).Result()
	defer rdb.Close()
	if err != nil {
		//panic(err)
		return false, Exp
	}

	for _, val := range r {

		d := models.ExperienceResponseDB{}

		json.Unmarshal(json.RawMessage(val), &d)

		Exp = append(Exp, d)

	}

	return true, Exp
}

//
func GetDeveloperLinksRD(key string) (bool, []models.LinksRequestReponse) {
	rdb := redis.NewClient(&redis.Options{
		Addr:       os.Getenv("REDISSERVER_HOST") + ":" + os.Getenv("REDISSERVER_PORT"),
		Password:   os.Getenv("REDISSERVER_PASSWORD"), // no password set
		DB:         0,                                 /*Transaction Redis DB*/
		MaxConnAge: 0,
	})

	Links := []models.LinksRequestReponse{}
	Newkey := strings.ToUpper(key)
	r, err := rdb.LRange(ctx, "LINKS:"+Newkey+":", 0, -1).Result()
	defer rdb.Close()
	if err != nil {
		//panic(err)
		return false, Links
	}

	for _, val := range r {

		d := models.LinksRequestReponse{}

		json.Unmarshal(json.RawMessage(val), &d)

		Links = append(Links, d)

	}

	return true, Links
}

////Set Redis Developer Profile
func SaveDeveloperResumeDesc(User models.ResumedescRedis, Key string) error {
	rdb := redis.NewClient(&redis.Options{
		Addr:       os.Getenv("REDISSERVER_HOST") + ":" + os.Getenv("REDISSERVER_PORT"),
		Password:   os.Getenv("REDISSERVER_PASSWORD"), // no password set
		DB:         0,
		MaxConnAge: 0,
	})
	b, err := json.Marshal(User)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	/**/
	Newkey := strings.ToUpper(Key)
	err = rdb.Set(ctx, "RESUME:DESCRIPTION:"+Newkey, b, 0).Err()
	return err
}

func GetDeveloperResumeDesc(Key string) (bool, *models.ResumedescRedis) {
	var cp *models.ResumedescRedis

	rdb := redis.NewClient(&redis.Options{
		Addr:       os.Getenv("REDISSERVER_HOST") + ":" + os.Getenv("REDISSERVER_PORT"),
		Password:   os.Getenv("REDISSERVER_PASSWORD"), // no password set
		DB:         0,
		MaxConnAge: 0, // use default DB
	})

	//err := rdb.Set(ctx, "TAXIMONEY:TAXIPROFILE:"+taxino, taxino, 0).Err()

	val, err := rdb.Get(ctx, "RESUME:DESCRIPTION:"+Key).Result()

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

package services

import (
	"context"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"

	"github.com/Modifa/DevFinde_Main/models"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4/pgxpool"
)

type DB struct {
}

//Register user
func (db *DB) RegisterDeveloper(functionnamewithschema string, m interface{}) ([]models.DeveloperResponseDB, error) {
	User := []models.DeveloperResponseDB{}
	u := ConVertInterface(functionnamewithschema, m)
	ctx := context.Background()
	db1, _ := pgxpool.Connect(ctx, os.Getenv("PostgresConString"))
	defer db1.Close()
	//
	err := pgxscan.Select(ctx, db1, &User, u)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			fmt.Println(pgErr.Message) // => syntax error at end of input
			fmt.Println(pgErr)         // => syntax error at end of input
			fmt.Println(pgErr.Code)    // => 42601
		}
	}
	return User, nil
}

func (db *DB) GetDeveloperLinks(functionnamewithschema string, m interface{}) ([]models.LinksRequestReponse, error) {
	User := []models.LinksRequestReponse{}
	u := ConVertInterface(functionnamewithschema, m)
	ctx := context.Background()
	db1, _ := pgxpool.Connect(ctx, os.Getenv("PostgresConString"))
	defer db1.Close()
	//
	err := pgxscan.Select(ctx, db1, &User, u)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			fmt.Println(pgErr.Message) // => syntax error at end of input
			fmt.Println(pgErr)         // => syntax error at end of input
			fmt.Println(pgErr.Code)    // => 42601
		}
	}
	return User, nil
}

//
func (db *DB) GetDeveloperExperience(functionnamewithschema string, m interface{}) ([]models.ExperienceResponseDB, error) {
	User := []models.ExperienceResponseDB{}
	u := ConVertInterface(functionnamewithschema, m)
	ctx := context.Background()
	db1, _ := pgxpool.Connect(ctx, os.Getenv("PostgresConString"))
	defer db1.Close()
	//
	err := pgxscan.Select(ctx, db1, &User, u)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			fmt.Println(pgErr.Message) // => syntax error at end of input
			fmt.Println(pgErr)         // => syntax error at end of input
			fmt.Println(pgErr.Code)    // => 42601
		}
	}
	return User, nil
}

//Get Developer Profile
func (db *DB) GetDeveloperProfile(functionnamewithschema string, m interface{}) ([]models.DeveloperProfile, error) {
	User := []models.DeveloperProfile{}
	u := ConVertInterface(functionnamewithschema, m)
	ctx := context.Background()
	db1, _ := pgxpool.Connect(ctx, os.Getenv("PostgresConString"))
	defer db1.Close()
	//
	err := pgxscan.Select(ctx, db1, &User, u)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			fmt.Println(pgErr.Message) // => syntax error at end of input
			fmt.Println(pgErr)         // => syntax error at end of input
			fmt.Println(pgErr.Code)    // => 42601
		}
	}
	return User, nil
}

//
func (db *DB) GetDeveloperResumeDesc(functionnamewithschema string, m interface{}) ([]models.ResumedescRes, error) {
	User := []models.ResumedescRes{}
	u := ConVertInterface(functionnamewithschema, m)
	ctx := context.Background()
	db1, _ := pgxpool.Connect(ctx, os.Getenv("PostgresConString"))
	defer db1.Close()
	//
	err := pgxscan.Select(ctx, db1, &User, u)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			fmt.Println(pgErr.Message) // => syntax error at end of input
			fmt.Println(pgErr)         // => syntax error at end of input
			fmt.Println(pgErr.Code)    // => 42601
		}
	}
	return User, nil
}

//ResumeResponse
func (db *DB) GetResume(functionnamewithschema string, m interface{}) ([]models.ResumeResponse, error) {
	User := []models.ResumeResponse{}
	u := ConVertInterface(functionnamewithschema, m)
	ctx := context.Background()
	db1, _ := pgxpool.Connect(ctx, os.Getenv("PostgresConString"))
	defer db1.Close()
	//
	err := pgxscan.Select(ctx, db1, &User, u)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			fmt.Println(pgErr.Message) // => syntax error at end of input
			fmt.Println(pgErr)         // => syntax error at end of input
			fmt.Println(pgErr.Code)    // => 42601
		}
	}
	return User, nil
}

//Education
func (db *DB) GetEducation(functionnamewithschema string, m interface{}) ([]models.Education, error) {
	User := []models.Education{}
	u := ConVertInterface(functionnamewithschema, m)
	ctx := context.Background()
	db1, _ := pgxpool.Connect(ctx, os.Getenv("PostgresConString"))
	defer db1.Close()
	//
	err := pgxscan.Select(ctx, db1, &User, u)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			fmt.Println(pgErr.Message) // => syntax error at end of input
			fmt.Println(pgErr)         // => syntax error at end of input
			fmt.Println(pgErr.Code)    // => 42601
		}
	}
	return User, nil
}

//Add
func (db *DB) SAVEONDB(functionnamewithschema string, m interface{}) (models.DBIDResponse, error) {
	User := models.DBIDResponse{}
	u := ConVertInterface(functionnamewithschema, m)
	ctx := context.Background()
	db1, _ := pgxpool.Connect(ctx, os.Getenv("PostgresConString"))
	defer db1.Close()
	//
	err := pgxscan.Select(ctx, db1, &User, u)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			fmt.Println(pgErr.Message) // => syntax error at end of input
			fmt.Println(pgErr)         // => syntax error at end of input
			fmt.Println(pgErr.Code)    // => 42601
		}
	}
	return User, nil
}

//Convert Interface and return Query string
func ConVertInterface(funcstr string, m interface{}) string {
	q := "select * from " + funcstr + "("

	if m != nil {

		v := reflect.ValueOf(m)
		typeOfS := v.Type()
		for i := 0; i < v.NumField(); i++ {

			switch typeOfS.Field(i).Type.Name() {
			case "int", "int16", "int32", "int64", "int8":
				str := v.Field(i).Interface().(int64)
				strInt64 := strconv.FormatInt(str, 10)
				q += strInt64 + ","
			case "float64":
				str := v.Field(i).Interface().(float64)
				s := fmt.Sprintf("%f", str)
				q += s + ","
			case "bool":
				q += "'" + strconv.FormatBool(v.Field(i).Interface().(bool)) + "',"
			default:
				if v.Field(i).Interface().(string) == "" {
					q += "null,"
				} else {
					q += "'" + v.Field(i).Interface().(string) + "',"
				}
			}
		}

		q = q[0 : len(q)-len(",")]
	}

	q += ")"

	return q
}

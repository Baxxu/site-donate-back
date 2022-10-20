package sql

import (
	"context"
	_ "embed"
	"github.com/Baxxu/site-donate-back/keys"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

var (
	DataBase = dataBase{}

	//go:embed CreateTableUsers.sql
	createTableUsers string

	//go:embed CreateTableSessions.sql
	createTableSessions string

	//go:embed CreateFuncGetUserIdWithTelegramId.sql
	createFuncGetUserIdWithTelegramId string

	//go:embed CreateFuncRefreshTokenValidate.sql
	createFuncSessionTokenValidate string
)

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)

	var err error

	DataBase.Pool, err = pgxpool.New(context.Background(), keys.DataBaseUrl)
	if err != nil {
		log.Printf("Unable to connect to database\n%s\n", err)
		panic(err)
	}

	_, err = DataBase.Pool.Exec(context.Background(), createTableUsers)
	if err != nil {
		log.Printf("%s", err)
		panic(err)
	}

	_, err = DataBase.Pool.Exec(context.Background(), createTableSessions)
	if err != nil {
		log.Printf("%s", err)
		panic(err)
	}

	_, err = DataBase.Pool.Exec(context.Background(), createFuncGetUserIdWithTelegramId)
	if err != nil {
		log.Printf("%s", err)
		panic(err)
	}

	_, err = DataBase.Pool.Exec(context.Background(), createFuncSessionTokenValidate)
	if err != nil {
		log.Printf("%s", err)
		panic(err)
	}
}

type dataBase struct {
	Pool *pgxpool.Pool
}

/*
CREATE DATABASE "siteDonation"
    WITH
    OWNER = postgres
    ENCODING = 'UTF8'
    CONNECTION LIMIT = -1
    IS_TEMPLATE = False;
*/

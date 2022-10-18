package sql

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

func (db *DataBase) Connect(DataBaseUrl string) {
	var err error

	db.Pool, err = pgxpool.New(context.Background(), DataBaseUrl)
	if err != nil {
		log.Printf("Unable to connect to database\n%s\n", err)
		panic(err)
	}

	_, err = db.Pool.Exec(context.Background(), CreateTableUsers)
	if err != nil {
		log.Printf("%s", err)
		panic(err)
	}

	_, err = db.Pool.Exec(context.Background(), CreateTableSessions)
	if err != nil {
		log.Printf("%s", err)
		panic(err)
	}

	_, err = db.Pool.Exec(context.Background(), CreateFuncGetUserIdWithTelegramId)
	if err != nil {
		log.Printf("%s", err)
		panic(err)
	}

	_, err = db.Pool.Exec(context.Background(), CreateFuncSessionTokenValidate)
	if err != nil {
		log.Printf("%s", err)
		panic(err)
	}
}

/*
CREATE DATABASE "siteDonation"
    WITH
    OWNER = postgres
    ENCODING = 'UTF8'
    CONNECTION LIMIT = -1
    IS_TEMPLATE = False;
*/

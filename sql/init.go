package sql

import (
	_ "embed"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

var (
	//go:embed CreateTableUsers.sql
	CreateTableUsers string

	//go:embed CreateTableSessions.sql
	CreateTableSessions string

	//go:embed CreateFuncGetUserIdWithTelegramId.sql
	CreateFuncGetUserIdWithTelegramId string

	//go:embed CreateFuncRefreshTokenValidate.sql
	CreateFuncSessionTokenValidate string
)

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
}

type DataBase struct {
	Pool *pgxpool.Pool
}

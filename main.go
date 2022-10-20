package main

import (
	"github.com/Baxxu/site-donate-back/api"
	"github.com/Baxxu/site-donate-back/auth"
	. "github.com/Baxxu/site-donate-back/sql"
	"log"
	"net/http"
	"time"
)

// TODO ДУДОС ЗАЩИТА

// TODO если у пользователя больше 100 сессий, то удалять все и добавлять новую

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
}

func main() {
	defer DataBase.Pool.Close()

	mux := http.NewServeMux()

	mux.HandleFunc("/auth/Telegram", auth.Telegram)
	mux.HandleFunc("/auth/GetAccessToken", auth.GetAccessToken)
	mux.HandleFunc("/auth/DeleteAllSessions", auth.DeleteAllSessions)
	mux.HandleFunc("/auth/DeleteSession", auth.DeleteSession)

	//mux.HandleFunc("/api/MakeTransaction", api.MakeTransaction)
	mux.HandleFunc("/api/Test", api.Test)

	server := http.Server{
		Addr:           "127.0.0.1:8080",
		Handler:        mux,
		MaxHeaderBytes: 4096,
		ReadTimeout:    time.Second * 10,
		WriteTimeout:   time.Second * 10,
		IdleTimeout:    time.Second * 60,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

/*
func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}
*/

package main

import (
	"log"
	"github.com/armon/go-socks5"
	"os"
	"io/ioutil"
	"strings"
	"github.com/caarlos0/env"
	"strconv"
)

type config struct {
	File         string        `env:"FILE" envDefault:"users"`
	Port         int           `env:"PORT" envDefault:"1111"`
}

func readUsersFromFile(path string) map[string]string {
	us, err := ioutil.ReadFile(path)
	if err != nil {
		log.Panic("Can't get users", err.Error())
	}

	users := make(map[string]string)

	for _, l := range strings.Split(string(us), "\n") {
		if l != "" {
			userPass := strings.Split(l, ":")
			users[userPass[0]] = userPass[1]
		}
	}

	return users
}

func main() {
	cfg := config{}
	err := env.Parse(&cfg)
	if err != nil {
		log.Printf("%+v\n", err)
	}
	log.Printf("Params run: %+v\n", cfg)

	log.Println("Init users..")
	users := make(socks5.StaticCredentials)
	for u, p := range readUsersFromFile(cfg.File) {
		users[u] = p
		log.Printf("User: %s", u)
	}
	auth := socks5.UserPassAuthenticator{Credentials: users}

	log.Println("Configuration..")
	srvConfig := &socks5.Config{
		AuthMethods: []socks5.Authenticator{auth},
		Logger:      log.New(os.Stdout, "", log.LstdFlags),
	}

	srv, err := socks5.New(srvConfig)
	if err != nil {
		log.Panic(err.Error())
	}

	log.Println("Start server")
	log.Println("Listen on 0.0.0.0:" + strconv.Itoa(cfg.Port))
	srv.ListenAndServe("tcp", "0.0.0.0:" + strconv.Itoa(cfg.Port))
}

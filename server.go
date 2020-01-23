package main

import (
	socks5 "github.com/vlakam/go-socks5"
	"log"
	"os"
)

func main() {
	socks5conf := &socks5.Config{
		AuthMethods: []socks5.Authenticator{
			&socks5.UserPassAuthenticator{
				Credentials: socks5.StaticCredentials{
					os.Getenv("SOCKS_USER"): os.Getenv("SOCKS_PASS"),
				},
			},
		},
	}
	log.Printf("Auth: %s:%s", os.Getenv("SOCKS_USER"), os.Getenv("SOCKS_PASS"))

	server, err := socks5.New(socks5conf)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	log.Print("Starting server on 1080 port")
	err = server.ListenAndServe("tcp", "0.0.0.0:1080")
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}
}

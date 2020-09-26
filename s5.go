package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"

	"github.com/armon/go-socks5"
)

func ProxyServer(username string, password string, port string) {

	creadentials := socks5.StaticCredentials{
		username: password,
	}
	authenticator := socks5.UserPassAuthenticator{Credentials: creadentials}

	// Create a SOCKS5 server
	config := &socks5.Config{
		AuthMethods: []socks5.Authenticator{authenticator},
		Logger:      log.New(os.Stdout, "", log.LstdFlags),
	}
	server, err := socks5.New(config)
	if err != nil {
		panic(err)
	}

	// Create SOCKS5 proxy on localhost port 1080
	address := fmt.Sprintf("0.0.0.0:%s", port)
	// fmt.Println(address)
	if err := server.ListenAndServe("tcp", address); err != nil {
		panic(err)
	}
}
func Daemon() {
	args := os.Args
	filePath := os.Args[0]
	cmd := exec.Command(filePath, args[1], args[2], args[3])
	cmd.Stdin = nil
	cmd.Stdout = nil
	cmd.Stderr = nil
	cmd.SysProcAttr = &syscall.SysProcAttr{Setsid: true}

	err := cmd.Start()
	if err == nil {
		_ = cmd.Process.Release()
		os.Exit(0)
	} else {
		fmt.Print(err)
	}
}
func main() {
	args := os.Args
	if len(args) != 4 {
		fmt.Println("./s5 username password port")
		return
	}
	if os.Getppid() != 1 {
		fmt.Println("./s5 is running")
		fmt.Println(args)
		Daemon()
	}
	fmt.Println("./s5 is start proxyServer")

	ProxyServer(args[1], args[2], args[3])
}

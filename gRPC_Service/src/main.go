package main

import (
	"flag"
	"fmt"
	"github.com/kardianos/service"
	"log"
	"os"
)

var (

	/**
	* gRPC 监听端口
	*/
	addr = flag.String("host", "0.0.0.0:9002", "")
)

type program struct{}

func (p *program) Start(s service.Service) error {
	go p.run()
	return nil
}

func (p *program) run() {

	log.Println("StartRPC()...")
	StartRPC()

}

func (p *program) Stop(s service.Service) error {
	return nil
}

func main() {

	svcConfig := &service.Config{
		Name:        "UAMUserService",
		DisplayName: "UAMUserService",
		Description: "UAM WindowsAD.UserManager gRPCService",
	}

	prg := &program{}
	s,err := service.New(prg,svcConfig)
	if err != nil {
		log.Println(err)
	}

	if len(os.Args) > 1 {
		if os.Args[1] == "install" {
			s.Install()
			fmt.Println("install service success")
			return
		}

		if os.Args[1] == "remove" {
			s.Uninstall()
			fmt.Println("remove service success")
			return
		}
	}

	err = s.Run()
	if err != nil {
		log.Println(err)
	}

	return
}

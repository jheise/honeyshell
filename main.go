package main

import (
	// standard
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	// external
	log "github.com/sirupsen/logrus"
)

func main() {
	// Setup args
	flag.Parse()
	args := flag.Args()
	logPath := args[0]
	ipInfo := strings.Split(args[1], ":")

	// Setup logging
	f, err := os.OpenFile(logPath,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Setup logging
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(f)
	log.SetLevel(log.InfoLevel)

	// start loop
	running := true
	for running {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("> ")
		cmd, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		cmd = strings.TrimSpace(cmd)
		log.WithFields(log.Fields{"commands": cmd, "ipaddr": ipInfo[0], "port": ipInfo[1]}).Info("cmd received")
		if strings.Compare(cmd, "exit") == 0 {
			running = false
		}
	}

}

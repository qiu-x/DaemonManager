package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
	"time"
	"bufio"
	"strings"
)

const SERVICE_PATH = "./services"
const CONFIG_PATH = "./conf"
const LOG_PATH = "./logs"

// Color escape sequences
const (
	colorReset = "\033[0m"
	colorRed = "\033[31m"
	colorGreen = "\033[32m"
	colorBlue = "\033[34m"
)

func initialize() {
	enable_file, err := os.Open(CONFIG_PATH + "/" + "enabled")
	if err != nil {
		fmt.Println(err)
	}
	defer enable_file.Close()
	scanner := bufio.NewScanner(enable_file)
	var service_files []string
	for scanner.Scan() {
		line := scanner.Text()
		service_files = append(service_files, strings.TrimSpace(line))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	for _, file := range service_files {
		var service Service
		if _, err := toml.DecodeFile(file, &service.Info); err != nil {
			fmt.Println("error: ", file, "not found")
		}
		fmt.Println("Starting:", service.Info.Name)
		go service.Run()
	}

	// TODO: handle ipc
	time.Sleep(1000 * time.Second)
}

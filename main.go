package main

import "log"

func main() {
	cfg := ParseConfig()
	if len(cfg.ApiList) == 0 {
		log.Println()
		return
	}
	for _, api := range cfg.ApiList {

	}
}

func sendApi() {

}

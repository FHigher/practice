package main

import (
	"fmt"
	"log"
	"os/exec"
)

func getData(dataChan chan []byte) {
	var data []byte
	var err error
	cmd := exec.Command("sudo sh", "-c", "lotus chain list | awk -F '[:,]' '{print$5}' | sort | uniq -c | sort -n")
	if data, err = cmd.CombinedOutput(); nil != err {
		log.Println("getData CombinedOutput err " + err.Error())
	}

	dataChan <- data
}

func main() {
	dataChan := make(chan []byte)

	getData(dataChan)
	fmt.Println(<-dataChan)
}

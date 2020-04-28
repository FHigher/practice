package main

import (
	"archive/tar"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)

	files, err := ioutil.ReadDir("files")
	if nil != err {
		log.Fatalln(err)
	}

	for _, file := range files {
		hdr, err := tar.FileInfoHeader(file, "")
		if nil != err {
			log.Fatalln(err)
		}
		fmt.Println(hdr, hdr.FileInfo().IsDir())

		if err := tw.WriteHeader(hdr); err != nil {
			log.Fatalln(err)
		}
		if !file.IsDir() {
			body, err := ioutil.ReadFile("files/" + file.Name())
			if nil != err {
				log.Fatalln(err)
			}
			if _, err := tw.Write(body); nil != err {
				log.Fatalln(err)
			}
		}

	}

	if err := tw.Close(); err != nil {
		log.Fatalln(err)
	}

	file, _ := os.Create("files.tar")

	n, err := io.Copy(file, buf)
	fmt.Println(n, err)
	file.Close()
}

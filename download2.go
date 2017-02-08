package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func download(url, name string) error {
	resp, err := http.Get(url)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	_, err = io.Copy(f, resp.Body)
	f.Close()
	if err != nil {
		os.Remove(name)
		return err
	}
	return nil
}

func run() error {
	f, err := os.Open("data2.json")
	if err != nil {
		return err
	}
	defer f.Close()
	d := json.NewDecoder(f)
	var data [][]string
	err = d.Decode(&data)
	if err != nil {
		return err
	}
	for i, item := range data {
		fmt.Printf("%d/%d %s\n", i+1, len(data), item[0])
		name := fmt.Sprintf("src2/%03d-%s.png", i+1, item[0])
		err := download(item[1], name)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

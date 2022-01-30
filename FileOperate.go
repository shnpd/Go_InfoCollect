package main

import "os"

func WriteFile(file string, jsonByte []byte) {
	fp, err := os.OpenFile(file, os.O_RDWR, 0666)
	if err != nil && os.IsNotExist(err) {
		os.Create(file)
	}
	defer fp.Close()
	_, err = fp.Write(jsonByte)
	if err != nil {
		panic(err)
	}
}

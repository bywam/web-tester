package utils

import (
	"io/ioutil"
	"log"
	"os"
)

func ReadConfigFile(fileName string) []byte {
	filePath := "./data/" + fileName + ".json"

	_, err := os.Stat(filePath)
	if err != nil {
		log.Fatal("No config file: " + fileName)
	}

	fileCnt, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer fileCnt.Close()

	byteVal, _ := ioutil.ReadAll(fileCnt)

	return byteVal
}

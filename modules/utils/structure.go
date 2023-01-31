package utils

import (
	"encoding/json"
	"log"
)

type T0Config struct {
	Urls []Url `json:"urls"`
}

type Url struct {
	Url       string   `json:"url"`
	Checks    []string `json:"checks"`
	ChecksCnt int      `json:"min_checks_cnt"`
}

func StructureConfigT0(fileCnt []byte) *T0Config {
	var urlItems T0Config
	err := json.Unmarshal(fileCnt, &urlItems)

	if err != nil {
		log.Fatal(err)
	}

	return &urlItems
}

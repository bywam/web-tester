package tasks

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
	"web-tester/modules/utils"
)

func prepareConfigT0(taskName string) *utils.T0Config {
	taskCnt := utils.ReadConfigFile(taskName)
	return utils.StructureConfigT0(taskCnt)
}

type T0Process struct {
	Type     string
	Url      string
	Response string
	Content  string
}

func taskProcess(process T0Process) {
	res, err := utils.HttpRequest(http.MethodGet, process.Url)
	if err != nil {
		fmt.Printf(process.Url+" Task ERROR \n - %s", err)
	}

	if process.Type == "url-status-code-check" {
		if res.StatusCode == 200 {
			fmt.Println(process.Url + " - Status Code 200 - PASS")
		} else {
			fmt.Println(process.Url + " - Status Code 200 - FAIL")
		}
	}

	if process.Type == "url-response-content-check" {
		resBody, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Printf(process.Url+" Task ERROR \n - %s", err)
		}
		if strings.Contains("ok", string(resBody)) {
			fmt.Println(process.Url + " - Text ok - PASS")
		} else {
			fmt.Println(process.Url + " - Text ok - FAIL")
		}
	}

	// todo save result to DB
}

func RunTasksT0(taskName string) {
	taskItems := prepareConfigT0(taskName)
	sleepInterval := 300

	for _, t := range taskItems.Urls {
		checkedUrl := t.Url
		checkCnt := t.ChecksCnt

		var process T0Process
		process.Url = checkedUrl

		if utils.Contains("status_code", t.Checks) {
			fmt.Printf("Status code |200| check at: %s with %d times.\n", checkedUrl, checkCnt)
			process.Type = "url-status-code-check"

			n := 0
			for n < checkCnt {
				taskProcess(process)
				time.Sleep(time.Millisecond * time.Duration(sleepInterval))
				n++
			}
			fmt.Println("")
			fmt.Println("-----")

		}

		if utils.Contains("text", t.Checks) {
			fmt.Printf("Content text |ok| check at: %s with %d times.\n", checkedUrl, checkCnt)
			process.Type = "url-response-content-check"

			n := 0
			for n < checkCnt {
				taskProcess(process)
				time.Sleep(time.Millisecond * time.Duration(sleepInterval))
				n++
			}
			fmt.Println("")
			fmt.Println("-----")
		}
	}
}

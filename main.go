package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type ProcesssDetailResponse struct {
	Event struct {
		CreatedDate string `json:"created_date"`
		Data        struct {
			Data        string                 `json:"data"`
			Cookies     map[string]string      `json:"cookies"`
			Files       map[string]interface{} `json:"files"`
			Form        map[string]interface{} `json:"form"`
			Headers     map[string]string      `json:"headers"`
			Method      string                 `json:"Method"`
			QueryParams map[string]string      `json:"query_params"`
			Url         string                 `json:"url"`
		}
		Id          string                 `json:"id"`
		Meta        map[string]interface{} `json:"meta"`
		TriggerType string                 `json:"trigger_type"`
	} `json:"event"`
	LairTrigger struct {
		ArchivedDate            string                 `json:"archived_date"`
		Command                 string                 `json:""`
		CreatedDate             string                 `json:"created_date"`
		Data                    map[string]interface{} `json:"data"`
		LairId                  string                 `json:"lair_id"`
		Settings                map[string]string      `json:"settings"`
		TestEvent               string                 `json:"test_event"`
		TriggerId               string                 `json:"trigger_id"`
		Type                    string                 `json:"type"`
		WorkspaceId             string                 `json:"workspace_id"`
		WorskpasceIntegrationId string                 `json:"worskpasce_integration_id"`
	} `json:"lair_trigger"`
	Process struct {
		Command       string `json:"command"`
		CompletedDate string `json:"completed_date"`
		CreatedDate   string `json:"created_date"`
		EventId       string `json:"event_id"`
		Id            string `json:"id"`
		LairId        string `json:"lair_id"`
		Port          string `json:"port"`
		ServiceId     string `json:"service_id"`
		Status        string `json:"status"`
		TriggerId     string `json:"trigger_id"`
		WorkspaceId   string `json:"workspace_id"`
	} `json:"process"`
}

type LayerDetailResponse struct {
	WorkspaceId string `json:`
}

func main() {
	res, err := GetUserByApplicationKey()
	fmt.Println(res)
	fmt.Println(err)

}

func getProcessDetail() (ProcesssDetailResponse, error) {

	fmt.Println(getProcessId())
	endpoint := fmt.Sprintf("processes/%s/detail", getProcessId())
	resp, err := makeRequest("GET", endpoint, os.Getenv("WAYSCRIPT_EXECUTION_USER_TOKEN"))

	defer resp.Body.Close()
	//var ResponseMap map[string]interface{}
	responseMap := &ProcesssDetailResponse{}
	body, err := ioutil.ReadAll(resp.Body)
	//strbody := string(body)
	//bodyText, err := strconv.Unquote(strbody)
	if err != nil {
		fmt.Println(err.Error())
		return *responseMap, err
	}
	//fmt.Println(bodyText)
	fmt.Println(string(body))
	err = json.Unmarshal(body, &responseMap)
	if err != nil {
		fmt.Println(err.Error())
		return *responseMap, err
	}

	fmt.Println(responseMap)
	fmt.Println((err))
	fmt.Println("~~~event~~~~")
	fmt.Println((responseMap.Event))

	return *responseMap, nil
}
func getProcessId() string {
	//return "3592c0e1-7c00-4740-a9da-d8816897ae6b"
	return os.Getenv("WS_PROCESS_ID")
}
func makeRequest(verb string, endpoint string, authToken string) (*http.Response, error) {

	client := &http.Client{}
	url := fmt.Sprintf("%s/%s", os.Getenv("WAYSCRIPT_ORIGIN"), endpoint)
	req, err := http.NewRequest(verb, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("authorization", fmt.Sprintf("Bearer %s", authToken))
	req.Header.Add("content-type", "application/json")

	resp, err := client.Do(req)
	return resp, err
}

// func GetEvent()(Resp) {
//res, err :=
// }

func GetUserByApplicationKey() (string, error) {

	res, err := getProcessDetail()
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	//endpoint := fmt.Sprintf("lairs/%s", res.Process.LairId)
	//get Layer detail
	endpoint := fmt.Sprintf("workspaces/%s/users/self", res.Process.WorkspaceId)
	res2, err := makeRequest("GET", endpoint, os.Getenv("WAYSCRIPT_EXECUTION_USER_APPLICATION_KEY"))
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	body, err := ioutil.ReadAll(res2.Body)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	fmt.Println(string(body))

	return string(body), nil

	//

}

// func SendHttpTriggerResponse() {

// }

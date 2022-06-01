package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type ProcesssDetailResponse struct {
	Event struct {
		CreatedDate time.Time `json:"created_date"`
		Data        struct {
			Data        string                      `json:"data"`
			Cookies     map[string]string           `json:"cookies"`
			Files       map[interface{}]interface{} `json:"files"`
			Form        map[interface{}]interface{} `json:"form"`
			Headers     map[string]string           `json:"headers"`
			Method      string                      `json:"Method"`
			QueryParams map[string]string           `json:"query_params"`
			Url         string                      `json:"url"`
		}
		Id          string                      `json:"id"`
		Meta        map[interface{}]interface{} `json:"meta"`
		TriggerType string                      `json:"trigger_type"`
	} `json:"event"`
	LairTrigger struct {
		ArchivedDate            time.Time                   `json:"archived_date"`
		Command                 string                      `json:""`
		CreatedDate             time.Time                   `json:"created_date"`
		Data                    map[interface{}]interface{} `json:"data"`
		LairId                  string                      `json:"lair_id"`
		Settings                map[string]string           `json:"settings"`
		TestEvent               string                      `json:"test_event"`
		TriggerId               string                      `json:"trigger_id"`
		Type                    string                      `json:"type"`
		WorkspaceId             string                      `json:"workspace_id"`
		WorskpasceIntegrationId string                      `json:"worskpasce_integration_id"`
	} `json:"lair_trigger"`
	Process struct {
		Command       string    `json:"command"`
		CompletedDate time.Time `json:"completed_date"`
		CreatedDate   time.Time `json:"created_date"`
		EventId       string    `json:"event_id"`
		Id            string    `json:"id"`
		LairId        string    `json:"lair_id"`
		Port          string    `json:"port"`
		ServiceId     string    `json:"service_id"`
		Status        string    `json:"status"`
		TriggerId     string    `json:"trigger_id"`
		WorkspaceId   string    `json:"workspace_id"`
	} `json:"process"`
}

func main() {

	fmt.Println(getProcessId())
	endpoint := fmt.Sprintf("processes/%s/detail", getProcessId())
	resp, err := makeRequest("GET", endpoint, "")

	defer resp.Body.Close()
	//var ResponseMap map[string]interface{}
	responseMap := &ProcesssDetailResponse{}
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &responseMap)
	if err != nil {
		fmt.Println(err.Error())
	}

	eventMap := ResponseMap["event"].(map[string]interface{})
	fmt.Println(eventMap)
	fmt.Println(eventMap["created_at"])
	fmt.Println((err))

}
func getProcessId() string {
	//return "3592c0e1-7c00-4740-a9da-d8816897ae6b"
	return os.Getenv("WS_PROCESS_ID")
}
func makeRequest(verb string, endpoint string, body string) (*http.Response, error) {

	client := &http.Client{}
	url := fmt.Sprintf("%s/%s", os.Getenv("WAYSCRIPT_ORIGIN"), endpoint)
	req, err := http.NewRequest(verb, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("authorization", fmt.Sprintf("Bearer %s", os.Getenv("WAYSCRIPT_EXECUTION_USER_TOKEN")))
	req.Header.Add("content-type", "application/json")

	resp, err := client.Do(req)
	return resp, err
}

/*
func GetEvent() {


}

func GetUserByApplicationID() {

}

func SendHttpTriggerResponse() {

}
*/

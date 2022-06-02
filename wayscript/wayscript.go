package wayscript

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type ProcesssDetailResponse struct {
	Event       `json:"event"`
	LairTrigger `json:"lair_trigger"`
	Process     `json:"process"`
}
type Event struct {
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
}
type LairTrigger struct {
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
}
type Process struct {
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
}

func GetProcessDetail() (ProcesssDetailResponse, error) {
	endpoint := fmt.Sprintf("processes/%s/detail", os.Getenv("WS_PROCESS_ID"))
	resp, err := makeRequest(http.MethodGet, endpoint, os.Getenv("WAYSCRIPT_EXECUTION_USER_TOKEN"))

	defer resp.Body.Close()
	responseMap := &ProcesssDetailResponse{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return *responseMap, err
	}
	err = json.Unmarshal(body, &responseMap)
	if err != nil {
		return *responseMap, err
	}

	return *responseMap, nil
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

func GetEvent() (Event, error) {
	processsDetail, err := GetProcessDetail()
	if err != nil {
		return Event{}, err
	}
	return processsDetail.Event, nil

}

func GetUserByApplicationKey() (string, error) {
	processDetail, err := GetProcessDetail()
	if err != nil {
		return "", err
	}

	endpoint := fmt.Sprintf("workspaces/%s/users/self", processDetail.Process.WorkspaceId)
	res, err := makeRequest(http.MethodGet, endpoint, os.Getenv("WAYSCRIPT_EXECUTION_USER_APPLICATION_KEY"))
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func SendHttpTriggerResponse(jsonData []byte) (string, error) {
	client := &http.Client{}
	url := fmt.Sprintf("%s/webhooks/http-trigger/response/%s", os.Getenv("WAYSCRIPT_ORIGIN"), os.Getenv("WS_PROCESS_ID"))

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	req.Header.Add("authorization", fmt.Sprintf("Bearer %s", os.Getenv("WAYSCRIPT_EXECUTION_USER_TOKEN")))
	req.Header.Add("content-type", "application/json")
	res, err := client.Do(req)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

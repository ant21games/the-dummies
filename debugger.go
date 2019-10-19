package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type TrainingQuestion struct {
	Question     string   `json:"question"`
	PlayerId     string   `json:"player_id"`
	QuestionId   string   `json:"question_id"`
	Alternatives []string `json:"alternatives"`
}

func AskQuestion(question TrainingQuestion, confg Configuration) (err error) {
	setProps := debugCmd{
		Cmd: "ask",
	}
	setProps.Args = map[string]commandArgs{
		"ask": map[string]interface{}{
			"player_id":    question.PlayerId,
			"alternatives": question.Alternatives,
			"question_id":  question.QuestionId,
			"question":     question.Question,
		},
	}
	_, err = sendDebugMsg(setProps, confg)
	return
}

func sendDebugMsg(msg debugCmd, confg Configuration) (newState GameMessage, err error) {
	uri := new(url.URL)
	uri.Scheme = "http"
	uri.Host = fmt.Sprintf("%s:%s", confg.WSHost, confg.WSPort)
	uri.Path = fmt.Sprintf("/%s/debug", confg.UUID)
	jsonValue, _ := json.Marshal(msg)

	resp, err := http.Post(uri.String(), "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return GameMessage{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
		var msg GameMessage
		err := json.Unmarshal(body, &msg)
		if err != nil {
			return GameMessage{}, err
		}
		newState = msg
	}
	return
}

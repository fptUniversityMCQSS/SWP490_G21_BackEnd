package ultity

import (
	"SWP490_G21_Backend/model"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func SendFileRequest(url string, method string, path string) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, errFile1 := os.Open(path)
	defer file.Close()
	part1,
		errFile1 := writer.CreateFormFile("file", filepath.Base(path))
	_, errFile1 = io.Copy(part1, file)
	if errFile1 != nil {
		fmt.Println(errFile1)
		return
	}
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

type QuestionRequest struct {
	Qn      int64           `json:"qn"`
	Content string          `json:"content"`
	Options []OptionRequest `json:"options"`
}

type OptionRequest struct {
	Key     string `json:"key"`
	Content string `json:"content"`
}

func SendQuestions(url string, method string, questions []*model.Question) {
	var questionRequests []QuestionRequest
	for _, question := range questions {
		var optionRequests []OptionRequest
		for _, option := range question.Options {
			o := OptionRequest{
				Key:     option.Key,
				Content: option.Content,
			}
			optionRequests = append(optionRequests, o)
		}
		q := QuestionRequest{
			Qn:      question.Number,
			Content: question.Content,
			Options: optionRequests,
		}
		questionRequests = append(questionRequests, q)
	}
	jsonQuestions, err := json.Marshal(questionRequests)
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Println(string(jsonQuestions))
	payload := strings.NewReader(string(jsonQuestions))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

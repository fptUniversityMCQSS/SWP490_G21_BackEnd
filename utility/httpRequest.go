package utility

import (
	"SWP490_G21_Backend/model"
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func SendFileRequest(url string, method string, path string) error {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, err := os.Open(path)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)
	part1, err := writer.CreateFormFile("file", filepath.Base(path))
	_, err = io.Copy(part1, file)
	if err != nil {
		return err
	}
	err = writer.Close()
	if err != nil {
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(res.Body)

	reader := bufio.NewReader(res.Body)
	str := ""
	for {
		b, err := reader.ReadByte()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		str += string(b)
		if reader.Buffered() <= 0 {
			println(str)
			str = ""
		}
	}
	return nil
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

func SendQuestions(url string, method string, questions []*model.Question) (*http.Response, error) {
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

		return nil, err
	}
	FileLog.Println(string(jsonQuestions))
	payload := strings.NewReader(string(jsonQuestions))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {

		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {

		return nil, err
	}
	return res, nil
}

func DeleteKnowledge(url string, method string, name string) error {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("name", name)
	err := writer.Close()
	if err != nil {
		return err
	}
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(res.Body)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	FileLog.Println(string(body))
	return nil
}

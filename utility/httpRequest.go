package utility

import (
	"SWP490_G21_Backend/model"
	"bufio"
	"bytes"
	"encoding/json"
	"io"
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
		log.Println(errFile1)
		return
	}
	err := writer.Close()
	if err != nil {
		log.Println(err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		log.Println(err)
		return
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer res.Body.Close()

	reader := bufio.NewReader(res.Body)
	str := ""
	for {
		b, err := reader.ReadByte()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal("Error reading HTTP response: ", err.Error())
		}
		str += string(b)
		if reader.Buffered() <= 0 {
			println(str)
			str = ""
		}
	}
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
		log.Println(err)
		return nil, err
	}
	log.Println(string(jsonQuestions))
	payload := strings.NewReader(string(jsonQuestions))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return res, nil
	//defer func(Body io.ReadCloser) {
	//	err := Body.Close()
	//	if err != nil {
	//		return
	//	}
	//}(res.Body)
	//
	//reader := bufio.NewReader(res.Body)
	//return reader, nil
}

package logic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"translate/model"
)

func (*logicBasic) Translate(r model.TranslateData) (*model.TranslateData, interface{}, error) {
	var raw interface{}
	URL := "https://translation.googleapis.com/language/translate/v2"

	client := &http.Client{}
	req, err := http.NewRequest("POST", URL, nil)
	if err != nil {
		return nil, nil, err
	}

	q := req.URL.Query()
	q.Add("key", os.Getenv("APIKEY"))
	q.Add("target", r.TranslatedType)
	q.Add("q", r.SourceText)
	q.Add("format", "text")
	req.URL.RawQuery = q.Encode()
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(body, &raw)
	if err != nil {
		return nil, nil, err
	}

	var dataResp model.TranslateBasicResp
	err = json.Unmarshal(body, &dataResp)
	if err != nil {
		return nil, raw, err
	}

	if len(dataResp.Data.Translations) >= 1 {
		r.TranslatedText = dataResp.Data.Translations[0].TranslatedText
		r.SourceType = dataResp.Data.Translations[0].DetectedSourceLanguage
	}

	return &r, raw, nil
}

func (*logicAdvance) Translate(r model.TranslateData) (*model.TranslateData, interface{}, error) {
	var reqBody model.BodyAdvanceReq
	var raw interface{}

	projectId := os.Getenv("PROJECTID")
	tokenOAuth2 := os.Getenv("TOKENOAUTH2")

	URL := fmt.Sprintf("https://translation.googleapis.com/v3/projects/%s:translateText", projectId)

	// reqBody.SourceLanguageCode = r.SourceType
	reqBody.TargetLanguageCode = r.TranslatedType
	reqBody.Contents = []string{r.SourceText}
	reqBody.MimeType = "text/plain"
	b, err := json.Marshal(&reqBody)
	if err != nil {
		return nil, nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(b))
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", "Bearer "+tokenOAuth2)
	req.Header.Set("x-goog-user-project", projectId)

	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(body, &raw)
	if err != nil {
		return nil, nil, err
	}

	var dataResp model.TranslateAdvanceResp
	err = json.Unmarshal(body, &dataResp.Data)
	if err != nil {
		return nil, raw, err
	}

	if len(dataResp.Data.Translations) >= 1 {
		r.TranslatedText = dataResp.Data.Translations[0].TranslatedText
		r.SourceType = dataResp.Data.Translations[0].DetectedLanguageCode
	}

	return &r, raw, nil
}

package model

type TranslateBasicResp struct {
	Data struct {
		Translations []TranslateBasicRespDetail `json:"translations"`
	} `json:"data"`
}

type TranslateBasicRespDetail struct {
	TranslatedText         string `json:"translatedText"`
	DetectedSourceLanguage string `json:"detectedSourceLanguage"`
}

type TranslateAdvanceResp struct {
	Data struct {
		Translations []TranslateAdvanceRespDetail `json:"translations"`
	} `json:"data"`
}

type TranslateAdvanceRespDetail struct {
	TranslatedText         string `json:"translatedText"`
	DetectedSourceLanguage string `json:"detectedLanguageCodetr"`
}

type TranslateData struct {
	SourceText     string `json:"Source_text"`
	SourceType     string `json:"Source_type"`
	TranslatedText string `json:"translated_text"`
	TranslatedType string `json:"translated_type"`
}

type Response struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Bix     interface{} `json:"bix,omitempty"`
}

// advance
type BodyAdvanceReq struct {
	SourceLanguageCode string   `json:"sourceLanguageCode"`
	TargetLanguageCode string   `json:"targetLanguageCode"`
	Contents           []string `json:"contents"`
	MimeType           string   `json:"mimeType"`
}

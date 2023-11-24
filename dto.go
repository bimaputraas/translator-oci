package main

type BasicResponse struct {
	Data struct {
		Translations []BasicResponseDetail `json:"translations"`
	} `json:"data"`
}

type BasicResponseDetail struct {
	TranslatedText         string `json:"translatedText"`
	DetectedSourceLanguage string `json:"detectedSourceLanguage"`
}

type BasicPayloadResponseDetail struct {
	SourceText     string `json:"Source_text"`
	SourceType     string `json:"Source_type"`
	TranslatedText string `json:"translated_text"`
	TranslatedType string `json:"translated_type"`
}

type BasicPayloadResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

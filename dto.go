package main

type BasicTranslationResponse struct {
	Data struct {
		Translations []BasicTranslation `json:"translations"`
	} `json:"data"`
}

type BasicTranslation struct {
	TranslatedText         string `json:"translatedText"`
	DetectedSourceLanguage string `json:"detectedSourceLanguage"`
}

type BasicPayloadResponse struct {
	SourceText     string `json:"Source_text"`
	SourceType     string `json:"Source_type"`
	TranslatedText string `json:"translated_text"`
	TranslatedType string `json:"translated_type"`
}

type PayloadResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

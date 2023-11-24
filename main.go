package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	r := gin.Default()

	r.GET("/basic", BasicHandler)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func BasicHandler(c *gin.Context) {
	text := c.Query("text")
	targetLanguage := c.Query("target")
	encodedText := url.QueryEscape(text)

	baseURL := fmt.Sprintf("https://translation.googleapis.com/language/translate/v2")
	query := fmt.Sprintf("?q=%vtarget=%s&key=%s", &encodedText, targetLanguage, os.Getenv("APIKEY"))
	url := fmt.Sprintf("%s%s", baseURL, query)

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	raw := string(body)

	var translationResponse = new(BasicTranslationResponse)
	err = json.Unmarshal(body, &translationResponse)
	if err != nil {
		c.JSON(http.StatusOK, PayloadResponse{
			Message: err.Error(),
			Data:    raw,
		})
	}

	fmt.Println(raw)

	payloadResponse := PayloadResponse{
		Message: "ok",
		Data: BasicPayloadResponse{
			SourceText:     text,
			SourceType:     translationResponse.Data.Translations[0].DetectedSourceLanguage,
			TranslatedText: translationResponse.Data.Translations[0].TranslatedText,
			TranslatedType: targetLanguage,
		},
	}
	c.JSON(http.StatusOK, payloadResponse)
}

// func main() {
// 	apiKey := "AIzaSyClk_14JOCthiw9OtOy9gffzALV74BGhx4"
// 	text := "Hello, how are you?"

// 	url := fmt.Sprintf("https://translation.googleapis.com/language/translate/v2?key=%s&q=%s&target=en", apiKey, text)
// 	// https://translation.googleapis.com/language/translate/v2?key=AIzaSyClk_14JOCthiw9OtOy9gffzALV74BGhx4&q=Hello, how are you?&target=en

// 	resp, err := http.Get(url)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer resp.Body.Close()

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println(string(body))
// }

// Bahasa Spanyol: es
// Bahasa Prancis: fr
// Bahasa Jerman: de
// Bahasa Mandarin (Tionghoa): zh
// Bahasa Jepang: ja
// Bahasa Korea: ko
// Bahasa Arab: ar
// Bahasa Rusia: ru
// Bahasa Italia: it
// Bahasa Portugis: pt
// Bahasa Indonesia: id
// Bahasa Inggris: en

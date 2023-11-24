package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
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
	URL := "https://translation.googleapis.com/language/translate/v2"

	client := &http.Client{}
	req, err := http.NewRequest("POST", URL, nil)
	if err != nil {
		c.JSON(500,BasicPayloadResponse{
			Message: err.Error(),
		})
	}

	q := req.URL.Query()
    q.Add("key", os.Getenv("APIKEY"))
    q.Add("target", targetLanguage)
    q.Add("q", text)
    q.Add("format", "text")
    req.URL.RawQuery = q.Encode()
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(500,BasicPayloadResponse{
			Message: err.Error(),
		})
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(500,BasicPayloadResponse{
			Message: err.Error(),
		})
	}

	var data interface{}
	err = json.Unmarshal(body,&data)
	if err != nil {
		c.JSON(500,BasicPayloadResponse{
			Message: err.Error(),
			Data: data,
		})
	}

	c.JSON(http.StatusOK, data)
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

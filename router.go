package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/otiai10/gosseract"
	"github.com/rs/cors"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"sync"
	"tantk/ocrtest/config"
	"time"
)

type ctrl struct{
	mu sync.RWMutex
}

//Router create router
func (c *ctrl)Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/captcha", c.SolveCaptcha)
	return router
}

func (c *ctrl)SolveCaptcha(w http.ResponseWriter, r *http.Request) {
	c.mu.Lock()
	switch r.Method {
	case "POST":
		if r.Header.Get("Content-type") == "application/json" {
			var (
				tempRaw       = "temp.png"
				tempProcessed = "tempDone.png"
			)
			var body = new(struct {
				Url string `json:"Url"`
			})
			reqBody, err := ioutil.ReadAll(r.Body)
			json.Unmarshal(reqBody, &body)
			check(err)
			imgRes, err := http.Get(body.Url)
			check(err)
			file, err := os.Create(tempRaw)
			check(err)
			_, err = io.Copy(file, imgRes.Body)
			check(err)
			file.Close()
			Preprocess(tempRaw, tempProcessed)
			client := gosseract.NewClient()
			defer client.Close()
			client.Languages = []string{"eng"}
			client.SetImage(tempProcessed)
			text, err := client.Text()
			check(err)
			var resData = new(struct {
				Solved string `json:"Solved"`
			})
			reg, err := regexp.Compile("[^a-zA-Z0-9]+")
			check(err)
			testCleaned := reg.ReplaceAllString(text, "")
			resData.Solved=testCleaned
			json.NewEncoder(w).Encode(resData)
		}
	}
	c.mu.Unlock()
}

//StartServer :
func (c *ctrl)StartServer() {

	router := c.Router()
	//allow all domain for local testing
	cross := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"*"},
		AllowedMethods:   []string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"},
		AllowCredentials: true,
	})
	conf := config.GetConfig()
	handler := cross.Handler(router)

	server := &http.Server{
		Addr:         conf.RESTport,
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	fmt.Println("Listening at port ", conf.RESTport)
	server.ListenAndServe()
}

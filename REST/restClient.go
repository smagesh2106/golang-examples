package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	BaseURL  string
	BasePORT int
}

type RequestBody struct {
	Name string `json:"name"`
	Job  string `json:"job"`
}

type ResponseBody struct {
	Name      string `json:"name"`
	Job       string `json:"job"`
	Id        string `json:"id"`
	CreatedAt string `json:"createdAt"`
}

type DataResp struct {
	Id         int    `json:"id"`
	Email      string `json:"email,omitempty"`
	First_name string `json:"first_name,omitempty"`
	Last_name  string `json:"last_name"`
	Avatar     string `json:"avatar"`
}

type SupportResp struct {
	Url  string `json:"url"`
	Text string `json:"text"`
}

type ResponseBody2 struct {
	Data    DataResp    `json:"data"`
	Support SupportResp `json:"support"`
}

type userActions interface {
	CreateUser(requestBody *RequestBody) (*ResponseBody, error)
	GetUser(id string) (*ResponseBody2, error)
}

func (c *Client) CreateUser(requestBody *RequestBody) (*ResponseBody, error) {
	path := "/api/users"

	url := fmt.Sprintf("%s:%d%s", c.BaseURL, c.BasePORT, path)
	requestBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(requestBodyBytes))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)

	//resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Decode response body
	var responseBody ResponseBody
	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		fmt.Println("Error decoding response body:", err)
		return nil, err
	}

	return &responseBody, nil
}

func (c *Client) GetUser(id string) (*ResponseBody2, error) {
	var err error = nil
	var resp *http.Response = nil
	client := &http.Client{}
	path := "/api/users/2"

	urlPath := fmt.Sprintf("%s:%d%s", c.BaseURL, c.BasePORT, path)
	req, _ := http.NewRequest("GET", urlPath, nil)
	req.Header.Add("Accept", "application/json")
	//resp, err = client.Do(req)

	//url := fmt.Sprintf("%s:%d/api/users/%s", c.BaseURL, c.BasePORT, id)
	//fmt.Println(url)

	timeOut := 3 * time.Second
	deadLine := time.Now().Add(timeOut)

	//retry until timeout
	for tries := 0; time.Now().Before(deadLine); tries++ {
		//if resp, err = http.Get(url); err == nil {
		if resp, err = client.Do(req); err == nil {
			defer resp.Body.Close()
			break
		}
		fmt.Println("Error getting user:", err)
		time.Sleep(500 * time.Millisecond)
	}

	if resp == nil || err != nil {
		return nil, errors.New("Timeout")
	}

	var responseBody ResponseBody2
	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		fmt.Println("Error decoding response body:", err)
		return nil, err
	}

	return &responseBody, nil
}

func main() {

	client := &Client{
		BaseURL: "https://reqres.in",

		BasePORT: 443,
	}
	/*
	   url := "https://reqres.in/api/users"

	   // Create request body
	   requestBody := RequestBody{Name: "John Doe", Job: "Leader"}
	   requestBodyBytes, err := json.Marshal(requestBody)
	   if err != nil {
	       fmt.Println("Error encoding request body:", err)
	       return
	   }

	   // Make POST request
	   resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBodyBytes))
	   if err != nil {
	       fmt.Println("Error making POST request:", err)
	       return
	   }
	   defer resp.Body.Close()

	   // Decode response body
	   var responseBody ResponseBody
	   err = json.NewDecoder(resp.Body).Decode(&responseBody)
	   if err != nil {
	       fmt.Println("Error decoding response body:", err)
	       return
	   }

	   fmt.Println("Response message:", responseBody.CreatedAt)
	*/
	responseBody, err := client.CreateUser(&RequestBody{
		Name: "Paul Doe",
		Job:  "Leader",
	})
	if err != nil {
		fmt.Println("Error decoding response body:", err)
		return
	}

	fmt.Println("Response (POST) message:", responseBody.CreatedAt)

	responseBody2, err := client.GetUser("2")
	if err != nil {
		fmt.Println("Error decoding response body:", err)
		return
	}

	fmt.Println("Response (GET) message: ", responseBody2.Data.Email)
}

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
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
	url := fmt.Sprintf("%s:%d/api/users", c.BaseURL, c.BasePORT)
	requestBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBodyBytes))
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
	url := fmt.Sprintf("%s:%d/api/users/%s", c.BaseURL, c.BasePORT, id)
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
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
		BaseURL:  "https://reqres.in",
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

	fmt.Println("Response message:", responseBody.CreatedAt)

	responseBody2, err := client.GetUser("2")
	if err != nil {
		fmt.Println("Error decoding response body:", err)
		return
	}

	fmt.Println("Response message: ", responseBody2.Data.Email)
}

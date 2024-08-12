package httpz

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
)

type Header struct {
	Key   string
	Value string
}

type GetRequest struct {
	URL     string
	Headers []Header
}

type PostRequest struct {
	URL     string
	Headers []Header
	Body    map[string]interface{}
}

type DeleteRequest struct {
	URL     string
	Headers []Header
}

type PutRequest struct {
	URL     string
	Headers []Header
	Body    map[string]interface{}
}

type PatchRequest struct {
	URL     string
	Headers []Header
	Body    map[string]interface{}
}

func SendPostRequest[T any](req PostRequest) (T, error) {
	client := &http.Client{}
	reqURL, err := url.Parse(req.URL)
	if err != nil {
		log.Println("error parsing url", err)
		var null T
		return null, err
	}
	reqURL.RawQuery = reqURL.Query().Encode()
	req.URL = reqURL.String()

	reqHeaders := http.Header{}
	for _, header := range req.Headers {
		reqHeaders.Set(header.Key, header.Value)
	}
	reqHeaders.Set("Content-Type", "application/json")
	reqHeaders.Set("Accept", "application/json")

	reqBody, err := json.Marshal(req.Body)
	if err != nil {
		log.Println("error marshalling request body", err)
		var null T
		return null, err
	}

	resp, err := client.Post(req.URL, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		log.Println("error sending request", err)
		var null T
		return null, err
	}
	if resp.StatusCode > 299 {
		log.Println("server responded with status ", resp.StatusCode)
		var null T
		return null, err
	}
	defer resp.Body.Close()
	var response_value T

	bodyBytes, err := io.ReadAll(resp.Body) // Use io.ReadAll
	if err != nil {
		log.Println("error reading response body", err)
		var null T
		return null, err
	}

	err = json.Unmarshal(bodyBytes, &response_value) // Unmarshal into responseValue directly
	if err != nil {
		log.Println("error unmarshalling response body:", err)
		log.Println(string(bodyBytes))
		var null T
		return null, err
	}

	return response_value, nil
}
func SendGetRequest[T any](req GetRequest) (T, error) {
	client := &http.Client{}
	reqURL, err := url.Parse(req.URL)
	if err != nil {
		log.Println("error parsing url", err)
		var null T
		return null, err
	}
	reqURL.RawQuery = reqURL.Query().Encode()
	req.URL = reqURL.String()

	reqHeaders := http.Header{}
	for _, header := range req.Headers {
		reqHeaders.Set(header.Key, header.Value)
	}
	reqHeaders.Set("Content-Type", "application/json")
	reqHeaders.Set("Accept", "application/json")

	resp, err := client.Get(req.URL)
	if err != nil {
		log.Println("error sending request", err)
		var null T
		return null, err
	}
	if resp.StatusCode > 299 {
		log.Println("server responded with status ", resp.StatusCode)
		var null T
		return null, err
	}
	defer resp.Body.Close()
	var response_value T

	bodyBytes, err := io.ReadAll(resp.Body) // Use io.ReadAll
	if err != nil {
		log.Println("error reading response body", err)
		var null T
		return null, err
	}

	err = json.Unmarshal(bodyBytes, &response_value) // Unmarshal into responseValue directly
	if err != nil {
		log.Println("error unmarshalling response body:", err)
		log.Println(string(bodyBytes))
		var null T
		return null, err
	}

	return response_value, nil
}

func SendDeleteRequest[T any](req DeleteRequest) (T, error) {
	client := &http.Client{}
	reqURL, err := url.Parse(req.URL)
	if err != nil {
		log.Println("error parsing url", err)
		var null T
		return null, err
	}
	reqURL.RawQuery = reqURL.Query().Encode()
	req.URL = reqURL.String()
	reqHeaders := http.Header{}
	for _, header := range req.Headers {
		reqHeaders.Set(header.Key, header.Value)
	}
	reqHeaders.Set("Content-Type", "application/json")
	reqHeaders.Set("Accept", "application/json")
	resp, err := client.Do(&http.Request{
		Method: "DELETE",
		URL:    reqURL,
		Header: reqHeaders,
	})

	if err != nil {
		log.Println("error sending request", err)
		var null T
		return null, err
	}
	if resp.StatusCode > 299 {
		log.Println("server responded with status ", resp.StatusCode)
		var null T
		return null, err
	}
	defer resp.Body.Close()
	var response_value T
	bodyBytes, err := io.ReadAll(resp.Body) // Use io.ReadAll
	if err != nil {
		log.Println("error reading response body", err)
		var null T
		return null, err
	}
	err = json.Unmarshal(bodyBytes, &response_value) // Unmarshal into responseValue directly
	if err != nil {
		log.Println("error unmarshalling response body:", err)
		log.Println(string(bodyBytes))
		var null T
		return null, err
	}
	return response_value, nil
}

func SendPutRequest[T any](req PutRequest) (T, error) {
	client := &http.Client{}
	reqURL, err := url.Parse(req.URL)
	if err != nil {
		log.Println("error parsing url", err)
		var null T
		return null, err
	}
	reqURL.RawQuery = reqURL.Query().Encode()
	req.URL = reqURL.String()
	reqHeaders := http.Header{}
	for _, header := range req.Headers {
		reqHeaders.Set(header.Key, header.Value)
	}
	reqHeaders.Set("Content-Type", "application/json")
	reqHeaders.Set("Accept", "application/json")
	reqBody, err := json.Marshal(req.Body)
	if err != nil {
		log.Println("error marshalling request body", err)
		var null T
		return null, err
	}
	resp, err := client.Do(&http.Request{
		Method: "PUT",
		URL:    reqURL,
		Header: reqHeaders,
		Body:   io.NopCloser(bytes.NewBuffer(reqBody)),
	})
	if err != nil {
		log.Println("error sending request", err)
		var null T
		return null, err
	}
	if resp.StatusCode > 299 {
		log.Println("server responded with status ", resp.StatusCode)
		var null T
		return null, err
	}
	defer resp.Body.Close()
	var response_value T
	bodyBytes, err := io.ReadAll(resp.Body) // Use io.ReadAll
	if err != nil {
		log.Println("error reading response body", err)
		var null T
		return null, err
	}
	err = json.Unmarshal(bodyBytes, &response_value) // Unmarshal into responseValue directly
	if err != nil {
		log.Println("error unmarshalling response body:", err)
		log.Println(string(bodyBytes))
		var null T
		return null, err
	}
	return response_value, nil
}

func SendPatchRequest[T any](req PatchRequest) (T, error) {
	client := &http.Client{}
	reqURL, err := url.Parse(req.URL)
	if err != nil {
		log.Println("error parsing url", err)
		var null T
		return null, err
	}
	reqURL.RawQuery = reqURL.Query().Encode()
	req.URL = reqURL.String()
	reqHeaders := http.Header{}
	for _, header := range req.Headers {
		reqHeaders.Set(header.Key, header.Value)
	}
	reqHeaders.Set("Content-Type", "application/json")
	reqHeaders.Set("Accept", "application/json")
	reqBody, err := json.Marshal(req.Body)
	if err != nil {
		log.Println("error marshalling request body", err)
		var null T
		return null, err
	}
	resp, err := client.Do(&http.Request{
		Method: "PATCH",
		URL:    reqURL,
		Header: reqHeaders,
		Body:   io.NopCloser(bytes.NewBuffer(reqBody)),
	})
	if err != nil {
		log.Println("error sending request", err)
		var null T
		return null, err
	}
	if resp.StatusCode > 299 {
		log.Println("server responded with status ", resp.StatusCode)
		var null T
		return null, err
	}
	defer resp.Body.Close()
	var response_value T
	bodyBytes, err := io.ReadAll(resp.Body) // Use io.ReadAll
	if err != nil {
		log.Println("error reading response body", err)
		var null T
		return null, err
	}
	err = json.Unmarshal(bodyBytes, &response_value) // Unmarshal into responseValue directly
	if err != nil {
		log.Println("error unmarshalling response body:", err)
		log.Println(string(bodyBytes))
		var null T
		return null, err
	}
	return response_value, nil
}

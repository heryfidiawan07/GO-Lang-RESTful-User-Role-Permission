package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

func Endpoint() []string {
	return []string{
		"https://jsonplaceholder.typicode.com/posts/1",
		"https://jsonplaceholder.typicode.com/posts/2",
		"https://jsonplaceholder.typicode.com/posts/3",
		"https://jsonplaceholder.typicode.com/posts/4",
		"https://jsonplaceholder.typicode.com/posts/5",
		"https://jsonplaceholder.typicode.com/posts/6",
		"https://jsonplaceholder.typicode.com/posts/7",
		"https://jsonplaceholder.typicode.com/posts/8",
		"https://jsonplaceholder.typicode.com/posts/9",
		"https://jsonplaceholder.typicode.com/posts/10",
	}
}

func fetchData(url string) (map[string]interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}
	return data, nil
}

// 1. GoRoutine (tanpa WaitGroup, pakai channel)
func GoRoutine(c *gin.Context) {
	endpoints := Endpoint()
	results := []map[string]interface{}{}
	ch := make(chan map[string]interface{})

	// jalankan goroutine
	for _, url := range endpoints {
		go func(u string) {
			data, err := fetchData(u)
			if err == nil {
				ch <- data
			} else {
				ch <- map[string]interface{}{"error": err.Error()}
			}
		}(url)
	}

	// ambil hasil sesuai jumlah endpoint
	for range endpoints {
		results = append(results, <-ch)
	}

	c.JSON(200, gin.H{
		"status":  true,
		"data":    results,
		"message": nil,
	})
}

// 3. GoRoutineWait (pakai sync.WaitGroup)
func GoRoutineWaitGroup(c *gin.Context) {
	endpoints := Endpoint()
	results := make([]map[string]interface{}, len(endpoints))
	var wg sync.WaitGroup

	for i, url := range endpoints {
		wg.Add(1)
		go func(idx int, u string) {
			defer wg.Done()
			data, err := fetchData(u)
			if err == nil {
				results[idx] = data
			} else {
				results[idx] = map[string]interface{}{"error": err.Error()}
			}
		}(i, url)
	}

	wg.Wait()

	c.JSON(200, gin.H{
		"status": true, "data": results, "message": nil,
	})
}

// 3. Synchronous (no goroutine)
func Synchronous(c *gin.Context) {
	endpoints := Endpoint()
	results := []map[string]interface{}{}

	for _, url := range endpoints {
		data, err := fetchData(url)
		if err == nil {
			results = append(results, data)
		} else {
			results = append(results, map[string]interface{}{"error": err.Error()})
		}
	}

	c.JSON(200, gin.H{
		"status": true, "data": results, "message": nil,
	})
}

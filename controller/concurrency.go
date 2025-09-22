package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

// GoRoutine
func fetchDataGoRoutine(url string, ch chan map[string]interface{}) {
	resp, err := http.Get(url)
	if err != nil {
		ch <- map[string]interface{}{
			"url":   url,
			"error": err.Error(),
		}
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var data interface{}
	json.Unmarshal(body, &data)

	ch <- map[string]interface{}{
		"url":  url,
		"data": data,
	}
}

func GoRoutine(c *gin.Context) {
	urls := []string{
		"https://jsonplaceholder.typicode.com/posts",
		"https://jsonplaceholder.typicode.com/comments",
		"https://jsonplaceholder.typicode.com/todos",
	}

	ch := make(chan map[string]interface{})
	for _, url := range urls {
		go fetchDataGoRoutine(url, ch)
	}

	var results []map[string]interface{}
	for range urls {
		result := <-ch
		results = append(results, result)
	}

	c.JSON(200, gin.H{
		"status":  true,
		"message": "All requests completed with goroutine+channel",
		"data":    results,
	})
}

// Javascript
// Promise.all()
// *************

// const task = async () => {
// 	// simulate a workload
// 	// sleep(3000);
// 	return Math.floor(Math.random() * Math.floor(100));
// };

// const [a, b, c] = await Promise.all(
// 	task(),
// 	task(),
// 	task()
// );
// console.log(a, b, c);

// PromiseAll
func fetchDataPromiseAll(url string, wg *sync.WaitGroup, results chan<- map[string]interface{}) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		results <- map[string]interface{}{
			"url":   url,
			"error": err.Error(),
		}
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var data interface{}
	json.Unmarshal(body, &data)

	results <- map[string]interface{}{
		"url":  url,
		"data": data,
	}
}

func PromiseAll(c *gin.Context) {
	urls := []string{
		"https://jsonplaceholder.typicode.com/posts",
		"https://jsonplaceholder.typicode.com/comments",
		"https://jsonplaceholder.typicode.com/todos",
	}

	var wg sync.WaitGroup
	results := make(chan map[string]interface{}, len(urls))

	for _, url := range urls {
		wg.Add(1)
		go fetchDataPromiseAll(url, &wg, results)
	}

	wg.Wait()
	close(results)

	var allResults []map[string]interface{}
	for result := range results {
		allResults = append(allResults, result)
	}

	c.JSON(200, gin.H{
		"status":  true,
		"message": "All requests completed with WaitGroup",
		"data":    allResults,
	})
}

// JavaScript
// Async Await
// **********

// const worker = async () => {
// // simulate a workload
// sleep(3000);
// 	return Math.floor(Math.random() * Math.floor(100));
// };

// const r = await worker();
// console.log(r);

// AsyncAwait
func fetchDataSequential(url string) map[string]interface{} {
	resp, err := http.Get(url)
	if err != nil {
		return map[string]interface{}{
			"url":   url,
			"error": err.Error(),
		}
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var data interface{}
	json.Unmarshal(body, &data)

	return map[string]interface{}{
		"url":  url,
		"data": data,
	}
}

func AsyncAwait(c *gin.Context) {
	urls := []string{
		"https://jsonplaceholder.typicode.com/posts",
		"https://jsonplaceholder.typicode.com/comments",
		"https://jsonplaceholder.typicode.com/todos",
	}

	var results []map[string]interface{}
	for _, url := range urls {
		result := fetchDataSequential(url) // satu per satu
		results = append(results, result)
	}

	c.JSON(200, gin.H{
		"status":  true,
		"message": "All requests completed sequentially",
		"data":    results,
	})
}

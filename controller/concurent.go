package controller

import (
	// "fmt"
	"math/rand"

	"github.com/gin-gonic/gin"
)

func task() <-chan int32 {
	r := make(chan int32)

	go func() {
		defer close(r)
		r <- rand.Int31n(100)
	}()

	return r
}

func PromiseAll(c *gin.Context) {
	one, two, three := <-task(), <-task(), <-task()
	// a, b, c := <-one, <-two, <-three
	var result map[string]interface{}
	result = map[string]interface{}{
		"total":  one + two + three,
		"detail": []int32{one, two, three},
	}

	c.JSON(200, gin.H{"status": true, "data": result, "message": nil})
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

func worker(text string) <-chan string {
	r := make(chan string)

	go func() {
		defer close(r)
		r <- text
	}()

	return r
}

func AsyncAwait(c *gin.Context) {
	sayHello := <-worker("Hello")
	full := <-worker(sayHello + " World !")

	c.JSON(200, gin.H{"status": true, "data": full, "message": nil})
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

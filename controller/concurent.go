package controller

import (
	"app/config"
	"app/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

func task(start int, end int) <-chan []models.Vehicles {
	r := make(chan []models.Vehicles)

	go func() {
		defer close(r)

		var vehicles []models.Vehicles
		config.DB.Where("id >= ? AND id <= ?", start, end).Find(&vehicles)
		r <- vehicles
	}()

	return r
}

func results() []interface{} {
	// one, two, three, four, five := <-task(1, 2000), <-task(2001, 4000), <-task(4001, 6000), <-task(6001, 8000), <-task(8001, 10000)
	// // a, b, c, d, e := <-one, <-two, <-three, <-four, <-five

	// var result = []interface{}{
	// 	// a, b, c, d, e,
	// 	one, two, three, four, five,
	// }
	// //
	// return result

	var vehicles []models.Vehicles
	total := config.DB.Find(&vehicles).RowsAffected

	start := 1
	end := int(total / 10)

	// data := make(map[int]interface{})
	var data []interface{}

	for i := 1; i <= 10; i++ {
		if int(end) >= int(total) {
			end = int(total)
		}

		// data[i] = <-task(start, end)
		data = append([]interface{}{<-task(start, end)}, data...)
		//
		fmt.Println("start:(", start, ") - end:(", end, ") - of:(", int(total), ") - index=>", i)

		if int(end) >= int(total) {
			break
		}

		start = start + int(total/5)
		end = end + int(total/5)
	}

	return data
}

func ConcurentIndex(c *gin.Context) {
	// c.JSON(200, gin.H{"status": true, "data": results(), "message": nil})
	c.JSON(200, gin.H{"status": true, "data": results(), "message": nil})
}

package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var cartItems = []map[string]interface{}{}
var generatedTimetables = []map[string]interface{}{
	{
		"timetable_id": "uuid-1234",
		"courses": []map[string]interface{}{
			{
				"lecture_id":  "uuid-5678",
				"course_name": "운영체제",
				"time": []map[string]string{
					{"day": "Mon", "start_time": "13:00", "end_time": "14:15"},
					{"day": "Wed", "start_time": "13:00", "end_time": "14:15"},
				},
			},
			{
				"lecture_id":  "uuid-91011",
				"course_name": "데이터베이스",
				"time": []map[string]string{
					{"day": "Mon", "start_time": "13:00", "end_time": "14:15"},
					{"day": "Wed", "start_time": "13:00", "end_time": "14:15"},
				},
			},
		},
	},
}

// POST /carts
func AddToCart(c *gin.Context) {
	var body struct {
		LectureID string `json:"lecture_id"`
		Priority  int    `json:"priority"`
	}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid request"})
		return
	}
	item := map[string]interface{}{
		"item_id":    len(cartItems) + 1,
		"lecture_id": body.LectureID,
		"priority":   body.Priority,
	}
	cartItems = append(cartItems, item)

	c.JSON(http.StatusOK, gin.H{
		"msg":  "cart item POST 성공",
		"data": gin.H{"cart_item_id": item["item_id"], "cart_item_count": len(cartItems)},
	})
}

// PATCH /carts
func UpdateCart(c *gin.Context) {
	var body struct {
		Updates []struct {
			ItemID   int `json:"item_id"`
			Priority int `json:"priority"`
		} `json:"updates"`
	}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid request"})
		return
	}

	for _, update := range body.Updates {
		for _, item := range cartItems {
			if item["item_id"] == update.ItemID {
				item["priority"] = update.Priority
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "cart item POST 성공",
		"data": gin.H{"cart_item_id": body.Updates[0].ItemID, "cart_item_count": len(cartItems)},
	})
}

// DELETE /carts/:lectureId
func DeleteFromCart(c *gin.Context) {
	lectureId := c.Param("lectureId")
	newCart := []map[string]interface{}{}
	for _, item := range cartItems {
		if item["lecture_id"] != lectureId {
			newCart = append(newCart, item)
		}
	}
	cartItems = newCart
	c.JSON(http.StatusOK, gin.H{"message": "삭제 완료"})
}

// GET /carts
func GetCart(c *gin.Context) {
	response := []gin.H{}
	for _, item := range cartItems {
		response = append(response, gin.H{
			"lecture_id":  item["lecture_id"],
			"course_name": "운영체제", // 예시
			"priority":    item["priority"],
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "조회 성공",
		"data":    response,
	})
}

// GET /timetables
func GenerateTimetables(c *gin.Context) {
	minCredit := c.Query("min_credit")
	maxCredit := c.Query("max_credit")

	//query
	c.Writer.WriteString("min_credit: " + minCredit + ", max_credit: " + maxCredit + "\n")

	c.JSON(http.StatusOK, gin.H{
		"filtered_timetables": generatedTimetables,
	})
}

// POST /timetables/filter
func FilterTimetables(c *gin.Context) {
	var body struct {
		Options []string `json:"option"`
	}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid filter request"})
		return
	}

	c.JSON(http.StatusOK, []interface{}{generatedTimetables[0]})
}

// GET /timetables/:id
func GetTimetableDetail(c *gin.Context) {
	id := c.Param("id")
	for _, tt := range generatedTimetables {
		if tt["timetable_id"] == id {
			c.JSON(http.StatusOK, gin.H{"courses": tt["courses"]})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "시간표를 찾을 수 없습니다"})
}

package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SearchSubject(c *gin.Context) {
	c.JSON(http.StatusOK, []gin.H{
		{
			"lecture_id":  1,
			"course_name": "자료구조",
			"professor":   "김철수",
			"time": []gin.H{
				{"day": "Tue", "start_time": "10:00", "end_time": "11:15"},
				{"day": "Thu", "start_time": "10:00", "end_time": "11:15"},
			},
			"location":       "정보B101",
			"subject_code":   "CS101",
			"ai_description": "~~~에 대해서 배우는 과목",
			"credit":         3,
			"target_grade":   "2학년",
			"category":       "전공선택",
			"syllabus_url":   "https://example.com/syllabus/cs101",
			"capacity":       50,
			"note":           "수업은 영어로 진행되지 않음",
			"language":       "한국어",
		},
		{
			"lecture_id":  2,
			"course_name": "알고리즘",
			"professor":   "이영희",
			"time": []gin.H{
				{"day": "Mon", "start_time": "13:00", "end_time": "14:15"},
				{"day": "Wed", "start_time": "13:00", "end_time": "14:15"},
			},
			"location":       "정보A202",
			"subject_code":   "CS102",
			"ai_description": "~~~에 대해서 배우는 과목",
			"credit":         3,
			"target_grade":   "2학년",
			"category":       "전공선택",
			"syllabus_url":   "https://example.com/syllabus/cs102",
			"capacity":       60,
			"note":           "수업은 영어로 진행되지 않음",
			"language":       "한국어",
		},
	})
}

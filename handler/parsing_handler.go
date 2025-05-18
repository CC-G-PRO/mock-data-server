package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IsValidPdf(c *gin.Context) {
	file, _ := c.FormFile("file")
	if file == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"valid":   false,
			"message": "파일이 업로드되지 않았습니다.",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"valid":   true,
		"message": "유효한 졸업사정진단표 파일입니다.",
	})
}

func ParsePdf(c *gin.Context) {
	file, _ := c.FormFile("file")
	if file == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"valid":   false,
			"message": "파일이 업로드되지 않았습니다.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "파일 업로드 및 분석 완료",
		"report_id": 1,
		"data": gin.H{
			"basic_info": gin.H{
				"student_number":  "2023105682",
				"student_name":    "홍길동",
				"department":      "컴퓨터공학부",
				"grade":           "3",
				"enroll_semester": "5",
				"evaluation_date": "2025/03/31 10:27:50",
			},
			"graduation_info": gin.H{
				"credit": gin.H{
					"total_credits_earned":   "78(96)",
					"total_credits_required": "140",
					"valid":                  false,
				},
				"grades": gin.H{
					"earned":   "4.143",
					"required": "1.7",
					"valid":    true,
				},
				"english": gin.H{
					"earned":   "7",
					"required": "3",
					"valid":    true,
				},
				"paper": gin.H{
					"earned":   "0",
					"required": "1",
					"valid":    false,
				},
			},
			"liberal_arts_info": []gin.H{
				{
					"category": "배분이수교과(2019~2023)",
					"required": 12,
					"earned":   6,
					"valid":    false,
				},
				{
					"category": "자유이수",
					"required": 10,
					"earned":   3,
					"valid":    false,
				},
				{
					"category": "필수교과",
					"required": 17,
					"earned":   17,
					"valid":    true,
				},
			},
			"major_info": gin.H{
				"advanced_major": "컴퓨터공학과",
				"reference_year": 2023,
				"major_basic": gin.H{
					"earned":   15,
					"required": 18,
				},
				"major_": gin.H{
					"earned":   36,
					"required": 45,
				},
				"required_plus_elective": gin.H{
					"earned":   51,
					"required": 75,
				},
				"passed": false,
				"industry_required": gin.H{
					"earned":   3,
					"required": 12,
				},
			},
		},
	})
}

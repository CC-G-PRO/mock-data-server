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
		"valid":     true,
		"message":   "파일 업로드 및 분석 완료",
		"report_id": 1,
		"parsed": gin.H{
			"basic_info": gin.H{
				"학번":          "2023105682",
				"성명":          "홍길동",
				"학과":          "컴퓨터공학부",
				"학년":          "3",
				"등록횟수":        "5",
				"학적상태":        "재학",
				"입학구분":        "신입학",
				"외국인여부":       "N",
				"최종변동":        "2023/03/01",
				"최종판정":        "졸업유예",
				"자료 생성 일시":    "2025/03/31 10:27:50",
				"이수 과목 기준 일시": "2025/04/16 23:18:20",
				"자료 출력 일시":    "2025/04/16 23:18:20",
			},
			"graduation_info": gin.H{
				"수강학점": gin.H{
					"취득": "78(96)",
					"기준": "140",
					"판정": false,
				},
				"취득학점": gin.H{
					"취득": "0",
					"기준": "1",
					"판정": false,
				},
				"전공": gin.H{
					"취득": "2",
					"기준": "3",
					"판정": false,
				},
				"교양": gin.H{
					"취득": "4.143",
					"기준": "1.7",
					"판정": true,
				},
				"졸업평점": gin.H{
					"취득": "7",
					"기준": "3",
					"판정": true,
				},
				"영어강의": gin.H{
					"취득": "0",
					"기준": "1",
					"판정": false,
				},
				"논문": gin.H{
					"취득": "전산(기타)",
					"기준": "해당없음",
					"판정": true,
				},
				"TOPIK": gin.H{
					"취득": "면제",
					"기준": "외국어",
					"판정": true,
				},
			},
			"liberal_arts_info": []gin.H{
				{
					"이수 구분":     "배분이수교과(2019~2023)",
					"영역(취득/기준)": "1/0",
					"학점(취득/기준)": "6/12",
					"판정":        false,
				},
				{
					"이수 구분":     "자유이수",
					"영역(취득/기준)": "2/2",
					"학점(취득/기준)": "10/3",
					"판정":        true,
				},
				{
					"이수 구분":     "필수교과",
					"영역(취득/기준)": "3/3",
					"학점(취득/기준)": "17/17",
					"판정":        true,
				},
			},
			"major_info": gin.H{
				"심화전공": "컴퓨터공학과",
				"기준연도": 2023,
				"전공기초": gin.H{
					"취득": 15,
					"기준": 18,
				},
				"전필": gin.H{
					"취득": 36,
					"기준": 45,
				},
				"필수+선택": gin.H{
					"취득": 51,
					"기준": 75,
				},
				"판정": false,
				"산학 필수": gin.H{
					"수강학점": 3,
					"기준학점": 12,
				},
			},
		},
	})
}

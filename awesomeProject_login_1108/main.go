package main

import (
	"bufio"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

// 从文件中读取数据,文件存放在“C:\Users\DELL\awesomeProject_login_1108\credentials.txt”
func checkCredentials(username, password string) int {
	file, err := os.Open("credentials.txt")
	if err != nil {
		fmt.Println(err)
		valid := 400
		return valid
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		cred := parseCredentials(line)
		if cred.Username == username && cred.Password == password {
			valid := 200
			return valid
		}
		if cred.Username != username {
			valid := 300
			return valid
		}
		if cred.Username == username && cred.Password != password {
			valid := 500
			return valid
		}
	}

	valid := 400
	return valid
}

type credentials struct {
	Username string
	Password string
}

func parseCredentials(line string) credentials {
	parts := []string{"", ""}
	_, err := fmt.Sscanf(line, "%s %s", &parts[0], &parts[1])
	if err != nil {
		fmt.Println(err)
		return credentials{}
	}
	return credentials{Username: parts[0], Password: parts[1]}
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/index.html")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})
	router.POST("/", func(ctx *gin.Context) {
		username := ctx.PostForm("username")
		password := ctx.PostForm("password")

		valid := checkCredentials(username, password)

		if valid == 200 {
			ctx.JSON(200, gin.H{
				"message": "ok, pass",
			})
		}
		if valid == 300 {
			ctx.JSON(300, gin.H{
				"message": "ok, username error",
			})
		} else if valid == 400 {
			ctx.JSON(400, gin.H{
				"message": "txt doesn't function",
			})
		} else if valid == 500 {
			ctx.JSON(500, gin.H{
				"message": "ok,password doesn't match",
			})
		}

	})

	router.Run(":8000")
}

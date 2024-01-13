// package main

// import (
// 	// "os"
// 	// "os/exec"
// 	// "runtime"

//     "github.com/gin-gonic/gin"
// )

// func main() {
//     // 创建Gin的默认引擎
//     r := gin.Default()

//     // 定义路由处理函数
//     r.GET("/", func(c *gin.Context) {
//         c.JSON(200, gin.H{
//             "message": "Hello, World!",
//         })
//     })

//     // 启动Web服务器，监听在默认端口（:8080）
//     r.Run()

// 	// openBrowser("http://localhost:8080")
// }

// func openBrowser(url string) {
// 	var err error

// 	switch runtime.GOOS {
// 	case "darwin":
// 		err = exec.Command("open", url).Start()
// 	case "linux":
// 		err = exec.Command("xdg-open", url).Start()
// 	case "windows":
// 		err = exec.Command("cmd", "/c", "start", url).Start()
// 	default:
// 		err = os.ErrNotExist
// 	}

// 	if err != nil {
// 		// 处理错误，例如输出到控制台
// 		// 在实际应用中，你可能需要根据具体情况进行处理
// 		println("Error opening browser:", err)
// 	}
// }

package main

import (
	// "math/rand"
	"net/http"
	// "time"

	"github.com/gin-gonic/gin"
)


func main() {
    //初始化路由
    router := gin.Default()
    //加载templates中所有模板文件, 使用不同目录下名称相同的模板,注意:一定要放在配置路由之前才得行
    router.LoadHTMLGlob("templates/*")
    //router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")

    router.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", gin.H{ 
            "title": "首页", 
        })
    })
    
    router.Run(":8080")
}
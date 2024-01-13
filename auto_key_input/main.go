package main

import (
	// "os"
	// "bufio"
	"fmt"
	"github.com/micmonay/keybd_event"
	"time"
)

func main() {
	for {
	// 获取需要按下的键
	fmt.Print("请输入需要按下的键（例如 A): ")
	var keyToPress string
	fmt.Scanf("%s\n", &keyToPress)

	// 获取按键持续时间
	fmt.Print("请输入按键持续时间（秒）: ")
	var pressDuration int
	fmt.Scanln(&pressDuration)

	// 获取按键间隔
	fmt.Print("请输入按键间隔时间（毫秒）: ")
	var interval int
	fmt.Scanln(&interval)

	// 创建键盘事件实例
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		fmt.Println("无法创建键盘事件实例:", err)
		return
	}

	fmt.Println("请在3秒内切换到需要按键的窗口")
	time.Sleep(time.Duration(3000) * time.Millisecond)
	fmt.Println("开始运行")

	// 设置按下的键
	// Convert the input string to a key code
	var keyCode int
	switch keyToPress {
	case "A":
		keyCode = keybd_event.VK_A
	// Add more cases for other keys as needed
	default:
		fmt.Println("Invalid key")
		return
	}

// ...
kb.SetKeys(keyCode)

	// 执行按键循环
	startTime := time.Now()
	for time.Since(startTime) < time.Duration(pressDuration)*time.Second {
		// 模拟按下和释放键
		err := kb.Launching()
		if err != nil {
			fmt.Println("无法模拟按键:", err)
			return
		}

		// 间隔一段时间再次按键
		time.Sleep(time.Duration(interval) * time.Millisecond)
	}


	// 显示运行完成
	fmt.Println("运行完成")
	fmt.Println("请设置新的按键循环")
	
	}

}

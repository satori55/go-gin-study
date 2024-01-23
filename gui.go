package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/micmonay/keybd_event"
	"time"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Key Press Simulator")

	myWindow.Resize(fyne.NewSize(400, 300))

	// 创建 GUI 组件
	keyEntry := widget.NewEntry()
	keyEntry.SetPlaceHolder("Enter key to press (e.g., A)")

	durationEntry := widget.NewEntry()
	durationEntry.SetPlaceHolder("Enter duration in seconds")

	intervalEntry := widget.NewEntry()
	intervalEntry.SetPlaceHolder("Enter interval in milliseconds")

	logLabel := widget.NewLabel("Status")

	// 按键模拟函数
	simulateKeyPress := func() {
		keyToPress := keyEntry.Text
		var pressDuration int
		fmt.Sscanf(durationEntry.Text, "%d", &pressDuration)
		var interval int
		fmt.Sscanf(intervalEntry.Text, "%d", &interval)

		logLabel.SetText("switch to target window in 3 seconds")
		time.Sleep(time.Duration(3000) * time.Millisecond)
		logLabel.SetText("Starting...")

		kb, err := keybd_event.NewKeyBonding()
		if err != nil {
			logLabel.SetText(fmt.Sprintf("Error: %s", err))
			return
		}

		// 设置按键
		var keyCode int
		switch keyToPress {
		case "A":
			keyCode = keybd_event.VK_A
		// ... 其他按键
		default:
			logLabel.SetText("Invalid key")
			return
		}

		kb.SetKeys(keyCode)

		// 按键模拟循环
		startTime := time.Now()
		for time.Since(startTime) < time.Duration(pressDuration)*time.Second {
			err := kb.Launching()
			if err != nil {
				logLabel.SetText(fmt.Sprintf("Error: %s", err))
				return
			}
			time.Sleep(time.Duration(interval) * time.Millisecond)
		}
		logLabel.SetText("Completed")
	}

	// 创建按钮以触发按键模拟
	startButton := widget.NewButton("Start", func() {
		simulateKeyPress()
	})

	// 设置窗口内容
	myWindow.SetContent(container.NewVBox(
		keyEntry,
		durationEntry,
		intervalEntry,
		startButton,
		logLabel,
	))

	myWindow.ShowAndRun()
}

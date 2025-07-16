package utils

import (
	"fmt"
	"github.com/gen2brain/beeep"
)

func NotifySuccess(msg string) {
	err := beeep.Notify("✅ autoupd", msg, "")
	if err != nil {
		fmt.Println("Notification error:", err)
	}
}

func NotifyFailure(msg string) {
	err := beeep.Notify("❌ autoupd", msg, "")
	if err != nil {
		fmt.Println("Notification error:", err)
	}
}

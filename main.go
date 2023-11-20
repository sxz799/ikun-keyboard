package main

import (
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"ikun-keyboard/util"
	"log"
	"strings"
	"time"
)

var lastNum int

func paintedEggShell(num int) {
	if num == 1 {
		lastNum = 0
	}
	if lastNum+1 == num {
		lastNum = num
		if lastNum == 4 {
			log.Println("小黑子,露出鸡脚了吧!")
			time.Sleep(time.Second * 1)
			go util.PlaySound("JNTM")
			lastNum = 0
		}
	} else {
		lastNum = 0
	}
}

func main() {
	robotgo.EventHook(hook.KeyDown, []string{}, func(event hook.Event) {
		key := strings.ToUpper(string(event.Keychar))
		log.Println("PRESS " + key)
		switch key {
		case "J":
			go util.PlaySound("J")
			paintedEggShell(1)
		case "N":
			go util.PlaySound("N")
			paintedEggShell(2)
		case "T":
			go util.PlaySound("T")
			paintedEggShell(3)
		case "M":
			go util.PlaySound("M")
			paintedEggShell(4)
		}

	})
	s := robotgo.EventStart()
	<-robotgo.EventProcess(s)
}

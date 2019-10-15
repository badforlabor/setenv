/**
 * Auth :   liubo
 * Date :   2019/10/15 23:37
 * Comment:
 */

package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
)

func init() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	termbox.SetCursor(0, 0)
	termbox.HideCursor()
}
func pause() {
	fmt.Println("请按任意键继续...")
Loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			break Loop
		}
	}
}

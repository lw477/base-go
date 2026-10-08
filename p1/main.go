package main

import (
	"fmt"
	"strings"
)

func main() {
	target := "Go语言"

	guesses := []rune{'G', '语'}

	displayText, isWin := RevealWord(target, guesses)

	fmt.Println("目标词汇长度（字符数）:", len([]rune(target)))
	fmt.Println("游戏界面显示：", displayText)
	fmt.Println("是否全部猜中:", isWin)
}

// RevealWord 根据玩家已猜中的字符，生成游戏展示文本
func RevealWord(target string, guessed []rune) (string, bool) {
	targetRunes := []rune(target)

	var display strings.Builder
	allGuess := true

	for i, ch := range targetRunes {
		if i > 0 {
			display.WriteString(" ")
		}
		found := false
		for _, g := range guessed {
			if ch == g {
				found = true
				break
			}
		}

		if found {
			display.WriteRune(ch)
		} else {
			display.WriteString("_")
			allGuess = false
		}
	}
	return display.String(), allGuess
}

// 为什么需要rune？字节数与字符个数的区别

package commands

import (
	"bufio"
	"fmt"
	"os"
	"playdoll/doll"
)

func Dialogue() {
	Beggining()
	Greeting()
	FillIn()
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	switch scanner.Text() {
	case "Дракулаура":
		n, s, phrases := Draculaura(doll.Doll{})
		Choice(n, s, phrases[0])
		h := phrases[1]
		h2 := phrases[2]
		st := "- любовь? -"
		var exit bool
		Pause()
		for !exit {
			Communication(st, &exit, h, h2)
		}
	case "Клодин":
		n, s, phrase := Clawdeen(doll.Doll{})
		fmt.Println(n, s, phrase[0])
		st := "- что ты носишь? -"
		h := phrase[1]
		h2 := phrase[2]
		var exit bool
		Pause()
		for !exit {
			Communication(st, &exit, h, h2)
		}
	case "Френки":
		n, s, phrase := Frankie(doll.Doll{})
		fmt.Println(n, s, phrase[0])
		st := "- что делаешь? -"
		h := phrase[1]
		h2 := phrase[2]
		var exit bool
		Pause()
		for !exit {
			Communication(st, &exit, h, h2)
		}
	}
}

func Communication(st string, e *bool, kd string, h string) {
	Variants(st)
	FillIn()
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	switch scanner.Text() {
	case st:
		fmt.Println(h)
		*e = false
	case "как дела":
		fmt.Println(kd)
		*e = false
	case "пока":
		Bye()
		*e = true
	default:
		DonGetIt()
		*e = false
	}
}

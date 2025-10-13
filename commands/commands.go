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
	Fillin()
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	switch scanner.Text() {
	case "Дракулаура":
		n, s, phrase := Draculaura(doll.Doll{})
		Choice(n, s, phrase[0])
	case "Клодин":
		n, s, phrase := Clawdeen(doll.Doll{})
		fmt.Println(n, s, phrase[0])
	case "Френки":
		n, s, phrase := Frankie(doll.Doll{})
		fmt.Println(n, s, phrase[0])
	}
}

func Communication() {

}

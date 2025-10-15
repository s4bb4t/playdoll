package main

import (
	"bufio"
	"fmt"
	"os"
	"playdoll/doll"
	"strings"
)

const (
	dollChooseTemplate = `Пользователь окликнул куклу: `
	question           = `Он остановился напротив нее, немного помялся, и наконец произнес: `
	answer             = `"%s" - сказала %s %s.`
)

func main() {
	dolls := map[string]*doll.Doll{
		"дракулаура": initDrakulaura(),
	}

	sc := bufio.NewScanner(os.Stdin)
	fmt.Print(dollChooseTemplate)
	sc.Scan()

	choice := strings.ToLower(strings.TrimSpace(sc.Text())) // прочитай что делают эти функции. Выпиши и используй их в дальшейнем
	d, exists := dolls[choice]
	if exists {
		fmt.Print(question)
		sc.Scan()
		fmt.Printf(answer, d.Talk(sc.Text()), d.Name(), d.Surname())
	} else {
		fmt.Println("Но никто не ответил...")
	}

}

func initDrakulaura() *doll.Doll {
	dDialogues := map[string]string{
		"привет":   "О, хеллооо! Я так рада тебя видеть, просто кровь кипит!",
		"как дела": "Ужасно прекрасно! Пью смузи из клубники и читаю модный журнал 🩸",
		"любовь":   "Ой, я вся краснею! Любовь — это как закат, только вечный!",
	}
	return doll.New("Дракулаура", "Жопа", dDialogues)
}

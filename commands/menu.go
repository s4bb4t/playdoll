package commands

import (
	"fmt"
	"playdoll/doll"
)

func Beggining() {
	fmt.Println("Вы зашли в комнату, в которой стоят три куклы. Нужно сделать выбор - с кем Вы хотели бы поговрить в первую очередь!\nКто тебя заинтересовал больше всего? Просто назови ее имя!\n")
}

func HiDraculaura() {
	fmt.Println("-Дракулаура- (милая, эмоциональная, немного вамп-драматичная)")
}
func HiClawdeen() {
	fmt.Println("-Клодин- (дерзкая, уверенная, немного “fashion queen”)")
}
func HiFrankie() {
	fmt.Println("-Френки- (немного наивная, любопытная, «на электричестве»)\n")
}

func Greeting() {
	HiDraculaura()
	HiClawdeen()
	HiFrankie()
}

func FillIn() {
	fmt.Print("Ваш выбор: ")
}

func Draculaura(d *doll.Doll) (name string, surname string, phrases []string) {
	d.SetName("Дракулаура")
	d.SetSurname("Вамп")
	d.SetPhrase("- О, хеллооо! Я так рада тебя видеть, просто кровь кипит!")
	d.SetPhrase("- Ужасно прекрасно! Пью смузи из клубники и читаю модный журнал 🩸")
	d.SetPhrase("- Ой, я вся краснею! Любовь — это как закат, только вечный!")
	return d.Name, d.Surname, d.Phrase
}

func Clawdeen(d doll.Doll) (string, string, []string) {
	d.SetName("Клодин")
	d.SetSurname("Вульф")
	d.SetPhrase("- Эй, волчица всегда на стиле! Ты как?")
	d.SetPhrase("- О, милашка, как всегда — идеально, просто луна сияет для меня!")
	d.SetPhrase("- Только мой фирменный фиолетовый и немного когтей. Стиль — это оружие!")
	return d.Name, d.Surname, d.Phrase
}

func Frankie(d doll.Doll) (string, string, []string) {
	d.SetName("Фрэнки")
	d.SetSurname("Штейн")
	d.SetPhrase("- Здарова! Надеюсь, ток между нами пробежал ⚡")
	d.SetPhrase("- Ой, сегодня чуть не перегорела, но всё супер, я подзарядилась!")
	d.SetPhrase("- Проверяю свои винтики и думаю, как улучшить контакт с людьми 🤖")
	return d.Name, d.Surname, d.Phrase
}

func Choice(n string, s string, answer string) {
	fmt.Println("Пользователь окликнул куклу по имени:", n, "\n\nВы остановились напротив нее, немного помялись, и наконец произнесли: Привет!\n", answer, " - сказала ", n, s)
}

func Pause() {
	fmt.Println("\nОго! Кукла осказалась действительно живой... Что же сказать дальше?")
}

func Variants(s string) {
	fmt.Println("\n- kак дела? -")
	fmt.Println(s)
	fmt.Println("- пока! -\n")
}

func Bye() {
	fmt.Println("- Уже? Ладно... Пока!")
}
func DonGetIt() {
	fmt.Println("- Ой, что-то я совсем ничего не поняла, повтори-ка еще раз!")
}

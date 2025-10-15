package doll

import (
	"strings"
	"unicode"
)

type Doll struct {
	name      string
	surname   string
	dialogues map[string]string
}

const unknownPhraseAnswer = "Не знаю что скзаать"

func New(name, surname string, dialogues map[string]string) *Doll {
	return &Doll{
		name:      name,
		surname:   surname,
		dialogues: dialogues,
	}
}

func (d *Doll) Name() string {
	return d.name
}

func (d *Doll) Surname() string {
	return d.surname
}

func (d *Doll) Talk(phrase string) string {
	ans, exists := d.dialogues[clearKeyword(phrase)]
	if exists {
		return ans
	}
	return unknownPhraseAnswer
}

// clearKeyword нужна затем, чтобы убрать лишние символы из строки.
// Она позволит распознать не только "привет", но и "Привет!", "пРиВет!!!!??" и тд.
func clearKeyword(s string) string {
	var keyword string

	p := strings.ToLower(strings.TrimSpace(s)) // прочитай что делают эти функции. Выпиши и используй их в дальшейнем
	for _, r := range p {
		if unicode.IsLetter(r) { // прочитай что делает эта функция. Выпиши и используй дальшейнем
			keyword += string(r)
		}
	}

	return keyword
}

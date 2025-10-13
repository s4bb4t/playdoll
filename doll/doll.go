package doll

type Doll struct {
	Name    string
	Surname string
	Phrase  []string
}

func New() *Doll {
	return &Doll{
		Phrase: make([]string, 0),
	}
}
func (d *Doll) SetName(s string) {
	d.Name = s
}

func (d *Doll) SetSurname(s string) {
	d.Surname = s
}
func (d *Doll) SetPhrase(phrase string) {
	d.Phrase = append(d.Phrase, phrase)
}

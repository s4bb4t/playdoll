package main

var (
	ManyLight orelOrReshka
	OneBig    course
)

type orelOrReshka struct {
	Orel []bool // true - orel
}

type course struct {
	Image   []byte
	Url     string
	Desc    string
	Video1  []byte
	Video2  []byte
	Video3  []byte
	Video4  []byte
	Trailer []byte
}

func init() {
	ManyLight.Orel = make([]bool, 10_000_000)

	OneBig = course{
		Image:   make([]byte, 8<<20),
		Url:     "jopa.com/image",
		Desc:    "lorem ipsum dolor sit amet lorem ipsum dolor sit amet lorem ipsum dolor sit amet lorem ipsum dolor sit amet lorem ipsum dolor sit amet lorem ipsum dolor sit amet lorem ipsum dolor sit amet lorem ipsum dolor sit amet lorem ipsum dolor sit amet lorem ipsum dolor sit amet lorem ipsum dolor sit amet lorem ipsum dolor sit amet ",
		Video1:  make([]byte, 425<<20),
		Video2:  make([]byte, 251<<20),
		Video3:  make([]byte, 502<<20),
		Video4:  make([]byte, 610<<20),
		Trailer: make([]byte, 120<<20),
	}
}

# Exapmle

```go
1. 2+2 = ?
 Ответ: 3 
```


```go
m := map[string]string {
    "2+2=?": "4",
    "6-2=?": "4",
}

examAnswer := "4"
m2 := map[string]string {
    "4": "4"
}

examQuestion := "2+2=?"
childAnsw := "3"

if m[examQuestion] == childAnsw {
	fmt.Print('sdal')
}

if m[examAnswer] == childAnsw {
fmt.Print('sdal')
}


```


```go

type Person struct {
	Age int
	Height int
	Name string
	Surmane string
	
	Hoobies ???
}

hoobie1 := "walleybool"
hoobie2 := "football"
hoobie3 := "hiking"
hoobie4 := "farting"
hoobie5 := "origami"

```

```go

type Person struct {
    Height1 int
    Height2 int
    Height3 int
    Height4 int
    Height5 int
    Height6 int
    ...
}

type Person struct {
	AgeHeigth map[int]int
}

mapa := map[int]int

```




```go
question1 := "2+2=?"
answ1 := "4"

question2 := "4+2=?"
answ2 := "6"
...

question228 := "sosal?"
answ288 := "da"

type Exam struct {
    QuestionAnswers ???
}

```
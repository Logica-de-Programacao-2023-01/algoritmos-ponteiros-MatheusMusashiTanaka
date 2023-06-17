package main

import "fmt"

type Livro struct {
	Titulo string
	Autor  string
}

func anonimo(l *Livro) {

	if l.Autor == "Anônimo" {

		l.Titulo = "Desconhecido"

	}

}

func main() {

	l := Livro{

		Titulo: "Os sete maridos de Evelyn Hugo",
		Autor:  "Taylor Jenkins Reid",
	}

	l2 := Livro{
		Titulo: "ahdvshajdg",
		Autor:  "Anônimo",
	}

	anonimo(&l)
	anonimo(&l2)

	fmt.Println(l)
	fmt.Println(l2)

}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Servidor struct {
	Nome      string
	Profissao string
	Idade     int
}

func main() {
	arquivo, err := os.Create("teste.csv")
	if err != nil {
		log.Panic(err)
	}
	defer arquivo.Close()

	servidores := []Servidor{{
		Nome:      "Vinícius",
		Profissao: "Desenvolvedor Back-End",
		Idade:     19,
	}, {
		Nome:      "Levid",
		Profissao: "Desenvolvedor Front-End",
		Idade:     19,
	},
	}
	writer := bufio.NewWriter(arquivo)
	writer.Flush()

	for _, servidor := range servidores {
		servidorSlice := []string{servidor.Nome, servidor.Profissao, fmt.Sprintf("%d", servidor.Idade)}
		servidorBytes := convert(servidorSlice)
		writer.Write(servidorBytes)
	}
}
func convert(arr []string) []byte {
	var bytess []byte
	for i := 0; i < len(arr); i++ {
		b := []byte(arr[i])
		for j := 0; j < len(b); j++ {
			bytess = append(bytess, b[j])
		}
	}
	return bytess
}

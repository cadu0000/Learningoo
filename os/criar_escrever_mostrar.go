package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	novoArquivo, err := os.Create("arquivo_teste.txt") //cria o arquivo
	if err != nil {
		fmt.Println(err)
		return
	}
	defer novoArquivo.Close() // fecha o arquivo após a execução de todo o programa
    
	novoTexto := strings.NewReader("Enviando texto para o arquivooo") //envia para o arquivo
	io.Copy(novoArquivo, novoTexto)

	novoArquivo, err = os.Open("arquivo_teste.txt") //abre o arquivo
	if err != nil {
		fmt.Println(err)
		return
	}

	conteudoArquivo, err := ioutil.ReadAll(novoArquivo) // lê o conteúdo e armazena em uma variável em formato sb
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(conteudoArquivo))
}

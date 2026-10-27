package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const DefaultTimeout = 1 * time.Second

func main() {
	args := os.Args[1:]

	var cep string

	if len(args) > 0 {
		cep = args[0]
	} else {
		cep = "01310100"
	}

	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	defer cancel()

	chBrasilApi := make(chan []byte)
	chViaCep := make(chan []byte)

	urlBrasilApi := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)
	urlViaCep := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)

	go buscarCep(ctx, chBrasilApi, urlBrasilApi)
	go buscarCep(ctx, chViaCep, urlViaCep)

	for {
		select {
		case body := <-chBrasilApi:
			fmt.Println("BrasilApi teve a resposta mais rápida. Abaixo os dados do CEP:")
			fmt.Println(string(body))
			return
		case body := <-chViaCep:
			fmt.Println("ViaCep teve a resposta mais rápida. Abaixo os dados do CEP:")
			fmt.Println(string(body))
			return
		case <-time.After(DefaultTimeout):
			log.Println("Tempo limite atingido ao buscar os dados do CEP")
			return
		}
	}
}

func buscarCep(ctx context.Context, chOutput chan<- []byte, url string) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	chOutput <- body
}

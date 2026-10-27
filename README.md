# PGE Multithreading - Consulta de CEP

Projeto implementado com base nos requisitos da atividade "Multithreading" no curso Pós Go Expert da FullCycle.

O objetivo é realizar duas requisições simultâneas para APIs diferentes de consulta de CEP ([BrasilAPI](https://brasilapi.com.br/) e [ViaCEP](https://viacep.com.br/)) e exibir o resultado da API que responder primeiro

## Como Executar

É possível executar o programa informando um CEP como argumento ou utilizando o CEP padrão definido no código.

### 1) Com CEP padrão

Caso nenhum argumento seja passado, será utilizado o CEP `01310100` (Avenida Paulista - SP).

```bash
go run main.go
```

### 2) Informando um CEP

```bash
go run main.go 01001000
```

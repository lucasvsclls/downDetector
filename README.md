# downDetector

Um projeto desenvolvido em **Go** para praticar programação concorrente e comunicação entre goroutines utilizando channels.

## Sobre

O `downDetector` verifica a disponibilidade de múltiplos sites simultaneamente e informa no terminal se cada endereço está acessível ou pode estar indisponível.

O projeto utiliza **goroutines** para realizar as verificações de forma concorrente, permitindo que várias requisições HTTP sejam executadas simultaneamente.

Após cada verificação, o endereço é enviado através de um **channel** e uma nova verificação é agendada após alguns segundos.

## Conceitos explorados

- Goroutines
- Channels
- Concorrência em Go
- Funções anônimas e closures
- Comunicação entre goroutines
- Requisições HTTP com `net/http`
- Range sobre channels
- `time.Sleep`
- Tratamento de erros

## Como funciona

O programa possui uma lista de URLs que precisam ser verificadas.

Para cada endereço, uma goroutine é criada:

```go
go checkLink(link, c)
```

A função `checkLink` realiza uma requisição HTTP utilizando `http.Get` e envia o endereço para o channel após a verificação.

O programa então recebe os endereços através do channel e agenda uma nova verificação:

```go
for l := range c {
    go func(link string) {
        time.Sleep(4 * time.Second)
        checkLink(link, c)
    }(l)
}
```

Dessa forma, o programa continua monitorando os endereços e realizando novas verificações periodicamente.

## Tecnologias

- Go
- Go Standard Library
- `net/http`
- `time`
- Git

## Como executar

Clone o repositório:

```bash
git clone https://github.com/lucasvsclls/downDetector.git
```

Acesse o diretório:

```bash
cd downDetector
```

Execute o projeto:

```bash
go run .
```

## Exemplo

A saída no terminal será semelhante a:

```text
http://www.google.com is available
http://www.golang.org is available
http://www.amazon.com is available
http://www.microsoft.com is available
http://www.facebook.com is available
```

Caso uma URL não esteja acessível:

```text
http://example.com might be down
```

## Objetivo

O objetivo do projeto é praticar os principais mecanismos de concorrência da linguagem Go, especialmente **goroutines e channels**, utilizando um problema simples de monitoramento de disponibilidade de sites.

O projeto faz parte dos meus estudos de Go e desenvolvimento backend.

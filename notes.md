1. Instale o GO no seu sistema operacional.
2. Instale a extensão oficial do GO (golang.go)
3. CTRL+SHIFT+P "GO" e rode o comando "GO: Install/Update Tools"
4. Selecione todas as alternativas e instale
5. Agora rode o comando no seu diretorio: "go mod init github.com/devfullcycle/imersao22/go-gateway"

Isso gera um arquivo go.mod

6. Crie um arquivo main.go na pasta cmd/app

```go
package main

func main() {
	println("Hello, world!")
}
```

7. Rode com o comando `go run cmd/app/main.go`
8. Faça o build com o comando `go build cmd/app/main.go`

Este comando vai gerar um executavel `main.exe` para todas as plataformas.
Uma convenção mais ou menos comum é criar uma pasta internal/ com toda a aplicação, mas nem todo mundo segue.

9. Crie uma pasta /internal/domain

Nela vai ter a entidade `account.go`

GO não é orientado a objetos, é orientado a dados:

```go
type Account struct {
	ID string
	Name string
	Email string
	ApiKey string
	Balance float64
	CreatedAt time.Time
	UpdatedAt time.Time
}
```

A função construtora retorna um ponteiro de `*Account` (o valor da struct fica apontado na memória), na prática, se você mudar a struct todos os lugares que usam ela `&Account´ também vai mudar também.

Agora preciso mudar o Balance de acordo com as transações, a função `AddBalance` vai estar atribuida ao struct Account, portanto ela vai servir como um método e a struct como uma classe.

Você pode fazer import de um lib direto do github "github.com/google/uuid" por exemplo e instalar com o comando `go mod tidy`

**CONCORRÊNCIA (RACE CONDITION)**

Se varias transações ocorrerem e alterarem o Balanace simutaneamente podemos garantir que não ocorra erro nos calculos (Race Condition) adicionando ao struct uma propriedade `mu`:

```go
type Account struct {
	mu 				sync.RWMutex // Read/Whuite Mutex
}

func (a *Account) AddBalance(amount float64) {
	a.mu.Lock()
	defer a.mu.Unlock()

	// Ninguém consegue realizar a mesma operação até esta ser finalizada
}
```

Para acessar o banco de dados, na `domain/repository.go`

`
$ go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
$ migrate -database "postgresql://postgres:postgres@localhost:5432/fullcycle-gateway?sslmode=disable" -path migrations up
`


RUN 

$ go mod tidy
$ go run cmd/app/main.go
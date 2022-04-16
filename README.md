
### Tech Stack
- Golang
- Docker
- Mysql

### O que é ?
Uma aplicação de uma simples rede social chamada devbook. \
Essa aplicação foi desenvolvida através de um curso que fiz na Udemy: [Aprenda Golang do Zero! Desenvolva uma APLICAÇÃO COMPLETA!](https://www.udemy.com/course/aprenda-golang-do-zero-desenvolva-uma-aplicacao-completa/) 

### Como rodar *sem* docker ?
- Instale o [go](https://go.dev/)
- Instale o [mysql](https://www.mysql.com/downloads/)
- Criar a base dados com as tabelas, conforme arquivo ./api/sql/ddl.sql
- Dentro da pasta .api/ execute o comando `go run main.go`
- Dentro da pasta .webapp/ execute o comando `go run main.go`
- Acesse: > **login** **`http://localhost:3000`**
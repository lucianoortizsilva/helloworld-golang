### Tecnologias
- <img src="https://badges.aleen42.com/src/golang.svg" alt="badge"/> 
- <img src="https://badges.aleen42.com/src/docker.svg" alt="badge"/> 
- [MySQL](https://www.mysql.com/downloads/)

### O que é ?
Uma aplicação de uma simples rede social chamada devbook. \
Essa aplicação foi desenvolvida através de um curso que fiz na Udemy: [Aprenda Golang do Zero! Desenvolva uma APLICAÇÃO COMPLETA!](https://www.udemy.com/course/aprenda-golang-do-zero-desenvolva-uma-aplicacao-completa/) 

### Como rodar sem docker ?
- Instale o [Go](https://go.dev/)
- Instale o [MySQL](https://www.mysql.com/downloads/)
- Criar a base dados com as tabelas, conforme arquivo ./api/sql/ddl.sql
- Dentro da pasta .api/ execute o comando `go run main.go`
- Dentro da pasta .webapp/ execute o comando `go run main.go`
- Acesse: > **login** **`http://localhost:3000`**

### Como rodar com docker ?
- Instale o [Docker](https://www.docker.com/)
- Dentro do diretório raiz execute: `docker-compose up`
- Após subir os 3 containers, crie a base de dados acessando o container `mysql-db`:
    - docker container ls (Copie o container_id do mysql-db)
    - docker exec -it `Colar o container_id` bash
    - mysql -u `Informe usuário` -p `Informe senha`
    - Execute os scripts da pasta ./api/sql/ddl.sql
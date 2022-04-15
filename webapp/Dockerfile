FROM golang:1.18.1-alpine3.14

LABEL authors="lucianoortizsilva@gmail.com"

#Criando um diretório chamado 'app'
#Isso instrui o Docker a usar esse diretório como destino padrão para todos os comandos subsequentes. 
#Dessa forma, não precisamos digitar caminhos de arquivo completos, mas podemos usar caminhos relativos com base nesse diretório.
WORKDIR /app

#O comando COPY  recebe dois parâmetros. 
#O primeiro parâmetro informa ao Docker quais arquivos você deseja copiar para a imagem. 
#O último parâmetro informa ao Docker para onde você deseja que esse arquivo seja copiado.
COPY go.mod /app
COPY go.sum /app

COPY /assets /app/assets
COPY /src /app/src
COPY /views /app/views

COPY .env /app
COPY main.go /app

#Para instalar os módulos Go em um diretório dentro da imagem
RUN go mod download

#Criando um arquivo binário com nome de 'dev-book-api'
RUN go build -o /dev-book-webapp

#Porta
EXPOSE 3000

#Informando qual arquivo executar ao iniciar im container
CMD ["/dev-book-webapp"]

# Comando para criar imagem:
# docker build --tag dev-book-webapp:1.0 .
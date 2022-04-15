### Instruções para usar o mysql via docker ###

Dentro da pasta api, crie os diretórios:
.docker-db-mysql

Após executar o docker-compose up, faça:

> docker container ls [anote o container_id do mysql]

> docker container inspect [informe o container_id do mysql], anote o "IPADDRESS"

> docker exec -it [container_id do mysql] bash

#TENTAR ESSE:
> winpty docker exec -it [container_id] bash

#SENÃO ESSE:
> docker exec -it [container_id] bash

> mysql -u ortiz -p admin

Execute os scripts da pasta /api/sql*
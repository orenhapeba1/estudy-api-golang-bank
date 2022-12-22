# API RESTFull BANK - GOLANG

Teste de desenvolvendo API RESTFull utilizando Framework Gin na linguagem GO.


## Tecnologias Utilizadas
- Golang 1.18
- Gin framework
- Jwt
- Docker
- MySQL

## PRÉ-REQUISITOS


```bash

Docker Versão 3+ instalada na máquina

```

## Instalação

Clone o repositório para sua máquina

```bash

https://github.com/orenhapeba1/estudy-api-golang-bank.git

```

depois de clonado suba o container 

```bash

docker-compose up -d

```

precisamos alterar as permissões do MySQL para poder subir a base de dados
```bash

docker exec -it mysql mysql -uroot -pdocker

```
depois de efetuar login no mysql executar o comando abaixo para alterar a permissao de acesso do root
```bash

RENAME USER `root`@`localhost` TO `root`@`%`;
USE db;
source /root/db/db.sql;

```


## Execução
Depois de configurado, basta executar o comando 
```bash

docker-compose up

```

## CONSULTA
Para realizar consultas, cadastros edições e deletes na API aconselho utilizar o postman, arquivo de configuração de todas as rotas esta incluso basta importar e utilizar (estudy-api-golang-bank.postman_collection.json)

- localhost:5000/api/v1/login (POST)
   - (para efetuar login e obter o token)
- localhost:5000/api/v1/accounts/ (GET)
  - (necessita de configuração do token na header Ex: Authorization: Bearer + Token Recebido)
- localhost:5000/api/v1/accounts/ (POST) 
  - cria nova conta (não é necessario token)
- localhost:5000/api/v1/accounts/ (PUT) 
  - edita a propria conta (necessita de configuração do token na header Ex: Authorization: Bearer + Token Recebido)
- localhost:5000/api/v1/accounts/{account_number} (PUT)
    - edita a conta informada na url (necessita de configuração do token na header Ex: Authorization: Bearer + Token Recebido)
- localhost:5000/api/v1/accounts/{account_number} (DELETE)
    - deleta a conta informada na url junto com saldo e extrato (necessita de configuração do token na header Ex: Authorization: Bearer + Token Recebido) (não é possivel deletar a propria conta)
- localhost:5000/api/v1/balance/ (GET)
    - Verifica dados completos da própria conta juntamente com extrato (necessita de configuração do token na header Ex: Authorization: Bearer + Token Recebido)
- localhost:5000/api/v1/balance/{account_number} (GET)
    - Verifica dados completos da conta informada na url juntamente com extrato (necessita de configuração do token na header Ex: Authorization: Bearer + Token Recebido)
- localhost:5000/api/v1/balance/ (POST)
    - Movimenta a propria conta fazendo entradas e saidas (necessita de configuração do token na header Ex: Authorization: Bearer + Token Recebido)
- localhost:5000/api/v1/balance/{account_number} (POST)
    - Movimenta a conta informada fazendo entradas e saidas (necessita de configuração do token na header Ex: Authorization: Bearer + Token Recebido)


## Contributing

Solicitações pull são bem-vindas. Para mudanças importantes, abra um problema primeiro
para discutir o que você gostaria de mudar.

Certifique-se de atualizar os testes conforme apropriado.

## License

[MIT](https://choosealicense.com/licenses/mit/)

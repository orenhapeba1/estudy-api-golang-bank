# API RESTFull BANK - GOLANG

Teste de desenvolvendo API RESTFull utilizando Framework Gin na linguagem GO.


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

docker exec -it estudy-api-golang-bank-mysql-1 mysql -uroot -pdocker

```
depois de efetuar login no mysql executar o comando abaixo para alterar a permissao de acesso do root
```bash

RENAME USER `root`@`localhost` TO `root`@`%`;
USE db;
source /root/db/db.sql;

```



## Contributing

Solicitações pull são bem-vindas. Para mudanças importantes, abra um problema primeiro
para discutir o que você gostaria de mudar.

Certifique-se de atualizar os testes conforme apropriado.

## License

[MIT](https://choosealicense.com/licenses/mit/)

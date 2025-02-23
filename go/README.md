# Imersão Fullcycle 21 - Home Broker

## Descrição

Repositório do Golang (simulador de bolsa de valores).

## Requerimentos

Ter Docker instalado e configurar o arquivo de hosts.

## Configurar /etc/hosts

O RabbitMQ está sendo executado no `docker-compose.yaml` da aplicação Golang, assim como o Django está em outro `docker-compose.yaml`, os containers estão em redes diferentes.
Usaremos a estratégia do `host.docker.internal` para comunicação entre os containers.

Para isto é necessário configurar um endereços que todos os containers Docker consigam acessar.

Acrescente no seu /etc/hosts (para Windows o caminho é C:\Windows\system32\drivers\etc\hosts):

```
127.0.0.1 host.docker.internal
```

Em todos os sistemas operacionais é necessário abrir o programa para editar o *hosts* como Administrator da máquina ou root.

Obs.: Se estiver usando o Docker Desktop, pode ser que o `host.docker.internal` já esteja configurado, então remova a linha do arquivo hosts e acrescente a recomendada acima.


## Rodar a aplicação

Esta configuração permite rodar todas as aplicações de uma vez só, por isso não é necessário rodar o comando `docker-compose up` em cada diretório.

Da raiz do projeto rode o comando:

```bash
docker-compose exec golang sh
go run cmd/trade/main.go
```

Isto irá rodar o simulador de bolsa de valores que já estará conectado com o Kafka.

Se quiser rodar apenas o simulador de bolsa de valores, rode o comando na pasta **go**:

```bash
docker-compose up
go run cmd/trade/main.go
```


# Imersão Fullcycle 21 - Home Broker

## Descrição

Repositório do Nest.js (back-end do home broker)

## Requerimentos

Ter Node.js e Docker instalados.

## Rodar a aplicação

Instale as dependências:

```bash
npm install
```

Rode o comando para levantar o servidor http:

```bash
npm run start:dev
```

Para rodar o comando que consome mensagens do Kafka:

```bash
npm run start:dev -- --entryFile _cmd/kafka/kafka.cmd
```

As especificações das chamadas HTTP estão no arquivo `api.http`. Você pode usar a extensão `REST Client` do VSCode para fazer as chamadas.

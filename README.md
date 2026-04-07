# Clean Architecture: Listagem de Orders (REST, gRPC e GraphQL)

## Execução

Comando único de execução:

```bash
docker compose up -d --build
```

Alternativa, caso seu Docker exija privilégios:

```bash
sudo make up
```

## Portas

- Web REST: `8080`
- GraphQL: `8081`
- gRPC: `50051`

Projeto em Go seguindo os princípios de Clean Architecture para cadastro e listagem de orders por três interfaces simultâneas:

- REST
- GraphQL
- gRPC

O caso de uso de listagem é único e reutilizado pelas três entradas da aplicação.

## Tecnologias

- Go
- Clean Architecture
- REST
- GraphQL
- gRPC
- SQLite
- Docker
- Docker Compose

Ao iniciar:

- a imagem da aplicação é construída
- o ambiente é carregado a partir do `.env.local`
- o banco SQLite é persistido em volume Docker
- as migrations são aplicadas automaticamente na inicialização
- a aplicação sobe com os serviços REST, GraphQL e gRPC

## Interfaces disponíveis

### REST

- Criar order: `POST /orders`
- Listar orders: `GET /orders`

Arquivo pronto para teste: [api.http](api.http)

### GraphQL

Endpoint: `http://localhost:8081/query`

Playground: `http://localhost:8081`

Query para listar orders:

```graphql
query ListOrders {
    ListOrders {
        id
        productId
        quantity
    }
}
```

Mutation para criar order:

```graphql
mutation CreateOrder {
    CreateOrder(input: { productId: 1, quantity: 2 }) {
        id
        productId
        quantity
    }
}
```

### gRPC

Service: `order.OrderService`

Métodos:

- `CreateOrder`
- `ListOrders`

Exemplo com `grpcurl` para listar:

```bash
grpcurl -plaintext localhost:50051 order.OrderService/ListOrders
```

Exemplo com `grpcurl` para criar:

```bash
grpcurl -plaintext \
  -d '{"product_id":1,"quantity":2}' \
  localhost:50051 \
  order.OrderService/CreateOrder
```

## Estrutura da solução

- `internal/domain`: entidades, contratos e casos de uso
- `internal/servers/rest`: entrega HTTP REST
- `internal/servers/graphql`: entrega GraphQL
- `internal/servers/grpc`: entrega gRPC
- `internal/infra/db/migrations`: migrations do banco
- `proto`: definição protobuf e código gerado do gRPC

## Banco de dados e migrations

O projeto utiliza SQLite persistido em volume Docker.

As migrations ficam em:

- `internal/infra/db/migrations`

Na subida do container, a aplicação executa automaticamente as migrations antes de atender as requisições.

## Arquivos auxiliares

- `.env.local`: configuração padrão das portas
- `api.http`: requisições prontas para criar e listar orders via REST
- `docker-compose.yml`: orquestração da aplicação
- `Dockerfile`: build e empacotamento da aplicação

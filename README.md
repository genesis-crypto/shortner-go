# __Encurtador de URL__
Este projeto consiste em uma aplicação que encurta urls, e que permite usuários realizarem a gestão dos links em questão. O projeto foi desenvolvido como parte da disciplina de Desenvolvimento Avançado de Sistemas.

## __Requisitos Funcionais__
* Permitir a gestão de usuários
* Permitir a gestão de links
* Permitir ser redirecionado acessando o link encurtado

## __Topologia de Arquivos__

O projeto está estruturado da seguinte forma:

```
shortner-go/
  ├── cmd/
  │   └── shortner-go/
  │         ├── .env
  │         └── main.go
  ├── configs/
  │   ├── grafana/
  │   │     └── config.monitoring
  │   ├── prometheus/
  │   │     └── prometheus.yml
  │   └── config.go
  ├── internal/
  │   ├── dto/
  │   │    └── dto.go
  │   ├── entities/
  │   │     ├── link_entity.go
  │   │     └── user_entity.go
  │   ├── handlers/
  │   │     ├── link_handler.go
  │   │     └── user_handler.go
  │   └── infra/
  │         ├── cache/
  │         │     └── redis.go
  │         └── database/
  │               ├── interface.go
  │               ├── link_db.go
  │               └── user_db.go
  ├── pkg/
  │   └── shortner/
  │         └── shortner.go
  ├── docker-compose.yml
  ├── Dockerfile
  ├── go.mod
  ├── go.sum
  ├── nginx.conf
  ├── nginx.dockerfile
  ├── prometheus.yml
  ├── README.md
  └── script.sh
```

## __Detalhes de implementação__

### __Tecnologias Utilizadas__
- Go 1.19.5
- MySQL 5.7
- RabbitMQ 3
- Prometheus
- Grafana
- Nginx
- Redis 6.2

### __Pré-Requisitos__
- Docker
- Docker Compose

### __Instruções de Uso:__

### __1. Clonando o Repositório__
```sh
git clone https://github.com/seu-usuario/shortner-go.git
```
### __1. Configurando o Ambiente__
Crie um arquivo .env na pasta cmd/shortner-go/ com as seguintes variáveis:
```sh
DB_DRIVER=mysql
DB_HOST=host.docker.internal
DB_PORT=3306
DB_USER=root
DB_PASSWORD=root
DB_NAME=shortner-go
WEB_SERVER_PORT=8080
JWT_SECRET=secret
JWT_EXPIRESIN=300
REDIS_HOST=cache
REDIS_PORT=6379
REDIS_PASSWORD=eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81
QUEUE_HOST=amqp
QUEUE_USER=guest
QUEUE_PASSWORD=guest
QUEUE_PORT=5672
QUEUE_NAME=message-broker
```

### __3. Iniciando o Ambiente__
Execute o comando abaixo para iniciar o ambiente:

```sh
cd shortner-go/
docker-compose up -d
```

### 4. Acessando o Painel do Grafana
Acesse o painel do Grafana em http://localhost:3000 com as seguintes credenciais:
```sh
username: admin
password: admin
```

### Configurar Datasource Prometheus
```http://host.docker.internal:9090```

## __Contato__
Pedro Cardozo - `p-cardozo@hotmail.com` ou `609455@univem.edu.br`
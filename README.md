# Tutorial de Configuração da API Go Gin

Este guia detalha os passos para configurar e rodar sua API Go Gin com um banco de dados PostgreSQL.

## Pré-requisitos

Antes de começar, certifique-se de ter o seguinte instalado:

* **Go**: Versão 1.18 ou superior.

* **Docker** e **Docker Compose**: Para rodar o banco de dados e a API em contêineres.

## 1. Configuração do Arquivo `.env`

O arquivo `.env` contém variáveis de ambiente essenciais para a conexão com o banco de dados e outras configurações da API.

### 1.1. Gerar o arquivo `.env`

Navegue até a pasta `api_go` no seu terminal e execute o seguinte comando para criar o arquivo `.env` a partir do `exemple.env`:

```

cp api\_go/exemple.env api\_go/.env

```

### 1.2. Conteúdo do arquivo `.env`

Edite o arquivo `api_go/.env` e insira o conteúdo abaixo. Escolha a seção apropriada dependendo se você vai rodar a API com Docker Compose ou localmente.


# \--- Configurações para rodar com Docker Compose (RECOMENDADO) ---

# Descomente esta parte e comente a seção "Para rodar localmente"

SECRET\_KEY=ad1m4e3z1S42fDb4RPbzfd1xCb7d37Dj6af06eDe047e86
DB\_HOST=db \# 'db' é o nome do serviço do banco no Docker Compose
DB\_PORT=5432
DB\_USER=user
DB\_PASS=user
DB\_NAME=api\_go
DB\_SSLMODE=disable

# \--- Configurações para rodar localmente (sem Docker Compose para o banco) ---

# Descomente esta parte e comente a seção "Para rodar com Docker Compose"

# SECRET\_KEY=ad1m4e3z1S42fDb4RPbzfd1xCb7d37Dj6af06eDe047e86

DB\_HOST=localhost \# 'localhost' se o banco estiver rodando diretamente na sua máquina
DB\_PORT=5432
DB\_USER=user
DB\_PASS=user
DB\_NAME=api\_go
DB\_SSLMODE=disable

**Importante**: Mantenha apenas uma das seções (Docker Compose ou Local) descomentada no seu arquivo `.env` para evitar conflitos.

## 2. Rodar a Aplicação com Docker Compose (Recomendado)

Esta é a maneira mais fácil de subir a API e o banco de dados juntos, garantindo um ambiente consistente.

Navegue até a pasta raiz do seu projeto (onde o arquivo `docker-compose.yml` está localizado) e execute o comando único:

```
docker compose up --build

```

**O que este comando faz:**

* `docker compose up`: Inicia os serviços definidos no `docker-compose.yml`.

* `--build`: Reconstrói as imagens dos serviços, garantindo que você esteja usando a versão mais recente do seu código.

Este comando irá:

* Construir a imagem da sua API Go.

* Iniciar um contêiner PostgreSQL.

* Criar uma rede interna para que a API e o banco possam se comunicar.

* Criar um volume para persistir os dados do banco de dados.

## 3. Rodar a Aplicação Localmente (Sem Docker Compose)

Se você preferir rodar o banco de dados e a API separadamente, siga estes passos.

### 3.1. Criar e Iniciar o Banco de Dados Localmente (se ainda não tiver)

Você pode usar o Docker para rodar apenas o contêiner do PostgreSQL, mantendo-o separado da sua API Go.

Primeiro, crie um volume Docker para persistir os dados do seu banco:

```

docker volume create api\_go\_data

```

Em seguida, inicie o contêiner PostgreSQL:

```
docker run -d  
\--name api\_go\_db  
\-e POSTGRES\_USER=user  
\-e POSTGRES\_PASSWORD=user  
\-e POSTGRES\_DB=api\_go  
\-v api\_go\_data:/var/lib/postgresql/data  
\-p 5432:5432  
postgres:15
```
### 3.2. Entrar na Pasta da API e Rodar

Certifique-se de que o banco de dados PostgreSQL esteja rodando (seja via Docker ou uma instalação local direta) e que seu arquivo `api_go/.env` esteja configurado para `DB_HOST=localhost`.

Navegue até a pasta `api_go` e execute sua aplicação Go:

```
cd api\_go
go run main.go
```
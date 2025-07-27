# go-gin-api-example

✅ 1. Criar o .env a partir de exemple.env
🧾 Comando para gerar .env:
bash
Copiar
Editar
cp api_go/exemple.env api_go/.env
🔁 Você pode manter exemple.env como base e o .env real será lido pela sua API.

✅ 2. Conteúdo do arquivo .env
Coloque isso dentro de api_go/.env

env
Copiar
Editar
# Para rodar localmente (com banco local)
# Descomente esta parte e comente a de baixo para rodar fora do Docker

#SECRET_KEY=ad1m4e3z1S42fDb4RPbzfd1xCb7d37Dj6af06eDe047e86
#DB_HOST=localhost
#DB_PORT=5432
#DB_USER=user
#DB_PASS=user
#DB_NAME=api_go
#DB_SSLMODE=disable

# Para rodar com Docker Compose
SECRET_KEY=ad1m4e3z1S42fDb4RPbzfd1xCb7d37Dj6af06eDe047e86
DB_HOST=db
DB_PORT=5432
DB_USER=user
DB_PASS=user
DB_NAME=api_go
DB_SSLMODE=disable
✅ 3. Rodar com Docker Compose
🔨 Comando único:
bash
Copiar
Editar
docker compose up --build
Isso sobe a API + banco Postgres + rede + volume, usando as variáveis da api_go/.env.

✅ 4. Rodar localmente no seu PC (sem Docker)
💾 Passo 1: Criar o banco localmente (se ainda não tiver)
bash
Copiar
Editar
docker volume create api_go_data

docker run -d \
  --name api_go \
  -e POSTGRES_USER=user \
  -e POSTGRES_PASSWORD=user \
  -e POSTGRES_DB=api_go \
  -v api_go_data:/var/lib/postgresql/data \
  -p 5432:5432 \
  postgres:15
💻 Passo 2: Entrar na pasta da API e rodar
bash
Copiar
Editar
cd api_go
go run main.go
🧠 Dicas extras
Para se conectar ao banco via terminal:

bash
Copiar
Editar
psql -h localhost -U user -d api_go
Se estiver no VS Code, você pode criar uma Run Configuration que usa esse .env.
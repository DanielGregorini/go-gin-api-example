# go-gin-api-example

# 1 Cria o volume
docker volume create api_go_data

# 2 Sobe o container
docker run -d \
  --name api_go \
  -e POSTGRES_USER=user \
  -e POSTGRES_PASSWORD=user \
  -e POSTGRES_DB=api_go \
  -v api_go_data:/var/lib/postgresql/data \
  -p 5432:5432 \
  postgres:15


psql -h localhost -U user -d api_go
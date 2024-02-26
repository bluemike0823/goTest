docker run --name postgres-container -d -p 5432:5432 -v ~/Postgres:/var/lib/postgresql/data -e POSTGRES_DB=postgres -e POSTGRES_USER=admin -e POSTGRES_PASSWORD='aadd2255' postgres:latest

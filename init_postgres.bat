@REM docker run --name postgres-container -v ~/Postgres:/var/lib/postgresql/data -e POSTGRES_DB=postgres -e POSTGRES_USER=dbuser -e POSTGRES_PASSWORD=mypassword -p 5432:5432 -d postgres:latest
docker run --name postgres-container -e POSTGRES_PASSWORD=aadd2255 -e POSTGRES_USER=admin -e POSTGRES_DB=gotest -p 5432:5432 -v /data:/var/lib/postgresql/data -d postgres

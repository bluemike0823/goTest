docker build -t gotest-image .

@REM docker login localhost:8080


docker network rm gotest_default
docker compose up -d
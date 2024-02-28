docker stop postgres-container
docker network rm gotest-network
docker network create gotest-network
docker network connect gotest-network postgres-container
docker network connect gotest-network gotest-container
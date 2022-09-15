1. Install Docker on your host

https://docs.docker.com/get-docker/

2. Go to broker directory. Then, run emqx broker container
cd broker/
docker-compose up -d

3. Go to pubsub directory. Then install mqtt go module
cd pubsub/
go mod tidy

4. Run subscriber code in pubsub directory
go run simplepubsub.go sub

4. Run publisher code in pubsub directory
go run simplepubsub.go pub
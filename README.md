git clone git@github.com:sofianhw/go-jetstream-order.git

docker build . -t gotainer

docker run --name goexec -v $(pwd):/app -d -i -t gotainer /bin/sh

docker exec -it goexec go run monitor/main.go

docker exec -it goexec go run push-order/main.go

docker exec -it goexec go run update-status/main.go


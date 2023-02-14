git clone git@github.com:sofianhw/go-jetstream-order.git

cd go-jetstream-order

docker build . -t gotainer

docker run --name goexec -v $(pwd):/app -d -i -t gotainer /bin/sh

docker exec -it goexec go run monitor/main.go

docker exec -it goexec go run push-order/main.go

docker exec -it goexec go run update-status/main.go


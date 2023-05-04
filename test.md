docker run -d --hostname my-rabbit --name some-rabbit -e RABBITMQ_DEFAULT_USER=guest -e RABBITMQ_DEFAULT_PASS=guest rabbitmq:3-management

docker build -t chatserver -f Dockerfile .
docker run -p 8080:8080 chatserver

docker build -t chatboot -f chatboot.Dockerfile .
docker run chatboot

docker build -t processor -f processor.Dockerfile .
docker run processor

sudo systemctl restart docker.socket docker.service

sudo -u rabbitmq rabbitmqctl stop

sudo chmod 666 /var/run/docker.sock

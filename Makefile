NAME=hello-world-server
TAGSERVER=raushan2016/helloworld-server
TAGCONSOLE=raushan2016/helloworld-console
VER=v1.0

all: clean build

build:
	go get github.com/gorilla/mux
	CGO_ENABLED=0 GOOS=linux go build -ldflags "-s" -a -installsuffix cgo -o hello-world
	docker build -t $(TAGCONSOLE) -f console-Dockerfile .
	docker build -t $(TAGSERVER) -f httpServer-Dockerfile .

run:
	docker run -d -p 80:80 -e PORT=80 --name=$(NAME) $(TAGSERVER)
	docker run -ti --rm --link $(NAME):$(NAME) qorbani/curl

clean:
	-docker stop $(NAME)
	-docker rm $(NAME)

push:
	docker push $(TAGCONSOLE)
	docker push $(TAGSERVER)


FROM golang:1.21

WORKDIR /web

COPY . .

RUN go mod download && go mod tidy

RUN go build -o bin ./main.go

CMD ["./bin"]

EXPOSE 8080

# 만약 ENTRYPOINT 를 사용하여 컨테이너 수행 명령을 정의한 경우,
# 해당 컨테이너가 수행될 때 반드시 ENTRYPOINT 에서 지정한 명령을 수행되도록 지정 된다.

#  하지만, CMD를 사용하여 수행 명령을 경우에는,
#  컨테이너를 실행할때 인자값을 주게 되면 Dockerfile 에 지정된 CMD 값을 대신하여 
#지정한 인자값으로 변경하여 실행되게 된다.
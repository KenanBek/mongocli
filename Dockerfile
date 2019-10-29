FROM golang:latest
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go test -v ./...
RUN go build -o mongocli mongocli.go
RUN touch ~/mongocli.yml
RUN echo -e "Server: db\nPort: 27017" > ~/mongocli.yml
RUN touch mongocli.yml
ENV PATH="/app:${PATH}"

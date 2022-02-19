FROM golang:1.16
WORKDIR /backend
COPY . .
RUN  go mod download
EXPOSE 4000
CMD ["go","run","main.go"]
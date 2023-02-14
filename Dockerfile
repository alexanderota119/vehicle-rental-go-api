FROM golang:1.20-alpine

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -v -o /app/vehicle-rental

EXPOSE 3000
ENTRYPOINT [ "/app/vehicle-rental"]
CMD [ "serve" ]
FROM golang:1.19-bullseye AS build

WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o server main.go

FROM alpine:3.16 AS runtime

COPY --from=build /app/server .

CMD [ "./server" ]

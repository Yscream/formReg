FROM golang:1.17.5 as build
COPY . /app/
WORKDIR /app/cmd
RUN CGO_ENABLED=0 GOOS=linux go build -o /registration main.go

FROM alpine:3.15
WORKDIR /app
COPY --from=build /registration /bin/registration
COPY --from=build /app/assets /assets
ENTRYPOINT [ "/bin/registration" ]

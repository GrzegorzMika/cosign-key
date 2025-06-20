FROM golang:1.24.4 AS development

WORKDIR /cosign-key

COPY go.mod ./
RUN go mod download && go mod verify

COPY main.go ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./build/cosign-key main.go
RUN chmod a+x /cosign-key

FROM scratch AS app

EXPOSE 8080

USER 1234

COPY --from=development /cosign-key/build/cosign-key /cosign-key

CMD [ "/cosign-key" ]
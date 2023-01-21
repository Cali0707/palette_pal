FROM golang:1.19.2-alpine as builder
COPY . go/src/palette_pal
WORKDIR go/src/palette_pal
EXPOSE 3000
ENTRYPOINT ["go", "run", "main.go"]

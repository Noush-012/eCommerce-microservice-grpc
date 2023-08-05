# Build stage

FROM golang:1.20.6-alpine3.18 AS build-stage
LABEL maintainer="Noushad Ibrahim <noushadibrahim012@gmail.com>"
WORKDIR /home/app
COPY ./ /home/app/
RUN apk update && apk add --no-cache git
RUN cd /home/app && go mod tidy
RUN go build -o ./api ./cmd

# Run Stage

FROM gcr.io/distroless/static-debian11
COPY --from=build-stage /home/app/api /api
COPY .env /.env
CMD ["/api"]



# Build stage

FROM golang:1.20.6-alpine3.18 AS build-stage
WORKDIR /home/app
COPY ./ /home/app/
RUN cd /home/app && go mod tidy
RUN go build -0 ./api ./cmd

# Run Stage

FROM rcr.io/distroless/static-debian11
COPY --from=build-stage /home/app/api /api
CMD ["/api"]



FROM golang:1.22.1-alpine3.19 as build

WORKDIR /src

COPY . .

RUN apk add --no-cache make

RUN make

FROM alpine:3.19.1

RUN apk --no-cache add curl

ENV PORT=8080
ENV GIN_MODE=release
EXPOSE 8080

COPY --from=build ./src/dist/bank-api /bank-api

ENTRYPOINT [ "./bank-api" ]


# build stage
FROM golang:alpine AS build-env
ADD . /src
RUN cd /src && go build -o ii ./cmd/install-it/

# final stage
FROM ubuntu
WORKDIR /app
COPY --from=build-env /src/ii /app/

RUN ./ii wget

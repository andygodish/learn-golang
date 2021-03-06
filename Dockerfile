FROM golang:1.17-alpine as dev

WORKDIR /work

FROM golang:1.17-alpine as build

WORKDIR /app
COPY ./json/* /app/
RUN go build -o app

FROM golang:1.17-alpine as runtime

WORKDIR /app

COPY --from=build /app/app /app/
COPY ./json/videos.json /app/
CMD ./app 


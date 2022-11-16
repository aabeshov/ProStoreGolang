FROM golang:latest as build-stage
WORKDIR /GolangwithFrame
#COPY ./ ./
COPY go.* .
COPY . .
RUN apt-get update
RUN apt-get -y install postgresql-client

# make wait-for-postgres.sh executable
RUN chmod +x wait-for-postgres.sh

RUN go mod download
RUN go build -o server ./app/server.go



#FROM alpine:latest
#RUN apk --no-cache add ca-certificates libc6-compat
#WORKDIR /GolangwithFrame/
#COPY --from=build-stage /GolangwithFrame/app-start .
CMD [ "./server" ]


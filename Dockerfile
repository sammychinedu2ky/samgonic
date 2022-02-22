FROM golang:latest

LABEL maintainer="Samson <sammychinedu2ky@gmail.com>"

WORKDIR /app

COPY go.sum .

# RUN go mod download

COPY . .

ENV PORT 5000

RUN go build

RUN find . -name "*.go" -type f -delete 

EXPOSE ${PORT}

CMD [ "./samgonic" ]
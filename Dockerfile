FROM golang:latest
ENV CARBOFRA_DATA_URL=${CARBOFRA_DATA_URL}
ENV CARBOFRA_MONGO_URL=${CARBOFRA_MONGO_URL}
WORKDIR /go/carbofrapuller
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
CMD ["carbofrapuller"]
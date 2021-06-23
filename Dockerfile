FROM golang
RUN mkdir /app && mkdir /app/data
ADD . /app/
WORKDIR /app
RUN make build
ENTRYPOINT ["/app/out/bin/sberapi-mock", "start", "--port", "8080"]
EXPOSE 8080
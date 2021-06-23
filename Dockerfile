FROM golang AS builder

RUN mkdir /app && mkdir /app/data
ADD . /app/
WORKDIR /app
RUN make build

FROM gcr.io/distroless/base-debian10

COPY --from=builder /app/out/bin/sberapi-mock /app/sberapi-mock

WORKDIR /app
EXPOSE 8080
ENTRYPOINT ["/app/sberapi-mock", "start", "--port", "8080"]
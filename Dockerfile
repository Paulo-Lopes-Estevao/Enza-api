FROM golang as builder

WORKDIR /go/src/apiworld

ENV GOOS=linux
ENV PATH="/go/bin:${PATH}"
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOARCH=amd64

COPY . /go/src/apiworld/


RUN go build .

WORKDIR /go/src/apiworld

FROM scratch

COPY --from=builder /go/src/apiworld/main ./main
COPY --from=builder /go/src/apiworld/ /
COPY --from=builder /go/src/apiworld/entities /entities
COPY --from=builder /go/src/apiworld/services /services
COPY --from=builder /go/src/apiworld/utils /utils

EXPOSE 2000

CMD [ "./main" ]
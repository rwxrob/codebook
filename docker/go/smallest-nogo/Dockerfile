FROM golang AS gobuild
ADD hello.go /
WORKDIR /
RUN go build hello.go 

FROM scratch
COPY --from=gobuild hello /
CMD ["/hello"]

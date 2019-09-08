from golang:1.13.0-alpine3.10

run apk add --no-cache git

workdir /go/src/app
copy . .

env GO111MODULE on

run go get -d -v ./...
run CGO_ENABLED=0 go install -ldflags '-w -s' -v ./...

from scratch

copy --from=0 /go/bin/toy-tax-calculator /bin/toy-tax-calculator

cmd ["toy-tax-calculator"]

FROM golang:1.18-alpine AS build

# Build outside of GOPATH, it's simpler when using Go Modules.
WORKDIR /src

# Copy everything else.
COPY . .

# Build a static binary.
RUN CGO_ENABLED=0 GOOS=linux go build -mod vendor -a -buildvcs=false -installsuffix cgo -o app .

# Verify if the binary is truly static.
RUN ldd /src/app 2>&1 | grep -q 'Not a valid dynamic program'

LABEL app=api-feira
LABEL builder=true
LABEL maintainer='<martins.rso@gmail.com>'

FROM alpine
COPY --from=build /src/app ./app
COPY --from=build /src/config ./config
COPY --from=build /src/db ./db
ENTRYPOINT ["./app"]
EXPOSE 8888
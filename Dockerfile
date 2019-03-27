##
## Build container
##
FROM golang:1.12-alpine AS build

# Ensure we can download dependencies
RUN apk add --no-cache git

COPY . /go/src/http-test

RUN go get /go/src/http-test/app && go build -o /bin/myapp http-test/app 

##
## Runtime image
##
## Golang bits not required to run the executable
##
FROM alpine

EXPOSE 3001

COPY --from=build /bin/myapp /bin/myapp

RUN mkdir -vp /assets
WORKDIR /assets/static
COPY --from=build /go/src/http-test/assets /assets

CMD [ "/bin/myapp" ]

# ENTRYPOINT [ "/bin/myapp" ]
# CMD [ "One", "Two", "Three :)" ]


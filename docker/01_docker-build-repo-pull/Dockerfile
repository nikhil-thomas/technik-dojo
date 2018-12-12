# example multistage docker build
# to fetch source code from private repository
# to build source code
# to create deployable image
#
# the code reposiotry used in this example as payload is
# https://github.com/nikhil-thomas/sebastian
#
# To build
# $ docker build --build-arg SSH_PRIVATE_KEY="$(cat <ssh_private_key_path>)"  -t  <name:tag> .
#
#==========:==========:==========:==========:==========:==========#
# Step1: Get source code
#
# soruce code is fetched in a separate build stage
# to protect ssh private key
# ssh-private key is a build argument
#
# build arggument (unline ENV) will be available only in image and not available in containers
#
FROM ubuntu as pullsource
# can be optimized using a smaller base image

RUN apt-get update
RUN apt-get install -y git

ARG SSH_PRIVATE_KEY
RUN mkdir /root/.ssh/
RUN echo "${SSH_PRIVATE_KEY}" > /root/.ssh/id_rsa
RUN chmod 600 /root/.ssh/id_rsa

RUN touch /root/.ssh/known_hosts
RUN ssh-keyscan github.com >> /root/.ssh/known_hosts

RUN git clone git@github.com:nikhil-thomas/sebastian.git
#==========:==========:==========:==========:==========:==========#

#==========:==========:==========:==========:==========:==========#
# Step2: Build Go server
#
# in this stage the application is built
# in first stage code is pulled using ssh private key
# then only the code repository is copied to the next stage
# hence avoiding the presence of ssh-private key in final image
#
FROM golang:1.9.2 AS golangbuildstage

COPY --from=pullsource /sebastian /go/src/github.com/nikhil-thomas/sebastian

WORKDIR /go/src/github.com/nikhil-thomas/sebastian/cmd/sebastian

# Add dep tool for restoring dependencies
RUN go get -u github.com/golang/dep/...

# Restore go dependencies
RUN dep ensure -vendor-only -v

RUN CGO_ENABLED=0 GOOS=linux go build -o sebastian
#==========:==========:==========:==========:==========:==========#

#==========:==========:==========:==========:==========:==========#
# Step3: Build server deployment container
#
# in this stage only the executable built in the
# previous stage is copied
# hence avoiding the presence of uncompiled source code in final image
#
FROM alpine:3.5

ENV GOPATH /go

WORKDIR /go/src/github.com/nikhil-thomas/sebastian/cmd/sebastian

# copy the executable
COPY --from=golangbuildstage /go/src/github.com/nikhil-thomas/sebastian/cmd/sebastian/sebastian .

# configuration, keys and/or other staic assets should be copied here
# if there are any

CMD ["./sebastian"]

EXPOSE 5000
#==========:==========:==========:==========:==========:==========#

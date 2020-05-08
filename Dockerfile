FROM golang:latest
# Author
MAINTAINER dangminhtruong
# Create working folder
RUN mkdir /app
COPY . /app
WORKDIR /app
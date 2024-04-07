FROM debian:bullseye-slim
LABEL maintainer="SimonShen <shenzh1990@gmail.com>"

RUN sed -i 's/deb.debian.org/mirrors.aliyun.com/g' /etc/apt/sources.list &&  \
    sed -i 's/security.debian.org/mirrors.aliyun.com/g' /etc/apt/sources.list &&  \
    apt-get update \
    && apt-get install -y \
      ca-certificates \
      curl gcc g++  \
      libleptonica-dev \
      libtesseract-dev=4.1.1-2.1 \
      tesseract-ocr=4.1.1-2.1 \
      tesseract-ocr-jpn \
      tesseract-ocr-chi-sim


RUN tesseract --version

RUN curl -L https://studygolang.com/dl/golang/go1.22.0.linux-amd64.tar.gz -o go.tar.gz \
    && tar -C /usr/local -xzf go.tar.gz \
    && rm go.tar.gz


ENV GO111MODULE=on
ENV GOPATH=/go
ENV PATH=$GOPATH/bin:/usr/local/go/bin:$PATH
ENV GOPROXY=https://goproxy.cn,direct

ADD . $GOPATH/src/go-ocr
WORKDIR $GOPATH/src/go-ocr
COPY  ./ ./
RUN go mod tidy
RUN go install .
#RUN go get -v ./...
#RUN go install .


EXPOSE 8888
CMD ["go-ocr"]

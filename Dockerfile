FROM golang
WORKDIR /tmp/build1

RUN git clone https://github.com/kansal-mukul/eirinix-annotate && \
    cd eirinix-annotate && \
    go build

ENTRYPOINT ["/tmp/build1/eirinix-annotate/eirinix-annotate"]

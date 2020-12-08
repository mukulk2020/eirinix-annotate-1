FROM golang
WORKDIR /tmp/build
RUN rm -r eirinix-annotate
RUN git clone https://github.com/kansal-mukul/eirinix-annotate && \
    cd eirinix-annotate && \
    go build

ENTRYPOINT ["/tmp/build/eirinix-annotate/eirinix-annotate"]

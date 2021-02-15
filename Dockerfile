FROM golang

ENV PRODUCT_ID=1234
ENV PRODUCT_METRIC=VIRTUAL_PROCESSOR_CORE_TEST
ENV PRODUCT_NAME="my_pRODUCT"
ENV PRODUCT_VERSION="1.0"
ENV PRODUCT_CHARGED_CONTAINERS=ALL

WORKDIR /tmp/build

COPY ./ .
COPY ./ /usr/local/go/src/eirinix-annotate/
RUN go build && \
ls
ENTRYPOINT ["/tmp/build/eirinix-annotate"]

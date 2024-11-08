FROM golang:1.23.3

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"

ARG HOST_OS
RUN if [ "$HOST_OS" = "linux" ]; then \
        usermod -u 1000 www-data && \
        mkdir -p /var/www/.cache && \
        chown -R www-data:www-data /go && \
        chown -R www-data:www-data /var/www/.cache && \
        chown -R www-data:www-data /go/src; \
    fi

RUN if [ "$HOST_OS" = "linux" ]; then \
        USER www-data; \
    fi

CMD ["tail", "-f", "/dev/null"]
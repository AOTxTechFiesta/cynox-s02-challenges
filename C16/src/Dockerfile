FROM ubuntu:22.04

RUN apt-get update && apt-get install -y socat

RUN useradd -ms /bin/sh escape_protocol
WORKDIR /home/escape_protocol

COPY ./escape_protocol ./
COPY ./flag.txt /

RUN chown -R escape_protocol:escape_protocol /home/escape_protocol && \
     chmod 750 /home/escape_protocol && \
     chmod 550 /home/escape_protocol/escape_protocol && \
     chown escape_protocol:escape_protocol /flag.txt && \
     chmod 440 /flag.txt

EXPOSE 9999

USER escape_protocol
CMD ["socat", "-T60", "TCP-LISTEN:9999,reuseaddr,fork","EXEC:/home/escape_protocol/escape_protocol"]
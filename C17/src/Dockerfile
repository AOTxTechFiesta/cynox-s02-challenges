FROM ubuntu:22.04

RUN apt-get update && apt-get install -y socat

RUN useradd -ms /bin/sh omnitrix
WORKDIR /home/omnitrix

COPY ./omnitrix ./
COPY ./flag.txt /

RUN chown -R omnitrix:omnitrix /home/omnitrix && \
     chmod 750 /home/omnitrix && \
     chmod 550 /home/omnitrix/omnitrix && \
     chown omnitrix:omnitrix /flag.txt && \
     chmod 440 /flag.txt

EXPOSE 9999

USER omnitrix
CMD ["socat", "-T60", "TCP-LISTEN:9999,reuseaddr,fork","EXEC:/home/omnitrix/omnitrix"]
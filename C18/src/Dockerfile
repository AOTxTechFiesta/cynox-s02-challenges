FROM debian:stable-slim

RUN apt-get update && \
    apt-get install -y --no-install-recommends socat procps psmisc \
    && apt-get clean && rm -rf /var/lib/apt/lists/*

RUN mkdir -p /home/bin
RUN cp /usr/bin/bash /home/bin/bash
RUN cp /usr/bin/socat /home/bin/socat

RUN mkdir -p /home/terminal
RUN cp /root/.bashrc /home/terminal/.bashrc
RUN cp /root/.profile /home/terminal/.profile

RUN echo "cyn0x{t3rm1nal_0ffl1n3_n0_c0mm4nd5_y3t_fl4g_w45_4lw4y5_th3r3}" > /home/terminal/.bash_history
RUN echo "Never gonna give you up" >> /home/terminal/.bash_history
RUN echo "Never gonna let you down" >> /home/terminal/.bash_history

RUN chmod 555 /home/terminal && \
    groupadd terminal && \
    useradd -g terminal -d /home/terminal -s /home/bin/bash terminal && \
    chown terminal:terminal /home/terminal

RUN echo "cyn0x{t3rm1nal_0ffl1n3_n0_c0mm4nd5_y3t_fl4g_w45_4lw4y5_th3r3}" > /home/terminal/.flag

RUN chmod 600 /home/terminal/.flag && \
    chown terminal:terminal /home/terminal/.flag

RUN chmod 755 /home/bin/bash
RUN chmod 777 /home/bin/socat

RUN rm -rf /usr/bin

USER terminal

EXPOSE 9999

CMD ["/home/bin/socat", "TCP-LISTEN:9999,reuseaddr,fork", "EXEC:/home/bin/bash,pty,stderr"]
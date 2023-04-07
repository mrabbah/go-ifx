# FROM mcr.microsoft.com/vscode/devcontainers/go:0-1.18-bullseye
FROM ubuntu:22.04

RUN apt-get update && apt-get install -y wget gcc

WORKDIR /tmp

RUN wget https://golang.org/dl/go1.17.5.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go1.17.5.linux-amd64.tar.gz \
    && rm go1.17.5.linux-amd64.tar.gz

ENV PATH="/usr/local/go/bin:${PATH}" 

RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y install --no-install-recommends unixodbc unixodbc-dev

ENV USERHOME /root 
ENV INFORMIXDIR /root/odbc
ENV LD_LIBRARY_PATH ${INFORMIXDIR}/lib:${INFORMIXDIR}/lib/cli:${INFORMIXDIR}/lib/esql

COPY ./odbc/OneDB-Linux64-ODBC-Driver.tar.gz /tmp
RUN  tar -xf OneDB-Linux64-ODBC-Driver.tar.gz -C ./
RUN  cp -r  /tmp/onedb-odbc-driver ${INFORMIXDIR}

COPY ./odbc/odbcinst.ini /etc/odbcinst.ini
COPY ./odbc/odbc.ini ${USERHOME}/.odbc.ini
COPY ./odbc/sqlhosts ${INFORMIXDIR}/etc/sqlhosts

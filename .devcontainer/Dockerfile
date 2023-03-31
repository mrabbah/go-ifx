FROM mcr.microsoft.com/vscode/devcontainers/go:0-1.18-bullseye

# [Optional] Uncomment this section to install additional OS packages.
RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y install --no-install-recommends unixodbc unixodbc-dev

COPY IBMInformixODBCDriverforUnix.deb /tmp/IBMInformixODBCDriverforUnix.deb
RUN dpkg -i /tmp/IBMInformixODBCDriverforUnix.deb

# [Optional] Uncomment the next line to use go get to install anything else you need
# RUN go get -x <your-dependency-or-tool>

# [Optional] Uncomment this line to install global node packages.
# RUN su vscode -c "source /usr/local/share/nvm/nvm.sh && npm install -g <your-package-here>" 2>&1

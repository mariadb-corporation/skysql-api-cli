# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.194.0/containers/go/.devcontainer/base.Dockerfile

# [Choice] Go version: 1, 1.16, 1.17
ARG VARIANT="1.17"
FROM mcr.microsoft.com/vscode/devcontainers/go:0-${VARIANT}

# [Choice] Node.js version: none, lts/*, 16, 14, 12, 10
ARG NODE_VERSION="none"
RUN if [ "${NODE_VERSION}" != "none" ]; then su vscode -c "umask 0002 && . /usr/local/share/nvm/nvm.sh && nvm install ${NODE_VERSION} 2>&1"; fi


ENV GOPRIVATE=github.com/mariadb-corporation
ENV GO111MODULE=on

# [Optional] Uncomment the next line to use go get to install anything else you need
RUN curl -vkL -o - https://github.com/intuit/auto/releases/download/v10.28.0/auto-linux.gz | gunzip > /usr/local/bin/auto && \
    chmod a+x /usr/local/bin/auto && \
    git config --global url."ssh://git@github.com:".insteadOf "https://github.com/" && \
    go get github.com/spf13/cobra/cobra && \
    chown -R vscode /go/pkg && \
    apt-get update && export DEBIAN_FRONTEND=noninteractive && \
    apt-get -y install --no-install-recommends bash-completion

# [Optional] Uncomment this line to install global node packages.
# RUN su vscode -c "source /usr/local/share/nvm/nvm.sh && npm install -g <your-package-here>" 2>&1

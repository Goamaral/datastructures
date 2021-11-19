FROM mcr.microsoft.com/vscode/devcontainers/go:1.17-bullseye

ARG WORKSPACE_PATH

RUN chown -R vscode /go

USER vscode

RUN go install github.com/go-delve/delve/cmd/dlv@master && mv $GOPATH/bin/dlv $GOPATH/bin/dlv-dap

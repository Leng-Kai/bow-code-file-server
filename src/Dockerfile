FROM golang:alpine AS build-env
ADD . /src
RUN cd /src && go build -o file-server main.go

FROM alpine
ENV PORT=9090 DOCS_PATH=/content/
WORKDIR /app
COPY --from=build-env /src/file-server /app/
EXPOSE ${PORT}
ENTRYPOINT ./file-server
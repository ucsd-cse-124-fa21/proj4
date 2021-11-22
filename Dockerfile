# syntax=docker/dockerfile:1
FROM golang:1.17.3-bullseye
COPY bin/SurfstoreClientExec /bin/SurfstoreClientExec
CMD /bin/SurfstoreClientExec

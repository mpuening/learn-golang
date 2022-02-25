##
## Build Go App
##
FROM golang:1.17 AS gobuild
WORKDIR /app
COPY ./ ./
RUN make build


##
## Build NPM Front End
##
FROM node:14 AS npmbuild
WORKDIR /app
COPY /frontend /app
RUN npm run build

##
## Deploy
##
FROM golang:1.17
WORKDIR /app
RUN cd /app
COPY --from=gobuild /app/bin/server /app/bin/server
COPY --from=npmbuild /app/dist /app/dist
EXPOSE 8080
CMD ["/app/bin/server", "--port=8080", "--dist=/app/dist"]
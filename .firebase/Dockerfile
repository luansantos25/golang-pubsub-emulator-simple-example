FROM node:alpine
RUN apk update && apk add bash
RUN apk add openjdk11

RUN npm install -g firebase-tools

WORKDIR /app

CMD [ "firebase", "--project=poc-project", "emulators:start" ]
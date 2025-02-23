FROM node:23.6.1-slim

RUN apt update && \
    apt install openssl procps -y && \
    npm install -g @nestjs/cli@11.0.2

USER node

WORKDIR /home/node/app

CMD [ "tail", "-f", "/dev/null" ]
FROM node:18.17.1

ENV APP_ROOT /web
ARG MAIN_DOMAIN
ENV NUXT_MAIN_DOMAIN ${MAIN_DOMAIN}
ENV NUXT_SERVER_DOMAIN ${SERVER_DOMAIN}

WORKDIR ${APP_ROOT}
ADD . ${APP_ROOT}

ADD package.json .
ADD package-lock.json .

RUN npm ci

RUN npm run build

CMD node .output/server/index.mjs

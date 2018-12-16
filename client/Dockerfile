FROM node:11-alpine as builder
COPY . /tmp/client
WORKDIR /tmp/client
RUN npm install
RUN npm build
ARG REQHACK_BASEHOST
ARG REQHACK_RANDOM
ARG REQHACK_USE_SSL
RUN node /tmp/client/make_nginx_config.js > /tmp/client/default.conf

FROM openresty/openresty:1.13.6.2-alpine as nginx
WORKDIR /var/www/app
COPY --from=builder /tmp/client/dist/ /var/www/app/
COPY --from=builder /tmp/client/default.conf /etc/nginx/conf.d/default.conf

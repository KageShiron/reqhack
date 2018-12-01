FROM node:11-alpine as builder
COPY . /tmp/client
WORKDIR /tmp/client
ARG REQHACK_BASEHOST
ARG REQHACK_RANDOM
ARG REQHACK_USE_SSL
RUN node /tmp/client/make_nginx_config.js > /tmp/client/default.conf
# > /tmp/client/default.conf
RUN yarn build

FROM nginx:1.15-alpine as nginx
WORKDIR /var/www/app
COPY --from=builder /tmp/client/dist/ /var/www/app/
COPY --from=builder /tmp/client/default.conf /etc/nginx/conf.d/default.conf

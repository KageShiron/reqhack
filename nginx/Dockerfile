FROM node:latest-alpine as builder
WORKDIR /tmp/client
RUN yarn build

FROM nginx:latest-alpine
WORKDIR /var/www/app
COPY --from=builder /tmp/app/dist/ /var/www/app/
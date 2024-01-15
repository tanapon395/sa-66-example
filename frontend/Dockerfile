FROM node:16 as builder
WORKDIR /app
COPY frontend ./
RUN yarn install
RUN yarn build

FROM nginx:alpine
COPY frontend/nginx.conf /etc/nginx/nginx.conf
WORKDIR /usr/share/nginx/html
COPY --from=builder /app/build .
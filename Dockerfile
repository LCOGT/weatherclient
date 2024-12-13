# Builder container
FROM node:11-alpine as builder
COPY . .
RUN npm install && npm run build

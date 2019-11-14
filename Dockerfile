# Builder container
FROM node:11-alpine as builder
COPY . .
RUN npm install && npm run build

# Production Deployment container
FROM nginx:1.16-alpine
COPY --from=builder dist/ /usr/share/nginx/html/dist
COPY --from=builder index.html index.js /usr/share/nginx/html/

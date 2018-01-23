FROM node:9.4-alpine
MAINTAINER Austin Riba <ariba@lco.global>

EXPOSE 80

COPY . .
RUN npm install && npm run build

RUN mkdir /srv/weather && cp -r dist/ index.html index.js /srv/weather

CMD ["node", "/srv/weather/index.js"]

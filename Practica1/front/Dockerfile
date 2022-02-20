FROM node:17-alpine3.14

RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app

COPY package*.json ./

RUN npm install --silent

COPY . .
EXPOSE 3000
CMD ["npm", "start"]
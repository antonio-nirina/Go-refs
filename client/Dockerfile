FROM node:14-slim

WORKDIR /usr/src
COPY package*.json ./
RUN npm ci -qy
COPY . .
EXPOSE 3000

CMD ["npm", "start"]
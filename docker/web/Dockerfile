FROM node:23.11-slim

WORKDIR /app

COPY ../../web/package*.json ./
RUN npm install

COPY ../../web .

# Uncomment the line below for production builds
#RUN npm run build
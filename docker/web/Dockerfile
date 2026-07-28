FROM node:23.11-slim AS deps

WORKDIR /app

COPY web/package*.json ./
RUN npm ci

FROM deps AS dev

COPY web .

CMD ["npm", "run", "dev"]

FROM deps AS builder

ARG VITE_API_HOST=http://localhost:8000
ARG VITE_WS_HOST=ws://localhost:8000

ENV VITE_API_HOST=$VITE_API_HOST
ENV VITE_WS_HOST=$VITE_WS_HOST

COPY web .
RUN npm run build

FROM nginx:1.27-alpine AS runner

COPY docker/web/nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=builder /app/dist /usr/share/nginx/html

EXPOSE 3000

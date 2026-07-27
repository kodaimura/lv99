FROM node:23.11-slim AS deps

WORKDIR /app

COPY web/package*.json ./
RUN npm ci

FROM deps AS dev

COPY web .

CMD ["npm", "run", "dev"]

FROM deps AS builder

ARG API_HOST=http://api:8000
ARG NEXT_PUBLIC_API_HOST=http://localhost:8000
ARG NEXT_PUBLIC_WS_HOST=ws://localhost:8000

ENV API_HOST=$API_HOST
ENV NEXT_PUBLIC_API_HOST=$NEXT_PUBLIC_API_HOST
ENV NEXT_PUBLIC_WS_HOST=$NEXT_PUBLIC_WS_HOST
ENV NEXT_TELEMETRY_DISABLED=1

COPY web .
RUN npm run build

FROM node:23.11-slim AS runner

WORKDIR /app

ENV NODE_ENV=production
ENV NEXT_TELEMETRY_DISABLED=1
ENV PORT=3000
ENV HOSTNAME=0.0.0.0

COPY --from=builder /app/public ./public
COPY --from=builder /app/.next/standalone ./
COPY --from=builder /app/.next/static ./.next/static

CMD ["node", "server.js"]

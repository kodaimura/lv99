# lv99 web

Vite + React + React Router で構成されたフロントエンドです。

## Development

```bash
npm install
npm run dev
```

デフォルトでは `http://localhost:3000` で起動します。
API は `VITE_API_HOST`、WebSocket は `VITE_WS_HOST` で指定できます。

```bash
VITE_API_HOST=http://localhost:8000 VITE_WS_HOST=ws://localhost:8000 npm run dev
```

## Build

```bash
npm run build
npm run preview
```

production で `VITE_API_HOST` / `VITE_WS_HOST` を未指定にすると、同一オリジンの `/api` と `/ws` を使います。

"use client";
import { useCallback, useEffect, useRef, useState } from "react";

type Msg = string | ArrayBufferLike | Blob | ArrayBufferView;
type Options = {
  retryBaseMs?: number;    // 初期リトライ間隔 (ms)
  retryMaxMs?: number;     // 最大リトライ間隔 (ms)
  pingEveryMs?: number;    // アプリ層 ping 間隔 (ms)。0/未指定なら無効
  pongTimeoutMs?: number;  // アプリ層 pong 待ちタイムアウト (ms)
  makePing?: () => any;    // 送る ping の形（既定: {type:"ping"}）
  isPong?: (data: any) => boolean; // pong 判定（既定: {type:"pong"}）
};

export const useReliableWebSocket = (url: string, opts: Options = {}) => {
  const {
    retryBaseMs = 500,
    retryMaxMs = 8000,
    // デフォルトは無効（サーバが JSON pong を返さない場合が多いため）
    pingEveryMs = 0,
    pongTimeoutMs = 10000,
    makePing = () => ({ type: "ping", t: Date.now() }),
    isPong = (d) => d && d.type === "pong",
  } = opts;

  const wsRef = useRef<WebSocket | null>(null);
  const [ready, setReady] = useState(false);
  const [attempt, setAttempt] = useState(0);
  const queueRef = useRef<Msg[]>([]);
  const onMessageRef = useRef<((ev: MessageEvent) => void) | null>(null);

  const pingTimer = useRef<number | null>(null);
  const pongTimer = useRef<number | null>(null);
  const reconnectTimer = useRef<number | null>(null);

  // StrictMode の初回マウント→即アンマウント対策:
  // CONNECTING 中にクリーンアップで即 close すると警告が出るので、
  // いったんフラグだけ立てて、onopen で正常 close する。
  const willCloseOnOpenRef = useRef(false);

  // アンマウント後に再接続スケジューリングしないためのフラグ
  const aliveRef = useRef(true);

  const cleanupTimers = () => {
    if (pingTimer.current) { clearInterval(pingTimer.current); pingTimer.current = null; }
    if (pongTimer.current) { clearTimeout(pongTimer.current); pongTimer.current = null; }
    if (reconnectTimer.current) { clearTimeout(reconnectTimer.current); reconnectTimer.current = null; }
  };

  const flush = useCallback(() => {
    const ws = wsRef.current;
    if (!ws || ws.readyState !== WebSocket.OPEN) return;
    while (queueRef.current.length) ws.send(queueRef.current.shift()!);
  }, []);

  const send = useCallback((data: Msg) => {
    const ws = wsRef.current;
    if (ws && ws.readyState === WebSocket.OPEN) {
      ws.send(data);
    } else {
      queueRef.current.push(data); // 開くまでキュー
    }
  }, []);

  const startHeartbeat = useCallback(() => {
    // アプリ層 ping を使わない場合は何もしない
    if (!pingEveryMs || pingEveryMs <= 0) return;

    if (pingTimer.current) clearInterval(pingTimer.current);
    pingTimer.current = window.setInterval(() => {
      const ws = wsRef.current;
      if (!ws || ws.readyState !== WebSocket.OPEN) return;

      try { ws.send(JSON.stringify(makePing())); } catch { }

      // pong待ちタイマー
      if (pongTimer.current) clearTimeout(pongTimer.current);
      pongTimer.current = window.setTimeout(() => {
        // pong なし → 接続が死んでいる前提で閉じて再接続
        try { ws.close(); } catch { }
      }, pongTimeoutMs);
    }, pingEveryMs) as unknown as number;
  }, [makePing, pingEveryMs, pongTimeoutMs]);

  // connect を参照で安定化（再接続スケジュールから呼ぶ）
  const connectRef = useRef<() => void>(() => { });

  const scheduleReconnect = useCallback(() => {
    if (!aliveRef.current) return; // アンマウント後は再接続しない
    setReady(false);
    const next = Math.min(retryBaseMs * 2 ** attempt, retryMaxMs);
    const jitter = Math.floor(Math.random() * (next * 0.2)); // ±20%ジッタ
    const delay = Math.max(250, next - jitter);
    if (reconnectTimer.current) clearTimeout(reconnectTimer.current);
    reconnectTimer.current = window.setTimeout(() => {
      if (!aliveRef.current) return;
      setAttempt((a) => a + 1);
      connectRef.current();
    }, delay);
  }, [attempt, retryBaseMs, retryMaxMs]);

  const connect = useCallback(() => {
    // 二重接続防止
    if (wsRef.current && (
      wsRef.current.readyState === WebSocket.OPEN ||
      wsRef.current.readyState === WebSocket.CONNECTING
    )) {
      return;
    }

    const ws = new WebSocket(url);
    wsRef.current = ws;
    willCloseOnOpenRef.current = false;

    ws.onopen = () => {
      // StrictMode 初回アンマウントの「あとで閉じる」指示があれば、正常 close して終わり
      if (willCloseOnOpenRef.current) {
        try { ws.close(); } catch { }
        return;
      }
      setReady(true);
      setAttempt(0);
      flush();
      startHeartbeat();
    };

    ws.onmessage = (ev) => {
      if (pingEveryMs && pingEveryMs > 0) {
        try {
          const data = JSON.parse((ev.data as string) ?? "");
          if (isPong(data)) {
            if (pongTimer.current) { clearTimeout(pongTimer.current); pongTimer.current = null; }
            return;
          }
        } catch {
          // JSON じゃない場合は素通し
        }
      }
      onMessageRef.current?.(ev);
    };

    ws.onerror = () => { /* 接続失敗などで発生するが、特に何もしない */ };

    ws.onclose = () => {
      setReady(false);
      cleanupTimers();
      if (wsRef.current === ws) wsRef.current = null;
      scheduleReconnect();
    };
  }, [flush, isPong, pingEveryMs, scheduleReconnect, startHeartbeat, url]);

  // connect を安定参照に格納
  useEffect(() => { connectRef.current = connect; }, [connect]);

  useEffect(() => {
    aliveRef.current = true;

    // オンライン復帰やタブ復帰で積極的に再接続
    const onOnline = () => connectRef.current();
    const onVisible = () => { if (document.visibilityState === "visible") connectRef.current(); };

    connectRef.current();
    window.addEventListener("online", onOnline);
    document.addEventListener("visibilitychange", onVisible);

    return () => {
      aliveRef.current = false;

      window.removeEventListener("online", onOnline);
      document.removeEventListener("visibilitychange", onVisible);
      cleanupTimers();

      const ws = wsRef.current;
      wsRef.current = null;

      if (!ws) return;

      // ★ CONNECTING 中は「確立後に正常 close」させる（警告回避）
      if (ws.readyState === WebSocket.CONNECTING) {
        willCloseOnOpenRef.current = true;
        return;
      }

      // OPEN / CLOSING は普通に閉じる
      if (ws.readyState === WebSocket.OPEN || ws.readyState === WebSocket.CLOSING) {
        try { ws.close(); } catch { }
      }
    };
  }, []);

  const onMessage = useCallback((fn: (ev: MessageEvent) => void) => {
    onMessageRef.current = fn;
    return () => { if (onMessageRef.current === fn) onMessageRef.current = null; };
  }, []);

  return { ready, send, onMessage };
};

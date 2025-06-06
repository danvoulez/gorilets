export function connectWs() {
  const ws = new WebSocket('ws://localhost:8080');
  ws.onopen = () => console.log('WS connected');
  ws.onclose = () => console.log('WS disconnected');
  ws.onerror = err => console.error('WS error', err);
}

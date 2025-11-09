// バージョン管理（キャッシュ破棄の足がかりに）
const SW_VERSION = "v1";

// 即時反映
self.addEventListener("install", (e) => {
    self.skipWaiting();
});

self.addEventListener("activate", (e) => {
    // 既存クライアントを持ち上げる（任意）
    e.waitUntil(self.clients.claim());
});

self.addEventListener('push', (event) => {
    const data = event.data.json();
    self.registration.showNotification(data.title, {
        body: data.body,
    });
});
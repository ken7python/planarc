// バージョン管理（キャッシュ破棄の足がかりに）
const SW_VERSION = "v1";

// 即時反映
self.addEventListener("install", (e) => {
    console.log("Service Worker インストール:", SW_VERSION);
    self.skipWaiting();
});

self.addEventListener("activate", (e) => {
    console.log("Service Worker アクティブ化:", SW_VERSION);
    // 既存クライアントを持ち上げる（任意）
    e.waitUntil(self.clients.claim());
});

self.addEventListener('push', (event) => {
    console.log("プッシュイベント受信:", event);

    try {
        // データの存在確認
        if (!event.data) {
            console.warn("プッシュデータがありません");
            return;
        }

        const data = event.data.json();
        console.log("プッシュデータ:", data);

        // デフォルト値を設定
        const title = data.title || "通知";
        const body = data.body || "メッセージがあります";

        event.waitUntil(
            self.registration.showNotification(title, {
                body: body,
                // icon: data.icon || '/icon-192x192.png', // オプション
                // badge: data.badge || '/badge-72x72.png', // オプション
                tag: data.tag || 'default',
                requireInteraction: false,
            })
        );

        console.log("通知表示完了");
    } catch (error) {
        console.error("プッシュ通知エラー:", error);

        // エラー時でも通知を表示
        event.waitUntil(
            self.registration.showNotification("通知", {
                body: "メッセージを受信しました"
            })
        );
    }
});

// 通知クリック時の処理（オプション）
self.addEventListener('notificationclick', (event) => {
    console.log("通知がクリックされました:", event);
    event.notification.close();

    // アプリを開く
    event.waitUntil(
        clients.openWindow('/')
    );
});
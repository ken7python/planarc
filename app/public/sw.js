// バージョン管理（キャッシュ破棄の足がかりに）
const SW_VERSION = "v2";

// VitePWAのプリキャッシュ機能（本番ビルド時に自動注入）
if (typeof self.__WB_MANIFEST !== 'undefined') {
  // Workboxのプリキャッシュが利用可能な場合
  console.log("Workbox プリキャッシュ利用可能");
}

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
    console.log("🚀 プッシュイベント受信:", event);
    console.log("プッシュデータ存在確認:", !!event.data);

    try {
        let title = "PlanArc通知";
        let body = "新しいメッセージがあります";

        // データの存在確認
        if (event.data) {
            try {
                const data = event.data.json();
                console.log("📦 プッシュデータ:", data);
                title = data.title || title;
                body = data.body || body;
            } catch (jsonError) {
                console.warn("JSONパースエラー:", jsonError);
                console.log("生データ:", event.data.text());
                body = event.data.text() || body;
            }
        } else {
            console.warn("⚠️ プッシュデータがありません - デフォルト通知を表示");
        }

        console.log("📱 通知表示開始 - タイトル:", title, "本文:", body);

        const notificationOptions = {
            body: body,
            icon: '/pwa-192x192.png',
            badge: '/pwa-192x192.png',
            tag: 'planarc-notification-' + Date.now(), // ユニークなタグで確実に表示
            requireInteraction: false,
            silent: false, // 音を鳴らす
            actions: [
                {
                    action: 'open',
                    title: '開く'
                }
            ],
            data: { // 追加データ
                timestamp: Date.now(),
                url: '/'
            }
        };

        console.log("📋 通知オプション:", notificationOptions);
        const notificationPromise = self.registration.showNotification(title, notificationOptions);

        event.waitUntil(
            notificationPromise.then(() => {
                console.log("✅ 通知表示完了");
            }).catch((error) => {
                console.error("❌ 通知表示エラー:", error);
            })
        );

    } catch (error) {
        console.error("❌ プッシュ処理エラー:", error);
        console.error("エラースタック:", error.stack);

        // エラー時でも通知を表示
        event.waitUntil(
            self.registration.showNotification("PlanArc - エラー通知", {
                body: `通知処理中にエラーが発生しました: ${error.message}`,
                icon: '/pwa-192x192.png',
                tag: 'error-notification'
            }).catch((showError) => {
                console.error("❌❌ 緊急通知表示エラー:", showError);
            })
        );
    }
});

// 通知クリック時の処理
self.addEventListener('notificationclick', (event) => {
    console.log("🖱️ 通知がクリックされました:", event);
    event.notification.close();

    // アプリを開く
    event.waitUntil(
        clients.openWindow('/').catch((error) => {
            console.error("ウィンドウを開けませんでした:", error);
        })
    );
});
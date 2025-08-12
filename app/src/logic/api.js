// 初回のみコミットして、２回目移行は更新しないこと
// export const api = "https://api.planarc.kencode.tech/api"
// export const api = "http://localhost:8080/api"

export const api = import.meta.env.VITE_API_URL || "http://localhost:8080/api";
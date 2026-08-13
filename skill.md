# Aturan Pengembangan Proyek (Project Rules)

Dokumen ini mendefinisikan standar teknologi dan panduan arsitektur utama untuk proyek ini. Setiap penulisan kode **wajib** berpedoman pada tumpukan teknologi (tech stack) berikut:

## 1. Backend (Golang + Gin)
- **Framework Utama:** Gunakan `github.com/gin-gonic/gin` untuk semua kebutuhan routing dan pembuatan HTTP API.
- **Struktur Direktori:** Terapkan pola *Clean Architecture* atau *Domain-Driven Design* (pisahkan antara `Handler/Controller`, `Service/Usecase`, dan `Repository`).
- **API Response:** Selalu kembalikan format JSON yang konsisten, misalnya:
  ```json
  {
    "status": "success|error",
    "message": "Deskripsi",
    "data": {}
  }
  ```
- **Validasi:** Gunakan validasi *struct tag* bawaan Gin (`binding:"required"`) beserta metode `ShouldBindJSON()` untuk menangani input dari client.

## 2. Frontend (Nuxt 4 + Vue 3)
- **Framework:** Wajib menggunakan Nuxt 4 dengan Vue 3 Composition API (`<script setup>`).
- **Struktur Nuxt 4:** Gunakan direktori `app/` untuk membungkus `pages`, `components`, `layouts`, dan `composables`.
- **Integrasi Backend:** Frontend hanya bertindak sebagai SPA/SSR client yang mengkonsumsi API Gin dari backend.

## 3. UI & Styling (Shadcn UI Vue)
- **Sistem Komponen:** Gunakan ekosistem **Shadcn UI (Vue)** yang di-bangun di atas Reka UI / Radix Vue dan Tailwind CSS.
- **Premium Aesthetics:** Jangan gunakan gaya *default* yang membosankan. Terapkan prinsip desain premium:
  - Gunakan palet warna HSL kustom.
  - Terapkan *glassmorphism* (contoh: `bg-background/80 backdrop-blur-md`).
  - Tambahkan *micro-animations* yang halus pada interaksi elemen (seperti efek *hover*, *active*, atau transisi masuk).

## 4. Data Fetching (TanStack Vue Query v5)
- **Standar v5:** Gunakan `@tanstack/vue-query` versi 5.
- **Aturan `useQuery`:** JANGAN menggunakan *callback* `onSuccess`, `onError`, atau `onSettled` di dalam `useQuery` karena sudah dilarang di v5. *Callback* tersebut HANYA boleh digunakan di `useMutation`.
- **Query Key Factory:** Jangan menulis kunci *query* secara *hardcode* dalam bentuk *array* di berbagai file. Buat satu file utilitas (misal: `utils/queryKeys.ts`) untuk mendefinisikan semua *keys*.
- **Enkapulasi:** Selalu bungkus logika pemanggilan `useQuery` dan `useMutation` ke dalam *custom composable* (contoh: `useGetProducts()`), lalu panggil *composable* tersebut di dalam komponen Vue.

---
name: tanstack-vue-query
description: Panduan penggunaan TanStack Query (Vue Query v5) di dalam project Nuxt/Vue. Trigger saat membuat hooks data fetching, integrasi API, caching, atau state management asinkron.
---

# TanStack Query (Vue Query v5) Best Practices

Gunakan skill ini ketika menulis kode yang berhubungan dengan *data fetching*, mutasi data, dan manajemen state asinkron menggunakan `@tanstack/vue-query` di ekosistem Vue/Nuxt.

## 1. Versi yang Digunakan
Pastikan selalu menggunakan pola **TanStack Query v5**.
- **JANGAN** menggunakan pola v4 (misalnya, `onSuccess`, `onError`, `onSettled` di dalam `useQuery` sudah dihilangkan di v5).
- Callback tersebut hanya berlaku untuk `useMutation`.

## 2. Struktur Query Keys (Query Key Factory)
Selalu gunakan objek *Query Key Factory* untuk mencegah *typo* dan mempermudah invalidasi cache. Jangan pernah menulis *array hardcode* di berbagai tempat.

```typescript
// utils/queryKeys.ts
export const queries = {
  products: {
    all: () => ['products'] as const,
    list: (filters: string) => [...queries.products.all(), { filters }] as const,
    detail: (id: string) => [...queries.products.all(), id] as const,
  },
  transactions: {
    all: () => ['transactions'] as const,
  }
}
```

## 3. Custom Composables
Selalu bungkus `useQuery` dan `useMutation` di dalam custom composables. Jangan memanggilnya langsung di dalam komponen Vue.

**Pengambilan Data (Query):**
```typescript
import { useQuery } from '@tanstack/vue-query'
import { queries } from '~/utils/queryKeys'

export function useProductList(filters: Ref<string>) {
  return useQuery({
    queryKey: computed(() => queries.products.list(filters.value)),
    queryFn: async () => {
      const { data } = await $fetch('/api/products', { query: { f: filters.value } })
      return data
    },
    staleTime: 5 * 60 * 1000, // 5 menit
  })
}
```
*Catatan: Pastikan reaktivitas terjaga dengan melemparkan `computed` ke `queryKey` jika bergantung pada parameter dinamis.*

**Mutasi Data (Mutation):**
```typescript
import { useMutation, useQueryClient } from '@tanstack/vue-query'
import { queries } from '~/utils/queryKeys'

export function useCreateProduct() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: (newProduct: ProductInput) => {
      return $fetch('/api/products', { method: 'POST', body: newProduct })
    },
    onSuccess: () => {
      // Invalidasi cache agar data terbaru langsung dimuat
      queryClient.invalidateQueries({ queryKey: queries.products.all() })
    },
    onError: (error) => {
      // Tangani error, misal panggil toast notification
    }
  })
}
```

## 4. Integrasi Nuxt 
Karena ini adalah project Nuxt, hindari *data fetching* ganda antara server dan client. 
- Jika data sangat penting untuk SEO, gunakan `useAsyncData` atau `useFetch` bawaan Nuxt.
- Jika data bersifat interaktif, sering berubah, dan di-*fetch* di *client-side* (seperti Dashboard, POS, Admin panel), gunakan `@tanstack/vue-query`.
- Selalu sediakan *Nuxt Plugin* untuk menginisialisasi `VueQueryPlugin`.

## 5. Optimistic Updates
Gunakan *Optimistic Updates* untuk UI yang responsif (misalnya saat klik tombol Like atau Add to Cart) dengan memanipulasi cache via `queryClient.setQueryData` sebelum mutasi selesai di `onMutate`.

## 6. Jangan Menyalin State
Jangan pernah menyalin data dari `useQuery` ke dalam local `ref` atau `reactive`. Selalu gunakan data langsung dari `data.value` milik TanStack Query sebagai sumber kebenaran tunggal (*Single Source of Truth*).

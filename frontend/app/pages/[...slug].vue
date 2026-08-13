<template>
  <div class="min-h-[70vh] flex flex-col items-center justify-center text-center p-6">
    <div class="w-16 h-16 rounded-2xl bg-emerald-100 dark:bg-emerald-900/30 text-emerald-600 dark:text-emerald-400 flex items-center justify-center text-3xl font-bold mb-4">
      404
    </div>
    <h1 class="text-2xl font-bold text-zinc-900 dark:text-zinc-100">Halaman Tidak Ditemukan</h1>
    <p class="text-zinc-500 dark:text-zinc-400 mt-2 max-w-md text-sm">
      Maaf, halaman atau path yang Anda tuju belum tersedia atau telah dipindahkan.
    </p>

    <div class="flex items-center gap-3 mt-6">
      <NuxtLink 
        to="/" 
        class="px-4 py-2 bg-emerald-600 hover:bg-emerald-700 text-white rounded-lg text-sm font-medium transition-colors shadow-sm"
      >
        Kembali ke Dashboard
      </NuxtLink>
      <NuxtLink 
        to="/pos" 
        class="px-4 py-2 bg-zinc-100 dark:bg-zinc-800 hover:bg-zinc-200 dark:hover:bg-zinc-700 text-zinc-800 dark:text-zinc-200 rounded-lg text-sm font-medium transition-colors"
      >
        Buka POS Kasir
      </NuxtLink>
    </div>
  </div>
</template>

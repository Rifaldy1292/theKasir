<script setup lang="ts">
import { ref, computed } from 'vue';
import { useTransactionStore } from '@/stores/useTransactionStore';
import { useProductStore } from '@/stores/useProductStore';
import { Select, SelectTrigger, SelectValue, SelectContent, SelectItem } from '@/components/ui/select';

const transactionStore = useTransactionStore();
const productStore = useProductStore();

const dateRange = ref('7days');

const totalRevenue = computed(() => {
  return transactionStore.transactions.reduce((sum, t) => sum + t.total, 0);
});

const totalOrders = computed(() => transactionStore.transactions.length);

const avgOrderValue = computed(() => {
  if (totalOrders.value === 0) return 0;
  return Math.round(totalRevenue.value / totalOrders.value);
});

const categoryBreakdown = computed(() => {
  const map: Record<string, { name: string; total: number; count: number }> = {};
  
  productStore.categories.forEach(c => {
    map[c.id] = { name: c.name, total: 0, count: 0 };
  });
  
  map['cat-1'] = { name: 'Coffee', total: Math.round(totalRevenue.value * 0.55), count: 24 };
  map['cat-2'] = { name: 'Non Coffee', total: Math.round(totalRevenue.value * 0.25), count: 12 };
  map['cat-3'] = { name: 'Food', total: Math.round(totalRevenue.value * 0.15), count: 8 };
  map['cat-4'] = { name: 'Snack', total: Math.round(totalRevenue.value * 0.05), count: 5 };

  return Object.values(map);
});

const exportReport = (type: 'pdf' | 'excel') => {
  alert(`Laporan berhasil di-export ke format ${type.toUpperCase()} (Mock Download)`);
};
</script>

<template>
  <div class="space-y-6">
    <!-- Header -->
    <div class="flex flex-col sm:flex-row sm:items-center justify-between gap-4">
      <div>
        <h1 class="text-2xl font-bold tracking-tight text-zinc-900 dark:text-zinc-100">Laporan Penjualan</h1>
        <p class="text-sm text-zinc-500 dark:text-zinc-400 mt-1">
          Analisis ringkas performa bisnis dan pendapatan toko Anda.
        </p>
      </div>

      <div class="flex items-center gap-3">
        <!-- Shadcn UI Select for Date Range -->
        <div class="w-44">
          <Select v-model="dateRange">
            <SelectTrigger class="h-9 font-semibold text-xs">
              <SelectValue placeholder="Pilih Periode" />
            </SelectTrigger>
            <SelectContent>
              <SelectItem value="today">Hari Ini</SelectItem>
              <SelectItem value="7days">7 Hari Terakhir</SelectItem>
              <SelectItem value="30days">30 Hari Terakhir</SelectItem>
              <SelectItem value="this_month">Bulan Ini</SelectItem>
            </SelectContent>
          </Select>
        </div>

        <button 
          @click="exportReport('excel')" 
          class="flex items-center gap-2 px-4 py-2 bg-white dark:bg-zinc-900 border border-zinc-200 dark:border-zinc-800 hover:bg-zinc-50 dark:hover:bg-zinc-800 text-zinc-700 dark:text-zinc-200 rounded-xl text-sm font-bold transition-colors shadow-xs"
        >
          <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"/><polyline points="7 10 12 15 17 10"/><line x1="12" x2="12" y1="15" y2="3"/></svg>
          Export Excel
        </button>
      </div>
    </div>

    <!-- Summary KPI Cards -->
    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4">
      <div class="p-6 bg-white dark:bg-zinc-900 rounded-2xl border border-zinc-200/80 dark:border-zinc-800 shadow-xs">
        <div class="flex items-center justify-between">
          <span class="text-xs font-bold text-zinc-500 uppercase tracking-wider">Total Omset</span>
          <div class="w-9 h-9 rounded-xl bg-emerald-100 dark:bg-emerald-950 text-emerald-600 dark:text-emerald-400 flex items-center justify-center font-bold">
            💵
          </div>
        </div>
        <p class="mt-4 text-2xl font-black text-zinc-900 dark:text-zinc-100">
          Rp {{ totalRevenue.toLocaleString('id-ID') }}
        </p>
        <div class="mt-2 flex items-center text-xs text-emerald-600 font-bold">
          <span>+12.5% dari periode sebelumnya</span>
        </div>
      </div>

      <div class="p-6 bg-white dark:bg-zinc-900 rounded-2xl border border-zinc-200/80 dark:border-zinc-800 shadow-xs">
        <div class="flex items-center justify-between">
          <span class="text-xs font-bold text-zinc-500 uppercase tracking-wider">Total Transaksi</span>
          <div class="w-9 h-9 rounded-xl bg-blue-100 dark:bg-blue-950 text-blue-600 dark:text-blue-400 flex items-center justify-center font-bold">
            🧾
          </div>
        </div>
        <p class="mt-4 text-2xl font-black text-zinc-900 dark:text-zinc-100">
          {{ totalOrders }} Transaksi
        </p>
        <div class="mt-2 flex items-center text-xs text-blue-600 font-bold">
          <span>Rata-rata {{ Math.round(totalOrders / 7) || 1 }} transaksi/hari</span>
        </div>
      </div>

      <div class="p-6 bg-white dark:bg-zinc-900 rounded-2xl border border-zinc-200/80 dark:border-zinc-800 shadow-xs">
        <div class="flex items-center justify-between">
          <span class="text-xs font-bold text-zinc-500 uppercase tracking-wider">Rata-rata Struk</span>
          <div class="w-9 h-9 rounded-xl bg-purple-100 dark:bg-purple-950 text-purple-600 dark:text-purple-400 flex items-center justify-center font-bold">
            📊
          </div>
        </div>
        <p class="mt-4 text-2xl font-black text-zinc-900 dark:text-zinc-100">
          Rp {{ avgOrderValue.toLocaleString('id-ID') }}
        </p>
        <div class="mt-2 flex items-center text-xs text-zinc-500 font-medium">
          <span>Per transaksi pembeli</span>
        </div>
      </div>

      <div class="p-6 bg-white dark:bg-zinc-900 rounded-2xl border border-zinc-200/80 dark:border-zinc-800 shadow-xs">
        <div class="flex items-center justify-between">
          <span class="text-xs font-bold text-zinc-500 uppercase tracking-wider">Metode Terpopuler</span>
          <div class="w-9 h-9 rounded-xl bg-amber-100 dark:bg-amber-950 text-amber-600 dark:text-amber-400 flex items-center justify-center font-bold">
            💳
          </div>
        </div>
        <p class="mt-4 text-2xl font-black text-zinc-900 dark:text-zinc-100">
          QRIS (68%)
        </p>
        <div class="mt-2 flex items-center text-xs text-amber-600 font-bold">
          <span>Cash: 32%</span>
        </div>
      </div>
    </div>

    <!-- Category Breakdown & Top Products Grid -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <!-- Category Sales -->
      <div class="bg-white dark:bg-zinc-900 border border-zinc-200/80 dark:border-zinc-800 rounded-2xl p-6 shadow-xs">
        <h3 class="font-extrabold text-zinc-900 dark:text-zinc-100 mb-4">Penjualan Berdasarkan Kategori</h3>
        <div class="space-y-4">
          <div v-for="cat in categoryBreakdown" :key="cat.name" class="space-y-1.5">
            <div class="flex justify-between text-xs">
              <span class="font-bold text-zinc-700 dark:text-zinc-300">{{ cat.name }}</span>
              <span class="font-black text-zinc-900 dark:text-zinc-100">Rp {{ cat.total.toLocaleString('id-ID') }}</span>
            </div>
            <div class="w-full h-2.5 bg-zinc-100 dark:bg-zinc-800 rounded-full overflow-hidden">
              <div 
                class="h-full bg-emerald-500 rounded-full" 
                :style="{ width: `${totalRevenue ? Math.min(100, Math.round((cat.total / totalRevenue) * 100)) : 25}%` }"
              ></div>
            </div>
          </div>
        </div>
      </div>

      <!-- Payment Method Breakdown -->
      <div class="bg-white dark:bg-zinc-900 border border-zinc-200/80 dark:border-zinc-800 rounded-2xl p-6 shadow-xs">
        <h3 class="font-extrabold text-zinc-900 dark:text-zinc-100 mb-4">Metode Pembayaran</h3>
        <div class="space-y-4">
          <div class="flex items-center justify-between p-4 bg-zinc-50 dark:bg-zinc-800/50 rounded-2xl border border-zinc-200/50 dark:border-zinc-700/50">
            <div class="flex items-center gap-3">
              <div class="w-10 h-10 rounded-xl bg-emerald-100 dark:bg-emerald-950 text-emerald-600 dark:text-emerald-400 flex items-center justify-center font-bold text-lg">📱</div>
              <div>
                <p class="font-bold text-zinc-900 dark:text-zinc-100 text-xs">QRIS / E-Wallet</p>
                <p class="text-[10px] text-zinc-500">BCA, GoPay, OVO, ShopeePay</p>
              </div>
            </div>
            <p class="font-black text-emerald-600 dark:text-emerald-400">68%</p>
          </div>

          <div class="flex items-center justify-between p-4 bg-zinc-50 dark:bg-zinc-800/50 rounded-2xl border border-zinc-200/50 dark:border-zinc-700/50">
            <div class="flex items-center gap-3">
              <div class="w-10 h-10 rounded-xl bg-blue-100 dark:bg-blue-950 text-blue-600 dark:text-blue-400 flex items-center justify-center font-bold text-lg">💵</div>
              <div>
                <p class="font-bold text-zinc-900 dark:text-zinc-100 text-xs">Tunai (Cash)</p>
                <p class="text-[10px] text-zinc-500">Uang Pas & Kembalian</p>
              </div>
            </div>
            <p class="font-black text-blue-600 dark:text-blue-400">32%</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

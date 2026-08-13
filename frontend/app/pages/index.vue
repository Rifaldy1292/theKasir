<script setup lang="ts">
import { computed } from 'vue';
import { useAppState } from '@/composables/useAppState';
import { useTransactionStore } from '@/stores/useTransactionStore';
import { useProductStore } from '@/stores/useProductStore';

const { currentRole, currentWorkspace } = useAppState();
const transactionStore = useTransactionStore();
const productStore = useProductStore();

const currentWsId = computed(() => currentWorkspace.value?.id || 'ws-1');

const todayOmset = computed(() => transactionStore.todaySales(currentWsId.value));
const totalTrxCount = computed(() => transactionStore.totalTransactions(currentWsId.value));
const workspaceProducts = computed(() => productStore.getProductsByWorkspace(currentWsId.value));
const recentTrxList = computed(() => transactionStore.recentTransactions(currentWsId.value).slice(0, 5));

const stats = computed(() => [
  { 
    label: "Omset Hari Ini", 
    value: `Rp ${todayOmset.value.toLocaleString('id-ID')}`, 
    icon: '💰', 
    trend: '+14.2% vs kemarin',
    bgColor: 'from-emerald-500/10 to-teal-500/10 border-emerald-500/20 text-emerald-600'
  },
  { 
    label: 'Total Transaksi', 
    value: `${totalTrxCount.value} Transaksi`, 
    icon: '🧾', 
    trend: 'Rata-rata 18/jam',
    bgColor: 'from-blue-500/10 to-indigo-500/10 border-blue-500/20 text-blue-600'
  },
  { 
    label: 'Katalog Produk', 
    value: `${workspaceProducts.value.length} Produk`, 
    icon: '📦', 
    trend: 'Stok Terkendali',
    bgColor: 'from-purple-500/10 to-pink-500/10 border-purple-500/20 text-purple-600'
  },
  { 
    label: 'Produk Terlaris', 
    value: 'Es Kopi Susu', 
    icon: '🔥', 
    trend: '42 Porsi Terjual',
    bgColor: 'from-amber-500/10 to-orange-500/10 border-amber-500/20 text-amber-600'
  },
]);

// Top selling products preview data
const topProducts = [
  { name: 'Es Kopi Susu', category: 'Coffee', sales: 42, revenue: 756000, percent: 85 },
  { name: 'Americano Hot', category: 'Coffee', sales: 28, revenue: 420000, percent: 65 },
  { name: 'Matcha Latte', category: 'Non Coffee', sales: 19, revenue: 418000, percent: 45 },
  { name: 'Roti Bakar Cokelat', category: 'Food', sales: 14, revenue: 210000, percent: 30 },
];
</script>

<template>
  <div class="space-y-8 font-sans">
    <!-- Header Title & Workspace Badge -->
    <div class="flex flex-col sm:flex-row sm:items-center justify-between gap-4">
      <div>
        <div class="flex items-center gap-2 mb-1">
          <span class="px-2.5 py-0.5 rounded-full bg-emerald-100 text-emerald-700 dark:bg-emerald-950 dark:text-emerald-300 text-xs font-bold uppercase tracking-wider">
            {{ currentWorkspace?.type || 'BUSINESS WORKSPACE' }}
          </span>
        </div>
        <h1 class="text-3xl font-black tracking-tight text-zinc-900 dark:text-zinc-100">
          Dashboard {{ currentWorkspace?.name }}
        </h1>
        <p class="text-sm text-zinc-500 dark:text-zinc-400 mt-1">
          Ringkasan performa penjualan dan aktivitas toko Anda secara real-time.
        </p>
      </div>

      <div class="flex items-center gap-3">
        <NuxtLink 
          to="/pos" 
          class="px-5 py-2.5 bg-gradient-to-r from-emerald-600 to-teal-600 hover:from-emerald-500 hover:to-teal-500 text-white font-extrabold rounded-xl text-sm transition-all shadow-lg shadow-emerald-600/20 flex items-center gap-2"
        >
          <span>🖥️</span> Buka Kasir POS
        </NuxtLink>
      </div>
    </div>

    <!-- Owner / Admin Dashboard View -->
    <template v-if="currentRole === 'owner' || currentRole === 'admin'">
      <!-- Stats KPI Cards -->
      <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-5">
        <div 
          v-for="stat in stats" 
          :key="stat.label"
          class="p-6 bg-white dark:bg-zinc-900 rounded-2xl border border-zinc-200/80 dark:border-zinc-800 shadow-sm hover:shadow-md transition-all duration-200 relative overflow-hidden group"
        >
          <div class="flex items-center justify-between">
            <span class="text-xs font-bold text-zinc-500 uppercase tracking-wider">{{ stat.label }}</span>
            <div class="w-10 h-10 rounded-xl bg-gradient-to-br flex items-center justify-center text-xl shadow-sm" :class="stat.bgColor">
              {{ stat.icon }}
            </div>
          </div>

          <p class="mt-4 text-2xl font-black text-zinc-900 dark:text-zinc-100 tracking-tight">
            {{ stat.value }}
          </p>

          <div class="mt-2 flex items-center text-xs font-bold text-emerald-600 dark:text-emerald-400">
            <span>{{ stat.trend }}</span>
          </div>
        </div>
      </div>

      <!-- Main Analytics Layout (Sales Chart & Top Products) -->
      <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
        <!-- Sales Performance Card -->
        <div class="lg:col-span-2 p-6 bg-white dark:bg-zinc-900 rounded-2xl border border-zinc-200/80 dark:border-zinc-800 shadow-sm flex flex-col justify-between space-y-6">
          <div class="flex items-center justify-between">
            <div>
              <h3 class="font-extrabold text-lg text-zinc-900 dark:text-zinc-100">Grafik Penjualan Mingguan</h3>
              <p class="text-xs text-zinc-500 mt-0.5">Tren omset penjualan 7 hari terakhir</p>
            </div>

            <div class="flex items-center gap-1 bg-zinc-100 dark:bg-zinc-800 p-1 rounded-xl text-xs font-bold text-zinc-600 dark:text-zinc-400">
              <span class="px-2.5 py-1 bg-white dark:bg-zinc-700 text-zinc-900 dark:text-white rounded-lg shadow-xs">7 Hari</span>
              <span class="px-2.5 py-1 hover:text-zinc-900 dark:hover:text-white cursor-pointer">30 Hari</span>
            </div>
          </div>

          <!-- Simulated Chart Bar Visuals -->
          <div class="h-64 flex items-end justify-between gap-3 pt-6 px-2">
            <div v-for="(val, idx) in [45, 65, 30, 85, 90, 55, 100]" :key="idx" class="flex-1 flex flex-col items-center gap-2 group">
              <div 
                class="w-full bg-gradient-to-t from-emerald-600 to-teal-400 rounded-t-xl group-hover:from-emerald-500 group-hover:to-teal-300 transition-all duration-300 relative"
                :style="{ height: `${val}%` }"
              >
                <div class="opacity-0 group-hover:opacity-100 absolute -top-8 left-1/2 -translate-x-1/2 bg-zinc-900 text-white text-[10px] font-bold px-2 py-1 rounded shadow-lg transition-opacity whitespace-nowrap">
                  Rp {{ (val * 40000).toLocaleString('id-ID') }}
                </div>
              </div>
              <span class="text-[11px] font-bold text-zinc-400">
                {{ ['Sen', 'Sel', 'Rab', 'Kam', 'Jum', 'Sab', 'Min'][idx] }}
              </span>
            </div>
          </div>
        </div>

        <!-- Leaderboard Best Seller Products -->
        <div class="p-6 bg-white dark:bg-zinc-900 rounded-2xl border border-zinc-200/80 dark:border-zinc-800 shadow-sm flex flex-col space-y-4">
          <h3 class="font-extrabold text-lg text-zinc-900 dark:text-zinc-100">Produk Terlaris</h3>

          <div class="space-y-4 flex-1">
            <div v-for="item in topProducts" :key="item.name" class="space-y-1.5">
              <div class="flex justify-between items-center text-xs">
                <div>
                  <p class="font-bold text-zinc-900 dark:text-zinc-100">{{ item.name }}</p>
                  <p class="text-[10px] text-zinc-400">{{ item.category }} • {{ item.sales }} terjual</p>
                </div>
                <span class="font-black text-emerald-600 dark:text-emerald-400">Rp {{ item.revenue.toLocaleString('id-ID') }}</span>
              </div>
              <div class="w-full h-2 bg-zinc-100 dark:bg-zinc-800 rounded-full overflow-hidden">
                <div class="h-full bg-gradient-to-r from-emerald-500 to-teal-400 rounded-full" :style="{ width: `${item.percent}%` }"></div>
              </div>
            </div>
          </div>

          <NuxtLink to="/reports" class="text-center text-xs font-bold text-emerald-600 dark:text-emerald-400 hover:underline pt-2 border-t border-zinc-100 dark:border-zinc-800">
            Lihat Laporan Lengkap →
          </NuxtLink>
        </div>
      </div>

      <!-- Recent Activity Table -->
      <div class="bg-white dark:bg-zinc-900 rounded-2xl border border-zinc-200/80 dark:border-zinc-800 shadow-sm p-6 space-y-4">
        <div class="flex items-center justify-between">
          <h3 class="font-extrabold text-lg text-zinc-900 dark:text-zinc-100">Transaksi Terbaru</h3>
          <NuxtLink to="/transactions" class="text-xs font-bold text-emerald-600 dark:text-emerald-400 hover:underline">
            Semua Transaksi →
          </NuxtLink>
        </div>

        <div class="overflow-x-auto">
          <table class="w-full text-left text-xs whitespace-nowrap">
            <thead class="bg-zinc-50 dark:bg-zinc-800/50 border-b border-zinc-200 dark:border-zinc-800 text-zinc-500 font-bold uppercase">
              <tr>
                <th class="px-4 py-3">ID Transaksi</th>
                <th class="px-4 py-3">Pelanggan</th>
                <th class="px-4 py-3">Kasir</th>
                <th class="px-4 py-3">Tanggal & Waktu</th>
                <th class="px-4 py-3">Total Belanja</th>
                <th class="px-4 py-3">Status</th>
              </tr>
            </thead>
            <tbody class="divide-y divide-zinc-200/60 dark:divide-zinc-800 font-medium">
              <tr v-for="trx in recentTrxList" :key="trx.id" class="hover:bg-zinc-50 dark:hover:bg-zinc-800/40 transition-colors">
                <td class="px-4 py-3 font-mono font-bold text-zinc-900 dark:text-zinc-100">{{ trx.id }}</td>
                <td class="px-4 py-3 text-zinc-700 dark:text-zinc-300 font-semibold">{{ trx.customerName || 'Guest' }}</td>
                <td class="px-4 py-3 text-zinc-500">{{ trx.cashier }}</td>
                <td class="px-4 py-3 text-zinc-400">{{ new Date(trx.date).toLocaleString('id-ID') }}</td>
                <td class="px-4 py-3 font-extrabold text-emerald-600 dark:text-emerald-400">Rp {{ trx.total.toLocaleString('id-ID') }}</td>
                <td class="px-4 py-3">
                  <span class="px-2.5 py-0.5 rounded-full text-[10px] font-bold bg-emerald-100 text-emerald-700 dark:bg-emerald-950 dark:text-emerald-300">
                    {{ trx.status }}
                  </span>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </template>

    <!-- Cashier View Welcome Card -->
    <template v-else>
      <div class="p-10 bg-white dark:bg-zinc-900 rounded-3xl border border-zinc-200/80 dark:border-zinc-800 shadow-xl flex flex-col items-center justify-center text-center space-y-4">
        <div class="w-20 h-20 rounded-3xl bg-gradient-to-br from-emerald-500 to-teal-600 text-white flex items-center justify-center text-4xl shadow-xl shadow-emerald-500/30">
          👋
        </div>
        <h2 class="text-2xl font-black text-zinc-900 dark:text-zinc-100">Selamat Datang, Kasir!</h2>
        <p class="text-sm text-zinc-500 dark:text-zinc-400 max-w-md">
          Anda login sebagai Kasir di workspace <span class="font-bold text-zinc-900 dark:text-zinc-100">{{ currentWorkspace?.name }}</span>. Buka terminal POS untuk memproses transaksi pesanan pelanggan.
        </p>
        <NuxtLink to="/pos" class="mt-4 px-6 py-3 bg-gradient-to-r from-emerald-600 to-teal-600 hover:from-emerald-500 hover:to-teal-500 text-white font-extrabold rounded-2xl text-sm transition-all shadow-xl shadow-emerald-600/30">
          Buka Terminal Kasir POS NOW
        </NuxtLink>
      </div>
    </template>
  </div>
</template>

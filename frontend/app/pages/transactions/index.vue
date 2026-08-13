<script setup lang="ts">
import { ref, computed } from 'vue';
import { useTransactionStore, type Transaction } from '@/stores/useTransactionStore';
import { useAppState } from '@/composables/useAppState';
import { Table, TableHeader, TableBody, TableRow, TableHead, TableCell } from '@/components/ui/table';

const transactionStore = useTransactionStore();
const { currentWorkspace } = useAppState();

const selectedTransaction = ref<Transaction | null>(null);

const workspaceTransactions = computed(() => {
  return transactionStore.recentTransactions(currentWorkspace.value?.id || 'ws-1');
});

const openReceipt = (trx: Transaction) => {
  selectedTransaction.value = trx;
};

const printReceipt = () => {
  window.print();
};
</script>

<template>
  <div class="space-y-6">
    <div class="flex items-center justify-between no-print">
      <div>
        <h1 class="text-2xl font-bold tracking-tight text-zinc-900 dark:text-zinc-100">Riwayat Transaksi</h1>
        <p class="text-sm text-zinc-500 dark:text-zinc-400 mt-1">
          Daftar seluruh transaksi penjualan di toko Anda. Klik baris transaksi untuk melihat struk detail.
        </p>
      </div>
    </div>

    <!-- Shadcn UI Table -->
    <div class="bg-white dark:bg-zinc-900 border border-zinc-200 dark:border-zinc-800 rounded-2xl shadow-xs overflow-hidden no-print">
      <Table>
        <TableHeader>
          <TableRow>
            <TableHead>No. Transaksi</TableHead>
            <TableHead>Pelanggan</TableHead>
            <TableHead>Tanggal & Waktu</TableHead>
            <TableHead>Kasir</TableHead>
            <TableHead>Jumlah Barang</TableHead>
            <TableHead>Total Pembayaran</TableHead>
            <TableHead>Status</TableHead>
            <TableHead class="text-right">Struk</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          <TableRow 
            v-for="trx in workspaceTransactions" 
            :key="trx.id" 
            @click="openReceipt(trx)"
            class="cursor-pointer"
          >
            <TableCell class="font-bold text-zinc-900 dark:text-zinc-100 font-mono">{{ trx.id }}</TableCell>
            <TableCell class="text-zinc-700 dark:text-zinc-300 font-bold">{{ trx.customerName || 'Guest' }}</TableCell>
            <TableCell class="text-zinc-500 text-xs">{{ new Date(trx.date).toLocaleString('id-ID') }}</TableCell>
            <TableCell class="text-zinc-600 dark:text-zinc-400">{{ trx.cashier }}</TableCell>
            <TableCell class="text-zinc-600 dark:text-zinc-400">{{ trx.items }} item</TableCell>
            <TableCell class="text-zinc-900 dark:text-zinc-100 font-black text-emerald-600">Rp {{ trx.total.toLocaleString('id-ID') }}</TableCell>
            <TableCell>
              <span class="px-2.5 py-1 rounded-full text-xs font-bold bg-emerald-100 text-emerald-700 dark:bg-emerald-500/20 dark:text-emerald-400 border border-emerald-200 dark:border-emerald-500/20">
                {{ trx.status }}
              </span>
            </TableCell>
            <TableCell class="text-right">
              <span class="text-xs text-emerald-600 dark:text-emerald-400 hover:underline font-bold">
                Lihat Struk
              </span>
            </TableCell>
          </TableRow>
        </TableBody>
      </Table>
    </div>

    <!-- Receipt Detail Modal & Thermal Print Container -->
    <div v-if="selectedTransaction" class="fixed inset-0 z-50 bg-black/50 backdrop-blur-sm flex items-center justify-center p-4 print-container">
      <div class="bg-white dark:bg-zinc-900 border border-zinc-200 dark:border-zinc-800 rounded-3xl p-6 w-full max-w-sm shadow-2xl space-y-4 printable-receipt">
        <!-- Receipt Content -->
        <div class="text-center pb-3 border-b border-dashed border-zinc-300 dark:border-zinc-700 space-y-1">
          <h3 class="font-black text-lg text-zinc-900 dark:text-zinc-100 uppercase">{{ currentWorkspace?.name || 'KASIR PLATFORM' }}</h3>
          <p class="text-xs text-zinc-500">Struk Pembayaran Sah</p>
          <p class="text-xs text-zinc-400 font-mono mt-1">{{ selectedTransaction.id }} • {{ new Date(selectedTransaction.date).toLocaleString('id-ID') }}</p>
        </div>

        <div class="space-y-2 text-xs">
          <div class="flex justify-between text-zinc-500">
            <span>Pelanggan:</span>
            <span class="font-bold text-zinc-900 dark:text-zinc-100">{{ selectedTransaction.customerName || 'Guest' }}</span>
          </div>
          <div class="flex justify-between text-zinc-500">
            <span>Kasir:</span>
            <span class="font-bold text-zinc-900 dark:text-zinc-100">{{ selectedTransaction.cashier }}</span>
          </div>

          <div v-if="selectedTransaction.details && selectedTransaction.details.length > 0" class="py-2 border-y border-zinc-100 dark:border-zinc-800 space-y-1.5">
            <div v-for="item in selectedTransaction.details" :key="item.productId" class="flex justify-between items-center">
              <div>
                <p class="font-bold text-zinc-900 dark:text-zinc-100">{{ item.name }}</p>
                <p class="text-zinc-400 text-[10px]">{{ item.quantity }} x Rp {{ item.price.toLocaleString('id-ID') }}</p>
              </div>
              <p class="font-extrabold text-zinc-900 dark:text-zinc-100">Rp {{ (item.quantity * item.price).toLocaleString('id-ID') }}</p>
            </div>
          </div>
          <div v-else class="py-2 text-zinc-400 text-center italic">
            {{ selectedTransaction.items }} items purchased
          </div>

          <div class="pt-2 flex justify-between items-center text-sm font-extrabold border-t border-dashed border-zinc-300 dark:border-zinc-700">
            <span class="text-zinc-900 dark:text-zinc-100">Total:</span>
            <span class="text-emerald-600 dark:text-emerald-400 text-base font-black">Rp {{ selectedTransaction.total.toLocaleString('id-ID') }}</span>
          </div>
        </div>

        <div class="flex gap-2 pt-2 no-print">
          <button 
            @click="selectedTransaction = null"
            class="flex-1 py-2 text-xs font-bold text-zinc-600 dark:text-zinc-400 hover:bg-zinc-100 dark:hover:bg-zinc-800 rounded-xl transition-colors border border-zinc-200 dark:border-zinc-700"
          >
            Tutup
          </button>
          <button 
            @click="printReceipt"
            class="flex-1 py-2 text-xs font-bold bg-emerald-600 hover:bg-emerald-700 text-white rounded-xl transition-colors flex items-center justify-center gap-1.5 shadow-xs"
          >
            <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polyline points="6 9 6 2 18 2 18 9"/><path d="M6 18H4a2 2 0 0 1-2-2v-5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v5a2 2 0 0 1-2 2h-2"/><rect width="12" height="8" x="6" y="14"/></svg>
            Cetak Struk Thermal
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<style>
@media print {
  body * {
    visibility: hidden;
  }
  .print-container, .print-container * {
    visibility: visible;
  }
  .no-print {
    display: none !important;
  }
  .print-container {
    position: absolute;
    left: 0;
    top: 0;
    width: 58mm;
    padding: 0;
    margin: 0;
    background: transparent !important;
  }
  .printable-receipt {
    border: none !important;
    box-shadow: none !important;
    width: 58mm !important;
    padding: 4mm !important;
    font-size: 10px !important;
    color: black !important;
    background: white !important;
  }
}
</style>

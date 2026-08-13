<script setup lang="ts">
import { ref } from 'vue';
import { useAppState } from '@/composables/useAppState';

const { currentWorkspace } = useAppState();

const storeName = ref(currentWorkspace.value?.name || 'Kopi Senja');
const storeAddress = ref('Jl. Malioboro No. 45, Yogyakarta');
const storePhone = ref('+62 812-3456-7890');
const currency = ref('IDR');
const taxPercent = ref(11);
const servicePercent = ref(0);
const receiptHeader = ref('Terima kasih telah berkunjung!');
const receiptFooter = ref('Simpan struk ini sebagai bukti pembayaran sah.');

const isSaved = ref(false);

const saveSettings = () => {
  isSaved.value = true;
  setTimeout(() => {
    isSaved.value = false;
  }, 3000);
};
</script>

<template>
  <div class="space-y-6 max-w-4xl">
    <div>
      <h1 class="text-2xl font-bold tracking-tight text-zinc-900 dark:text-zinc-100">Pengaturan Toko</h1>
      <p class="text-sm text-zinc-500 dark:text-zinc-400 mt-1">
        Kelola informasi bisnis, pajak, dan konfigurasi struk kasir Anda.
      </p>
    </div>

    <!-- Alert Success -->
    <div v-if="isSaved" class="p-4 bg-emerald-50 dark:bg-emerald-950/40 border border-emerald-200 dark:border-emerald-800 text-emerald-800 dark:text-emerald-300 rounded-xl text-sm flex items-center gap-2">
      <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M22 11.08V12a10 10 0 1 1-5.93-9.14"/><polyline points="22 4 12 14.01 9 11.01"/></svg>
      <span>Pengaturan berhasil disimpan!</span>
    </div>

    <form @submit.prevent="saveSettings" class="space-y-6">
      <!-- General Store Profile -->
      <div class="bg-white dark:bg-zinc-900 border border-zinc-200 dark:border-zinc-800 rounded-xl p-6 shadow-sm space-y-4">
        <h2 class="text-lg font-semibold text-zinc-900 dark:text-zinc-100 border-b border-zinc-200 dark:border-zinc-800 pb-3">Profil Usaha</h2>
        
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div class="space-y-1.5">
            <label class="text-xs font-semibold text-zinc-600 dark:text-zinc-400 uppercase tracking-wider">Nama Toko</label>
            <input 
              v-model="storeName"
              type="text" 
              class="w-full px-3 py-2 bg-zinc-50 dark:bg-zinc-800 border border-zinc-200 dark:border-zinc-700 rounded-lg text-sm text-zinc-900 dark:text-zinc-100 focus:outline-none focus:ring-2 focus:ring-emerald-500/20 focus:border-emerald-500"
            />
          </div>

          <div class="space-y-1.5">
            <label class="text-xs font-semibold text-zinc-600 dark:text-zinc-400 uppercase tracking-wider">No. Telepon / WhatsApp</label>
            <input 
              v-model="storePhone"
              type="text" 
              class="w-full px-3 py-2 bg-zinc-50 dark:bg-zinc-800 border border-zinc-200 dark:border-zinc-700 rounded-lg text-sm text-zinc-900 dark:text-zinc-100 focus:outline-none focus:ring-2 focus:ring-emerald-500/20 focus:border-emerald-500"
            />
          </div>
        </div>

        <div class="space-y-1.5">
          <label class="text-xs font-semibold text-zinc-600 dark:text-zinc-400 uppercase tracking-wider">Alamat Lengkap</label>
          <textarea 
            v-model="storeAddress"
            rows="2"
            class="w-full px-3 py-2 bg-zinc-50 dark:bg-zinc-800 border border-zinc-200 dark:border-zinc-700 rounded-lg text-sm text-zinc-900 dark:text-zinc-100 focus:outline-none focus:ring-2 focus:ring-emerald-500/20 focus:border-emerald-500"
          ></textarea>
        </div>
      </div>

      <!-- Tax & Financial Settings -->
      <div class="bg-white dark:bg-zinc-900 border border-zinc-200 dark:border-zinc-800 rounded-xl p-6 shadow-sm space-y-4">
        <h2 class="text-lg font-semibold text-zinc-900 dark:text-zinc-100 border-b border-zinc-200 dark:border-zinc-800 pb-3">Pajak & Keuangan</h2>
        
        <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
          <div class="space-y-1.5">
            <label class="text-xs font-semibold text-zinc-600 dark:text-zinc-400 uppercase tracking-wider">Mata Uang</label>
            <select 
              v-model="currency"
              class="w-full px-3 py-2 bg-zinc-50 dark:bg-zinc-800 border border-zinc-200 dark:border-zinc-700 rounded-lg text-sm text-zinc-900 dark:text-zinc-100 focus:outline-none focus:ring-2 focus:ring-emerald-500/20 focus:border-emerald-500"
            >
              <option value="IDR">Rupiah (Rp)</option>
              <option value="USD">US Dollar ($)</option>
            </select>
          </div>

          <div class="space-y-1.5">
            <label class="text-xs font-semibold text-zinc-600 dark:text-zinc-400 uppercase tracking-wider">Pajak (PB1/PPN %)</label>
            <input 
              v-model.number="taxPercent"
              type="number" 
              min="0" max="100"
              class="w-full px-3 py-2 bg-zinc-50 dark:bg-zinc-800 border border-zinc-200 dark:border-zinc-700 rounded-lg text-sm text-zinc-900 dark:text-zinc-100 focus:outline-none focus:ring-2 focus:ring-emerald-500/20 focus:border-emerald-500"
            />
          </div>

          <div class="space-y-1.5">
            <label class="text-xs font-semibold text-zinc-600 dark:text-zinc-400 uppercase tracking-wider">Biaya Layanan (%)</label>
            <input 
              v-model.number="servicePercent"
              type="number" 
              min="0" max="100"
              class="w-full px-3 py-2 bg-zinc-50 dark:bg-zinc-800 border border-zinc-200 dark:border-zinc-700 rounded-lg text-sm text-zinc-900 dark:text-zinc-100 focus:outline-none focus:ring-2 focus:ring-emerald-500/20 focus:border-emerald-500"
            />
          </div>
        </div>
      </div>

      <!-- Receipt Print Settings -->
      <div class="bg-white dark:bg-zinc-900 border border-zinc-200 dark:border-zinc-800 rounded-xl p-6 shadow-sm space-y-4">
        <h2 class="text-lg font-semibold text-zinc-900 dark:text-zinc-100 border-b border-zinc-200 dark:border-zinc-800 pb-3">Pesan Struk Kasir</h2>
        
        <div class="space-y-4">
          <div class="space-y-1.5">
            <label class="text-xs font-semibold text-zinc-600 dark:text-zinc-400 uppercase tracking-wider">Pesan Atas (Header)</label>
            <input 
              v-model="receiptHeader"
              type="text" 
              class="w-full px-3 py-2 bg-zinc-50 dark:bg-zinc-800 border border-zinc-200 dark:border-zinc-700 rounded-lg text-sm text-zinc-900 dark:text-zinc-100 focus:outline-none focus:ring-2 focus:ring-emerald-500/20 focus:border-emerald-500"
            />
          </div>

          <div class="space-y-1.5">
            <label class="text-xs font-semibold text-zinc-600 dark:text-zinc-400 uppercase tracking-wider">Pesan Bawah (Footer)</label>
            <input 
              v-model="receiptFooter"
              type="text" 
              class="w-full px-3 py-2 bg-zinc-50 dark:bg-zinc-800 border border-zinc-200 dark:border-zinc-700 rounded-lg text-sm text-zinc-900 dark:text-zinc-100 focus:outline-none focus:ring-2 focus:ring-emerald-500/20 focus:border-emerald-500"
            />
          </div>
        </div>
      </div>

      <div class="flex justify-end">
        <button 
          type="submit"
          class="bg-emerald-600 hover:bg-emerald-700 text-white px-6 py-2.5 rounded-xl font-medium text-sm transition-colors shadow-sm active:scale-95"
        >
          Simpan Perubahan
        </button>
      </div>
    </form>
  </div>
</template>

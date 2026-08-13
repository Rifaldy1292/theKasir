<script setup lang="ts">
import { useAppState } from '@/composables/useAppState';
import { useWorkspaceStore, type Role } from '@/stores/useWorkspaceStore';
import { ref } from 'vue';

definePageMeta({
  layout: 'auth'
});

const workspaceStore = useWorkspaceStore();
const { setWorkspace } = useAppState();

const businessName = ref('');
const businessType = ref('COFFEE_SHOP');

const createWorkspace = () => {
  if (!businessName.value.trim()) return;
  
  const newId = `ws-${Date.now()}`;
  workspaceStore.addWorkspace({
    name: businessName.value,
    type: businessType.value,
    role: 'owner'
  });
  
  setWorkspace(newId);
  navigateTo('/');
};
</script>

<template>
  <div class="bg-white/80 dark:bg-zinc-900/80 backdrop-blur-xl border border-zinc-200 dark:border-zinc-800 rounded-2xl shadow-xl overflow-hidden">
    <div class="p-8">
      <div class="w-12 h-12 bg-emerald-100 dark:bg-emerald-900/30 text-emerald-600 dark:text-emerald-400 rounded-xl flex items-center justify-center mb-6">
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M3 9h18v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V9Z"/><path d="m3 9 2.45-4.9A2 2 0 0 1 7.24 3h9.52a2 2 0 0 1 1.8 1.1L21 9"/><path d="M12 3v6"/></svg>
      </div>
      <h1 class="text-2xl font-bold text-zinc-900 dark:text-zinc-100">Buat Workspace Baru</h1>
      <p class="text-sm text-zinc-500 dark:text-zinc-400 mt-2">
        Mulai atur bisnis Anda di Platform Kasir.
      </p>

      <form @submit.prevent="createWorkspace" class="mt-8 space-y-5">
        <div class="space-y-2">
          <label class="text-sm font-medium text-zinc-700 dark:text-zinc-300">Nama Bisnis / Toko</label>
          <input 
            v-model="businessName"
            type="text" 
            placeholder="contoh: Kopi Senja" 
            required
            class="w-full px-3.5 py-2.5 bg-zinc-50 dark:bg-zinc-800 border border-zinc-300 dark:border-zinc-700 rounded-xl text-zinc-900 dark:text-zinc-100 placeholder:text-zinc-400 focus:outline-none focus:ring-2 focus:ring-emerald-500/20 focus:border-emerald-500 transition-colors text-sm"
          />
        </div>
        
        <div class="space-y-2">
          <label class="text-sm font-medium text-zinc-700 dark:text-zinc-300">Kategori Bisnis</label>
          <select 
            v-model="businessType"
            class="w-full px-3.5 py-2.5 bg-zinc-50 dark:bg-zinc-800 border border-zinc-300 dark:border-zinc-700 rounded-xl text-zinc-900 dark:text-zinc-100 focus:outline-none focus:ring-2 focus:ring-emerald-500/20 focus:border-emerald-500 transition-colors text-sm"
          >
            <option value="COFFEE_SHOP">Coffee Shop / Kafe</option>
            <option value="LAUNDRY">Laundry</option>
            <option value="RETAIL">Retail / Minimarket / Warung</option>
            <option value="RESTAURANT">Restoran / Rumah Makan</option>
          </select>
        </div>

        <button 
          type="submit"
          class="w-full mt-6 bg-emerald-600 hover:bg-emerald-700 text-white font-medium py-3 rounded-xl transition-colors shadow-sm active:scale-95 text-sm"
        >
          Buat Workspace & Lanjut
        </button>
      </form>
      
      <div class="mt-6 text-center text-sm text-zinc-500">
        Sudah punya workspace? 
        <NuxtLink to="/workspace/select" class="text-emerald-600 dark:text-emerald-400 font-medium hover:underline">Pilih yang ada</NuxtLink>
      </div>
    </div>
  </div>
</template>

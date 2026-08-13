<script setup lang="ts">
import { useAppState } from '@/composables/useAppState';
import { useShiftStore } from '@/stores/useShiftStore';
import { useOnline } from '@vueuse/core';
import { computed, ref, onMounted, watchEffect } from 'vue';

const { currentWorkspace, currentRole } = useAppState();
const shiftStore = useShiftStore();

const isOnline = ref(true);
onMounted(() => {
  const onlineState = useOnline();
  watchEffect(() => {
    isOnline.value = onlineState.value;
  });
});

const cashierName = computed(() => currentRole.value === 'cashier' ? 'Budi (Kasir)' : 'Rifky (Owner)');
const currentShift = computed(() => shiftStore.activeShift(currentWorkspace.value?.id || 'ws-1', cashierName.value));
</script>

<template>
  <div class="min-h-screen bg-zinc-100 dark:bg-zinc-950 flex flex-col font-sans">
    <!-- POS Header -->
    <header class="h-16 flex items-center justify-between px-6 bg-white dark:bg-zinc-900 border-b border-zinc-200 dark:border-zinc-800 shrink-0">
      <div class="flex items-center gap-4">
        <NuxtLink to="/" class="p-2 hover:bg-zinc-100 dark:hover:bg-zinc-800 rounded-lg text-zinc-500 transition-colors" title="Kembali ke Dashboard">
          <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="m15 18-6-6 6-6"/></svg>
        </NuxtLink>
        <div class="font-bold text-lg text-zinc-900 dark:text-white flex items-center gap-2">
          {{ currentWorkspace?.name || 'Kasir POS' }}
        </div>
        <div class="px-2.5 py-1 rounded-md bg-zinc-100 dark:bg-zinc-800 text-xs font-medium text-zinc-600 dark:text-zinc-400">
          POS Terminal 01
        </div>

        <!-- Shift Indicator -->
        <div 
          class="px-2.5 py-1 rounded-md text-xs font-medium flex items-center gap-1.5"
          :class="currentShift ? 'bg-emerald-100 text-emerald-700 dark:bg-emerald-950/60 dark:text-emerald-300' : 'bg-amber-100 text-amber-700 dark:bg-amber-950/60 dark:text-amber-300'"
        >
          <span class="w-2 h-2 rounded-full" :class="currentShift ? 'bg-emerald-500 animate-pulse' : 'bg-amber-500'"></span>
          <span>Shift: {{ currentShift ? currentShift.id : 'Belum Buka' }}</span>
        </div>

        <!-- Online/Offline Indicator -->
        <ClientOnly>
          <div 
            class="flex items-center gap-1.5 px-2 py-0.5 rounded-md text-[11px] font-medium"
            :class="isOnline ? 'bg-zinc-100 text-zinc-600 dark:bg-zinc-800 dark:text-zinc-400' : 'bg-rose-100 text-rose-700 dark:bg-rose-950/60 dark:text-rose-300'"
          >
            {{ isOnline ? 'Online' : 'Offline' }}
          </div>
        </ClientOnly>
      </div>

      <div class="flex items-center gap-4">
        <div class="flex items-center gap-3">
          <div class="text-right">
            <p class="text-sm font-medium text-zinc-900 dark:text-zinc-100">{{ cashierName }}</p>
            <p class="text-xs text-zinc-500 dark:text-zinc-400 capitalize">Role: {{ currentRole }}</p>
          </div>
          <div class="w-10 h-10 rounded-full bg-emerald-100 dark:bg-emerald-900/30 flex items-center justify-center font-bold text-emerald-600 dark:text-emerald-400">
            {{ cashierName.charAt(0) }}
          </div>
        </div>
      </div>
    </header>

    <!-- POS Main Workspace -->
    <main class="flex-1 flex overflow-hidden">
      <slot />
    </main>

    <ClientOnly>
      <RoleSwitcher />
    </ClientOnly>
  </div>
</template>

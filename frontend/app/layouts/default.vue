<script setup lang="ts">
import { useAppState } from '@/composables/useAppState';
import { useOnline } from '@vueuse/core';

const { currentRole, currentWorkspace } = useAppState();
const isOnline = useOnline();

const navigation = [
  { name: 'Dashboard', path: '/', roles: ['owner', 'admin'] },
  { name: 'POS Kasir', path: '/pos', roles: ['owner', 'admin', 'cashier'] },
  { name: 'Transaksi', path: '/transactions', roles: ['owner', 'admin', 'cashier'] },
  { name: 'Pelanggan', path: '/customers', roles: ['owner', 'admin', 'cashier'] },
  { name: 'Produk', path: '/products', roles: ['owner', 'admin'] },
  { name: 'Karyawan', path: '/users', roles: ['owner', 'admin'] },
  { name: 'Laporan', path: '/reports', roles: ['owner', 'admin'] },
  { name: 'Pengaturan', path: '/settings', roles: ['owner'] },
];
</script>

<template>
  <div class="min-h-screen bg-zinc-50 dark:bg-zinc-950 flex">
    <!-- Sidebar -->
    <aside class="w-64 flex-shrink-0 bg-white dark:bg-zinc-900 border-r border-zinc-200 dark:border-zinc-800 flex flex-col sticky top-0 h-screen">
      <!-- Logo / Workspace Name -->
      <div class="h-16 flex items-center px-6 border-b border-zinc-200 dark:border-zinc-800">
        <div class="font-bold text-lg text-zinc-900 dark:text-white flex items-center gap-2">
          <div class="w-8 h-8 rounded bg-emerald-600 flex items-center justify-center text-white">
            <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M3 9h18v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V9Z"/><path d="m3 9 2.45-4.9A2 2 0 0 1 7.24 3h9.52a2 2 0 0 1 1.8 1.1L21 9"/><path d="M12 3v6"/></svg>
          </div>
          {{ currentWorkspace?.name || 'Kasir Apps' }}
        </div>
      </div>

      <!-- Navigation -->
      <nav class="flex-1 overflow-y-auto p-4 space-y-1">
        <template v-for="item in navigation" :key="item.path">
          <NuxtLink 
            v-if="item.roles.includes(currentRole)"
            :to="item.path"
            class="flex items-center gap-3 px-3 py-2.5 rounded-lg text-sm font-medium transition-colors text-zinc-600 dark:text-zinc-400 hover:bg-zinc-100/50 dark:hover:bg-zinc-800/50 hover:text-zinc-900 dark:hover:text-zinc-100"
            active-class="bg-zinc-100 dark:bg-zinc-800 text-zinc-900 dark:text-zinc-100"
          >
            {{ item.name }}
          </NuxtLink>
        </template>
      </nav>

      <!-- User Info -->
      <div class="p-4 border-t border-zinc-200 dark:border-zinc-800">
        <div class="flex items-center gap-3">
          <div class="w-10 h-10 rounded-full bg-zinc-200 dark:bg-zinc-800 flex items-center justify-center font-bold text-zinc-600 dark:text-zinc-400">
            R
          </div>
          <div class="flex-1 min-w-0">
            <p class="text-sm font-medium text-zinc-900 dark:text-zinc-100 truncate">Rifky</p>
            <p class="text-xs text-zinc-500 dark:text-zinc-500 truncate capitalize">{{ currentRole }}</p>
          </div>
        </div>
      </div>
    </aside>

    <!-- Main Content -->
    <main class="flex-1 flex flex-col min-w-0 overflow-hidden">
      <!-- Topbar -->
      <header class="h-16 flex items-center justify-between px-8 bg-white/80 dark:bg-zinc-900/80 backdrop-blur-md border-b border-zinc-200 dark:border-zinc-800 sticky top-0 z-10">
        <div class="flex items-center gap-3">
          <div class="px-3 py-1 rounded-full bg-zinc-100 dark:bg-zinc-800 text-xs font-medium text-zinc-600 dark:text-zinc-400 capitalize">
            Workspace: {{ currentWorkspace?.type }}
          </div>
          <!-- Online/Offline Indicator -->
          <div 
            class="flex items-center gap-1.5 px-2.5 py-1 rounded-full text-xs font-medium"
            :class="isOnline ? 'bg-emerald-100 text-emerald-700 dark:bg-emerald-950/60 dark:text-emerald-300' : 'bg-rose-100 text-rose-700 dark:bg-rose-950/60 dark:text-rose-300'"
          >
            <span class="w-2 h-2 rounded-full" :class="isOnline ? 'bg-emerald-500 animate-pulse' : 'bg-rose-500'"></span>
            {{ isOnline ? 'Online' : 'Offline Mode' }}
          </div>
        </div>
        <div class="flex items-center gap-4">
          <NuxtLink to="/workspace/select" class="text-sm font-medium text-zinc-600 dark:text-zinc-400 hover:text-zinc-900 dark:hover:text-zinc-100 transition-colors">
            Switch Workspace
          </NuxtLink>
        </div>
      </header>

      <!-- Page Content -->
      <div class="flex-1 overflow-auto p-8">
        <slot />
      </div>
    </main>

    <RoleSwitcher />
  </div>
</template>

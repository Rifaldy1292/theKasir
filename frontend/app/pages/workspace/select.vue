<script setup lang="ts">
import { useAppState } from '@/composables/useAppState';
import { useWorkspaceStore } from '@/stores/useWorkspaceStore';

definePageMeta({
  layout: 'auth'
});

const workspaceStore = useWorkspaceStore();
const { setWorkspace, currentRole } = useAppState();

const selectWorkspace = (id: string) => {
  setWorkspace(id);
  navigateTo('/');
};
</script>

<template>
  <div class="bg-white/80 dark:bg-zinc-900/80 backdrop-blur-xl border border-zinc-200 dark:border-zinc-800 rounded-2xl shadow-xl overflow-hidden">
    <div class="p-8">
      <h1 class="text-2xl font-bold text-zinc-900 dark:text-zinc-100">Select Workspace</h1>
      <p class="text-sm text-zinc-500 dark:text-zinc-400 mt-2">
        Choose a business workspace to continue.
      </p>

      <div class="mt-8 space-y-3">
        <button 
          v-for="ws in workspaceStore.workspaces" 
          :key="ws.id"
          @click="selectWorkspace(ws.id)"
          class="w-full flex items-center justify-between p-4 rounded-xl border border-zinc-200 dark:border-zinc-800 hover:border-emerald-500 hover:ring-1 hover:ring-emerald-500/20 bg-white dark:bg-zinc-950 transition-all text-left group"
        >
          <div class="flex items-center gap-4">
            <div class="w-10 h-10 rounded-lg bg-zinc-100 dark:bg-zinc-800 flex items-center justify-center text-zinc-600 dark:text-zinc-400 group-hover:bg-emerald-100 dark:group-hover:bg-emerald-900/30 group-hover:text-emerald-600 dark:group-hover:text-emerald-400 transition-colors">
              <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M3 9h18v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V9Z"/><path d="m3 9 2.45-4.9A2 2 0 0 1 7.24 3h9.52a2 2 0 0 1 1.8 1.1L21 9"/><path d="M12 3v6"/></svg>
            </div>
            <div>
              <div class="font-medium text-zinc-900 dark:text-zinc-100">{{ ws.name }}</div>
              <div class="text-xs text-zinc-500 dark:text-zinc-400 capitalize">{{ ws.type.toLowerCase().replace('_', ' ') }} • {{ ws.role }}</div>
            </div>
          </div>
          <svg class="text-zinc-300 dark:text-zinc-600 group-hover:text-emerald-500 transition-colors" xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="m9 18 6-6-6-6"/></svg>
        </button>
      </div>

      <div v-if="currentRole === 'owner'" class="mt-6 pt-6 border-t border-zinc-200 dark:border-zinc-800">
        <NuxtLink 
          to="/workspace/create"
          class="w-full flex items-center justify-center gap-2 py-3 rounded-lg border border-dashed border-zinc-300 dark:border-zinc-700 text-sm font-medium text-zinc-600 dark:text-zinc-400 hover:text-emerald-600 hover:border-emerald-500 hover:bg-emerald-50 dark:hover:bg-emerald-950/30 transition-all"
        >
          <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M5 12h14"/><path d="M12 5v14"/></svg>
          Create New Workspace
        </NuxtLink>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useAppState } from '@/composables/useAppState';

const { currentRole, setRole } = useAppState();

const roles = ['owner', 'admin', 'cashier'] as const;
</script>

<template>
  <div class="fixed bottom-4 right-4 z-50 flex flex-col gap-2 p-3 bg-white/80 dark:bg-zinc-900/80 backdrop-blur-md rounded-xl border border-zinc-200 dark:border-zinc-800 shadow-xl">
    <div class="text-xs font-semibold text-zinc-500 mb-1 px-1">Dev: Role Switcher</div>
    <div class="flex gap-1 bg-zinc-100 dark:bg-zinc-800 p-1 rounded-lg">
      <button
        v-for="role in roles"
        :key="role"
        @click="setRole(role)"
        class="px-3 py-1.5 text-xs font-medium rounded-md transition-all duration-200 capitalize"
        :class="[
          currentRole === role 
            ? 'bg-white dark:bg-zinc-700 text-zinc-900 dark:text-zinc-100 shadow-sm' 
            : 'text-zinc-500 hover:text-zinc-700 dark:hover:text-zinc-300 hover:bg-zinc-200/50 dark:hover:bg-zinc-700/50'
        ]"
      >
        {{ role }}
      </button>
    </div>
  </div>
</template>

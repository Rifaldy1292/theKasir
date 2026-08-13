<script setup lang="ts">
import { cn } from '@/lib/utils';

const props = defineProps<{
  class?: string;
}>();
</script>

<template>
  <thead :class="cn('[&_tr]:border-b bg-zinc-50/80 dark:bg-zinc-900/80 text-zinc-500 font-bold uppercase text-xs', props.class)">
    <slot />
  </thead>
</template>

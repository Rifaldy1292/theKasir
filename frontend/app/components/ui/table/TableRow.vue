<script setup lang="ts">
import { cn } from '@/lib/utils';

const props = defineProps<{
  class?: string;
}>();
</script>

<template>
  <tr :class="cn('border-b border-zinc-200/60 dark:border-zinc-800 transition-colors hover:bg-zinc-50/80 dark:hover:bg-zinc-800/40 data-[state=selected]:bg-zinc-100', props.class)">
    <slot />
  </tr>
</template>

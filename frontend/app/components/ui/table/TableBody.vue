<script setup lang="ts">
import { cn } from '@/lib/utils';

const props = defineProps<{
  class?: string;
}>();
</script>

<template>
  <tbody :class="cn('[&_tr:last-child]:border-0 divide-y divide-zinc-200/60 dark:divide-zinc-800', props.class)">
    <slot />
  </tbody>
</template>

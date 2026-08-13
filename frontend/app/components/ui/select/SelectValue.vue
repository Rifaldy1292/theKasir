<script setup lang="ts">
import { SelectValue } from 'reka-ui';

const props = defineProps<{
  placeholder?: string;
  asChild?: boolean;
  as?: any;
}>();
</script>

<template>
  <SelectValue v-bind="props">
    <slot />
  </SelectValue>
</template>

<script setup lang="ts">
import { SelectItem, SelectItemIndicator, SelectItemText, useForwardProps } from 'reka-ui';
import { Check } from '@lucide/vue';
import { cn } from '@/lib/utils';

const props = defineProps<{
  value: string;
  disabled?: boolean;
  textValue?: string;
  asChild?: boolean;
  as?: any;
  class?: string;
}>();

const forwardedProps = useForwardProps(props);
</script>

<template>
  <SelectItem
    v-bind="forwardedProps"
    :class="
      cn(
        'relative flex w-full cursor-default select-none items-center rounded-xl py-2 pl-8 pr-3 text-xs font-semibold outline-none focus:bg-emerald-50 focus:text-emerald-700 dark:focus:bg-emerald-950/50 dark:focus:text-emerald-300 data-[disabled]:pointer-events-none data-[disabled]:opacity-50 transition-colors',
        props.class
      )
    "
  >
    <span class="absolute left-2 flex h-3.5 w-3.5 items-center justify-center">
      <SelectItemIndicator>
        <Check class="h-4 w-4 text-emerald-600 dark:text-emerald-400 font-bold" />
      </SelectItemIndicator>
    </span>

    <SelectItemText>
      <slot />
    </SelectItemText>
  </SelectItem>
</template>

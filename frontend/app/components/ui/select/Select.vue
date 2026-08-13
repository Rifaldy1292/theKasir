<script setup lang="ts">
import { SelectRoot, useForwardPropsEmits } from 'reka-ui';

const props = defineProps<{
  open?: boolean;
  defaultOpen?: boolean;
  modelValue?: string | number | boolean | Record<string, any> | any[];
  defaultValue?: string | number | boolean | Record<string, any> | any[];
  dir?: 'ltr' | 'rtl';
  name?: string;
  autocomplete?: string;
  disabled?: boolean;
  required?: boolean;
}>();

const emits = defineEmits<{
  'update:modelValue': [value: any];
  'update:open': [value: boolean];
}>();

const forwarded = useForwardPropsEmits(props, emits);
</script>

<template>
  <SelectRoot v-bind="forwarded">
    <slot />
  </SelectRoot>
</template>

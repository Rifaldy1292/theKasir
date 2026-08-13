<template>
  <div class="min-h-screen bg-background text-foreground flex items-center justify-center p-4 selection:bg-primary selection:text-primary-foreground">
    <!-- Abstract premium shapes for background (Linear style) -->
    <div class="fixed inset-0 overflow-hidden pointer-events-none -z-10">
      <div class="absolute -top-[20%] -right-[10%] w-[50%] h-[50%] rounded-full bg-primary/5 blur-[120px]" />
      <div class="absolute -bottom-[20%] -left-[10%] w-[50%] h-[50%] rounded-full bg-primary/5 blur-[120px]" />
    </div>

    <!-- The Auth Content -->
    <div class="w-full max-w-[400px]">
      <slot />
    </div>
  </div>
</template>

<style scoped>
/* Ensure smooth fade in for auth pages */
div {
  animation: fade-in 0.6s ease-out forwards;
}

@keyframes fade-in {
  from { opacity: 0; transform: translateY(10px); }
  to { opacity: 1; transform: translateY(0); }
}
</style>

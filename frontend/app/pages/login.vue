<template>
  <NuxtLayout name="auth">
    <Card class="border-border/50 bg-card/60 backdrop-blur-xl shadow-2xl shadow-primary/5">
      <CardHeader class="space-y-1">
        <CardTitle class="text-2xl font-bold tracking-tight">theKasir</CardTitle>
        <CardDescription>
          Sign in to manage your workspaces.
        </CardDescription>
      </CardHeader>
      <CardContent class="grid gap-4">
        <div class="grid gap-2">
          <Label for="email">Email</Label>
          <Input id="email" type="email" placeholder="owner@kopi-senja.com" class="bg-background/50" />
        </div>
        <div class="grid gap-2">
          <Label for="password">Password</Label>
          <Input id="password" type="password" class="bg-background/50" />
        </div>
      </CardContent>
      <CardFooter class="flex flex-col gap-3">
        <Button class="w-full transition-all active:scale-[0.98]">
          Sign In
        </Button>
        <div class="text-sm text-center text-muted-foreground">
          Don't have an account? 
          <NuxtLink to="/register" class="text-primary hover:underline underline-offset-4">
            Register here
          </NuxtLink>
        </div>
      </CardFooter>
    </Card>
  </NuxtLayout>
</template>

<script setup lang="ts">
// We will integrate TanStack query login mutation here later
definePageMeta({
  layout: false // We are using manual layout wrapping above for better transition control if needed
})
</script>

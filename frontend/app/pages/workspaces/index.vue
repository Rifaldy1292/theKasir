<template>
  <NuxtLayout name="dashboard">
    <div class="space-y-6">
      <div class="flex items-center justify-between">
        <div>
          <h2 class="text-3xl font-bold tracking-tight">Workspaces</h2>
          <p class="text-muted-foreground">
            Select a workspace to manage its Point of Sale.
          </p>
        </div>
        <Button>
          Create Workspace
        </Button>
      </div>

      <!-- Loading State -->
      <div v-if="isLoading" class="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
        <div v-for="i in 3" :key="i" class="h-32 rounded-xl bg-muted/50 animate-pulse border border-border" />
      </div>

      <!-- Empty State -->
      <div v-else-if="!workspaces?.length" class="flex flex-col items-center justify-center p-12 text-center border rounded-xl bg-card/50 border-dashed">
        <div class="h-12 w-12 rounded-full bg-primary/10 flex items-center justify-center mb-4">
          <span class="text-xl">🏢</span>
        </div>
        <h3 class="text-lg font-semibold">No workspaces found</h3>
        <p class="text-sm text-muted-foreground mt-1 mb-4">You haven't joined or created any workspace yet.</p>
        <Button variant="outline">Create your first Workspace</Button>
      </div>

      <!-- Workspace Grid -->
      <div v-else class="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
        <Card 
          v-for="workspace in workspaces" 
          :key="workspace.id"
          class="group hover:border-primary/50 transition-colors cursor-pointer bg-card/60 backdrop-blur-sm"
          @click="selectWorkspace(workspace.id)"
        >
          <CardHeader>
            <CardTitle class="flex items-center justify-between">
              {{ workspace.name }}
              <span class="text-xs px-2 py-1 rounded-full bg-primary/10 text-primary font-normal">
                {{ workspace.role || 'Owner' }}
              </span>
            </CardTitle>
            <CardDescription class="capitalize">{{ workspace.business_type }}</CardDescription>
          </CardHeader>
        </Card>
      </div>
    </div>
  </NuxtLayout>
</template>

<script setup lang="ts">
import { useQuery } from '@tanstack/vue-query'
import { queryKeys } from '~/utils/queryKeys'

// Mock fetching workspaces using Vue Query (TanStack v5)
const api = useApi()

const { data: workspaces, isLoading } = useQuery({
  queryKey: queryKeys.workspaces.all(),
  queryFn: async () => {
    // For Phase 1 we mock the response until backend returns real data
    // const res = await api('/workspaces')
    // return res.data
    
    return [
      { id: 'ws_1', name: 'Kopi Senja', business_type: 'coffee_shop', role: 'Owner' },
      { id: 'ws_2', name: 'Laundry Bersih', business_type: 'laundry', role: 'Admin' }
    ]
  }
})

const selectWorkspace = (id: string) => {
  // Save selected workspace to state/cookie and navigate to POS
  navigateTo(`/workspaces/${id}/pos`)
}
</script>

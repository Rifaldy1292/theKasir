import { useState } from '#imports';
import { useWorkspaceStore, type Role, type Workspace } from '@/stores/useWorkspaceStore';
import { computed } from 'vue';

export function useAppState() {
  const currentRole = useState<Role>('app-current-role', () => 'owner');
  const currentWorkspaceId = useState<string>('app-current-ws-id', () => 'ws-1');

  const setRole = (role: Role) => {
    currentRole.value = role;
  };

  const setWorkspace = (workspaceId: string) => {
    currentWorkspaceId.value = workspaceId;
  };

  const currentWorkspace = computed<Workspace | undefined>(() => {
    const store = useWorkspaceStore();
    return store.workspaces.find(w => w.id === currentWorkspaceId.value) || store.workspaces[0];
  });

  return {
    currentRole,
    currentWorkspace,
    setRole,
    setWorkspace
  };
}

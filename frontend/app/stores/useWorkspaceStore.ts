import { defineStore } from 'pinia';
import { useLocalStorage } from '@vueuse/core';

export type Role = 'owner' | 'admin' | 'cashier';

export type Workspace = {
  id: string;
  name: string;
  type: string;
  role: Role;
};

const DEFAULT_WORKSPACES: Workspace[] = [
  { id: 'ws-1', name: 'Kopi Senja', type: 'COFFEE_SHOP', role: 'owner' },
  { id: 'ws-2', name: 'Laundry Bersih', type: 'LAUNDRY', role: 'admin' },
];

export const useWorkspaceStore = defineStore('workspace', {
  state: () => ({
    workspaces: useLocalStorage<Workspace[]>('thekasir-workspaces', DEFAULT_WORKSPACES),
  }),
  actions: {
    addWorkspace(workspace: Omit<Workspace, 'id'>) {
      const newId = `ws-${Date.now()}`;
      this.workspaces.push({ ...workspace, id: newId });
    }
  }
});

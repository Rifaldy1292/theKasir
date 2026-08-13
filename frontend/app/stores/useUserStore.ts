import { defineStore } from 'pinia';
import { useLocalStorage } from '@vueuse/core';

export type User = {
  id: string;
  workspaceId: string;
  name: string;
  email: string;
  role: 'owner' | 'admin' | 'cashier';
  status: string;
};

const DEFAULT_USERS: User[] = [
  { id: 'usr-1', workspaceId: 'ws-1', name: 'Rifky', email: 'rifky@kopi.com', role: 'owner', status: 'Active' },
  { id: 'usr-2', workspaceId: 'ws-1', name: 'Budi Kasir', email: 'budi@kopi.com', role: 'cashier', status: 'Active' },
  { id: 'usr-3', workspaceId: 'ws-1', name: 'Andi Admin', email: 'andi@kopi.com', role: 'admin', status: 'Active' },
  { id: 'usr-4', workspaceId: 'ws-2', name: 'Sari Laundry', email: 'sari@laundry.com', role: 'cashier', status: 'Active' },
];

export const useUserStore = defineStore('user', {
  state: () => ({
    users: useLocalStorage<User[]>('thekasir-users', DEFAULT_USERS),
  }),
  getters: {
    getUsersByWorkspace: (state) => {
      return (workspaceId: string) => {
        return state.users.filter(u => u.workspaceId === workspaceId);
      };
    }
  },
  actions: {
    addUser(user: Omit<User, 'id'>) {
      const newId = `usr-${Date.now()}`;
      this.users.push({ ...user, id: newId });
    },
    updateUser(id: string, updates: Partial<User>) {
      const index = this.users.findIndex(u => u.id === id);
      if (index !== -1) {
        this.users[index] = { ...this.users[index], ...updates };
      }
    },
    deleteUser(id: string) {
      this.users = this.users.filter(u => u.id !== id);
    }
  }
});

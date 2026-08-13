import { defineStore } from 'pinia';
import { useLocalStorage } from '@vueuse/core';

export type TransactionItem = {
  productId: string;
  name: string;
  price: number;
  quantity: number;
};

export type Transaction = {
  id: string;
  workspaceId: string;
  shiftId?: string;
  customerId?: string;
  customerName?: string;
  cashier: string;
  items: number;
  total: number;
  status: string;
  date: string;
  details: TransactionItem[];
};

const DEFAULT_TRANSACTIONS: Transaction[] = [
  { id: 'TRX-001', workspaceId: 'ws-1', customerName: 'Siti Aminah', cashier: 'Budi (Cashier)', items: 2, total: 36000, status: 'Paid', date: '2026-08-13T10:00:00Z', details: [] },
  { id: 'TRX-002', workspaceId: 'ws-1', customerName: 'Pelanggan Umum (Guest)', cashier: 'Rifky (Owner)', items: 1, total: 22000, status: 'Paid', date: '2026-08-13T11:30:00Z', details: [] },
  { id: 'TRX-003', workspaceId: 'ws-1', customerName: 'Bambang Tri', cashier: 'Budi (Cashier)', items: 3, total: 51000, status: 'Paid', date: '2026-08-13T12:15:00Z', details: [] },
];

export const useTransactionStore = defineStore('transaction', {
  state: () => ({
    transactions: useLocalStorage<Transaction[]>('thekasir-transactions', DEFAULT_TRANSACTIONS),
  }),
  getters: {
    getTransactionsByWorkspace: (state) => {
      return (workspaceId: string) => {
        return state.transactions.filter(t => t.workspaceId === workspaceId);
      };
    },
    todaySales: (state) => {
      return (workspaceId: string) => {
        const today = new Date().toISOString().split('T')[0];
        return state.transactions
          .filter(t => t.workspaceId === workspaceId && t.date.startsWith(today))
          .reduce((sum, t) => sum + t.total, 0);
      };
    },
    totalTransactions: (state) => {
      return (workspaceId: string) => state.transactions.filter(t => t.workspaceId === workspaceId).length;
    },
    recentTransactions: (state) => {
      return (workspaceId: string) => {
        return state.transactions
          .filter(t => t.workspaceId === workspaceId)
          .sort((a, b) => new Date(b.date).getTime() - new Date(a.date).getTime());
      };
    }
  },
  actions: {
    addTransaction(transaction: Omit<Transaction, 'id' | 'date'>) {
      const newId = `TRX-${String(this.transactions.length + 1).padStart(3, '0')}`;
      const newTrx: Transaction = {
        ...transaction,
        id: newId,
        date: new Date().toISOString()
      };
      this.transactions.push(newTrx);
      return newTrx;
    }
  }
});

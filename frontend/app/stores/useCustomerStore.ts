import { defineStore } from 'pinia';
import { useLocalStorage } from '@vueuse/core';

export type Customer = {
  id: string;
  workspaceId: string;
  name: string;
  phone: string;
  email?: string;
  address?: string;
  totalSpending: number;
  totalTransactions: number;
  createdAt: string;
};

const DEFAULT_CUSTOMERS: Customer[] = [
  {
    id: 'CUST-001',
    workspaceId: 'ws-1',
    name: 'Pelanggan Umum (Guest)',
    phone: '-',
    email: 'guest@kopi.com',
    address: '-',
    totalSpending: 0,
    totalTransactions: 0,
    createdAt: '2026-08-01T00:00:00Z'
  },
  {
    id: 'CUST-002',
    workspaceId: 'ws-1',
    name: 'Siti Aminah',
    phone: '081299887766',
    email: 'siti@gmail.com',
    address: 'Jl. Gejayan No. 12',
    totalSpending: 250000,
    totalTransactions: 5,
    createdAt: '2026-08-05T00:00:00Z'
  },
  {
    id: 'CUST-003',
    workspaceId: 'ws-1',
    name: 'Bambang Tri',
    phone: '085611223344',
    email: 'bambang@yahoo.com',
    address: 'Jl. Kaliurang Km 5',
    totalSpending: 180000,
    totalTransactions: 3,
    createdAt: '2026-08-10T00:00:00Z'
  }
];

export const useCustomerStore = defineStore('customer', {
  state: () => ({
    customers: useLocalStorage<Customer[]>('thekasir-customers', DEFAULT_CUSTOMERS),
  }),
  getters: {
    getCustomersByWorkspace: (state) => {
      return (workspaceId: string) => {
        return state.customers.filter(c => c.workspaceId === workspaceId);
      };
    }
  },
  actions: {
    addCustomer(customer: Omit<Customer, 'id' | 'totalSpending' | 'totalTransactions' | 'createdAt'>) {
      const newCustomer: Customer = {
        ...customer,
        id: `CUST-${Date.now().toString().slice(-4)}`,
        totalSpending: 0,
        totalTransactions: 0,
        createdAt: new Date().toISOString()
      };
      this.customers.push(newCustomer);
      return newCustomer;
    },
    updateCustomer(id: string, updates: Partial<Customer>) {
      const index = this.customers.findIndex(c => c.id === id);
      if (index !== -1) {
        this.customers[index] = { ...this.customers[index], ...updates };
      }
    },
    deleteCustomer(id: string) {
      this.customers = this.customers.filter(c => c.id !== id);
    },
    recordCustomerPurchase(id: string, amount: number) {
      const customer = this.customers.find(c => c.id === id);
      if (customer) {
        customer.totalSpending += amount;
        customer.totalTransactions += 1;
      }
    }
  }
});

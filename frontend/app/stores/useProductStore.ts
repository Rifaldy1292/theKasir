import { defineStore } from 'pinia';
import { useLocalStorage } from '@vueuse/core';

export type ItemType = 'PRODUCT' | 'SERVICE';

export type Product = {
  id: string;
  workspaceId: string;
  sku: string;
  name: string;
  categoryId: string;
  price: number;
  stock: number;
  itemType: ItemType; // PRODUCT has inventory, SERVICE (laundry, repair) has no stock deduction
  status: string;
};

export type Category = {
  id: string;
  workspaceId: string;
  name: string;
};

const DEFAULT_CATEGORIES: Category[] = [
  { id: 'cat-1', workspaceId: 'ws-1', name: 'Coffee' },
  { id: 'cat-2', workspaceId: 'ws-1', name: 'Non Coffee' },
  { id: 'cat-3', workspaceId: 'ws-1', name: 'Food' },
  { id: 'cat-4', workspaceId: 'ws-1', name: 'Snack' },
  { id: 'cat-5', workspaceId: 'ws-2', name: 'Layanan Laundry' },
];

const DEFAULT_PRODUCTS: Product[] = [
  { id: 'prod-1', workspaceId: 'ws-1', sku: 'KOP-001', name: 'Es Kopi Susu', categoryId: 'cat-1', price: 18000, stock: 120, itemType: 'PRODUCT', status: 'Active' },
  { id: 'prod-2', workspaceId: 'ws-1', sku: 'KOP-002', name: 'Americano', categoryId: 'cat-1', price: 15000, stock: 50, itemType: 'PRODUCT', status: 'Active' },
  { id: 'prod-3', workspaceId: 'ws-1', sku: 'NON-001', name: 'Matcha Latte', categoryId: 'cat-2', price: 22000, stock: 40, itemType: 'PRODUCT', status: 'Active' },
  { id: 'prod-4', workspaceId: 'ws-1', sku: 'FOD-001', name: 'Roti Bakar', categoryId: 'cat-3', price: 15000, stock: 30, itemType: 'PRODUCT', status: 'Active' },
  { id: 'prod-5', workspaceId: 'ws-2', sku: 'LND-001', name: 'Cuci Kering Setrika (per Kg)', categoryId: 'cat-5', price: 7000, stock: 0, itemType: 'SERVICE', status: 'Active' },
];

export const useProductStore = defineStore('product', {
  state: () => ({
    products: useLocalStorage<Product[]>('thekasir-products', DEFAULT_PRODUCTS),
    categories: useLocalStorage<Category[]>('thekasir-categories', DEFAULT_CATEGORIES),
  }),
  getters: {
    getProductsByWorkspace: (state) => {
      return (workspaceId: string) => {
        return state.products.filter(p => p.workspaceId === workspaceId);
      };
    },
    getCategoriesByWorkspace: (state) => {
      return (workspaceId: string) => {
        return state.categories.filter(c => c.workspaceId === workspaceId);
      };
    }
  },
  actions: {
    addProduct(product: Omit<Product, 'id'>) {
      const newId = `prod-${Date.now()}`;
      this.products.push({ ...product, id: newId });
    },
    updateProduct(id: string, updates: Partial<Product>) {
      const index = this.products.findIndex(p => p.id === id);
      if (index !== -1) {
        this.products[index] = { ...this.products[index], ...updates };
      }
    },
    deleteProduct(id: string) {
      this.products = this.products.filter(p => p.id !== id);
    },
    reduceStock(id: string, quantity: number) {
      const product = this.products.find(p => p.id === id);
      if (product && product.itemType === 'PRODUCT') {
        product.stock = Math.max(0, product.stock - quantity);
      }
    }
  }
});

import { defineStore } from 'pinia';
import { useLocalStorage } from '@vueuse/core';
import type { Product } from './useProductStore';

export type CartItem = {
  product: Product;
  quantity: number;
};

export const useCartStore = defineStore('cart', {
  state: () => ({
    items: useLocalStorage<CartItem[]>('thekasir-cart', []),
  }),
  getters: {
    totalItems: (state) => state.items.reduce((sum, item) => sum + item.quantity, 0),
    subtotal: (state) => state.items.reduce((sum, item) => sum + (item.product.price * item.quantity), 0),
    tax: (state) => {
      // 10% tax for example
      return Math.round(state.items.reduce((sum, item) => sum + (item.product.price * item.quantity), 0) * 0.1);
    },
    total: (state) => {
      const sub = state.items.reduce((sum, item) => sum + (item.product.price * item.quantity), 0);
      return sub + Math.round(sub * 0.1);
    }
  },
  actions: {
    addToCart(product: Product) {
      const existing = this.items.find(i => i.product.id === product.id);
      if (existing) {
        existing.quantity++;
      } else {
        this.items.push({ product, quantity: 1 });
      }
    },
    removeFromCart(productId: string) {
      this.items = this.items.filter(i => i.product.id !== productId);
    },
    updateQuantity(productId: string, quantity: number) {
      const item = this.items.find(i => i.product.id === productId);
      if (item) {
        if (quantity <= 0) {
          this.removeFromCart(productId);
        } else {
          item.quantity = quantity;
        }
      }
    },
    clearCart() {
      this.items = [];
    }
  }
});

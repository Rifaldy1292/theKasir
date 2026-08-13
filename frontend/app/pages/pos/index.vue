<script setup lang="ts">
import { ref, computed, onMounted, onUnmounted, watchEffect } from 'vue';
import { useProductStore } from '@/stores/useProductStore';
import { useCartStore } from '@/stores/useCartStore';
import { useTransactionStore } from '@/stores/useTransactionStore';
import { useCustomerStore } from '@/stores/useCustomerStore';
import { useShiftStore } from '@/stores/useShiftStore';
import { storeToRefs } from 'pinia';
import { useAppState } from '@/composables/useAppState';

import { Select, SelectTrigger, SelectValue, SelectContent, SelectItem } from '@/components/ui/select';

definePageMeta({
  layout: 'pos'
});

const productStore = useProductStore();
const cartStore = useCartStore();
const transactionStore = useTransactionStore();
const customerStore = useCustomerStore();
const shiftStore = useShiftStore();

const { currentRole, currentWorkspace } = useAppState();
const { items: cart, subtotal, tax, total } = storeToRefs(cartStore);

// Order Type State (Dine-in, Takeaway, Delivery)
const orderType = ref<'dine_in' | 'takeaway' | 'delivery'>('dine_in');

// Discount Percentage
const discountPercent = ref<number>(0);

// Shift Management State
const cashierName = computed(() => currentRole.value === 'cashier' ? 'Budi (Kasir)' : 'Rifky (Owner)');
const activeShift = computed(() => shiftStore.activeShift(currentWorkspace.value?.id || 'ws-1', cashierName.value));

const isOpenShiftModalOpen = ref(false);
const openingBalanceInput = ref(300000);

const isCloseShiftModalOpen = ref(false);
const actualCashInput = ref(0);
const closeShiftNotes = ref('');

// Customer Selection State
const selectedCustomerId = ref('CUST-001');
const workspaceCustomers = computed(() => customerStore.getCustomersByWorkspace(currentWorkspace.value?.id || 'ws-1'));
const selectedCustomer = computed(() => workspaceCustomers.value.find(c => c.id === selectedCustomerId.value) || workspaceCustomers.value[0]);

// Search & Filter
const activeCategory = ref<string | null>(null);
const searchQuery = ref('');
const searchInputRef = ref<HTMLInputElement | null>(null);

const workspaceProducts = computed(() => productStore.getProductsByWorkspace(currentWorkspace.value?.id || 'ws-1'));
const workspaceCategories = computed(() => productStore.getCategoriesByWorkspace(currentWorkspace.value?.id || 'ws-1'));

const filteredProducts = computed(() => {
  let result = workspaceProducts.value;
  if (activeCategory.value) {
    result = result.filter(p => p.categoryId === activeCategory.value);
  }
  if (searchQuery.value) {
    result = result.filter(p => 
      p.name.toLowerCase().includes(searchQuery.value.toLowerCase()) || 
      p.sku.toLowerCase().includes(searchQuery.value.toLowerCase())
    );
  }
  return result;
});

// Final Calculated Total with Discount
const discountAmount = computed(() => Math.round((subtotal.value * discountPercent.value) / 100));
const grandTotal = computed(() => Math.max(0, total.value - discountAmount.value));

// Payment Modal State
const isPaymentModalOpen = ref(false);
const paymentMethod = ref<'cash' | 'qris' | 'debit'>('qris');
const cashGiven = ref<number>(0);
const lastCompletedTransaction = ref<any>(null);
const isSuccessModalOpen = ref(false);

const changeAmount = computed(() => {
  if (paymentMethod.value !== 'cash') return 0;
  return Math.max(0, cashGiven.value - grandTotal.value);
});

// Quick Cash Preset Buttons
const setCashPreset = (amount: number) => {
  cashGiven.value = amount;
};

// Open Shift Action
const handleOpenShift = () => {
  shiftStore.openShift(currentWorkspace.value?.id || 'ws-1', cashierName.value, openingBalanceInput.value);
  isOpenShiftModalOpen.value = false;
};

// Close Shift Action
const handleCloseShift = () => {
  if (!activeShift.value) return;
  shiftStore.closeShift(activeShift.value.id, actualCashInput.value, closeShiftNotes.value);
  isCloseShiftModalOpen.value = false;
  alert('Shift berhasil ditutup!');
};

const openPayment = () => {
  if (cart.value.length === 0) return;
  if (!activeShift.value) {
    isOpenShiftModalOpen.value = true;
    return;
  }
  cashGiven.value = grandTotal.value;
  isPaymentModalOpen.value = true;
};

const confirmPayment = () => {
  if (cart.value.length === 0) return;
  if (paymentMethod.value === 'cash' && cashGiven.value < grandTotal.value) {
    alert('Uang pembayaran kurang dari total tagihan!');
    return;
  }

  const currentWsId = currentWorkspace.value?.id || 'ws-1';

  // Record Transaction
  const newTrx = {
    workspaceId: currentWsId,
    shiftId: activeShift.value?.id,
    customerId: selectedCustomer.value?.id,
    customerName: selectedCustomer.value?.name,
    cashier: cashierName.value,
    items: cartStore.totalItems,
    total: grandTotal.value,
    status: 'Paid',
    details: cart.value.map(item => ({
      productId: item.product.id,
      name: item.product.name,
      price: item.product.price,
      quantity: item.quantity
    }))
  };

  transactionStore.addTransaction(newTrx);

  // Record Cash Sale in Shift if cash payment
  if (paymentMethod.value === 'cash' && activeShift.value) {
    shiftStore.recordCashSale(activeShift.value.id, grandTotal.value);
  }

  // Update Customer Spending
  if (selectedCustomer.value) {
    customerStore.recordCustomerPurchase(selectedCustomer.value.id, grandTotal.value);
  }

  // Reduce stock for physical products only
  cart.value.forEach(item => {
    productStore.reduceStock(item.product.id, item.quantity);
  });

  lastCompletedTransaction.value = {
    ...newTrx,
    paymentMethod: paymentMethod.value,
    cashGiven: cashGiven.value,
    changeAmount: changeAmount.value,
    orderType: orderType.value,
    discountAmount: discountAmount.value,
    date: new Date().toISOString()
  };

  isPaymentModalOpen.value = false;
  isSuccessModalOpen.value = true;
  cartStore.clearCart();
  discountPercent.value = 0;
};

const startNewOrder = () => {
  isSuccessModalOpen.value = false;
  lastCompletedTransaction.value = null;
};

// Keyboard Shortcuts (F2: Search, F9: Checkout)
const handleKeyDown = (e: KeyboardEvent) => {
  if (e.key === 'F2') {
    e.preventDefault();
    searchInputRef.value?.focus();
  } else if (e.key === 'F9') {
    e.preventDefault();
    openPayment();
  }
};

onMounted(() => {
  window.addEventListener('keydown', handleKeyDown);
  if (!activeShift.value) {
    isOpenShiftModalOpen.value = true;
  }
});

onUnmounted(() => {
  window.removeEventListener('keydown', handleKeyDown);
});
</script>

<template>
  <div class="flex-1 flex overflow-hidden font-sans">
    <!-- Product Grid Area (Left Panel) -->
    <div class="flex-1 flex flex-col bg-zinc-100 dark:bg-zinc-950 overflow-hidden">
      <!-- Search & Filter Bar -->
      <div class="p-4 bg-white/90 dark:bg-zinc-900/90 backdrop-blur-md border-b border-zinc-200/80 dark:border-zinc-800/80 flex items-center justify-between gap-4 shrink-0 shadow-sm z-10">
        <div class="relative flex-1 max-w-lg">
          <svg class="absolute left-3.5 top-1/2 -translate-y-1/2 text-zinc-400" xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="11" cy="11" r="8"/><path d="m21 21-4.3-4.3"/></svg>
          <input 
            ref="searchInputRef"
            v-model="searchQuery"
            type="text" 
            placeholder="Cari produk nama atau SKU (Tekan F2)..." 
            class="w-full pl-10 pr-12 py-2.5 bg-zinc-100 dark:bg-zinc-800/80 border border-zinc-200 dark:border-zinc-700/60 focus:bg-white dark:focus:bg-zinc-900 border-transparent focus:border-emerald-500 rounded-xl text-zinc-900 dark:text-zinc-100 focus:outline-none focus:ring-4 focus:ring-emerald-500/10 transition-all text-sm font-medium"
          />
          <span class="absolute right-3 top-1/2 -translate-y-1/2 text-[10px] bg-zinc-200/80 dark:bg-zinc-700 text-zinc-600 dark:text-zinc-300 px-1.5 py-0.5 rounded font-mono font-bold">F2</span>
        </div>

        <!-- Category Pills -->
        <div class="flex gap-2 overflow-x-auto pb-1 hide-scrollbar">
          <button 
            @click="activeCategory = null"
            class="px-4 py-2 rounded-xl font-semibold text-xs whitespace-nowrap transition-all duration-200 flex items-center gap-1.5"
            :class="activeCategory === null 
              ? 'bg-emerald-600 text-white shadow-md shadow-emerald-600/20' 
              : 'bg-white dark:bg-zinc-800 text-zinc-600 dark:text-zinc-400 border border-zinc-200 dark:border-zinc-700/70 hover:bg-zinc-50 dark:hover:bg-zinc-700'"
          >
            <span>✨</span> Semua Items
          </button>
          <button 
            v-for="cat in workspaceCategories" 
            :key="cat.id"
            @click="activeCategory = cat.id"
            class="px-4 py-2 rounded-xl font-semibold text-xs whitespace-nowrap transition-all duration-200 flex items-center gap-1.5"
            :class="activeCategory === cat.id 
              ? 'bg-emerald-600 text-white shadow-md shadow-emerald-600/20' 
              : 'bg-white dark:bg-zinc-800 text-zinc-600 dark:text-zinc-400 border border-zinc-200 dark:border-zinc-700/70 hover:bg-zinc-50 dark:hover:bg-zinc-700'"
          >
            <span>{{ cat.name.includes('Coffee') ? '☕' : cat.name.includes('Food') ? '🥪' : cat.name.includes('Snack') ? '🍰' : '📦' }}</span>
            {{ cat.name }}
          </button>
        </div>
      </div>

      <!-- Products Grid -->
      <div class="flex-1 overflow-y-auto p-5">
        <div class="grid grid-cols-2 sm:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 gap-4">
          <button 
            v-for="product in filteredProducts" 
            :key="product.id"
            @click="cartStore.addToCart(product)"
            class="group bg-white dark:bg-zinc-900/90 rounded-2xl border border-zinc-200/80 dark:border-zinc-800/80 overflow-hidden hover:border-emerald-500/80 hover:ring-4 hover:ring-emerald-500/10 hover:shadow-xl hover:-translate-y-1 text-left transition-all duration-200 active:scale-95 flex flex-col justify-between"
          >
            <div>
              <div class="aspect-[4/3] bg-gradient-to-br from-zinc-100 to-zinc-200/60 dark:from-zinc-800 dark:to-zinc-900 flex items-center justify-center text-4xl group-hover:scale-110 transition-transform duration-300 relative overflow-hidden">
                <div class="absolute inset-0 bg-emerald-500/5 group-hover:bg-emerald-500/10 transition-colors"></div>
                {{ product.itemType === 'SERVICE' ? '🧺' : product.name.includes('Kopi') || product.name.includes('Coffee') ? '☕' : product.name.includes('Matcha') ? '🍵' : '🥪' }}
                
                <div class="absolute top-2 right-2">
                  <span 
                    class="px-2 py-0.5 rounded-full text-[10px] font-bold shadow-sm backdrop-blur-md"
                    :class="product.itemType === 'SERVICE' 
                      ? 'bg-blue-500/90 text-white' 
                      : product.stock <= 5 
                        ? 'bg-amber-500/90 text-white' 
                        : 'bg-emerald-500/90 text-white'"
                  >
                    {{ product.itemType === 'SERVICE' ? 'Jasa' : `${product.stock} pcs` }}
                  </span>
                </div>
              </div>

              <div class="p-3.5 space-y-1">
                <h3 class="font-bold text-zinc-900 dark:text-zinc-100 line-clamp-2 leading-tight text-sm group-hover:text-emerald-600 dark:group-hover:text-emerald-400 transition-colors">
                  {{ product.name }}
                </h3>
                <p class="text-[11px] text-zinc-400 font-mono tracking-wider">{{ product.sku }}</p>
              </div>
            </div>

            <div class="p-3.5 pt-0 flex items-center justify-between">
              <span class="text-xs text-zinc-400 font-semibold">Harga</span>
              <span class="text-emerald-600 dark:text-emerald-400 font-extrabold text-sm bg-emerald-50 dark:bg-emerald-950/50 px-2.5 py-1 rounded-lg border border-emerald-200/50 dark:border-emerald-800/50">
                Rp {{ product.price.toLocaleString('id-ID') }}
              </span>
            </div>
          </button>
        </div>
      </div>
    </div>

    <!-- Cart & Order Terminal Panel (Right Panel) -->
    <div class="w-[26rem] bg-white dark:bg-zinc-900 border-l border-zinc-200/80 dark:border-zinc-800/80 flex flex-col shrink-0 shadow-2xl z-20 relative">
      <!-- Order Type Tabs (Dine-In, Takeaway, Delivery) -->
      <div class="p-4 bg-zinc-50 dark:bg-zinc-950/60 border-b border-zinc-200 dark:border-zinc-800 space-y-3">
        <div class="grid grid-cols-3 gap-1.5 bg-zinc-200/60 dark:bg-zinc-800/80 p-1 rounded-xl">
          <button 
            @click="orderType = 'dine_in'"
            class="py-1.5 rounded-lg text-xs font-bold transition-all flex items-center justify-center gap-1"
            :class="orderType === 'dine_in' ? 'bg-white dark:bg-zinc-700 text-zinc-900 dark:text-white shadow-sm' : 'text-zinc-500 hover:text-zinc-800 dark:hover:text-zinc-200'"
          >
            <span>🍽️</span> Dine-In
          </button>
          <button 
            @click="orderType = 'takeaway'"
            class="py-1.5 rounded-lg text-xs font-bold transition-all flex items-center justify-center gap-1"
            :class="orderType === 'takeaway' ? 'bg-white dark:bg-zinc-700 text-zinc-900 dark:text-white shadow-sm' : 'text-zinc-500 hover:text-zinc-800 dark:hover:text-zinc-200'"
          >
            <span>🛍️</span> Takeaway
          </button>
          <button 
            @click="orderType = 'delivery'"
            class="py-1.5 rounded-lg text-xs font-bold transition-all flex items-center justify-center gap-1"
            :class="orderType === 'delivery' ? 'bg-white dark:bg-zinc-700 text-zinc-900 dark:text-white shadow-sm' : 'text-zinc-500 hover:text-zinc-800 dark:hover:text-zinc-200'"
          >
            <span>🛵</span> Delivery
          </button>
        </div>

        <!-- ClientOnly wrapped Shadcn UI Customer Select -->
        <div class="flex items-center justify-between gap-2 pt-1">
          <div class="flex-1">
            <ClientOnly>
              <Select v-model="selectedCustomerId">
                <SelectTrigger class="h-9">
                  <SelectValue placeholder="Pilih Pelanggan" />
                </SelectTrigger>
                <SelectContent>
                  <SelectItem v-for="cust in workspaceCustomers" :key="cust.id" :value="cust.id">
                    👤 {{ cust.name }} ({{ cust.phone }})
                  </SelectItem>
                </SelectContent>
              </Select>

              <template #fallback>
                <div class="h-9 w-full bg-zinc-100 dark:bg-zinc-800 rounded-xl flex items-center px-3 text-xs text-zinc-400">
                  Pelanggan Umum (Guest)
                </div>
              </template>
            </ClientOnly>
          </div>

          <button 
            v-if="activeShift" 
            @click="actualCashInput = activeShift.expectedCash; isCloseShiftModalOpen = true"
            class="px-3 py-2 bg-rose-50 dark:bg-rose-950/40 text-rose-600 dark:text-rose-400 border border-rose-200 dark:border-rose-800/60 rounded-xl text-xs font-bold hover:bg-rose-100 transition-colors whitespace-nowrap"
          >
            Tutup Shift
          </button>
        </div>
      </div>

      <!-- Cart Items List -->
      <div class="flex-1 overflow-y-auto p-4 space-y-3">
        <div v-if="cart.length === 0" class="h-full flex flex-col items-center justify-center text-zinc-400 text-center p-6 space-y-3">
          <div class="w-16 h-16 rounded-full bg-zinc-100 dark:bg-zinc-800 flex items-center justify-center text-3xl opacity-60">
            🛒
          </div>
          <div>
            <p class="font-bold text-zinc-800 dark:text-zinc-200 text-sm">Pesanan Kosong</p>
            <p class="text-xs text-zinc-400 mt-1">Klik item di sebelah kiri untuk menambah ke struk pesanan kasir</p>
          </div>
        </div>
        
        <div 
          v-for="item in cart" 
          :key="item.product.id" 
          class="flex items-center gap-3 p-3 bg-zinc-50 dark:bg-zinc-800/40 border border-zinc-200/60 dark:border-zinc-800 rounded-2xl transition-all hover:bg-white dark:hover:bg-zinc-800"
        >
          <div class="w-12 h-12 rounded-xl bg-gradient-to-br from-emerald-100 to-teal-100 dark:from-emerald-950 dark:to-teal-950 text-emerald-600 dark:text-emerald-300 flex items-center justify-center shrink-0 text-xl font-bold">
            {{ item.product.itemType === 'SERVICE' ? '🧺' : '☕' }}
          </div>

          <div class="flex-1 min-w-0">
            <div class="flex justify-between items-start">
              <h4 class="font-bold text-xs text-zinc-900 dark:text-zinc-100 truncate leading-snug">{{ item.product.name }}</h4>
              <button @click="cartStore.updateQuantity(item.product.id, 0)" class="text-zinc-400 hover:text-rose-500 transition-colors p-0.5">
                <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M18 6 6 18"/><path d="m6 6 12 12"/></svg>
              </button>
            </div>

            <div class="flex items-center justify-between mt-2">
              <p class="text-xs font-extrabold text-emerald-600 dark:text-emerald-400">
                Rp {{ (item.product.price * item.quantity).toLocaleString('id-ID') }}
              </p>

              <div class="flex items-center gap-2 bg-white dark:bg-zinc-900 border border-zinc-200 dark:border-zinc-700 rounded-xl px-2 py-1 shadow-xs">
                <button @click="cartStore.updateQuantity(item.product.id, item.quantity - 1)" class="w-5 h-5 flex items-center justify-center text-zinc-600 dark:text-zinc-300 hover:bg-zinc-100 dark:hover:bg-zinc-800 rounded transition-colors font-bold text-xs">-</button>
                <span class="text-xs font-extrabold w-4 text-center text-zinc-900 dark:text-zinc-100">{{ item.quantity }}</span>
                <button @click="cartStore.updateQuantity(item.product.id, item.quantity + 1)" class="w-5 h-5 flex items-center justify-center text-zinc-600 dark:text-zinc-300 hover:bg-zinc-100 dark:hover:bg-zinc-800 rounded transition-colors font-bold text-xs">+</button>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Quick Discount & Checkout Summary -->
      <div class="p-4 bg-zinc-50 dark:bg-zinc-950 border-t border-zinc-200/80 dark:border-zinc-800/80 space-y-3.5">
        <div class="flex items-center justify-between text-xs">
          <span class="font-bold text-zinc-600 dark:text-zinc-400">Diskon Produk:</span>
          <div class="flex gap-1.5">
            <button 
              v-for="d in [0, 5, 10, 15]" 
              :key="d"
              @click="discountPercent = d"
              class="px-2.5 py-1 rounded-lg font-bold text-[11px] transition-all"
              :class="discountPercent === d ? 'bg-emerald-600 text-white shadow-xs' : 'bg-zinc-200/70 dark:bg-zinc-800 text-zinc-600 dark:text-zinc-400 hover:bg-zinc-300'"
            >
              {{ d }}%
            </button>
          </div>
        </div>

        <div class="space-y-1.5 text-xs pt-1 border-t border-zinc-200/60 dark:border-zinc-800/60">
          <div class="flex justify-between text-zinc-500">
            <span>Subtotal</span>
            <span class="font-bold text-zinc-900 dark:text-zinc-100">Rp {{ subtotal.toLocaleString('id-ID') }}</span>
          </div>
          <div v-if="discountAmount > 0" class="flex justify-between text-emerald-600 dark:text-emerald-400 font-bold">
            <span>Diskon ({{ discountPercent }}%)</span>
            <span>- Rp {{ discountAmount.toLocaleString('id-ID') }}</span>
          </div>
          <div class="flex justify-between text-zinc-500">
            <span>Pajak PPN (11%)</span>
            <span class="font-bold text-zinc-900 dark:text-zinc-100">Rp {{ tax.toLocaleString('id-ID') }}</span>
          </div>
          <div class="flex justify-between items-center pt-2.5 border-t border-dashed border-zinc-300 dark:border-zinc-700">
            <span class="font-extrabold text-sm text-zinc-900 dark:text-white">TOTAL PEMBAYARAN</span>
            <span class="font-black text-2xl text-emerald-600 dark:text-emerald-400">Rp {{ grandTotal.toLocaleString('id-ID') }}</span>
          </div>
        </div>
        
        <button 
          @click="openPayment"
          :disabled="cart.length === 0"
          class="w-full py-4 rounded-2xl font-black text-base text-white transition-all duration-200 shadow-xl flex items-center justify-center gap-2 active:scale-98"
          :class="cart.length > 0 ? 'bg-gradient-to-r from-emerald-600 to-teal-600 hover:from-emerald-500 hover:to-teal-500 shadow-emerald-600/30' : 'bg-zinc-300 dark:bg-zinc-800 text-zinc-500 cursor-not-allowed'"
        >
          <span>BAYAR KASIR (Tekan F9)</span>
          <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><path d="m9 18 6-6-6-6"/></svg>
        </button>
      </div>
    </div>

    <!-- Modal Open Shift Requirement -->
    <div v-if="isOpenShiftModalOpen" class="fixed inset-0 z-50 bg-black/60 backdrop-blur-md flex items-center justify-center p-4">
      <div class="bg-white dark:bg-zinc-900 border border-zinc-200 dark:border-zinc-800 rounded-3xl p-7 w-full max-w-md shadow-2xl space-y-6">
        <div class="w-14 h-14 rounded-2xl bg-gradient-to-br from-emerald-500 to-teal-600 text-white flex items-center justify-center font-bold text-3xl shadow-lg shadow-emerald-500/30 mx-auto">
          🔓
        </div>

        <div class="text-center">
          <h3 class="text-2xl font-extrabold text-zinc-900 dark:text-zinc-100">Buka Shift Kasir</h3>
          <p class="text-xs text-zinc-500 dark:text-zinc-400 mt-1.5 max-w-xs mx-auto">
            Masukkan jumlah modal uang tunai laci kasir (Opening Balance) sebelum memulai operasional.
          </p>
        </div>

        <div class="space-y-3 bg-zinc-50 dark:bg-zinc-800/50 p-4 rounded-2xl border border-zinc-200/60 dark:border-zinc-700/60">
          <div class="flex justify-between text-xs text-zinc-500">
            <span>Nama Petugas Kasir:</span>
            <span class="font-bold text-zinc-900 dark:text-zinc-100">{{ cashierName }}</span>
          </div>
          <div>
            <label class="block text-xs font-bold text-zinc-600 dark:text-zinc-400 mb-1.5">Modal Uang Tunai / Laci (Rp)</label>
            <input 
              v-model.number="openingBalanceInput"
              type="number" 
              min="0"
              required
              class="w-full px-4 py-3 bg-white dark:bg-zinc-900 border border-zinc-300 dark:border-zinc-700 rounded-xl text-xl font-black text-emerald-600 focus:outline-none focus:ring-4 focus:ring-emerald-500/20"
            />
          </div>
        </div>

        <button 
          @click="handleOpenShift"
          class="w-full py-3.5 bg-gradient-to-r from-emerald-600 to-teal-600 hover:from-emerald-500 hover:to-teal-500 text-white font-extrabold rounded-2xl transition-all shadow-lg shadow-emerald-600/30 text-sm"
        >
          Mulai & Buka Shift Sekarang
        </button>
      </div>
    </div>

    <!-- Modal Close Shift -->
    <div v-if="isCloseShiftModalOpen && activeShift" class="fixed inset-0 z-50 bg-black/60 backdrop-blur-md flex items-center justify-center p-4">
      <div class="bg-white dark:bg-zinc-900 border border-zinc-200 dark:border-zinc-800 rounded-3xl p-7 w-full max-w-md shadow-2xl space-y-5">
        <h3 class="text-xl font-extrabold text-zinc-900 dark:text-zinc-100 border-b border-zinc-200 dark:border-zinc-800 pb-3">
          Tutup Shift Kasir (#{{ activeShift.id }})
        </h3>

        <div class="space-y-2 text-xs bg-zinc-50 dark:bg-zinc-800/50 p-4 rounded-2xl border border-zinc-200/60 dark:border-zinc-700/60">
          <div class="flex justify-between text-zinc-500">
            <span>Modal Awal Laci:</span>
            <span class="font-bold text-zinc-900 dark:text-zinc-100">Rp {{ activeShift.openingBalance.toLocaleString('id-ID') }}</span>
          </div>
          <div class="flex justify-between text-zinc-500">
            <span>Total Penjualan Tunai:</span>
            <span class="font-bold text-zinc-900 dark:text-zinc-100">Rp {{ activeShift.cashSales.toLocaleString('id-ID') }}</span>
          </div>
          <div class="flex justify-between text-zinc-500 pt-2 border-t border-zinc-200 dark:border-zinc-700 font-extrabold">
            <span class="text-zinc-900 dark:text-zinc-100">Total Uang Diharapkan:</span>
            <span class="text-emerald-600 text-sm">Rp {{ activeShift.expectedCash.toLocaleString('id-ID') }}</span>
          </div>
        </div>

        <div class="space-y-3">
          <div>
            <label class="block text-xs font-bold text-zinc-600 dark:text-zinc-400 mb-1">Uang Tunai Fisik di Kasir (Rp)</label>
            <input 
              v-model.number="actualCashInput"
              type="number" 
              min="0"
              class="w-full px-4 py-2.5 bg-zinc-50 dark:bg-zinc-800 border border-zinc-300 dark:border-zinc-700 rounded-xl text-lg font-black text-zinc-900 dark:text-zinc-100 focus:outline-none focus:ring-4 focus:ring-emerald-500/20"
            />
          </div>

          <div class="flex justify-between items-center text-xs font-bold p-3 rounded-xl" :class="actualCashInput - activeShift.expectedCash === 0 ? 'bg-emerald-50 text-emerald-700 border border-emerald-200' : 'bg-amber-50 text-amber-800 border border-amber-200'">
            <span>Selisih (Difference):</span>
            <span class="text-sm">Rp {{ (actualCashInput - activeShift.expectedCash).toLocaleString('id-ID') }}</span>
          </div>

          <div>
            <label class="block text-xs font-bold text-zinc-600 dark:text-zinc-400 mb-1">Catatan Shift (Opsional)</label>
            <textarea 
              v-model="closeShiftNotes"
              rows="2"
              placeholder="Catatan penutupan shift..."
              class="w-full px-3 py-2 bg-zinc-50 dark:bg-zinc-800 border border-zinc-200 dark:border-zinc-700 rounded-xl text-xs text-zinc-900 dark:text-zinc-100 focus:outline-none focus:ring-2 focus:ring-emerald-500/20"
            ></textarea>
          </div>
        </div>

        <div class="flex justify-end gap-2 pt-2">
          <button 
            type="button" 
            @click="isCloseShiftModalOpen = false"
            class="px-4 py-2.5 text-xs font-bold text-zinc-600 dark:text-zinc-400 hover:bg-zinc-100 dark:hover:bg-zinc-800 rounded-xl"
          >
            Batal
          </button>
          <button 
            type="button" 
            @click="handleCloseShift"
            class="px-6 py-2.5 text-xs bg-rose-600 hover:bg-rose-700 text-white font-extrabold rounded-xl shadow-lg shadow-rose-600/30"
          >
            Konfirmasi Tutup Shift
          </button>
        </div>
      </div>
    </div>

    <!-- Payment Choice Modal with Cash Presets -->
    <div v-if="isPaymentModalOpen" class="fixed inset-0 z-50 bg-black/60 backdrop-blur-md flex items-center justify-center p-4">
      <div class="bg-white dark:bg-zinc-900 border border-zinc-200 dark:border-zinc-800 rounded-3xl p-7 w-full max-w-lg shadow-2xl space-y-6">
        <div class="flex items-center justify-between border-b border-zinc-200 dark:border-zinc-800 pb-4">
          <h3 class="text-xl font-black text-zinc-900 dark:text-zinc-100">
            Pembayaran Kasir
          </h3>
          <span class="text-xs font-bold bg-emerald-100 text-emerald-700 px-3 py-1 rounded-full uppercase tracking-wider">
            {{ orderType }}
          </span>
        </div>

        <!-- Payment Method Tabs -->
        <div class="grid grid-cols-3 gap-3">
          <button 
            @click="paymentMethod = 'qris'"
            class="p-4 rounded-2xl border-2 flex flex-col items-center gap-2 transition-all text-xs font-extrabold"
            :class="paymentMethod === 'qris' ? 'border-emerald-500 bg-emerald-50 dark:bg-emerald-950/40 text-emerald-600 shadow-md' : 'border-zinc-200 dark:border-zinc-800 text-zinc-600 dark:text-zinc-400 hover:bg-zinc-50'"
          >
            <span class="text-3xl">📱</span>
            QRIS / GoPay
          </button>
          <button 
            @click="paymentMethod = 'cash'"
            class="p-4 rounded-2xl border-2 flex flex-col items-center gap-2 transition-all text-xs font-extrabold"
            :class="paymentMethod === 'cash' ? 'border-emerald-500 bg-emerald-50 dark:bg-emerald-950/40 text-emerald-600 shadow-md' : 'border-zinc-200 dark:border-zinc-800 text-zinc-600 dark:text-zinc-400 hover:bg-zinc-50'"
          >
            <span class="text-3xl">💵</span>
            Uang Tunai
          </button>
          <button 
            @click="paymentMethod = 'debit'"
            class="p-4 rounded-2xl border-2 flex flex-col items-center gap-2 transition-all text-xs font-extrabold"
            :class="paymentMethod === 'debit' ? 'border-emerald-500 bg-emerald-50 dark:bg-emerald-950/40 text-emerald-600 shadow-md' : 'border-zinc-200 dark:border-zinc-800 text-zinc-600 dark:text-zinc-400 hover:bg-zinc-50'"
          >
            <span class="text-3xl">💳</span>
            Kartu Debit
          </button>
        </div>

        <!-- Calculation Box -->
        <div class="p-5 bg-zinc-50 dark:bg-zinc-800/50 rounded-2xl space-y-3 border border-zinc-200/60 dark:border-zinc-700/60 text-sm">
          <div class="flex justify-between text-zinc-500">
            <span>Pelanggan:</span>
            <span class="font-bold text-zinc-900 dark:text-zinc-100">{{ selectedCustomer?.name }}</span>
          </div>
          <div class="flex justify-between items-center">
            <span class="text-zinc-500 font-bold">Total Tagihan:</span>
            <span class="font-black text-2xl text-emerald-600 dark:text-emerald-400">Rp {{ grandTotal.toLocaleString('id-ID') }}</span>
          </div>

          <!-- Cash Nominal Presets -->
          <div v-if="paymentMethod === 'cash'" class="pt-3 border-t border-zinc-200 dark:border-zinc-700 space-y-3">
            <div>
              <label class="block text-xs font-bold text-zinc-600 dark:text-zinc-400 mb-1.5">Nominal Uang Tunai Diterima</label>
              <input 
                v-model.number="cashGiven"
                type="number" 
                class="w-full px-4 py-3 bg-white dark:bg-zinc-900 border border-zinc-300 dark:border-zinc-700 rounded-xl font-black text-emerald-600 text-xl focus:outline-none focus:ring-4 focus:ring-emerald-500/20"
              />
            </div>

            <!-- Quick Cash Nominal Buttons -->
            <div class="grid grid-cols-4 gap-2">
              <button 
                @click="setCashPreset(grandTotal)"
                class="py-2 px-1 bg-white dark:bg-zinc-800 border border-zinc-200 dark:border-zinc-700 hover:border-emerald-500 rounded-xl text-xs font-bold text-zinc-800 dark:text-zinc-200 transition-colors"
              >
                Uang Pas
              </button>
              <button 
                @click="setCashPreset(50000)"
                class="py-2 px-1 bg-white dark:bg-zinc-800 border border-zinc-200 dark:border-zinc-700 hover:border-emerald-500 rounded-xl text-xs font-bold text-zinc-800 dark:text-zinc-200 transition-colors"
              >
                50.000
              </button>
              <button 
                @click="setCashPreset(100000)"
                class="py-2 px-1 bg-white dark:bg-zinc-800 border border-zinc-200 dark:border-zinc-700 hover:border-emerald-500 rounded-xl text-xs font-bold text-zinc-800 dark:text-zinc-200 transition-colors"
              >
                100.000
              </button>
              <button 
                @click="setCashPreset(200000)"
                class="py-2 px-1 bg-white dark:bg-zinc-800 border border-zinc-200 dark:border-zinc-700 hover:border-emerald-500 rounded-xl text-xs font-bold text-zinc-800 dark:text-zinc-200 transition-colors"
              >
                200.000
              </button>
            </div>

            <div class="flex justify-between items-center text-sm font-extrabold pt-1">
              <span>Uang Kembalian:</span>
              <span class="text-emerald-600 dark:text-emerald-400 text-xl font-black">Rp {{ changeAmount.toLocaleString('id-ID') }}</span>
            </div>
          </div>
        </div>

        <div class="flex justify-end gap-3 pt-2">
          <button 
            type="button" 
            @click="isPaymentModalOpen = false"
            class="px-5 py-3 text-sm font-bold text-zinc-600 dark:text-zinc-400 hover:bg-zinc-100 dark:hover:bg-zinc-800 rounded-2xl transition-colors"
          >
            Batal
          </button>
          <button 
            type="button" 
            @click="confirmPayment"
            class="px-7 py-3 text-sm bg-gradient-to-r from-emerald-600 to-teal-600 hover:from-emerald-500 hover:to-teal-500 text-white font-black rounded-2xl transition-all shadow-lg shadow-emerald-600/30"
          >
            SELESAIKAN PEMBAYARAN
          </button>
        </div>
      </div>
    </div>

    <!-- Success & Receipt Modal -->
    <div v-if="isSuccessModalOpen" class="fixed inset-0 z-50 bg-black/60 backdrop-blur-md flex items-center justify-center p-4">
      <div class="bg-white dark:bg-zinc-900 border border-zinc-200 dark:border-zinc-800 rounded-3xl p-7 w-full max-w-sm shadow-2xl space-y-5 text-center">
        <div class="w-16 h-16 bg-gradient-to-br from-emerald-500 to-teal-600 text-white rounded-full flex items-center justify-center mx-auto text-3xl font-black shadow-lg shadow-emerald-500/30 animate-bounce">
          ✓
        </div>

        <div>
          <h3 class="text-2xl font-black text-zinc-900 dark:text-zinc-100">Pembayaran Berhasil!</h3>
          <p class="text-xs text-zinc-400 mt-1">Transaksi telah tersimpan dalam sistem kasir.</p>
        </div>
        
        <div v-if="lastCompletedTransaction" class="p-4 bg-zinc-50 dark:bg-zinc-800/50 rounded-2xl text-left space-y-2 text-xs border border-zinc-200/60 dark:border-zinc-700/60">
          <div class="flex justify-between text-zinc-500">
            <span>Pelanggan:</span>
            <span class="font-bold text-zinc-900 dark:text-zinc-100">{{ lastCompletedTransaction.customerName }}</span>
          </div>
          <div class="flex justify-between text-zinc-500">
            <span>Tipe Pesanan:</span>
            <span class="font-bold text-zinc-900 dark:text-zinc-100 uppercase">{{ lastCompletedTransaction.orderType }}</span>
          </div>
          <div class="flex justify-between text-zinc-500">
            <span>Metode Bayar:</span>
            <span class="font-bold text-zinc-900 dark:text-zinc-100 uppercase">{{ lastCompletedTransaction.paymentMethod }}</span>
          </div>
          <div class="flex justify-between text-zinc-500 pt-1 border-t border-zinc-200 dark:border-zinc-700">
            <span>Total Tagihan:</span>
            <span class="font-bold text-zinc-900 dark:text-zinc-100">Rp {{ lastCompletedTransaction.total.toLocaleString('id-ID') }}</span>
          </div>
          <div v-if="lastCompletedTransaction.paymentMethod === 'cash'" class="flex justify-between text-zinc-500">
            <span>Kembalian:</span>
            <span class="font-bold text-emerald-600 text-sm">Rp {{ lastCompletedTransaction.changeAmount.toLocaleString('id-ID') }}</span>
          </div>
        </div>

        <button 
          @click="startNewOrder"
          class="w-full py-3.5 bg-gradient-to-r from-emerald-600 to-teal-600 hover:from-emerald-500 hover:to-teal-500 text-white font-extrabold text-sm rounded-2xl transition-all shadow-lg shadow-emerald-600/30"
        >
          Mulai Transaksi Baru
        </button>
      </div>
    </div>
  </div>
</template>

<style scoped>
.hide-scrollbar::-webkit-scrollbar {
  display: none;
}
.hide-scrollbar {
  -ms-overflow-style: none;
  scrollbar-width: none;
}
</style>

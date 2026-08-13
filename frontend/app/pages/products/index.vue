<script setup lang="ts">
import { ref } from 'vue';
import { useProductStore, type Product, type ItemType } from '@/stores/useProductStore';
import { Select, SelectTrigger, SelectValue, SelectContent, SelectItem } from '@/components/ui/select';

const productStore = useProductStore();

const isModalOpen = ref(false);
const editingId = ref<string | null>(null);

const formName = ref('');
const formSku = ref('');
const formCategoryId = ref('cat-1');
const formPrice = ref<number>(0);
const formStock = ref<number>(0);
const formItemType = ref<ItemType>('PRODUCT');
const formStatus = ref('Active');

const openAddModal = () => {
  editingId.value = null;
  formName.value = '';
  formSku.value = `SKU-${Math.floor(1000 + Math.random() * 9000)}`;
  formCategoryId.value = productStore.categories[0]?.id || 'cat-1';
  formPrice.value = 15000;
  formStock.value = 50;
  formItemType.value = 'PRODUCT';
  formStatus.value = 'Active';
  isModalOpen.value = true;
};

const openEditModal = (product: Product) => {
  editingId.value = product.id;
  formName.value = product.name;
  formSku.value = product.sku;
  formCategoryId.value = product.categoryId;
  formPrice.value = product.price;
  formStock.value = product.stock;
  formItemType.value = product.itemType || 'PRODUCT';
  formStatus.value = product.status;
  isModalOpen.value = true;
};

const saveProduct = () => {
  if (!formName.value.trim() || !formSku.value.trim()) return;

  if (editingId.value) {
    productStore.updateProduct(editingId.value, {
      name: formName.value,
      sku: formSku.value,
      categoryId: formCategoryId.value,
      price: formPrice.value,
      stock: formStock.value,
      itemType: formItemType.value,
      status: formStatus.value,
    });
  } else {
    productStore.addProduct({
      workspaceId: 'ws-1',
      name: formName.value,
      sku: formSku.value,
      categoryId: formCategoryId.value,
      price: formPrice.value,
      stock: formStock.value,
      itemType: formItemType.value,
      status: formStatus.value,
    });
  }

  isModalOpen.value = false;
};

const confirmDelete = (id: string, name: string) => {
  if (confirm(`Apakah Anda yakin ingin menghapus produk "${name}"?`)) {
    productStore.deleteProduct(id);
  }
};
</script>

<template>
  <div class="space-y-6">
    <div class="flex items-center justify-between">
      <div>
        <h1 class="text-2xl font-bold tracking-tight text-zinc-900 dark:text-zinc-100">Produk & Inventaris</h1>
        <p class="text-sm text-zinc-500 dark:text-zinc-400 mt-1">
          Kelola katalog produk, harga, dan ketersediaan stok barang.
        </p>
      </div>
      <button 
        @click="openAddModal"
        class="bg-emerald-600 hover:bg-emerald-700 text-white px-4 py-2 rounded-xl font-bold text-sm transition-colors shadow-xs flex items-center gap-2"
      >
        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M5 12h14"/><path d="M12 5v14"/></svg>
        Tambah Produk
      </button>
    </div>

    <div class="bg-white dark:bg-zinc-900 border border-zinc-200 dark:border-zinc-800 rounded-2xl shadow-xs overflow-hidden">
      <div class="overflow-x-auto">
        <table class="w-full text-left text-sm whitespace-nowrap">
          <thead class="bg-zinc-50 dark:bg-zinc-900/50 border-b border-zinc-200 dark:border-zinc-800 text-zinc-500 dark:text-zinc-400 font-bold">
            <tr>
              <th class="px-6 py-4">Nama Produk</th>
              <th class="px-6 py-4">SKU</th>
              <th class="px-6 py-4">Tipe Item</th>
              <th class="px-6 py-4">Harga</th>
              <th class="px-6 py-4">Stok</th>
              <th class="px-6 py-4">Status</th>
              <th class="px-6 py-4 text-right">Aksi</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-zinc-200 dark:divide-zinc-800">
            <tr v-for="product in productStore.products" :key="product.id" class="hover:bg-zinc-50/50 dark:hover:bg-zinc-800/20 transition-colors">
              <td class="px-6 py-4">
                <div class="font-bold text-zinc-900 dark:text-zinc-100">{{ product.name }}</div>
              </td>
              <td class="px-6 py-4 text-zinc-500 font-mono text-xs">{{ product.sku }}</td>
              <td class="px-6 py-4">
                <span class="px-2 py-0.5 rounded text-xs font-semibold" :class="product.itemType === 'SERVICE' ? 'bg-blue-100 text-blue-700' : 'bg-zinc-100 text-zinc-700'">
                  {{ product.itemType === 'SERVICE' ? 'Jasa (Laundry/Repair)' : 'Produk (Fisik)' }}
                </span>
              </td>
              <td class="px-6 py-4 text-zinc-900 dark:text-zinc-100 font-extrabold">Rp {{ product.price.toLocaleString('id-ID') }}</td>
              <td class="px-6 py-4">
                <span :class="product.itemType === 'SERVICE' ? 'text-zinc-400' : product.stock <= 5 ? 'text-amber-600 font-bold' : 'text-zinc-600 dark:text-zinc-400'">
                  {{ product.itemType === 'SERVICE' ? 'Tanpa Stok' : `${product.stock} unit` }}
                </span>
              </td>
              <td class="px-6 py-4">
                <span class="px-2.5 py-1 rounded-full text-xs font-bold bg-emerald-100 text-emerald-700 dark:bg-emerald-500/20 dark:text-emerald-400 border border-emerald-200 dark:border-emerald-500/20">
                  {{ product.status }}
                </span>
              </td>
              <td class="px-6 py-4 text-right space-x-2">
                <button 
                  @click="openEditModal(product)"
                  class="text-zinc-400 hover:text-emerald-600 transition-colors p-1"
                  title="Edit Produk"
                >
                  <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M17 3a2.85 2.83 0 1 1 4 4L7.5 20.5 2 22l1.5-5.5Z"/><path d="m15 5 4 4"/></svg>
                </button>
                <button 
                  @click="confirmDelete(product.id, product.name)"
                  class="text-zinc-400 hover:text-rose-600 transition-colors p-1"
                  title="Hapus Produk"
                >
                  <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M3 6h18"/><path d="M19 6v14c0 1-1 2-2 2H7c-1 0-2-1-2-2V6"/><path d="M8 6V4c0-1 1-2 2-2h4c1 0 2 1 2 2v2"/></svg>
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- Modal Form Add/Edit with Shadcn Select -->
    <div v-if="isModalOpen" class="fixed inset-0 z-50 bg-black/50 backdrop-blur-sm flex items-center justify-center p-4">
      <div class="bg-white dark:bg-zinc-900 border border-zinc-200 dark:border-zinc-800 rounded-3xl p-6 w-full max-w-md shadow-2xl space-y-4">
        <h3 class="text-lg font-bold text-zinc-900 dark:text-zinc-100">
          {{ editingId ? 'Edit Produk' : 'Tambah Produk Baru' }}
        </h3>

        <form @submit.prevent="saveProduct" class="space-y-4">
          <div>
            <label class="block text-xs font-bold text-zinc-600 dark:text-zinc-400 mb-1">Nama Produk</label>
            <input 
              v-model="formName" 
              type="text" 
              required 
              class="w-full px-3 py-2 bg-zinc-50 dark:bg-zinc-800 border border-zinc-200 dark:border-zinc-700 rounded-xl text-sm text-zinc-900 dark:text-zinc-100 focus:outline-none focus:ring-2 focus:ring-emerald-500/20 font-medium"
            />
          </div>

          <div class="grid grid-cols-2 gap-3">
            <div>
              <label class="block text-xs font-bold text-zinc-600 dark:text-zinc-400 mb-1">SKU</label>
              <input 
                v-model="formSku" 
                type="text" 
                required 
                class="w-full px-3 py-2 bg-zinc-50 dark:bg-zinc-800 border border-zinc-200 dark:border-zinc-700 rounded-xl text-sm text-zinc-900 dark:text-zinc-100 focus:outline-none focus:ring-2 focus:ring-emerald-500/20 font-mono"
              />
            </div>

            <div>
              <label class="block text-xs font-bold text-zinc-600 dark:text-zinc-400 mb-1">Kategori</label>
              <Select v-model="formCategoryId">
                <SelectTrigger class="h-10">
                  <SelectValue placeholder="Pilih Kategori" />
                </SelectTrigger>
                <SelectContent>
                  <SelectItem v-for="cat in productStore.categories" :key="cat.id" :value="cat.id">
                    {{ cat.name }}
                  </SelectItem>
                </SelectContent>
              </Select>
            </div>
          </div>

          <div class="grid grid-cols-2 gap-3">
            <div>
              <label class="block text-xs font-bold text-zinc-600 dark:text-zinc-400 mb-1">Tipe Item</label>
              <Select v-model="formItemType">
                <SelectTrigger class="h-10">
                  <SelectValue placeholder="Tipe Item" />
                </SelectTrigger>
                <SelectContent>
                  <SelectItem value="PRODUCT">Produk (Fisik)</SelectItem>
                  <SelectItem value="SERVICE">Jasa (Laundry/Repair)</SelectItem>
                </SelectContent>
              </Select>
            </div>

            <div>
              <label class="block text-xs font-bold text-zinc-600 dark:text-zinc-400 mb-1">Harga (Rp)</label>
              <input 
                v-model.number="formPrice" 
                type="number" 
                min="0"
                required 
                class="w-full px-3 py-2 bg-zinc-50 dark:bg-zinc-800 border border-zinc-200 dark:border-zinc-700 rounded-xl text-sm text-zinc-900 dark:text-zinc-100 font-bold focus:outline-none focus:ring-2 focus:ring-emerald-500/20"
              />
            </div>
          </div>

          <div>
            <label class="block text-xs font-bold text-zinc-600 dark:text-zinc-400 mb-1">Stok Awal</label>
            <input 
              v-model.number="formStock" 
              type="number" 
              min="0"
              required 
              :disabled="formItemType === 'SERVICE'"
              class="w-full px-3 py-2 bg-zinc-50 dark:bg-zinc-800 border border-zinc-200 dark:border-zinc-700 rounded-xl text-sm text-zinc-900 dark:text-zinc-100 focus:outline-none focus:ring-2 focus:ring-emerald-500/20 disabled:opacity-50"
            />
          </div>

          <div class="flex justify-end gap-2 pt-4 border-t border-zinc-200 dark:border-zinc-800">
            <button 
              type="button" 
              @click="isModalOpen = false"
              class="px-4 py-2 text-xs font-bold text-zinc-600 dark:text-zinc-400 hover:bg-zinc-100 dark:hover:bg-zinc-800 rounded-xl transition-colors"
            >
              Batal
            </button>
            <button 
              type="submit" 
              class="px-5 py-2 text-xs bg-emerald-600 hover:bg-emerald-700 text-white font-extrabold rounded-xl transition-colors"
            >
              Simpan
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

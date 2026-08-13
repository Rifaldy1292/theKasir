<script setup lang="ts">
import { ref, computed } from 'vue';
import { useCustomerStore, type Customer } from '@/stores/useCustomerStore';
import { useAppState } from '@/composables/useAppState';
import { Table, TableHeader, TableBody, TableRow, TableHead, TableCell } from '@/components/ui/table';

const customerStore = useCustomerStore();
const { currentWorkspace } = useAppState();

const searchQuery = ref('');
const isModalOpen = ref(false);
const editingId = ref<string | null>(null);

const formName = ref('');
const formPhone = ref('');
const formEmail = ref('');
const formAddress = ref('');

const workspaceCustomers = computed(() => {
  const list = customerStore.getCustomersByWorkspace(currentWorkspace.value?.id || 'ws-1');
  if (!searchQuery.value) return list;
  return list.filter(c => 
    c.name.toLowerCase().includes(searchQuery.value.toLowerCase()) || 
    c.phone.includes(searchQuery.value)
  );
});

const openAddModal = () => {
  editingId.value = null;
  formName.value = '';
  formPhone.value = '';
  formEmail.value = '';
  formAddress.value = '';
  isModalOpen.value = true;
};

const openEditModal = (c: Customer) => {
  editingId.value = c.id;
  formName.value = c.name;
  formPhone.value = c.phone;
  formEmail.value = c.email || '';
  formAddress.value = c.address || '';
  isModalOpen.value = true;
};

const saveCustomer = () => {
  if (!formName.value.trim()) return;

  if (editingId.value) {
    customerStore.updateCustomer(editingId.value, {
      name: formName.value,
      phone: formPhone.value,
      email: formEmail.value,
      address: formAddress.value
    });
  } else {
    customerStore.addCustomer({
      workspaceId: currentWorkspace.value?.id || 'ws-1',
      name: formName.value,
      phone: formPhone.value,
      email: formEmail.value,
      address: formAddress.value
    });
  }

  isModalOpen.value = false;
};

const confirmDelete = (id: string, name: string) => {
  if (confirm(`Hapus pelanggan "${name}"?`)) {
    customerStore.deleteCustomer(id);
  }
};
</script>

<template>
  <div class="space-y-6 font-sans">
    <div class="flex flex-col sm:flex-row sm:items-center justify-between gap-4">
      <div>
        <h1 class="text-2xl font-bold tracking-tight text-zinc-900 dark:text-zinc-100">Manajemen Pelanggan</h1>
        <p class="text-sm text-zinc-500 dark:text-zinc-400 mt-1">
          Kelola basis data pelanggan, nomor kontak, dan histori transaksi toko Anda.
        </p>
      </div>

      <button 
        @click="openAddModal"
        class="bg-emerald-600 hover:bg-emerald-700 text-white px-4 py-2 rounded-xl font-bold text-sm transition-colors shadow-xs flex items-center gap-2"
      >
        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M16 21v-2a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v2"/><circle cx="9" cy="7" r="4"/><line x1="19" x2="19" y1="8" y2="14"/><line x1="22" x2="16" y1="11" y2="11"/></svg>
        Tambah Pelanggan
      </button>
    </div>

    <!-- Search Input -->
    <div class="relative max-w-md">
      <svg class="absolute left-3.5 top-1/2 -translate-y-1/2 text-zinc-400" xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="11" cy="11" r="8"/><path d="m21 21-4.3-4.3"/></svg>
      <input 
        v-model="searchQuery"
        type="text" 
        placeholder="Cari pelanggan berdasarkan nama atau No. HP..." 
        class="w-full pl-10 pr-4 py-2.5 bg-white dark:bg-zinc-900 border border-zinc-200 dark:border-zinc-800 rounded-xl text-zinc-900 dark:text-zinc-100 focus:outline-none focus:ring-2 focus:ring-emerald-500/20 text-sm font-medium"
      />
    </div>

    <!-- Shadcn UI Table -->
    <div class="bg-white dark:bg-zinc-900 border border-zinc-200 dark:border-zinc-800 rounded-2xl shadow-xs overflow-hidden">
      <Table>
        <TableHeader>
          <TableRow>
            <TableHead>Nama Pelanggan</TableHead>
            <TableHead>No. Telepon / WA</TableHead>
            <TableHead>Email</TableHead>
            <TableHead>Jumlah Transaksi</TableHead>
            <TableHead>Total Belanja</TableHead>
            <TableHead class="text-right">Aksi</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          <TableRow v-for="cust in workspaceCustomers" :key="cust.id">
            <TableCell>
              <div class="font-bold text-zinc-900 dark:text-zinc-100">{{ cust.name }}</div>
              <div class="text-xs text-zinc-400 font-mono">{{ cust.id }}</div>
            </TableCell>
            <TableCell class="text-zinc-600 dark:text-zinc-300 font-mono text-xs font-semibold">{{ cust.phone }}</TableCell>
            <TableCell class="text-zinc-500">{{ cust.email || '-' }}</TableCell>
            <TableCell class="text-zinc-900 dark:text-zinc-100 font-bold">{{ cust.totalTransactions }}x</TableCell>
            <TableCell class="text-emerald-600 dark:text-emerald-400 font-black">Rp {{ cust.totalSpending.toLocaleString('id-ID') }}</TableCell>
            <TableCell class="text-right space-x-2">
              <button 
                @click="openEditModal(cust)"
                class="text-zinc-400 hover:text-emerald-600 transition-colors p-1"
                title="Edit Pelanggan"
              >
                <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M17 3a2.85 2.83 0 1 1 4 4L7.5 20.5 2 22l1.5-5.5Z"/><path d="m15 5 4 4"/></svg>
              </button>
              <button 
                v-if="cust.id !== 'CUST-001'"
                @click="confirmDelete(cust.id, cust.name)"
                class="text-zinc-400 hover:text-rose-600 transition-colors p-1"
                title="Hapus Pelanggan"
              >
                <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M3 6h18"/><path d="M19 6v14c0 1-1 2-2 2H7c-1 0-2-1-2-2V6"/><path d="M8 6V4c0-1 1-2 2-2h4c1 0 2 1 2 2v2"/></svg>
              </button>
            </TableCell>
          </TableRow>
        </TableBody>
      </Table>
    </div>

    <!-- Modal Form Add/Edit -->
    <div v-if="isModalOpen" class="fixed inset-0 z-50 bg-black/50 backdrop-blur-sm flex items-center justify-center p-4">
      <div class="bg-white dark:bg-zinc-900 border border-zinc-200 dark:border-zinc-800 rounded-3xl p-6 w-full max-w-md shadow-2xl space-y-4">
        <h3 class="text-lg font-bold text-zinc-900 dark:text-zinc-100">
          {{ editingId ? 'Edit Pelanggan' : 'Tambah Pelanggan Baru' }}
        </h3>

        <form @submit.prevent="saveCustomer" class="space-y-4">
          <div>
            <label class="block text-xs font-bold text-zinc-600 dark:text-zinc-400 mb-1">Nama Lengkap *</label>
            <input 
              v-model="formName" 
              type="text" 
              required 
              class="w-full px-3 py-2 bg-zinc-50 dark:bg-zinc-800 border border-zinc-200 dark:border-zinc-700 rounded-xl text-sm text-zinc-900 dark:text-zinc-100 font-medium focus:outline-none focus:ring-2 focus:ring-emerald-500/20"
            />
          </div>

          <div class="grid grid-cols-2 gap-3">
            <div>
              <label class="block text-xs font-bold text-zinc-600 dark:text-zinc-400 mb-1">No. HP / WhatsApp</label>
              <input 
                v-model="formPhone" 
                type="text" 
                class="w-full px-3 py-2 bg-zinc-50 dark:bg-zinc-800 border border-zinc-200 dark:border-zinc-700 rounded-xl text-sm text-zinc-900 dark:text-zinc-100 font-medium focus:outline-none focus:ring-2 focus:ring-emerald-500/20"
              />
            </div>

            <div>
              <label class="block text-xs font-bold text-zinc-600 dark:text-zinc-400 mb-1">Email</label>
              <input 
                v-model="formEmail" 
                type="email" 
                class="w-full px-3 py-2 bg-zinc-50 dark:bg-zinc-800 border border-zinc-200 dark:border-zinc-700 rounded-xl text-sm text-zinc-900 dark:text-zinc-100 focus:outline-none focus:ring-2 focus:ring-emerald-500/20"
              />
            </div>
          </div>

          <div>
            <label class="block text-xs font-bold text-zinc-600 dark:text-zinc-400 mb-1">Alamat</label>
            <textarea 
              v-model="formAddress" 
              rows="2"
              class="w-full px-3 py-2 bg-zinc-50 dark:bg-zinc-800 border border-zinc-200 dark:border-zinc-700 rounded-xl text-sm text-zinc-900 dark:text-zinc-100 focus:outline-none focus:ring-2 focus:ring-emerald-500/20"
            ></textarea>
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

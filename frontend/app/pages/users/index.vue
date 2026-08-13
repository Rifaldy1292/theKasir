<script setup lang="ts">
import { ref } from 'vue';
import { useUserStore, type User } from '@/stores/useUserStore';
import { Select, SelectTrigger, SelectValue, SelectContent, SelectItem } from '@/components/ui/select';

const userStore = useUserStore();

const isModalOpen = ref(false);
const editingId = ref<string | null>(null);

const formName = ref('');
const formEmail = ref('');
const formRole = ref<'owner' | 'admin' | 'cashier'>('cashier');
const formStatus = ref('Active');

const openAddModal = () => {
  editingId.value = null;
  formName.value = '';
  formEmail.value = '';
  formRole.value = 'cashier';
  formStatus.value = 'Active';
  isModalOpen.value = true;
};

const openEditModal = (user: User) => {
  editingId.value = user.id;
  formName.value = user.name;
  formEmail.value = user.email;
  formRole.value = user.role;
  formStatus.value = user.status;
  isModalOpen.value = true;
};

const saveUser = () => {
  if (!formName.value.trim() || !formEmail.value.trim()) return;

  if (editingId.value) {
    userStore.updateUser(editingId.value, {
      name: formName.value,
      email: formEmail.value,
      role: formRole.value,
      status: formStatus.value,
    });
  } else {
    userStore.addUser({
      workspaceId: 'ws-1',
      name: formName.value,
      email: formEmail.value,
      role: formRole.value,
      status: formStatus.value,
    });
  }

  isModalOpen.value = false;
};

const confirmDelete = (id: string, name: string) => {
  if (confirm(`Apakah Anda yakin ingin menghapus pengguna "${name}"?`)) {
    userStore.deleteUser(id);
  }
};
</script>

<template>
  <div class="space-y-6">
    <div class="flex items-center justify-between">
      <div>
        <h1 class="text-2xl font-bold tracking-tight text-zinc-900 dark:text-zinc-100">Karyawan & Pengguna</h1>
        <p class="text-sm text-zinc-500 dark:text-zinc-400 mt-1">
          Kelola hak akses pengguna, kasir, dan staf toko Anda.
        </p>
      </div>
      <button 
        @click="openAddModal"
        class="bg-emerald-600 hover:bg-emerald-700 text-white px-4 py-2 rounded-xl font-bold text-sm transition-colors shadow-xs flex items-center gap-2"
      >
        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M16 21v-2a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v2"/><circle cx="9" cy="7" r="4"/><line x1="19" x2="19" y1="8" y2="14"/><line x1="22" x2="16" y1="11" y2="11"/></svg>
        Undang / Tambah Karyawan
      </button>
    </div>

    <div class="bg-white dark:bg-zinc-900 border border-zinc-200 dark:border-zinc-800 rounded-2xl shadow-xs overflow-hidden">
      <div class="overflow-x-auto">
        <table class="w-full text-left text-sm whitespace-nowrap">
          <thead class="bg-zinc-50 dark:bg-zinc-900/50 border-b border-zinc-200 dark:border-zinc-800 text-zinc-500 dark:text-zinc-400 font-bold">
            <tr>
              <th class="px-6 py-4">Nama Pengguna</th>
              <th class="px-6 py-4">Email</th>
              <th class="px-6 py-4">Role (Peran)</th>
              <th class="px-6 py-4">Status</th>
              <th class="px-6 py-4 text-right">Aksi</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-zinc-200 dark:divide-zinc-800">
            <tr v-for="user in userStore.users" :key="user.id" class="hover:bg-zinc-50/50 dark:hover:bg-zinc-800/20 transition-colors">
              <td class="px-6 py-4">
                <div class="flex items-center gap-3">
                  <div class="w-8 h-8 rounded-full bg-emerald-100 text-emerald-700 dark:bg-emerald-900/50 dark:text-emerald-300 flex items-center justify-center font-bold text-xs uppercase">
                    {{ user.name.charAt(0) }}
                  </div>
                  <div class="font-bold text-zinc-900 dark:text-zinc-100">{{ user.name }}</div>
                </div>
              </td>
              <td class="px-6 py-4 text-zinc-500">{{ user.email }}</td>
              <td class="px-6 py-4">
                <span class="capitalize px-2.5 py-1 rounded-lg text-xs font-bold bg-zinc-100 dark:bg-zinc-800 text-zinc-800 dark:text-zinc-200">
                  {{ user.role }}
                </span>
              </td>
              <td class="px-6 py-4">
                <span class="px-2.5 py-1 rounded-full text-xs font-bold bg-emerald-100 text-emerald-700 dark:bg-emerald-500/20 dark:text-emerald-400 border border-emerald-200 dark:border-emerald-500/20">
                  {{ user.status }}
                </span>
              </td>
              <td class="px-6 py-4 text-right space-x-2">
                <button 
                  @click="openEditModal(user)"
                  class="text-zinc-400 hover:text-emerald-600 transition-colors p-1"
                  title="Edit Pengguna"
                >
                  <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M17 3a2.85 2.83 0 1 1 4 4L7.5 20.5 2 22l1.5-5.5Z"/><path d="m15 5 4 4"/></svg>
                </button>
                <button 
                  @click="confirmDelete(user.id, user.name)"
                  class="text-zinc-400 hover:text-rose-600 transition-colors p-1"
                  title="Hapus Pengguna"
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
          {{ editingId ? 'Edit Karyawan' : 'Tambah Karyawan Baru' }}
        </h3>

        <form @submit.prevent="saveUser" class="space-y-4">
          <div>
            <label class="block text-xs font-bold text-zinc-600 dark:text-zinc-400 mb-1">Nama Lengkap</label>
            <input 
              v-model="formName" 
              type="text" 
              required 
              class="w-full px-3 py-2 bg-zinc-50 dark:bg-zinc-800 border border-zinc-200 dark:border-zinc-700 rounded-xl text-sm text-zinc-900 dark:text-zinc-100 font-medium focus:outline-none focus:ring-2 focus:ring-emerald-500/20"
            />
          </div>

          <div>
            <label class="block text-xs font-bold text-zinc-600 dark:text-zinc-400 mb-1">Email</label>
            <input 
              v-model="formEmail" 
              type="email" 
              required 
              class="w-full px-3 py-2 bg-zinc-50 dark:bg-zinc-800 border border-zinc-200 dark:border-zinc-700 rounded-xl text-sm text-zinc-900 dark:text-zinc-100 focus:outline-none focus:ring-2 focus:ring-emerald-500/20"
            />
          </div>

          <div class="grid grid-cols-2 gap-3">
            <div>
              <label class="block text-xs font-bold text-zinc-600 dark:text-zinc-400 mb-1">Role (Peran)</label>
              <Select v-model="formRole">
                <SelectTrigger class="h-10">
                  <SelectValue placeholder="Pilih Role" />
                </SelectTrigger>
                <SelectContent>
                  <SelectItem value="owner">Owner</SelectItem>
                  <SelectItem value="admin">Admin</SelectItem>
                  <SelectItem value="cashier">Cashier (Kasir)</SelectItem>
                </SelectContent>
              </Select>
            </div>

            <div>
              <label class="block text-xs font-bold text-zinc-600 dark:text-zinc-400 mb-1">Status</label>
              <Select v-model="formStatus">
                <SelectTrigger class="h-10">
                  <SelectValue placeholder="Pilih Status" />
                </SelectTrigger>
                <SelectContent>
                  <SelectItem value="Active">Aktif</SelectItem>
                  <SelectItem value="Inactive">Nonaktif</SelectItem>
                </SelectContent>
              </Select>
            </div>
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

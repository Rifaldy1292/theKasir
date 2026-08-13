import { defineStore } from 'pinia';
import { useLocalStorage } from '@vueuse/core';

export type ShiftStatus = 'OPEN' | 'CLOSED';

export type Shift = {
  id: string;
  workspaceId: string;
  cashierName: string;
  openingBalance: number;
  cashSales: number;
  expectedCash: number;
  actualCash: number;
  difference: number;
  status: ShiftStatus;
  openedAt: string;
  closedAt?: string;
  notes?: string;
};

const DEFAULT_SHIFTS: Shift[] = [
  {
    id: 'SHIFT-001',
    workspaceId: 'ws-1',
    cashierName: 'Budi (Cashier)',
    openingBalance: 300000,
    cashSales: 1500000,
    expectedCash: 1800000,
    actualCash: 1800000,
    difference: 0,
    status: 'CLOSED',
    openedAt: '2026-08-13T08:00:00Z',
    closedAt: '2026-08-13T14:00:00Z',
    notes: 'Shift pagi lancar'
  }
];

export const useShiftStore = defineStore('shift', {
  state: () => ({
    shifts: useLocalStorage<Shift[]>('thekasir-shifts', DEFAULT_SHIFTS),
  }),
  getters: {
    activeShift: (state) => {
      return (workspaceId: string, cashierName: string) => {
        return state.shifts.find(s => s.workspaceId === workspaceId && s.cashierName === cashierName && s.status === 'OPEN');
      };
    },
    shiftHistory: (state) => {
      return (workspaceId: string) => {
        return state.shifts.filter(s => s.workspaceId === workspaceId);
      };
    }
  },
  actions: {
    openShift(workspaceId: string, cashierName: string, openingBalance: number) {
      const newShift: Shift = {
        id: `SHIFT-${Date.now().toString().slice(-4)}`,
        workspaceId,
        cashierName,
        openingBalance,
        cashSales: 0,
        expectedCash: openingBalance,
        actualCash: openingBalance,
        difference: 0,
        status: 'OPEN',
        openedAt: new Date().toISOString()
      };
      this.shifts.push(newShift);
      return newShift;
    },
    recordCashSale(shiftId: string, amount: number) {
      const shift = this.shifts.find(s => s.id === shiftId);
      if (shift && shift.status === 'OPEN') {
        shift.cashSales += amount;
        shift.expectedCash = shift.openingBalance + shift.cashSales;
      }
    },
    closeShift(shiftId: string, actualCash: number, notes?: string) {
      const shift = this.shifts.find(s => s.id === shiftId);
      if (shift && shift.status === 'OPEN') {
        shift.actualCash = actualCash;
        shift.difference = actualCash - shift.expectedCash;
        shift.status = 'CLOSED';
        shift.closedAt = new Date().toISOString();
        if (notes) shift.notes = notes;
      }
    }
  }
});

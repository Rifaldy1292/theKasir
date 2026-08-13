import { defineNuxtRouteMiddleware, navigateTo } from '#imports';
import { useAppState } from '@/composables/useAppState';

export default defineNuxtRouteMiddleware((to) => {
  const { currentRole } = useAppState();

  // Public / auth pages exempt from RBAC
  if (to.path.startsWith('/workspace')) {
    return;
  }

  // Cashier role restricted paths according to PRD section 5 & 20
  const cashierRestrictedPaths = ['/reports', '/settings', '/users', '/products'];

  if (currentRole.value === 'cashier') {
    const isRestricted = cashierRestrictedPaths.some(path => to.path.startsWith(path));
    if (isRestricted) {
      // Redirect cashier attempting restricted path back to POS
      return navigateTo('/pos');
    }
  }
});

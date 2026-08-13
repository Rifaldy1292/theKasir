export const queryKeys = {
  auth: {
    user: () => ['auth', 'user'] as const,
  },
  workspaces: {
    all: () => ['workspaces'] as const,
    detail: (id: string) => ['workspaces', id] as const,
  },
}

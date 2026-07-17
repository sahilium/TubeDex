import { writable } from 'svelte/store';
import type { User } from './api';

export const currentUser = writable<User | null>(null);
export const isAuthenticated = writable<boolean>(false);

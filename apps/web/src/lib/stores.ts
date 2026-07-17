import { writable } from 'svelte/store';
import type { User } from './api';

export const currentUser = writable<User | null>(null);
export const isAuthenticated = writable<boolean>(false);

function createTheme() {
	const { subscribe, set, update } = writable<'light' | 'dark'>('light');

	function apply(theme: 'light' | 'dark') {
		if (theme === 'dark') {
			document.documentElement.classList.add('dark');
		} else {
			document.documentElement.classList.remove('dark');
		}
	}

	function init() {
		const stored = localStorage.getItem('theme');
		if (stored === 'light' || stored === 'dark') {
			apply(stored);
			set(stored);
			return;
		}
		const prefersDark = window.matchMedia('(prefers-color-scheme: dark)').matches;
		const theme = prefersDark ? 'dark' : 'light';
		apply(theme);
		set(theme);
	}

	return {
		subscribe,
		toggle: () => {
			let current: 'light' | 'dark' = 'light';
			const unsub = subscribe(v => { current = v; });
			unsub();
			const next = current === 'dark' ? 'light' : 'dark';
			apply(next);
			localStorage.setItem('theme', next);
			set(next);
		},
		init,
	};
}

export const theme = createTheme();

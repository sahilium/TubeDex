<script lang="ts">
	import { currentUser, theme } from '$lib/stores';
	import { api } from '$lib/api';
	import { Moon, Sun } from 'lucide-svelte';

	async function logout() {
		await api.post('/auth/logout');
		window.location.href = '/';
	}
</script>

<div class="max-w-4xl mx-auto px-4 py-6">
	<h1 class="text-2xl font-bold mb-6">Settings</h1>

	<div class="p-6 border border-border rounded-xl mb-4">
		<h2 class="font-semibold mb-4">Account</h2>
		{#if $currentUser}
			<div class="flex items-center gap-4">
				<img
					src={$currentUser.avatar_url}
					alt={$currentUser.name}
					class="w-12 h-12 rounded-full"
				/>
				<div>
					<p class="font-medium">{$currentUser.name}</p>
					<p class="text-sm text-muted">{$currentUser.email}</p>
				</div>
			</div>
		{/if}
	</div>

	<div class="p-6 border border-border rounded-xl mb-4">
		<h2 class="font-semibold mb-4">Appearance</h2>
		<div class="flex items-center justify-between">
			<span class="text-sm">Dark mode</span>
			<button
				onclick={() => theme.toggle()}
				class="p-2 rounded-lg text-muted hover:text-gray-900 dark:hover:text-gray-100 hover:bg-gray-100 dark:hover:bg-gray-800 transition-colors"
				aria-label="Toggle dark mode"
			>
				{#if $theme === 'dark'}
					<Sun class="w-5 h-5" />
				{:else}
					<Moon class="w-5 h-5" />
				{/if}
			</button>
		</div>
		<p class="text-xs text-muted mt-2">Follows your system preference by default.</p>
	</div>

	<button
		onclick={logout}
		class="px-4 py-2 border border-red-200 dark:border-red-800 text-red-600 dark:text-red-400 rounded-lg hover:bg-red-50 dark:hover:bg-red-950 transition-colors text-sm font-medium"
	>
		Sign Out
	</button>
</div>

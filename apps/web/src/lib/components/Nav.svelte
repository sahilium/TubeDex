<script lang="ts">
	import { page } from '$app/stores';
	import { theme } from '$lib/stores';
	import { Home, Search, FolderIcon, RefreshCw, Settings, Moon, Sun } from 'lucide-svelte';

	const links = [
		{ href: '/dashboard', label: 'Home', icon: Home },
		{ href: '/search', label: 'Search', icon: Search },
		{ href: '/collections', label: 'Collections', icon: FolderIcon },
		{ href: '/sync', label: 'Sync', icon: RefreshCw },
		{ href: '/settings', label: 'Settings', icon: Settings },
	];
</script>

<nav class="fixed bottom-0 left-0 right-0 bg-white dark:bg-gray-900 border-t border-border md:top-0 md:bottom-auto md:left-0 md:right-auto md:w-64 md:h-full md:border-t-0 md:border-r md:overflow-y-auto">
	<div class="hidden md:flex md:items-center md:justify-between md:px-6 md:py-5 md:border-b md:border-border">
		<span class="text-lg font-bold">TubeDex</span>
		<button
			onclick={() => theme.toggle()}
			class="p-1.5 rounded-lg text-muted hover:text-gray-900 dark:hover:text-gray-100 hover:bg-gray-100 dark:hover:bg-gray-800 transition-colors"
			aria-label="Toggle dark mode"
		>
			{#if $theme === 'dark'}
				<Sun class="w-4 h-4" />
			{:else}
				<Moon class="w-4 h-4" />
			{/if}
		</button>
	</div>
	<ul class="flex justify-around md:flex-col md:p-4 md:gap-1">
		{#each links as link}
			<li>
				<a
					href={link.href}
					class="flex flex-col items-center gap-1 px-3 py-2 text-xs font-medium transition-colors md:flex-row md:text-sm md:rounded-lg md:px-3 md:py-2.5 {$page.url.pathname.startsWith(link.href) ? 'text-accent bg-blue-50 dark:bg-blue-900/30' : 'text-muted hover:text-gray-900 dark:hover:text-gray-100 hover:bg-gray-50 dark:hover:bg-gray-800'}"
				>
					<svelte:component this={link.icon} class="w-5 h-5 md:w-4 md:h-4" />
					<span>{link.label}</span>
				</a>
			</li>
		{/each}
	</ul>
</nav>

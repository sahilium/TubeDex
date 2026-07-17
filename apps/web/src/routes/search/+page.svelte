<script lang="ts">
	import { createQuery } from '@tanstack/svelte-query';
	import { api } from '$lib/api';
	import { Search as SearchIcon } from 'lucide-svelte';
	import type { SearchResult } from '$lib/api';

	let inputValue = $state('');
	let searchQuery = $state('');
	let debounceTimer: ReturnType<typeof setTimeout>;

	const results = createQuery(() => ({
		queryKey: ['search', searchQuery],
		queryFn: () => api.get('/search', { q: searchQuery }),
		enabled: searchQuery.length > 0,
	}));

	function handleInput(e: Event) {
		const value = (e.target as HTMLInputElement).value;
		inputValue = value;
		clearTimeout(debounceTimer);
		debounceTimer = setTimeout(() => {
			searchQuery = value;
		}, 200);
	}
</script>

<div class="max-w-4xl mx-auto px-4 py-6">
	<div class="relative mb-6">
		<SearchIcon class="absolute left-4 top-1/2 -translate-y-1/2 w-5 h-5 text-muted" />
		<input
			type="search"
			placeholder="Search channels, collections, notes..."
			oninput={handleInput}
			value={inputValue}
			class="w-full pl-12 pr-4 py-3 border border-border rounded-xl bg-white text-base focus:outline-none focus:ring-2 focus:ring-accent focus:border-transparent"
		/>
	</div>

	{#if !searchQuery}
		<p class="text-center text-muted py-12">Type to search your library</p>
	{:else if results.isLoading}
		<div class="flex justify-center py-12">
			<div class="w-6 h-6 border-2 border-accent border-t-transparent rounded-full animate-spin"></div>
		</div>
	{:else if results.data?.length === 0}
		<p class="text-center text-muted py-12">No results for "{searchQuery}"</p>
	{:else}
		<div class="space-y-2">
			{#each results.data ?? [] as result}
				<a
					href={result.type === 'channel' ? `/channel/${result.id}` : `/collections`}
					class="flex items-center gap-4 p-4 rounded-xl border border-border hover:border-gray-300 hover:bg-gray-50 transition-colors"
				>
					{#if result.avatar}
						<img src={result.avatar} alt={result.name} class="w-10 h-10 rounded-full flex-shrink-0" />
					{:else}
						<div class="w-10 h-10 rounded-full bg-gray-100 flex items-center justify-center flex-shrink-0">
							<span class="text-sm font-medium text-muted">{result.name[0]}</span>
						</div>
					{/if}
					<div class="flex-1 min-w-0">
						<h3 class="font-medium truncate">{result.name}</h3>
						<p class="text-sm text-muted truncate">
							{result.type === 'channel' ? `@${result.handle}` : result.type}
						</p>
					</div>
				</a>
			{/each}
		</div>
	{/if}
</div>

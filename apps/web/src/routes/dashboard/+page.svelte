<script lang="ts">
	import { createQuery } from '@tanstack/svelte-query';
	import { api } from '$lib/api';
	import ChannelCard from '$lib/components/ChannelCard.svelte';
	import type { ChannelWithDetails } from '$lib/api';

	let sort = $state('name');

	const query = createQuery(() => ({
		queryKey: ['subscriptions', sort],
		queryFn: () => api.get('/channels', { sort }),
	}));
</script>

<div class="max-w-4xl mx-auto px-4 py-6">
	<div class="flex items-center justify-between mb-6">
		<h1 class="text-2xl font-bold">Subscriptions</h1>
		<select
			bind:value={sort}
			class="text-sm border border-border rounded-lg px-3 py-1.5 bg-white dark:bg-gray-900"
		>
			<option value="name">A–Z</option>
			<option value="recent">Recent</option>
		</select>
	</div>

	{#if query.isLoading}
		<div class="flex justify-center py-12">
			<div class="w-6 h-6 border-2 border-accent border-t-transparent rounded-full animate-spin"></div>
		</div>
	{:else if query.isError}
		<p class="text-red-500 dark:text-red-400 text-center py-12">Failed to load subscriptions</p>
	{:else if query.data?.length === 0}
		<div class="text-center py-12">
			<p class="text-muted mb-4">No subscriptions yet.</p>
			<a href="/sync" class="text-accent hover:underline">Import your subscriptions</a>
		</div>
	{:else}
		<div class="grid gap-3">
			{#each query.data ?? [] as channel}
				<ChannelCard {channel} />
			{/each}
		</div>
	{/if}
</div>

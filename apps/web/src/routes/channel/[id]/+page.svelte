<script lang="ts">
	import { createQuery, createMutation } from '@tanstack/svelte-query';
	import { api } from '$lib/api';
	import { queryClient } from '$lib/queryClient';
	import { page } from '$app/stores';
	import { formatSubscribers } from '$lib/utils';
	import { FileText, Star } from 'lucide-svelte';
	import type { ChannelWithDetails, Collection } from '$lib/api';

	const channelId = $page.params.id;

	const channel = createQuery(() => ({
		queryKey: ['channel', channelId],
		queryFn: () => api.get(`/channels/${channelId}`),
	}));

	const collections = createQuery(() => ({
		queryKey: ['collections'],
		queryFn: () => api.get('/collections'),
	}));

	let noteText = $state('');
	let userRating = $state(0);

	const saveNote = createMutation(() => ({
		mutationFn: (body: string) => api.put('/notes', { channel_id: Number(channelId), body }),
		onSuccess: () => queryClient.invalidateQueries({ queryKey: ['channel', channelId] }),
	}));

	const saveRating = createMutation(() => ({
		mutationFn: (rating: number) => api.put('/ratings', { channel_id: Number(channelId), rating }),
		onSuccess: () => queryClient.invalidateQueries({ queryKey: ['channel', channelId] }),
	}));
</script>

<div class="max-w-4xl mx-auto px-4 py-6">
	{#if channel.isLoading}
		<div class="flex justify-center py-12">
			<div class="w-6 h-6 border-2 border-accent border-t-transparent rounded-full animate-spin"></div>
		</div>
	{:else if channel.data}
		{@const ch = channel.data}

		<div class="flex items-center gap-5 mb-8">
			<img src={ch.avatar} alt={ch.name} class="w-20 h-20 rounded-full" />
			<div>
				<h1 class="text-2xl font-bold">{ch.name}</h1>
				<p class="text-muted">@{ch.handle ?? 'unknown'}</p>
				<p class="text-sm text-muted">{formatSubscribers(ch.subscriber_count)} subscribers</p>
			</div>
		</div>

		<div class="space-y-6">
			{#if ch.description}
				<div class="p-4 border border-border rounded-xl">
					<p class="text-sm text-muted leading-relaxed">{ch.description}</p>
				</div>
			{/if}

			<div class="p-4 border border-border rounded-xl">
				<h2 class="font-semibold mb-3 flex items-center gap-2">
					<Star class="w-4 h-4 text-amber-500" />
					Rating
				</h2>
				<div class="flex gap-1">
					{#each [1, 2, 3, 4, 5] as star}
						<button
							onclick={() => { userRating = star; saveRating.mutate(star); }}
							class="text-2xl transition-colors {star <= (ch.rating ?? userRating) ? 'text-amber-400' : 'text-gray-200 dark:text-gray-600'} hover:text-amber-400"
						>
							★
						</button>
					{/each}
				</div>
			</div>

			<div class="p-4 border border-border rounded-xl">
				<h2 class="font-semibold mb-3 flex items-center gap-2">
					<FileText class="w-4 h-4" />
					Notes
				</h2>
				<textarea
					bind:value={noteText}
					placeholder="Write your notes about this channel..."
					class="w-full px-3 py-2 border border-border rounded-lg text-sm resize-none h-24 bg-white dark:bg-gray-900 focus:outline-none focus:ring-2 focus:ring-accent"
				></textarea>
				<button
					onclick={() => saveNote.mutate(noteText)}
					disabled={saveNote.isPending}
					class="mt-2 px-4 py-1.5 bg-accent text-white text-sm font-medium rounded-lg hover:bg-accent-hover disabled:opacity-50"
				>
					Save Note
				</button>
			</div>

			{#if collections.data}
				<div class="p-4 border border-border rounded-xl">
					<h2 class="font-semibold mb-3">Collections</h2>
					<div class="flex flex-wrap gap-2">
						{#each collections.data ?? [] as collection}
							<span
								class="inline-flex items-center gap-1 px-3 py-1 rounded-full text-sm font-medium"
								style="background-color: {collection.color}20; color: {collection.color}"
							>
								{collection.name}
							</span>
						{/each}
					</div>
				</div>
			{/if}
		</div>
	{/if}
</div>

<script lang="ts">
	import { createQuery, createMutation } from '@tanstack/svelte-query';
	import { api } from '$lib/api';
	import { queryClient } from '$lib/queryClient';
	import { Plus } from 'lucide-svelte';
	import type { Collection } from '$lib/api';

	const collections = createQuery(() => ({
		queryKey: ['collections'],
		queryFn: () => api.get('/collections'),
	}));

	let showForm = $state(false);
	let name = $state('');
	let color = $state('#3b82f6');

	const create = createMutation(() => ({
		mutationFn: (data: { name: string; color: string }) =>
			api.post('/collections', { ...data, icon: 'folder' }),
		onSuccess: () => {
			queryClient.invalidateQueries({ queryKey: ['collections'] });
			showForm = false;
			name = '';
			color = '#3b82f6';
		},
	}));

	const del = createMutation(() => ({
		mutationFn: (id: number) => api.delete(`/collections/${id}`),
		onSuccess: () => queryClient.invalidateQueries({ queryKey: ['collections'] }),
	}));
</script>

<div class="max-w-4xl mx-auto px-4 py-6">
	<div class="flex items-center justify-between mb-6">
		<h1 class="text-2xl font-bold">Collections</h1>
		<button
			onclick={() => (showForm = !showForm)}
			class="flex items-center gap-2 text-sm font-medium text-accent hover:text-accent-hover"
		>
			<Plus class="w-4 h-4" />
			New
		</button>
	</div>

	{#if showForm}
		<form
			onsubmit={(e) => { e.preventDefault(); create.mutate({ name, color }); }}
			class="flex items-center gap-3 mb-6 p-4 border border-border rounded-xl"
		>
			<input
				type="text"
				placeholder="Collection name"
				bind:value={name}
				required
				class="flex-1 px-3 py-2 border border-border rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-accent"
			/>
			<input
				type="color"
				bind:value={color}
				class="w-9 h-9 rounded-lg border border-border cursor-pointer"
			/>
			<button
				type="submit"
				disabled={create.isPending}
				class="px-4 py-2 bg-accent text-white text-sm font-medium rounded-lg hover:bg-accent-hover disabled:opacity-50"
			>
				{create.isPending ? '...' : 'Create'}
			</button>
		</form>
	{/if}

	{#if collections.isLoading}
		<div class="flex justify-center py-12">
			<div class="w-6 h-6 border-2 border-accent border-t-transparent rounded-full animate-spin"></div>
		</div>
	{:else if collections.data?.length === 0}
		<p class="text-center text-muted py-12">No collections yet. Create one to organize your channels.</p>
	{:else}
		<div class="space-y-2">
			{#each collections.data ?? [] as collection}
				<div class="flex items-center gap-4 p-4 rounded-xl border border-border">
					<div
						class="w-4 h-4 rounded-full flex-shrink-0"
						style="background-color: {collection.color}"
					></div>
					<span class="flex-1 font-medium">{collection.name}</span>
					<button
						onclick={() => del.mutate(collection.id)}
						class="text-sm text-muted hover:text-red-500 transition-colors"
					>
						Delete
					</button>
				</div>
			{/each}
		</div>
	{/if}
</div>

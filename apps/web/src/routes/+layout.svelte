<script lang="ts">
	import { onMount } from 'svelte';
	import { QueryClientProvider } from '@tanstack/svelte-query';
	import { queryClient } from '$lib/queryClient';
	import { currentUser, isAuthenticated, theme } from '$lib/stores';
	import { api } from '$lib/api';
	import Nav from '$lib/components/Nav.svelte';
	import '../app.css';

	let { children } = $props();

	let ready = $state(false);

	onMount(async () => {
		theme.init();
		try {
			const user = await api.get('/me');
			currentUser.set(user);
			isAuthenticated.set(true);
		} catch {
			isAuthenticated.set(false);
		} finally {
			ready = true;
		}
	});
</script>

<QueryClientProvider client={queryClient}>
	{#if ready}
		<div class="min-h-screen bg-white dark:bg-gray-950">
			<main class="pb-20 md:pb-0 md:ml-64">
				{@render children()}
			</main>
			{#if $isAuthenticated}
				<Nav />
			{/if}
		</div>
	{:else}
		<div class="flex items-center justify-center min-h-screen">
			<div class="w-6 h-6 border-2 border-accent border-t-transparent rounded-full animate-spin"></div>
		</div>
	{/if}
</QueryClientProvider>

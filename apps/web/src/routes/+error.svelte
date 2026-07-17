<script lang="ts">
	import { page } from '$app/stores';
	import { Home, ArrowLeft } from 'lucide-svelte';
</script>

<div class="flex flex-col items-center justify-center min-h-screen px-6 text-center">
	{#if $page.error?.status === 404}
		<span class="text-7xl font-bold text-accent tracking-tight">404</span>
		<h1 class="text-2xl font-bold mt-4">Page not found</h1>
		<p class="text-muted mt-2 max-w-sm">
			The page you're looking for doesn't exist or has been moved.
		</p>
	{:else}
		<span class="text-7xl font-bold text-red-500 tracking-tight">500</span>
		<h1 class="text-2xl font-bold mt-4">Something went wrong</h1>
		<p class="text-muted mt-2 max-w-sm">
			An unexpected error occurred. Please try again later.
		</p>
	{/if}
	<div class="flex items-center gap-3 mt-8">
		<a
			href="/"
			class="inline-flex items-center gap-2 px-5 py-2.5 bg-accent text-white text-sm font-medium rounded-xl hover:bg-accent-hover transition-colors"
		>
			<Home class="w-4 h-4" />
			Home
		</a>
		<button
			onclick={() => history.back()}
			class="inline-flex items-center gap-2 px-5 py-2.5 border border-border text-sm font-medium rounded-xl hover:bg-gray-50 dark:hover:bg-gray-900 transition-colors"
		>
			<ArrowLeft class="w-4 h-4" />
			Go back
		</button>
	</div>
</div>

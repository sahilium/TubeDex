<script lang="ts">
	import { createQuery, createMutation } from '@tanstack/svelte-query';
	import { api } from '$lib/api';
	import { queryClient } from '$lib/queryClient';
	import { RefreshCw, CheckCircle, XCircle, Clock } from 'lucide-svelte';
	import type { SyncJob } from '$lib/api';

	const status = createQuery(() => ({
		queryKey: ['sync-status'],
		queryFn: () => api.get('/sync/status'),
		retry: false,
	}));

	const sync = createMutation(() => ({
		mutationFn: () => api.post<SyncJob>('/sync'),
		onSuccess: () => {
			queryClient.invalidateQueries({ queryKey: ['sync-status'] });
			setTimeout(() => {
				queryClient.invalidateQueries({ queryKey: ['sync-status'] });
			}, 2000);
		},
	}));

	function statusIcon(status: string) {
		switch (status) {
			case 'completed': return CheckCircle;
			case 'failed': return XCircle;
			case 'running': return RefreshCw;
			default: return Clock;
		}
	}

	function statusColor(status: string) {
		switch (status) {
			case 'completed': return 'text-green-500';
			case 'failed': return 'text-red-500';
			case 'running': return 'text-accent';
			default: return 'text-muted';
		}
	}
</script>

<div class="max-w-4xl mx-auto px-4 py-6">
	<h1 class="text-2xl font-bold mb-6">Sync</h1>

	<div class="p-6 border border-border rounded-xl mb-6">
		<p class="text-muted mb-4">
			Import your YouTube subscriptions. This will fetch all channels you're subscribed to and store them locally.
		</p>
		<button
			onclick={() => sync.mutate()}
			disabled={sync.isPending}
			class="flex items-center gap-2 px-6 py-3 bg-accent text-white font-medium rounded-lg hover:bg-accent-hover disabled:opacity-50 transition-colors"
		>
			<RefreshCw class="w-4 h-4 {sync.isPending ? 'animate-spin' : ''}" />
			{sync.isPending ? 'Syncing...' : 'Sync Now'}
		</button>
	</div>

	{#if status.data}
		<div class="p-4 border border-border rounded-xl">
			<h2 class="font-semibold mb-3">Last Sync</h2>
			<div class="flex items-center gap-3">
				{#each [status.data] as job}
					{@const Icon = statusIcon(job.status)}
					<Icon class="w-5 h-5 {statusColor(job.status)}" />
					<div>
						<p class="capitalize font-medium">{job.status}</p>
						<p class="text-sm text-muted">
							{new Date(job.started_at).toLocaleString()}
						</p>
					</div>
				{/each}
			</div>
		</div>
	{/if}
</div>

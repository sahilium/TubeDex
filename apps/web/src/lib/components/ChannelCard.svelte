<script lang="ts">
	import type { ChannelWithDetails } from '$lib/api';
	import { formatSubscribers } from '$lib/utils';

	let { channel }: { channel: ChannelWithDetails } = $props();
</script>

<a
	href="/channel/{channel.id}"
	class="flex items-center gap-4 p-4 rounded-xl border border-border hover:border-gray-300 hover:bg-gray-50 transition-colors"
>
	<img
		src={channel.avatar}
		alt={channel.name}
		class="w-12 h-12 rounded-full flex-shrink-0"
	/>
	<div class="flex-1 min-w-0">
		<h3 class="font-semibold truncate">{channel.name}</h3>
		<p class="text-sm text-muted truncate">
			{channel.handle ? `@${channel.handle}` : ''}
			{channel.handle ? ' · ' : ''}
			{formatSubscribers(channel.subscriber_count)} subscribers
		</p>
	</div>
	{#if channel.rating}
		<div class="flex items-center gap-0.5 flex-shrink-0">
			<span class="text-sm text-amber-500">{'★'.repeat(channel.rating)}</span>
			<span class="text-sm text-gray-300">{'★'.repeat(5 - channel.rating)}</span>
		</div>
	{/if}
</a>

<script lang="ts">
	import { currentUser } from '$lib/stores';
	import { api } from '$lib/api';

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

	<button
		onclick={logout}
		class="px-4 py-2 border border-red-200 text-red-600 rounded-lg hover:bg-red-50 transition-colors text-sm font-medium"
	>
		Sign Out
	</button>
</div>

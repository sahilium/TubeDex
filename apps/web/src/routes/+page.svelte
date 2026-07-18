<script lang="ts">
	import { PUBLIC_API_URL } from '$env/static/public';
	const apiUrl = PUBLIC_API_URL || '/api/v1';
</script>

<div class="flex flex-col min-h-screen">

	<!-- Nav -->
	<header class="sticky top-0 z-50 bg-white/80 dark:bg-gray-950/80 backdrop-blur-md border-b border-border">
		<div class="flex items-center justify-between px-6 py-4 max-w-6xl mx-auto w-full">
			<span class="text-lg font-bold tracking-tight">TubeDex</span>
			<a
				href="{apiUrl}/auth/login"
				class="text-sm font-medium text-accent hover:text-accent-hover transition-colors"
			>
				Sign in
			</a>
		</div>
	</header>

	<main>
		<!-- Hero -->
		<section class="px-6 py-24 sm:py-32">
			<div class="max-w-3xl mx-auto text-center">
				<div class="inline-flex items-center gap-2 px-3 py-1 rounded-full border border-border text-xs text-muted mb-8">
					<span class="w-1.5 h-1.5 rounded-full bg-green-500"></span>
					Your personal YouTube library
				</div>
				<h1 class="text-5xl sm:text-6xl lg:text-7xl font-bold tracking-tight leading-tight">
					Your YouTube<br />
					<span class="text-accent">personal library</span>
				</h1>
				<p class="text-muted text-lg sm:text-xl mt-6 max-w-lg mx-auto leading-relaxed">
					Organize, rate, and rediscover your subscriptions. Like Letterboxd, for YouTube.
				</p>
				<div class="mt-10 flex items-center justify-center gap-4">
					<a
						href="{apiUrl}/auth/login"
						class="inline-flex items-center gap-2 bg-accent text-white px-8 py-3.5 rounded-xl font-medium hover:bg-accent-hover transition-colors shadow-sm"
					>
						<svg class="w-5 h-5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
							<path d="M15 3h4a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2h-4" />
							<polyline points="10 17 15 12 10 7" />
							<line x1="15" y1="12" x2="3" y2="12" />
						</svg>
						Continue with Google
					</a>
				</div>
			</div>
		</section>

		<!-- Features -->
		<section class="px-6 py-20 sm:py-28 bg-gray-50 dark:bg-gray-900/50">
			<div class="max-w-5xl mx-auto">
				<h2 class="text-3xl sm:text-4xl font-bold text-center tracking-tight">Everything you need</h2>
				<p class="text-muted text-center mt-3 max-w-md mx-auto">Tools to turn your subscriptions into a curated library.</p>
				<div class="grid sm:grid-cols-3 gap-6 mt-12">
					<div class="p-6 rounded-2xl border border-border bg-white dark:bg-gray-900">
						<div class="w-10 h-10 rounded-xl bg-accent/10 flex items-center justify-center mb-4">
							<svg class="w-5 h-5 text-accent" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
								<rect x="3" y="3" width="7" height="7" />
								<rect x="14" y="3" width="7" height="7" />
								<rect x="14" y="14" width="7" height="7" />
								<rect x="3" y="14" width="7" height="7" />
							</svg>
						</div>
						<h3 class="font-semibold">Collections</h3>
						<p class="text-sm text-muted mt-2 leading-relaxed">
							Group channels into custom collections. Keep your topics organized and easy to browse.
						</p>
					</div>
					<div class="p-6 rounded-2xl border border-border bg-white dark:bg-gray-900">
						<div class="w-10 h-10 rounded-xl bg-accent/10 flex items-center justify-center mb-4">
							<svg class="w-5 h-5 text-accent" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
								<polygon points="12 2 15.09 8.26 22 9.27 17 14.14 18.18 21.02 12 17.77 5.82 21.02 7 14.14 2 9.27 8.91 8.26 12 2" />
							</svg>
						</div>
						<h3 class="font-semibold">Ratings</h3>
						<p class="text-sm text-muted mt-2 leading-relaxed">
							Rate channels out of five stars. Build a personal catalog of what you love.
						</p>
					</div>
					<div class="p-6 rounded-2xl border border-border bg-white dark:bg-gray-900">
						<div class="w-10 h-10 rounded-xl bg-accent/10 flex items-center justify-center mb-4">
							<svg class="w-5 h-5 text-accent" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
								<path d="M12 20h9" />
								<path d="M16.5 3.5a2.121 2.121 0 0 1 3 3L7 19l-4 1 1-4L16.5 3.5z" />
							</svg>
						</div>
						<h3 class="font-semibold">Notes</h3>
						<p class="text-sm text-muted mt-2 leading-relaxed">
							Write thoughts about any channel. Remember why you subscribed.
						</p>
					</div>
					<div class="p-6 rounded-2xl border border-border bg-white dark:bg-gray-900">
						<div class="w-10 h-10 rounded-xl bg-accent/10 flex items-center justify-center mb-4">
							<svg class="w-5 h-5 text-accent" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
								<circle cx="11" cy="11" r="8" />
								<path d="m21 21-4.3-4.3" />
							</svg>
						</div>
						<h3 class="font-semibold">Search</h3>
						<p class="text-sm text-muted mt-2 leading-relaxed">
							Search across channels, notes, and collections. Find anything instantly.
						</p>
					</div>
					<div class="p-6 rounded-2xl border border-border bg-white dark:bg-gray-900">
						<div class="w-10 h-10 rounded-xl bg-accent/10 flex items-center justify-center mb-4">
							<svg class="w-5 h-5 text-accent" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
								<path d="M21 12a9 9 0 1 1-9-9" />
								<path d="M21 3v6h-6" />
								<path d="M21 3l-9 9" />
							</svg>
						</div>
						<h3 class="font-semibold">Sync</h3>
						<p class="text-sm text-muted mt-2 leading-relaxed">
							Import your YouTube subscriptions in one click. Always up to date.
						</p>
					</div>
					<div class="p-6 rounded-2xl border border-border bg-white dark:bg-gray-900">
						<div class="w-10 h-10 rounded-xl bg-accent/10 flex items-center justify-center mb-4">
							<svg class="w-5 h-5 text-accent" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
								<path d="M2 3h6a4 4 0 0 1 4 4v14a3 3 0 0 0-3-3H2z" />
								<path d="M22 3h-6a4 4 0 0 0-4 4v14a3 3 0 0 1 3-3h7z" />
							</svg>
						</div>
						<h3 class="font-semibold">Dark mode</h3>
						<p class="text-sm text-muted mt-2 leading-relaxed">
							Comfortable browsing at any hour. Follows your system preference.
						</p>
					</div>
				</div>
			</div>
		</section>

		<!-- CTA -->
		<section class="px-6 py-20 sm:py-28">
			<div class="max-w-2xl mx-auto text-center">
				<h2 class="text-3xl sm:text-4xl font-bold tracking-tight">Ready to organize your YouTube life?</h2>
				<p class="text-muted mt-4 text-lg">Sign in with Google and import your subscriptions in seconds.</p>
				<div class="mt-8">
					<a
						href="{apiUrl}/auth/login"
						class="inline-flex items-center gap-2 bg-accent text-white px-8 py-3.5 rounded-xl font-medium hover:bg-accent-hover transition-colors shadow-sm"
					>
						<svg class="w-5 h-5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
							<path d="M15 3h4a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2h-4" />
							<polyline points="10 17 15 12 10 7" />
							<line x1="15" y1="12" x2="3" y2="12" />
						</svg>
						Get started free
					</a>
				</div>
			</div>
		</section>
	</main>

	<!-- Footer -->
	<footer class="border-t border-border px-6 py-10">
		<div class="max-w-6xl mx-auto flex flex-col items-center gap-6">
			<div class="flex items-center gap-6">
				<a
					href="https://github.com/sahilium/tubedex"
					target="_blank"
					rel="noopener noreferrer"
					class="text-muted hover:text-gray-900 dark:hover:text-gray-100 transition-colors"
					aria-label="GitHub"
				>
					<svg class="w-5 h-5" viewBox="0 0 24 24" fill="currentColor">
						<path d="M12 0C5.37 0 0 5.37 0 12c0 5.31 3.435 9.795 8.205 11.385.6.105.825-.255.825-.57 0-.285-.015-1.23-.015-2.235-3.015.555-3.795-.735-4.035-1.41-.135-.345-.72-1.41-1.23-1.695-.42-.225-1.02-.78-.015-.795.945-.015 1.62.87 1.845 1.23 1.08 1.815 2.805 1.305 3.495.99.105-.78.42-1.305.765-1.605-2.67-.3-5.46-1.335-5.46-5.925 0-1.305.465-2.385 1.23-3.225-.12-.3-.54-1.53.12-3.18 0 0 1.005-.315 3.3 1.23.96-.27 1.98-.405 3-.405s2.04.135 3 .405c2.295-1.56 3.3-1.23 3.3-1.23.66 1.65.24 2.88.12 3.18.765.84 1.23 1.905 1.23 3.225 0 4.605-2.805 5.625-5.475 5.925.435.375.81 1.095.81 2.22 0 1.605-.015 2.895-.015 3.3 0 .315.225.69.825.57A12.02 12.02 0 0 0 24 12c0-6.63-5.37-12-12-12z"/>
					</svg>
				</a>
				<a
					href="https://github.com/sahilium/tubedex"
					target="_blank"
					rel="noopener noreferrer"
					class="text-xs text-muted hover:text-gray-900 dark:hover:text-gray-100 transition-colors"
				>
					GitHub
				</a>
			</div>
			<div class="flex flex-col sm:flex-row items-center gap-2 text-xs text-muted">
				<span>&copy; {new Date().getFullYear()} TubeDex</span>
				<span class="hidden sm:inline">·</span>
				<span>Made with YouTube Data API</span>
				<span class="hidden sm:inline">·</span>
				<span>Not affiliated with Google</span>
			</div>
		</div>
	</footer>
</div>

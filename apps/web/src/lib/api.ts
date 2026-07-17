const API_BASE = '/api/v1';

interface RequestOptions {
	method?: string;
	body?: unknown;
	params?: Record<string, string>;
}

export class ApiError extends Error {
	constructor(public status: number, message: string) {
		super(message);
	}
}

async function request<T>(path: string, options: RequestOptions = {}): Promise<T> {
	const url = new URL(`${API_BASE}${path}`, window.location.origin);

	if (options.params) {
		for (const [key, value] of Object.entries(options.params)) {
			url.searchParams.set(key, value);
		}
	}

	const res = await fetch(url, {
		method: options.method || 'GET',
		headers: {
			'Content-Type': 'application/json',
		},
		credentials: 'include',
		body: options.body ? JSON.stringify(options.body) : undefined,
	});

	if (!res.ok) {
		const data = await res.json().catch(() => ({ error: 'Unknown error' }));
		throw new ApiError(res.status, data.error || 'Request failed');
	}

	if (res.status === 204) {
		return undefined as T;
	}

	return res.json();
}

export const api = {
	get: <T>(path: string, params?: Record<string, string>) =>
		request<T>(path, { params }),

	post: <T>(path: string, body?: unknown) =>
		request<T>(path, { method: 'POST', body }),

	put: <T>(path: string, body?: unknown) =>
		request<T>(path, { method: 'PUT', body }),

	patch: <T>(path: string, body?: unknown) =>
		request<T>(path, { method: 'PATCH', body }),

	delete: <T>(path: string, body?: unknown) =>
		request<T>(path, { method: 'DELETE', body }),
};

// Types
export interface User {
	id: number;
	google_id: string;
	email: string;
	name: string;
	avatar_url: string;
}

export interface Channel {
	id: number;
	youtube_channel_id: string;
	name: string;
	handle: string;
	description: string;
	avatar: string;
	banner: string;
	subscriber_count: number;
}

export interface ChannelWithDetails extends Channel {
	rating: number | null;
	note: string | null;
	collection_ids: number[];
}

export interface Collection {
	id: number;
	user_id: number;
	name: string;
	icon: string;
	color: string;
}

export interface SyncJob {
	id: number;
	user_id: number;
	started_at: string;
	finished_at: string | null;
	status: string;
	error: string;
}

export interface SearchResult {
	type: 'channel' | 'collection' | 'note';
	id: number;
	name: string;
	avatar: string;
	handle: string;
	description: string;
}

export interface Note {
	id: number;
	user_id: number;
	channel_id: number;
	body: string;
}

export interface Rating {
	id: number;
	user_id: number;
	channel_id: number;
	rating: number;
}

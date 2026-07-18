import { dev } from '$app/environment';
import { PUBLIC_API_URL } from '$env/static/public';

const API_BASE = dev ? '/api/v1' : (PUBLIC_API_URL || '/api/v1');

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
	const base = PUBLIC_API_URL || '/api/v1';
	const url = new URL(`${base}${path}`, window.location.origin);

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

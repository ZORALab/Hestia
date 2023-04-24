{{- /*
Copyright 2022 "Holloway" Chew, Kean Ho <kean.ho.chew@zoralab.com>
Copyright 2022 ZORALab Enterprise <tech@zoralab.com>


Licensed under the Apache License, Version 2.0 (the "License"); you may not use
this file except in compliance with the License. You may obtain a copy of the
License at

                 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed
under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
CONDITIONS OF ANY KIND, either express or implied. See the License for the
specific language governing permissions and limitations under the License.
*/ -}}
{{- /* INPUT PARAMETERS */ -}}
{{- /* . = Page data in Hestia Structure */ -}}




// Constant PWA Definitions
const POLICY_CACHE_ONLY = 'cache-only';
const POLICY_CACHE_FIRST = 'cache-first';
const POLICY_NETWORK_ONLY = 'network-only';




// Generated PWA Definitions
let OFFLINE_CACHE = '{{- safeJS .PWA.Caches.ID -}}';
let OFFLINE_RESOURCES = {};

{{- $ret := partial "Hestia/functions/data/caches/offline/Load" . -}}
{{- range $url, $policy := $ret }}
OFFLINE_RESOURCES['{{ safeJS (string $url) -}}'] = '{{- safeJS (string $policy) -}}';
{{- end }}




// PWA Caching Mechanisms
const __putInCache = async (request, response) => {
	const cache = await caches.open(OFFLINE_CACHE);
	await cache.put(request, response);
}

const _fetchNetworkOnly = async (request) => {
	try {
		const networkResponse = await fetch(request);
		__putInCache(request, networkResponse.clone());
		return networkResponse;
	} catch {
		return new Response("Network Error", {
			status: 408,
			handlers: { "Content-Type": "text-plain" },
		});
	}
}

const _fetchCacheOnly = async (request) => {
	const cachedResponse = await caches.match(request);
	if (cachedResponse) {
		return cachedResponse;
	}

	return new Response("Asset Not Found", {
		status: 404,
		handlers: { "Content-Type": "text-plain" },
	});
}

const _fetchNetworkFirst = async (request) => {
	try {
		// source from network
		const networkResponse = await fetch(request);
		__putInCache(request, networkResponse.clone());
		return networkResponse;
	} catch {
		// source from cache
		const cachedResponse = await caches.match(request);
		if (cachedResponse) {
			return cachedResponse;
		}
	}

	// both failed - return with error response
	return new Response("Network Error", {
		status: 408,
		handlers: { "Content-Type": "text-plain" },
	});
}


const _fetchCacheFirst = async (request) => {
	// source from cache
	const cachedResponse = await caches.match(request);
	if (cachedResponse) {
		return cachedResponse;
	}

	try {
		// fallback to network
		const networkResponse = await fetch(request);
		__putInCache(request, networkResponse.clone());
		return networkResponse;
	} catch {
		// both failed -return with error response
		return new Response("Network Error", {
			status: 408,
			handlers: { "Content-Type": "text-plain" },
		});
	}
}


self.addEventListener("fetch", event => {
	switch (OFFLINE_RESOURCES[event.request.url]) {
	case POLICY_CACHE_FIRST:
		event.respondWith(_fetchCacheFirst(event.request));
		break;
	case POLICY_CACHE_ONLY:
		event.respondWith(_fetchCacheOnly(event.request));
		break;
	case POLICY_NETWORK_ONLY:
		event.respondWith(_fetchNetworkOnly(event.request));
		break;
	default: // POLICY_NETWORK_FIRST
		event.respondWith(_fetchNetworkFirst(event.request));
		break;
	}
});




// PWA Install Mechanism
self.addEventListener("install", event => {
	event.waitUntil(async () => {
		var resources = [];

		for(let key in OFFLINE_RESOURCES) {
			resources.push(key);
		}

		const cache = await caches.open(OFFLINE_CACHE);
		return cache.addAll(resources);
	});
});




// PWA Activate Mechanism
self.addEventListener("activate", event => {
	event.waitUntil(caches
	.keys()
	.then(keyList => {
		return Promise.all(keyList.map(key => {
			if (key === OFFLINE_CACHE) {
				return;
			}

			return caches.delete(key);
		}));
	}));
});

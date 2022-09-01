// Copyright 2022 "Holloway" Chew, Kean Ho <kean.ho.chew@zoralab.com>
// Copyright 2022 ZORALab Enterprise <tech@zoralab.com>
//
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not
// use this file except in compliance with the License. You may obtain a copy of
// the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations under
// the License.




// Constant PWA Definitions
//const POLICY_CACHE_FIRST = 'cache-first'; // commented out as it's default.
const POLICY_CACHE_ONLY = 'cache-only';
const POLICY_NETWORK_FIRST = 'network-first';
const POLICY_NETWORK_ONLY = 'network-only';




// Generated PWA Definitions
let OFFLINE_CACHE = "pwa-assets";  // Cache suite name
let OFFLINE_RESOURCES = {};        // List of offline resources and their caching policies;

{{- range $i, $v := .PWA.Caches.List }}
OFFLINE_RESOURCES['{{ safeJS (string $v.URL) -}}'] = '{{- safeJS (string $v.Cache) -}}';
{{- end }}




// PWA Caching Mechanisms
function _fetchByNetwork(event) {
	event.respondWith(fetch(event.request));
}


function _fetchByCache(event) {
	event.respondWith(caches.match(event.request));
}


function _fetchByNetworkFirst(event) {
	event.respondWith(
		caches
		.match(event.request)
		.then(cachedResponse => {
			const networkResponse = fetch(event.request)
			.then(response => {
				caches
				.open(OFFLINE_CACHE)
				.then(cache => {
					cache.put(event.request,
						response.clone()
					);
				});
			});

			return networkResponse || cachedResponse;
		})
	)
}


function _fetchByCacheFirst(event) {
	event.respondWith(
		caches
		.match(event.request)
		.then(cachedResponse => {
			const networkResponse = fetch(event.request)
			.then(response => {
				caches
				.open(OFFLINE_CACHE)
				.then(cache => {
					cache.put(event.request,
						response.clone()
					);
				});
			});

			return cachedResponse || networkResponse;
		})
	)
}


self.addEventListener("fetch", event => {
	switch (OFFLINE_RESOURCES[event.request.url]) {
	case POLICY_NETWORK_ONLY:
		_fetchByNetwork(event);
		break;
	case POLICY_NETWORK_FIRST:
		_fetchByNetworkFirst(event);
		break;
	case POLICY_CACHE_ONLY:
		_fetchByCache(event);
		break;
	default: // POLICY_CACHE_FIRST
		_fetchByCacheFirst(event);
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
	console.log("[  PWA  ] Service worker activated.");
});

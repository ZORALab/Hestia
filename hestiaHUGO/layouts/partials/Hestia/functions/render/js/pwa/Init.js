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
if('serviceWorker' in navigator) {
	navigator.serviceWorker
	.register('{{- safeJS .PWA.URL.Worker -}}')
	.then(registration => {
		if (registration.active) {
			console.log('[  PWA  ] SERVICE WORKER: active');
			return
		} else if (registration.installing) {
			worker = registration.installing;
		} else if (registration.waiting) {
			worker = registration.waiting;
		}

		if (worker) {
			worker.addEventListener("statechange", (e) => {
				console.log('[  PWA  ] SERVICE WORKER: ' + e.target.state);
			});
		}
	});
}

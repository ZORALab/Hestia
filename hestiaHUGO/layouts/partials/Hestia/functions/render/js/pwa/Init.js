if('serviceWorker' in navigator) {
	navigator.serviceWorker
	.register('{{- safeJS .PWA.Worker.URL -}}')
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

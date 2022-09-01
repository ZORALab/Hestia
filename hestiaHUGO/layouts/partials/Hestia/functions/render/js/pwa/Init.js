if('serviceWorker' in navigator) {
	navigator.serviceWorker
	.register('{{- safeJS .PWA.Worker.URL -}}')
	.then(function(registration) {
		console.log('[   OK   ] - Service Worker Registered.');
	});


	navigator.serviceWorker
	.ready
	.then(function(registration) {
		console.log('[   OK   ] - Service Worker Ready.');
	});
}

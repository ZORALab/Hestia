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




{{- if .WASM.Embed }}
var source = new Response(
	{{- $ret := merge . (dict "Input" (dict "Path" .WASM.Embed)) -}}
	{{- $ret = partial "Hestia/functions/data/filesystem/ParseFile" $ret -}}
	{{- if $ret.Console -}}
		{{- $console := printf "render: .WASM.Embed: %s" $ret.Console.Message -}}
		{{- $console = dict "Console" (dict "Message" $console) -}}
		{{- $console = merge . $console -}}
		{{- partial "Hestia/functions/console/Errorf" $console -}}
		{{- $ret = "" -}}
	{{- else -}}
		{{- $ret = $ret.Output -}}
	{{- end }}
	Uint8Array.from('{{- safeJS (printf "%X" (readFile $ret)) -}}'
		.match(/.{1,2}/g)
		.map((byte) => {
			return parseInt(byte, 16);
		})
	),
	{headers: { 'Content-Type': 'application/wasm'}}
);
	{{- $ret = false -}}
{{- else }}
var source = fetch('{{- safeJS .WASM.URL -}}');
{{- end }}




{{- if .WASM.Setup }}
{{ safeJS .WASM.Setup -}}
{{- end }}




// Initialize WASM
if (!WebAssembly.instantiateStreaming) { // polyfill
	WebAssembly.instantiateStreaming = async (resp, importObject) => {
		const source = await (await resp).arrayBuffer();
		return await WebAssembly.instantiate(source, importObject);
	};
}


WebAssembly
.instantiateStreaming(source
	{{- if .WASM.Import -}}, {{ safeJS .WASM.Import -}}{{- end -}}
)
.then(result => {
	{{ safeJS .WASM.Init }}
});

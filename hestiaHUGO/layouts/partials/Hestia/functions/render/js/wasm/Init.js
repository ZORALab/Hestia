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
.instantiateStreaming(fetch('{{- safeJS .WASM.URL -}}')
	{{- if .WASM.Import -}}, {{ safeJS .WASM.Import -}}{{- end -}}
)
.then(result => {
	{{ safeJS .WASM.Init }}
});

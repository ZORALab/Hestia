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




{{- /* prepare working variables for this function */ -}}
{{- $dataList := "" -}}
{{- $path := "" -}}




{{- /* process pathing */ -}}
{{- $path = lower (printf "%s%s" .URL.Current.Absolute.Path "app.js") -}}




{{- /* render output */ -}}
{{- $dataList = string (partial "Hestia/functions/render/js/pwa/worker.js" .) -}}
{{- if not .IsServerMode -}}
	{{- $dataList = dict "Type" "application/javascript" "Data" $dataList -}}
	{{- $dataList = merge . (dict "Input" $dataList) -}}
	{{- $dataList = partial "Hestia/functions/data/string/Minify" $dataList -}}
	{{- $dataList = $dataList.Output -}}
{{- end -}}
{{- $dataList = merge . (dict "Input" (dict "Content" (string $dataList) "Path" $path)) -}}
{{- $dataList = partial "Hestia/functions/data/filesystem/WriteFile" $dataList -}}

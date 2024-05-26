{{- define "base_element" -}}
<div class="container">
    <div class="item" id="{{ .ChartID }}" style="width:{{ .Initialization.Width }};height:{{ .Initialization.Height }};"></div>
</div>
{{- end -}}

{{- define "base_script" -}}
<script type="text/javascript">
    "use strict";
    let goecharts_{{ .ChartID | safeJS }} = echarts.init(document.getElementById('{{ .ChartID | safeJS }}'), "{{ .Theme }}", { renderer: "{{  .Initialization.Renderer }}" });
    let option_{{ .ChartID | safeJS }} = {{ template "base_option" . }}
    goecharts_{{ .ChartID | safeJS }}.setOption(option_{{ .ChartID | safeJS }});

  {{- range  $listener := .EventListeners }}
    {{if .Query  }}
    goecharts_{{ $.ChartID | safeJS }}.on({{ $listener.EventName }}, {{ $listener.Query | safeJS }}, {{ injectInstance $listener.Handler "%MY_ECHARTS%"  $.ChartID | safeJS }});
    {{ else }}
    goecharts_{{ $.ChartID | safeJS }}.on({{ $listener.EventName }}, {{ injectInstance $listener.Handler "%MY_ECHARTS%"  $.ChartID | safeJS }})
    {{ end }}
  {{- end }}

    {{- range .JSFunctions.Fns }}
    {{ injectInstance . "%MY_ECHARTS%"  $.ChartID  | safeJS }}
    {{- end }}
</script>
{{- end -}}

{{- define "base_option" }}
    {{- .JSONNotEscaped | safeJS }}
{{- end }};

{{- define "table" }}
	<style>
		table, th, td {
			border:1px solid black;
		}
    h3 {
      align: center
    }
	</style>

  <center><h3>Eigenlayer TVL Breakdown</h3></center>
	<div class="container">
		<table>
		{{- range .Table }}
			<tr>
			{{- range . }}
				<td>{{ . }}</td>
			{{- end }}
			</tr>
		{{- end }}
		</table>
	</div>
{{- end }}

{{- define "base" }}
    {{- template "base_element" . }}
    {{- template "base_script" . }}
    {{- template "table" . }}
{{- end }}

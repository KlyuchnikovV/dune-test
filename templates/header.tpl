{{ define "header" }}
<head>
  <meta charset="utf-8">
  <title>{{ .PageTitle }} --> This is my own style template 🐶</title>
{{- range .JSAssets.Values }}
  <script src="{{ . }}"></script>
{{- end }}
{{- range .CSSAssets.Values }}
  <link href="{{ . }}" rel="stylesheet">
{{- end }}
</head>
{{ end }}

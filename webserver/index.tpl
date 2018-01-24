<!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8">
		<title>Heat Control</title>
	</head>
	<body>
		<h1>{{.Title}}</h1>

		<div id="sidebar">
			{{block "sidebar" .}}
			<h2>Links</h2>
			{{/* The dashes in the following template directives
			     ensure the generated HTML of this list contains no
			     extraneous spaces or line breaks. */}}
			<ul>
				{{- range .Links}}
				<li><a href="{{.URL}}">{{.Title}}</a></li>
				{{- end}}
			</ul>
			{{end}}
		</div>

		{{block "content" .}}
		<div id="content">
			{{.Body}}
		</div>
		{{end}}
	</body>
</html>
# filepath: /Volumes/VijaysSSD/private/go-learning-path/projects/shortit/deployments/templates/_helpers.tpl
{{- define "shortit.dockerconfig" -}}
{
  "auths": {
    "docker.io": {
      "username": "{{ .Values.dockerHub.username }}",
      "password": "{{ .Values.dockerHub.password }}",
      "email": "{{ .Values.dockerHub.email }}",
      "auth": "{{ printf "%s:%s" .Values.dockerHub.username .Values.dockerHub.password | b64enc }}"
    }
  }
}
{{- end -}}
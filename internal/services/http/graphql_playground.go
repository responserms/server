// Copyright (c) 2021 Contaim, LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package http

import (
	"fmt"
	"html/template"
	"net/http"
)

var playgroundTemplateContent = `<!DOCTYPE html>
 <html>
 <head>
	 <meta charset=utf-8/>
	 <meta name="viewport" content="user-scalable=no, initial-scale=1.0, minimum-scale=1.0, maximum-scale=1.0, minimal-ui">
	 <link rel="shortcut icon" href="https://graphcool-playground.netlify.com/favicon.png">
	 <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/graphql-playground-react@{{ .version }}/build/static/css/index.css"
		 integrity="{{ .cssSRI }}" crossorigin="anonymous"/>
	 <link rel="shortcut icon" href="https://cdn.jsdelivr.net/npm/graphql-playground-react@{{ .version }}/build/favicon.png"
		 integrity="{{ .faviconSRI }}" crossorigin="anonymous"/>
	 <script src="https://cdn.jsdelivr.net/npm/graphql-playground-react@{{ .version }}/build/static/js/middleware.js"
		 integrity="{{ .jsSRI }}" crossorigin="anonymous"></script>
	 <title>{{.title}}</title>
 </head>
 <body>
 <style type="text/css">
	 html { font-family: "Open Sans", sans-serif; overflow: hidden; }
	 body { margin: 0; background: #172a3a; }
 </style>
 <div id="root"></div>
 <script type="text/javascript">
	window.addEventListener('load', function (event) {
		const root = document.getElementById('root');
		root.classList.add('playgroundIn');
		const wsProto = location.protocol == 'https:' ? 'wss:' : 'ws:'
		GraphQLPlayground.init(root, {
			endpoint: location.protocol + '//' + location.host + '{{.endpoint}}',
			subscriptionsEndpoint: wsProto + '//' + location.host + '{{.endpoint }}',
			shareEnabled: true,
			settings: {
				'schema.polling.enable': false,
				'request.credentials': 'same-origin'
			}
		})
	})
 </script>
 </body>
 </html>`

func (s *Server) registerGraphQLPlayground() error {
	playgroundTemplate, err := template.New("playground").Parse(playgroundTemplateContent)
	if err != nil {
		return fmt.Errorf("register graphql playground: %w", err)
	}

	s.mux.HandleFunc("/api", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "text/html")

		err := playgroundTemplate.Execute(rw, map[string]string{
			"title":      "Response GraphQL Playground",
			"endpoint":   "/api/graphql",
			"version":    "1.7.20",
			"cssSRI":     "sha256-cS9Vc2OBt9eUf4sykRWukeFYaInL29+myBmFDSa7F/U=",
			"faviconSRI": "sha256-GhTyE+McTU79R4+pRO6ih+4TfsTOrpPwD8ReKFzb3PM=",
			"jsSRI":      "sha256-4QG1Uza2GgGdlBL3RCBCGtGeZB6bDbsw8OltCMGeJsA=",
		})

		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
		}
	})

	return nil
}

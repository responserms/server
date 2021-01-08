// Copyright (c) 2021 Contaim, LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/apollotracing"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/facebookincubator/ent-contrib/entgql"
	"github.com/gorilla/websocket"
	"github.com/responserms/server/graphql/resolvers"
	"github.com/responserms/server/graphql/server"
)

func (s *Server) registerGraphQL(svcs Backend) error {
	res, err := resolvers.New(s.logger, svcs)
	if err != nil {
		return fmt.Errorf("register graphql: %w", err)
	}

	graphql := handler.New(server.NewExecutableSchema(server.Config{
		Resolvers: res,
		// Directives: server.DirectiveRoot{
		// 	IsAuthenticated: directives.NewIsAuthenticated(),
		// 	IsResourceOwner: directives.NewIsResourceOwner(),
		// },
	}))

	// add extensions
	graphql.Use(extension.Introspection{})
	graphql.Use(apollotracing.Tracer{})
	graphql.Use(entgql.Transactioner{TxOpener: svcs.Database()})

	// add transports
	graphql.AddTransport(transport.POST{})
	graphql.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	})

	// this needs to auth in the future
	s.mux.Handle("/api/graphql", graphql)

	return nil
}

// Copyright (c) Abstract Machines
// SPDX-License-Identifier: Apache-2.0

package api

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"strings"

	"github.com/absmach/magistrala/re"
	"github.com/absmach/supermq"
	api "github.com/absmach/supermq/api/http"
	apiutil "github.com/absmach/supermq/api/http/util"
	mgauthn "github.com/absmach/supermq/pkg/authn"
	"github.com/absmach/supermq/pkg/errors"
	"github.com/go-chi/chi/v5"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

const (
	ruleIdKey       = "ruleID"
	reportIdKey     = "reportID"
	inputChannelKey = "input_channel"
	statusKey       = "status"
	actionKey       = "action"
	defAction       = "view"
)

// MakeHandler creates an HTTP handler for the service endpoints.
func MakeHandler(svc re.Service, authn mgauthn.Authentication, mux *chi.Mux, logger *slog.Logger, instanceID string) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(apiutil.LoggingErrorEncoder(logger, api.EncodeError)),
	}
	mux.Group(func(r chi.Router) {
		r.Use(api.AuthenticateMiddleware(authn, true))
		r.Route("/{domainID}", func(r chi.Router) {
			r.Route("/rules", func(r chi.Router) {
				r.Post("/", otelhttp.NewHandler(kithttp.NewServer(
					addRuleEndpoint(svc),
					decodeAddRuleRequest,
					api.EncodeResponse,
					opts...,
				), "create_rule").ServeHTTP)

				r.Get("/", otelhttp.NewHandler(kithttp.NewServer(
					listRulesEndpoint(svc),
					decodeListRulesRequest,
					api.EncodeResponse,
					opts...,
				), "list_rules").ServeHTTP)

				r.Route("/{ruleID}", func(r chi.Router) {
					r.Get("/", otelhttp.NewHandler(kithttp.NewServer(
						viewRuleEndpoint(svc),
						decodeViewRuleRequest,
						api.EncodeResponse,
						opts...,
					), "view_rule").ServeHTTP)

					r.Patch("/", otelhttp.NewHandler(kithttp.NewServer(
						updateRuleEndpoint(svc),
						decodeUpdateRuleRequest,
						api.EncodeResponse,
						opts...,
					), "update_rule").ServeHTTP)

					r.Patch("/tags", otelhttp.NewHandler(kithttp.NewServer(
						updateRuleTagsEndpoint(svc),
						decodeUpdateRuleTags,
						api.EncodeResponse,
						opts...,
					), "update_rule_tags").ServeHTTP)

					r.Patch("/schedule", otelhttp.NewHandler(kithttp.NewServer(
						updateRuleScheduleEndpoint(svc),
						decodeUpdateRuleScheduleRequest,
						api.EncodeResponse,
						opts...,
					), "update_rule_scheduler").ServeHTTP)

					r.Delete("/", otelhttp.NewHandler(kithttp.NewServer(
						deleteRuleEndpoint(svc),
						decodeDeleteRuleRequest,
						api.EncodeResponse,
						opts...,
					), "delete_rule").ServeHTTP)

					r.Post("/enable", otelhttp.NewHandler(kithttp.NewServer(
						enableRuleEndpoint(svc),
						decodeUpdateRuleStatusRequest,
						api.EncodeResponse,
						opts...,
					), "enable_rule").ServeHTTP)

					r.Post("/disable", otelhttp.NewHandler(kithttp.NewServer(
						disableRuleEndpoint(svc),
						decodeUpdateRuleStatusRequest,
						api.EncodeResponse,
						opts...,
					), "disable_rule").ServeHTTP)
				})
			})
		})
	})

	mux.Get("/health", supermq.Health("rule_engine", instanceID))
	mux.Handle("/metrics", promhttp.Handler())

	return mux
}

func decodeAddRuleRequest(_ context.Context, r *http.Request) (interface{}, error) {
	if !strings.Contains(r.Header.Get("Content-Type"), api.ContentType) {
		return nil, errors.Wrap(apiutil.ErrValidation, apiutil.ErrUnsupportedContentType)
	}
	var rule re.Rule
	if err := json.NewDecoder(r.Body).Decode(&rule); err != nil {
		return nil, errors.Wrap(apiutil.ErrValidation, errors.Wrap(errors.ErrMalformedEntity, err))
	}
	return addRuleReq{Rule: rule}, nil
}

func decodeViewRuleRequest(_ context.Context, r *http.Request) (interface{}, error) {
	id := chi.URLParam(r, ruleIdKey)
	return viewRuleReq{id: id}, nil
}

func decodeUpdateRuleRequest(_ context.Context, r *http.Request) (interface{}, error) {
	if !strings.Contains(r.Header.Get("Content-Type"), api.ContentType) {
		return nil, errors.Wrap(apiutil.ErrValidation, apiutil.ErrUnsupportedContentType)
	}
	var rule re.Rule
	if err := json.NewDecoder(r.Body).Decode(&rule); err != nil {
		return nil, errors.Wrap(apiutil.ErrValidation, errors.Wrap(errors.ErrMalformedEntity, err))
	}
	rule.ID = chi.URLParam(r, ruleIdKey)

	return updateRuleReq{Rule: rule}, nil
}

func decodeUpdateRuleTags(_ context.Context, r *http.Request) (interface{}, error) {
	if !strings.Contains(r.Header.Get("Content-Type"), api.ContentType) {
		return nil, errors.Wrap(apiutil.ErrValidation, apiutil.ErrUnsupportedContentType)
	}

	req := updateRuleTagsReq{
		id: chi.URLParam(r, ruleIdKey),
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, errors.Wrap(apiutil.ErrValidation, errors.Wrap(errors.ErrMalformedEntity, err))
	}

	return req, nil
}

func decodeUpdateRuleScheduleRequest(_ context.Context, r *http.Request) (interface{}, error) {
	if !strings.Contains(r.Header.Get("Content-Type"), api.ContentType) {
		return nil, errors.Wrap(apiutil.ErrValidation, apiutil.ErrUnsupportedContentType)
	}

	req := updateRuleScheduleReq{
		id: chi.URLParam(r, ruleIdKey),
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, errors.Wrap(apiutil.ErrValidation, errors.Wrap(errors.ErrMalformedEntity, err))
	}

	return req, nil
}

func decodeUpdateRuleStatusRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := updateRuleStatusReq{
		id: chi.URLParam(r, ruleIdKey),
	}

	return req, nil
}

func decodeListRulesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	offset, err := apiutil.ReadNumQuery[uint64](r, api.OffsetKey, api.DefOffset)
	if err != nil {
		return nil, errors.Wrap(apiutil.ErrValidation, err)
	}
	limit, err := apiutil.ReadNumQuery[uint64](r, api.LimitKey, api.DefLimit)
	if err != nil {
		return nil, errors.Wrap(apiutil.ErrValidation, err)
	}
	name, err := apiutil.ReadStringQuery(r, api.NameKey, "")
	if err != nil {
		return nil, errors.Wrap(apiutil.ErrValidation, err)
	}
	ic, err := apiutil.ReadStringQuery(r, inputChannelKey, "")
	if err != nil {
		return nil, errors.Wrap(apiutil.ErrValidation, err)
	}
	s, err := apiutil.ReadStringQuery(r, api.StatusKey, api.DefStatus)
	if err != nil {
		return nil, errors.Wrap(apiutil.ErrValidation, err)
	}
	dir, err := apiutil.ReadStringQuery(r, api.DirKey, api.DefDir)
	if err != nil {
		return nil, errors.Wrap(apiutil.ErrValidation, err)
	}
	st, err := re.ToStatus(s)
	if err != nil {
		return nil, errors.Wrap(apiutil.ErrValidation, err)
	}
	tag, err := apiutil.ReadStringQuery(r, api.TagKey, "")
	if err != nil {
		return nil, errors.Wrap(apiutil.ErrValidation, err)
	}

	return listRulesReq{
		PageMeta: re.PageMeta{
			Offset:       offset,
			Limit:        limit,
			Name:         name,
			InputChannel: ic,
			Status:       st,
			Dir:          dir,
			Tag:          tag,
		},
	}, nil
}

func decodeDeleteRuleRequest(_ context.Context, r *http.Request) (interface{}, error) {
	id := chi.URLParam(r, ruleIdKey)

	return deleteRuleReq{id: id}, nil
}

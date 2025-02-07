// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package oauth

import (
	"net/http"
	"net/url"

	osin "github.com/openshift/osin"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/events"
	"go.thethings.network/lorawan-stack/v3/pkg/webhandlers"
)

var (
	errInvalidLogoutRedirectURI = errors.DefineInvalidArgument(
		"invalid_logout_redirect_uri",
		"the redirect URI did not match the one(s) defined in the client",
	)
	errMissingAccessTokenIDParam = errors.DefinePermissionDenied(
		"missing_param_access_token_id",
		"access token ID was not provided",
	)
)

func (s *server) ClientLogout(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	accessTokenID := r.URL.Query().Get("access_token_id")
	redirectURI := s.config.UI.MountPath()
	if accessTokenID == "" {
		webhandlers.Error(w, r, errMissingAccessTokenIDParam.New())
		return
	}
	at, err := s.store.GetAccessToken(ctx, accessTokenID)
	if err != nil && !errors.IsNotFound(err) {
		webhandlers.Error(w, r, err)
		return
	}
	if at != nil {
		client, err := s.store.GetClient(ctx, at.ClientIds, []string{"logout_redirect_uris"})
		if err != nil {
			webhandlers.Error(w, r, err)
			return
		}
		if err = s.store.DeleteAccessToken(ctx, accessTokenID); err != nil {
			webhandlers.Error(w, r, err)
			return
		}
		events.Publish(evtAccessTokenDeleted.NewWithIdentifiersAndData(ctx, at.UserIds, nil))
		err = s.store.DeleteSession(ctx, at.UserIds, at.UserSessionId)
		if err != nil && !errors.IsNotFound(err) {
			webhandlers.Error(w, r, err)
			return
		}
		events.Publish(EvtUserLogout.NewWithIdentifiersAndData(ctx, at.UserIds, nil))
		redirectParam := r.URL.Query().Get("post_logout_redirect_uri")
		if redirectParam == "" {
			if len(client.LogoutRedirectUris) != 0 {
				redirectURI = client.LogoutRedirectUris[0]
			}
		} else {
			for _, uri := range client.LogoutRedirectUris {
				redirectURI, err = osin.ValidateUri(uri, redirectParam)
				if err == nil {
					break
				}
			}
			if err != nil {
				webhandlers.Error(w, r, errInvalidLogoutRedirectURI.WithCause(err))
				return
			}
		}
	}
	r, session, err := s.session.Get(w, r)
	if err != nil && !errors.IsUnauthenticated(err) && !errors.IsNotFound(err) {
		webhandlers.Error(w, r, err)
		return
	}
	if session != nil {
		events.Publish(evtUserSessionTerminated.NewWithIdentifiersAndData(ctx, session.GetUserIds(), nil))
		if err = s.store.DeleteSession(ctx, session.GetUserIds(), session.SessionId); err != nil {
			webhandlers.Error(w, r, err)
			return
		}
	}
	s.session.RemoveAuthCookie(w, r)
	url, err := url.Parse(redirectURI)
	if err != nil {
		webhandlers.Error(w, r, err)
		return
	}
	http.Redirect(w, r, url.String(), http.StatusFound)
}

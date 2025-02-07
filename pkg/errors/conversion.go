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

package errors

import (
	"context"
	"crypto/x509"
	"net"
	"net/url"

	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc/status"
)

var (
	errNetInvalidAddr    = DefineInvalidArgument("net_invalid_addr", "{message}", "message", "temporary", "timeout")
	errNetAddr           = DefineUnavailable("net_addr", "{message}", "message", "temporary", "timeout")
	errNetDNS            = DefineUnavailable("net_dns", "{message}", "message", "temporary", "timeout", "not_found")
	errNetUnknownNetwork = DefineNotFound("net_unknown_network", "{message}", "message", "temporary", "timeout")
	errNetOperation      = DefineUnavailable("net_operation", "{message}", "message", "op", "net", "source", "address", "timeout", "temporary")

	errRequest = Define("request", "request to `{url}` failed", "op")
	errURL     = DefineInvalidArgument("url", "invalid url `{url}`", "op")

	errX509UnknownAuthority   = DefineUnavailable("x509_unknown_authority", "unknown certificate authority")
	errX509Hostname           = DefineUnavailable("x509_hostname", "certificate authorized names do not match the requested name", "host")
	errX509CertificateInvalid = DefineUnavailable("x509_certificate_invalid", "certificate invalid", "detail", "reason")

	errContextCancelled        = DefineCanceled("context_canceled", "context canceled")
	errContextDeadlineExceeded = DefineDeadlineExceeded("context_deadline_exceeded", "context deadline exceeded")
)

// From returns an *Error if it can be derived from the given input.
// For a nil error, false will be returned.
func From(err error) (out *Error, ok bool) {
	if err == nil {
		return nil, false
	}
	defer func() {
		if out != nil {
			copy := *out
			out = &copy
		}
	}()
	// Do not refactor this into a map, as value-type
	// errors are not usable as indices in a map.
	switch {
	case err == context.Canceled:
		return build(errContextCancelled, 0), true
	case err == context.DeadlineExceeded:
		return build(errContextDeadlineExceeded, 0), true
	}
	switch err := err.(type) {
	case *Error:
		if err == nil {
			return nil, false // This is invalid.
		}
		return err, true
	case *Definition:
		if err == nil {
			return nil, false // This is invalid.
		}
		return build(err, 0), true
	case ErrorDetails: // Received over an API.
		var e Error
		setErrorDetails(&e, err)
		return &e, true
	case interface{ GRPCStatus() *status.Status }:
		return FromGRPCStatus(err.GRPCStatus()), true
	case validationError:
		e := build(errValidation, 0).WithAttributes(
			"field", err.Field(),
			"reason", err.Reason(),
			"name", err.ErrorName(),
		)
		if cause := err.Cause(); cause != nil {
			e = e.WithCause(cause)
		}
		return e, true
	case *net.DNSError:
		e := build(errNetDNS, 0).WithAttributes(
			"not_found", err.IsNotFound,
		).WithAttributes(
			netErrorDetails(err)...,
		)
		return e, true
	case *net.AddrError:
		return build(errNetAddr, 0).WithAttributes(netErrorDetails(err)...), true
	case net.InvalidAddrError:
		return build(errNetInvalidAddr, 0).WithAttributes(netErrorDetails(err)...), true
	case net.UnknownNetworkError:
		return build(errNetUnknownNetwork, 0).WithAttributes(netErrorDetails(err)...), true
	case *net.OpError:
		// Do not use netErrorDetails(err) as err.Error() will panic if err.Err is nil.
		e := build(errNetOperation, 0).WithAttributes(
			"op", err.Op,
			"net", err.Net,
			"timeout", err.Timeout(),
			"temporary", err.Temporary(),
		)
		if err.Addr != nil {
			e = e.WithAttributes("address", err.Addr.String())
		}
		if err.Source != nil {
			e = e.WithAttributes("source", err.Source.String())
		}
		if err.Err != nil {
			e = e.WithCause(err.Err)
		}
		return e, true
	case *url.Error:
		definition := errRequest
		if err.Op == "parse" {
			definition = errURL
		}
		e := build(definition, 0).WithAttributes(
			"url", err.URL,
			"op", err.Op,
		)
		if err.Err != nil {
			e = e.WithCause(err.Err)
		}
		return e, true
	case x509.CertificateInvalidError:
		return build(errX509CertificateInvalid, 0).WithAttributes(
			"detail", err.Detail,
			"reason", err.Reason,
		), true
	case x509.UnknownAuthorityError:
		return build(errX509UnknownAuthority, 0), true
	case x509.HostnameError:
		return build(errX509Hostname, 0).WithAttributes(
			"host", err.Host,
		), true
	}
	return nil, false
}

// ErrorDetails that can be carried over API.
type ErrorDetails interface {
	Error() string
	Namespace() string
	Name() string
	MessageFormat() string
	PublicAttributes() map[string]interface{}
	CorrelationID() string
	Cause() error
	Code() uint32
	Details() []proto.Message
}

func setErrorDetails(err *Error, details ErrorDetails) {
	if err.Definition == nil {
		err.Definition = &Definition{}
	}
	if namespace := details.Namespace(); namespace != "" {
		err.namespace = namespace
	}
	if name := details.Name(); name != "" {
		err.name = name
	}
	if messageFormat := details.MessageFormat(); messageFormat != "" {
		err.messageFormat = messageFormat
		err.messageFormatArguments = messageFormatArguments(messageFormat)
		err.parsedMessageFormat, _ = formatter.Parse(messageFormat)
	}
	if attributes := details.PublicAttributes(); len(attributes) != 0 {
		err.attributes = attributes
		for attr := range attributes {
			err.publicAttributes = append(err.publicAttributes, attr)
		}
	}
	if correlationID := details.CorrelationID(); correlationID != "" {
		err.correlationID = correlationID
	}
	if cause := details.Cause(); cause != nil {
		err.cause, _ = From(cause)
	}
	if code := details.Code(); code != 0 {
		err.code = code
	}
	err.details = append(err.details, details.Details()...)
}

func netErrorDetails(err net.Error) []interface{} {
	return []interface{}{
		"message", err.Error(),
		"temporary", err.Temporary(),
		"timeout", err.Timeout(),
	}
}

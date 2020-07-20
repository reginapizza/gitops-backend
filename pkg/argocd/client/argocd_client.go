// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/rhd-gitops-examples/gitops-backend/pkg/argocd/client/account_service"
	"github.com/rhd-gitops-examples/gitops-backend/pkg/argocd/client/application_service"
	"github.com/rhd-gitops-examples/gitops-backend/pkg/argocd/client/certificate_service"
	"github.com/rhd-gitops-examples/gitops-backend/pkg/argocd/client/cluster_service"
	"github.com/rhd-gitops-examples/gitops-backend/pkg/argocd/client/project_service"
	"github.com/rhd-gitops-examples/gitops-backend/pkg/argocd/client/repo_creds_service"
	"github.com/rhd-gitops-examples/gitops-backend/pkg/argocd/client/repository_service"
	"github.com/rhd-gitops-examples/gitops-backend/pkg/argocd/client/session_service"
	"github.com/rhd-gitops-examples/gitops-backend/pkg/argocd/client/settings_service"
	"github.com/rhd-gitops-examples/gitops-backend/pkg/argocd/client/version_service"
)

// Default argocd HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http", "https"}

// NewHTTPClient creates a new argocd HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Argocd {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new argocd HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Argocd {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new argocd client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Argocd {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(Argocd)
	cli.Transport = transport
	cli.AccountService = account_service.New(transport, formats)
	cli.ApplicationService = application_service.New(transport, formats)
	cli.CertificateService = certificate_service.New(transport, formats)
	cli.ClusterService = cluster_service.New(transport, formats)
	cli.ProjectService = project_service.New(transport, formats)
	cli.RepoCredsService = repo_creds_service.New(transport, formats)
	cli.RepositoryService = repository_service.New(transport, formats)
	cli.SessionService = session_service.New(transport, formats)
	cli.SettingsService = settings_service.New(transport, formats)
	cli.VersionService = version_service.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// Argocd is a client for argocd
type Argocd struct {
	AccountService account_service.ClientService

	ApplicationService application_service.ClientService

	CertificateService certificate_service.ClientService

	ClusterService cluster_service.ClientService

	ProjectService project_service.ClientService

	RepoCredsService repo_creds_service.ClientService

	RepositoryService repository_service.ClientService

	SessionService session_service.ClientService

	SettingsService settings_service.ClientService

	VersionService version_service.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Argocd) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.AccountService.SetTransport(transport)
	c.ApplicationService.SetTransport(transport)
	c.CertificateService.SetTransport(transport)
	c.ClusterService.SetTransport(transport)
	c.ProjectService.SetTransport(transport)
	c.RepoCredsService.SetTransport(transport)
	c.RepositoryService.SetTransport(transport)
	c.SessionService.SetTransport(transport)
	c.SettingsService.SetTransport(transport)
	c.VersionService.SetTransport(transport)
}
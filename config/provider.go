/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/upbound/upjet/pkg/config"

	autoscaler "github.com/haarchri/provider-castai/config/autoscaler"
	ekscluster "github.com/haarchri/provider-castai/config/ekscluster"
	nodeconfiguration "github.com/haarchri/provider-castai/config/nodeconfiguration"
	nodeconfigurationdefault "github.com/haarchri/provider-castai/config/nodeconfigurationdefault"
	nodetemplate "github.com/haarchri/provider-castai/config/nodetemplate"
	rebalancingjob "github.com/haarchri/provider-castai/config/rebalancingjob"
	rebalancingschedule "github.com/haarchri/provider-castai/config/rebalancingschedule"
)

const (
	resourcePrefix = "castai"
	modulePath     = "github.com/haarchri/provider-castai"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		ekscluster.Configure,
		autoscaler.Configure,
		nodeconfiguration.Configure,
		nodeconfigurationdefault.Configure,
		nodetemplate.Configure,
		rebalancingjob.Configure,
		rebalancingschedule.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

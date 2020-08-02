package generated

import (
	iam "github.com/crossplane/provider-terraform-gcp/generated/iam/v1alpha1"
	"github.com/crossplane/terraform-provider-runtime/pkg/plugin"
)

// TODO: output this list somewhere in the codegen pipeline
var ResourceImplementations = []*plugin.Implementation{
	iam.ServiceAccountImplementation(),
}

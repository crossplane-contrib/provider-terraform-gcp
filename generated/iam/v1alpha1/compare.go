package v1alpha1

import (
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane/terraform-provider-runtime/pkg/plugin"
)

type resourceMerger struct{}

func (r *resourceMerger) MergeResources(kube xpresource.Managed, prov xpresource.Managed) plugin.MergeDescription {
	k := kube.(*ServiceAccount)
	p := prov.(*ServiceAccount)
	md := plugin.MergeDescription{}

	if *k.Spec.ForProvider.DisplayName != *p.Spec.ForProvider.DisplayName {
		if k.Spec.ForProvider.DisplayName == nil {
			k.Spec.ForProvider.DisplayName = p.Spec.ForProvider.DisplayName
			md.LateInitializedSpec = true
		} else {
			md.NeedsProviderUpdate = true
		}
	}

	if k.Status.AtProvider.Email != p.Status.AtProvider.Email {
		k.Status.AtProvider.Email = p.Status.AtProvider.Email
		md.StatusUpdated = true
	}
	if k.Status.AtProvider.Name != p.Status.AtProvider.Name {
		k.Status.AtProvider.Name = p.Status.AtProvider.Name
		md.StatusUpdated = true
	}
	if k.Status.AtProvider.Project != p.Status.AtProvider.Project {
		k.Status.AtProvider.Project = p.Status.AtProvider.Project
		md.StatusUpdated = true
	}
	if k.Status.AtProvider.UniqueID != p.Status.AtProvider.UniqueID {
		k.Status.AtProvider.UniqueID = p.Status.AtProvider.UniqueID
		md.StatusUpdated = true
	}
	for key, v := range p.Annotations {
		if k.Annotations[key] != v {
			k.Annotations[key] = v
			md.AnnotationsUpdated = true
		}
	}

	return md
}

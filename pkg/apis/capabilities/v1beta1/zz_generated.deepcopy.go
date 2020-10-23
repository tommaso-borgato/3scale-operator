// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1beta1

import (
	common "github.com/3scale/3scale-operator/pkg/common"
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApicastHostedSpec) DeepCopyInto(out *ApicastHostedSpec) {
	*out = *in
	if in.Authentication != nil {
		in, out := &in.Authentication, &out.Authentication
		*out = new(AuthenticationSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApicastHostedSpec.
func (in *ApicastHostedSpec) DeepCopy() *ApicastHostedSpec {
	if in == nil {
		return nil
	}
	out := new(ApicastHostedSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApicastSelfManagedSpec) DeepCopyInto(out *ApicastSelfManagedSpec) {
	*out = *in
	if in.Authentication != nil {
		in, out := &in.Authentication, &out.Authentication
		*out = new(AuthenticationSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.StagingPublicBaseURL != nil {
		in, out := &in.StagingPublicBaseURL, &out.StagingPublicBaseURL
		*out = new(string)
		**out = **in
	}
	if in.ProductionPublicBaseURL != nil {
		in, out := &in.ProductionPublicBaseURL, &out.ProductionPublicBaseURL
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApicastSelfManagedSpec.
func (in *ApicastSelfManagedSpec) DeepCopy() *ApicastSelfManagedSpec {
	if in == nil {
		return nil
	}
	out := new(ApicastSelfManagedSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppKeyAppIDAuthenticationSpec) DeepCopyInto(out *AppKeyAppIDAuthenticationSpec) {
	*out = *in
	if in.AppID != nil {
		in, out := &in.AppID, &out.AppID
		*out = new(string)
		**out = **in
	}
	if in.AppKey != nil {
		in, out := &in.AppKey, &out.AppKey
		*out = new(string)
		**out = **in
	}
	if in.CredentialsLoc != nil {
		in, out := &in.CredentialsLoc, &out.CredentialsLoc
		*out = new(string)
		**out = **in
	}
	if in.Security != nil {
		in, out := &in.Security, &out.Security
		*out = new(SecuritySpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppKeyAppIDAuthenticationSpec.
func (in *AppKeyAppIDAuthenticationSpec) DeepCopy() *AppKeyAppIDAuthenticationSpec {
	if in == nil {
		return nil
	}
	out := new(AppKeyAppIDAuthenticationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationPlanSpec) DeepCopyInto(out *ApplicationPlanSpec) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.AppsRequireApproval != nil {
		in, out := &in.AppsRequireApproval, &out.AppsRequireApproval
		*out = new(bool)
		**out = **in
	}
	if in.TrialPeriod != nil {
		in, out := &in.TrialPeriod, &out.TrialPeriod
		*out = new(int)
		**out = **in
	}
	if in.SetupFee != nil {
		in, out := &in.SetupFee, &out.SetupFee
		*out = new(string)
		**out = **in
	}
	if in.CostMonth != nil {
		in, out := &in.CostMonth, &out.CostMonth
		*out = new(string)
		**out = **in
	}
	if in.PricingRules != nil {
		in, out := &in.PricingRules, &out.PricingRules
		*out = make([]PricingRuleSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Limits != nil {
		in, out := &in.Limits, &out.Limits
		*out = make([]LimitSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationPlanSpec.
func (in *ApplicationPlanSpec) DeepCopy() *ApplicationPlanSpec {
	if in == nil {
		return nil
	}
	out := new(ApplicationPlanSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthenticationSpec) DeepCopyInto(out *AuthenticationSpec) {
	*out = *in
	if in.UserKeyAuthentication != nil {
		in, out := &in.UserKeyAuthentication, &out.UserKeyAuthentication
		*out = new(UserKeyAuthenticationSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.AppKeyAppIDAuthentication != nil {
		in, out := &in.AppKeyAppIDAuthentication, &out.AppKeyAppIDAuthentication
		*out = new(AppKeyAppIDAuthenticationSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthenticationSpec.
func (in *AuthenticationSpec) DeepCopy() *AuthenticationSpec {
	if in == nil {
		return nil
	}
	out := new(AuthenticationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Backend) DeepCopyInto(out *Backend) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Backend.
func (in *Backend) DeepCopy() *Backend {
	if in == nil {
		return nil
	}
	out := new(Backend)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Backend) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackendList) DeepCopyInto(out *BackendList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Backend, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackendList.
func (in *BackendList) DeepCopy() *BackendList {
	if in == nil {
		return nil
	}
	out := new(BackendList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackendList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackendSpec) DeepCopyInto(out *BackendSpec) {
	*out = *in
	if in.MappingRules != nil {
		in, out := &in.MappingRules, &out.MappingRules
		*out = make([]MappingRuleSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = make(map[string]MetricSpec, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Methods != nil {
		in, out := &in.Methods, &out.Methods
		*out = make(map[string]MethodSpec, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ProviderAccountRef != nil {
		in, out := &in.ProviderAccountRef, &out.ProviderAccountRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackendSpec.
func (in *BackendSpec) DeepCopy() *BackendSpec {
	if in == nil {
		return nil
	}
	out := new(BackendSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackendStatus) DeepCopyInto(out *BackendStatus) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(int64)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(common.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackendStatus.
func (in *BackendStatus) DeepCopy() *BackendStatus {
	if in == nil {
		return nil
	}
	out := new(BackendStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackendUsageSpec) DeepCopyInto(out *BackendUsageSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackendUsageSpec.
func (in *BackendUsageSpec) DeepCopy() *BackendUsageSpec {
	if in == nil {
		return nil
	}
	out := new(BackendUsageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LimitSpec) DeepCopyInto(out *LimitSpec) {
	*out = *in
	in.MetricMethodRef.DeepCopyInto(&out.MetricMethodRef)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LimitSpec.
func (in *LimitSpec) DeepCopy() *LimitSpec {
	if in == nil {
		return nil
	}
	out := new(LimitSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MappingRuleSpec) DeepCopyInto(out *MappingRuleSpec) {
	*out = *in
	if in.Last != nil {
		in, out := &in.Last, &out.Last
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MappingRuleSpec.
func (in *MappingRuleSpec) DeepCopy() *MappingRuleSpec {
	if in == nil {
		return nil
	}
	out := new(MappingRuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MethodSpec) DeepCopyInto(out *MethodSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MethodSpec.
func (in *MethodSpec) DeepCopy() *MethodSpec {
	if in == nil {
		return nil
	}
	out := new(MethodSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricMethodRefSpec) DeepCopyInto(out *MetricMethodRefSpec) {
	*out = *in
	if in.BackendSystemName != nil {
		in, out := &in.BackendSystemName, &out.BackendSystemName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricMethodRefSpec.
func (in *MetricMethodRefSpec) DeepCopy() *MetricMethodRefSpec {
	if in == nil {
		return nil
	}
	out := new(MetricMethodRefSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricSpec) DeepCopyInto(out *MetricSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricSpec.
func (in *MetricSpec) DeepCopy() *MetricSpec {
	if in == nil {
		return nil
	}
	out := new(MetricSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenAPIRefSpec) DeepCopyInto(out *OpenAPIRefSpec) {
	*out = *in
	if in.ConfigMapRef != nil {
		in, out := &in.ConfigMapRef, &out.ConfigMapRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
	if in.URL != nil {
		in, out := &in.URL, &out.URL
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenAPIRefSpec.
func (in *OpenAPIRefSpec) DeepCopy() *OpenAPIRefSpec {
	if in == nil {
		return nil
	}
	out := new(OpenAPIRefSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Openapi) DeepCopyInto(out *Openapi) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Openapi.
func (in *Openapi) DeepCopy() *Openapi {
	if in == nil {
		return nil
	}
	out := new(Openapi)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Openapi) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenapiList) DeepCopyInto(out *OpenapiList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Openapi, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenapiList.
func (in *OpenapiList) DeepCopy() *OpenapiList {
	if in == nil {
		return nil
	}
	out := new(OpenapiList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenapiList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenapiSpec) DeepCopyInto(out *OpenapiSpec) {
	*out = *in
	in.OpenAPIRef.DeepCopyInto(&out.OpenAPIRef)
	if in.ProviderAccountRef != nil {
		in, out := &in.ProviderAccountRef, &out.ProviderAccountRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenapiSpec.
func (in *OpenapiSpec) DeepCopy() *OpenapiSpec {
	if in == nil {
		return nil
	}
	out := new(OpenapiSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenapiStatus) DeepCopyInto(out *OpenapiStatus) {
	*out = *in
	if in.ProductResourceName != nil {
		in, out := &in.ProductResourceName, &out.ProductResourceName
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	if in.BackendResourceNames != nil {
		in, out := &in.BackendResourceNames, &out.BackendResourceNames
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(common.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenapiStatus.
func (in *OpenapiStatus) DeepCopy() *OpenapiStatus {
	if in == nil {
		return nil
	}
	out := new(OpenapiStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PricingRuleSpec) DeepCopyInto(out *PricingRuleSpec) {
	*out = *in
	in.MetricMethodRef.DeepCopyInto(&out.MetricMethodRef)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PricingRuleSpec.
func (in *PricingRuleSpec) DeepCopy() *PricingRuleSpec {
	if in == nil {
		return nil
	}
	out := new(PricingRuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Product) DeepCopyInto(out *Product) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Product.
func (in *Product) DeepCopy() *Product {
	if in == nil {
		return nil
	}
	out := new(Product)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Product) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProductDeploymentSpec) DeepCopyInto(out *ProductDeploymentSpec) {
	*out = *in
	if in.ApicastHosted != nil {
		in, out := &in.ApicastHosted, &out.ApicastHosted
		*out = new(ApicastHostedSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.ApicastSelfManaged != nil {
		in, out := &in.ApicastSelfManaged, &out.ApicastSelfManaged
		*out = new(ApicastSelfManagedSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProductDeploymentSpec.
func (in *ProductDeploymentSpec) DeepCopy() *ProductDeploymentSpec {
	if in == nil {
		return nil
	}
	out := new(ProductDeploymentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProductList) DeepCopyInto(out *ProductList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Product, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProductList.
func (in *ProductList) DeepCopy() *ProductList {
	if in == nil {
		return nil
	}
	out := new(ProductList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProductList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProductSpec) DeepCopyInto(out *ProductSpec) {
	*out = *in
	if in.Deployment != nil {
		in, out := &in.Deployment, &out.Deployment
		*out = new(ProductDeploymentSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.MappingRules != nil {
		in, out := &in.MappingRules, &out.MappingRules
		*out = make([]MappingRuleSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.BackendUsages != nil {
		in, out := &in.BackendUsages, &out.BackendUsages
		*out = make(map[string]BackendUsageSpec, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = make(map[string]MetricSpec, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Methods != nil {
		in, out := &in.Methods, &out.Methods
		*out = make(map[string]MethodSpec, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ApplicationPlans != nil {
		in, out := &in.ApplicationPlans, &out.ApplicationPlans
		*out = make(map[string]ApplicationPlanSpec, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.ProviderAccountRef != nil {
		in, out := &in.ProviderAccountRef, &out.ProviderAccountRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProductSpec.
func (in *ProductSpec) DeepCopy() *ProductSpec {
	if in == nil {
		return nil
	}
	out := new(ProductSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProductStatus) DeepCopyInto(out *ProductStatus) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(int64)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(common.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProductStatus.
func (in *ProductStatus) DeepCopy() *ProductStatus {
	if in == nil {
		return nil
	}
	out := new(ProductStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecuritySpec) DeepCopyInto(out *SecuritySpec) {
	*out = *in
	if in.HostHeader != nil {
		in, out := &in.HostHeader, &out.HostHeader
		*out = new(string)
		**out = **in
	}
	if in.SecretToken != nil {
		in, out := &in.SecretToken, &out.SecretToken
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecuritySpec.
func (in *SecuritySpec) DeepCopy() *SecuritySpec {
	if in == nil {
		return nil
	}
	out := new(SecuritySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserKeyAuthenticationSpec) DeepCopyInto(out *UserKeyAuthenticationSpec) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.CredentialsLoc != nil {
		in, out := &in.CredentialsLoc, &out.CredentialsLoc
		*out = new(string)
		**out = **in
	}
	if in.Security != nil {
		in, out := &in.Security, &out.Security
		*out = new(SecuritySpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserKeyAuthenticationSpec.
func (in *UserKeyAuthenticationSpec) DeepCopy() *UserKeyAuthenticationSpec {
	if in == nil {
		return nil
	}
	out := new(UserKeyAuthenticationSpec)
	in.DeepCopyInto(out)
	return out
}

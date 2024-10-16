/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package applyconfiguration

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	testing "k8s.io/client-go/testing"
	v1alpha1 "sigs.k8s.io/kueue/apis/kueue/v1alpha1"
	v1beta1 "sigs.k8s.io/kueue/apis/kueue/v1beta1"
	visibilityv1alpha1 "sigs.k8s.io/kueue/apis/visibility/v1alpha1"
	visibilityv1beta1 "sigs.k8s.io/kueue/apis/visibility/v1beta1"
	internal "sigs.k8s.io/kueue/client-go/applyconfiguration/internal"
	kueuev1alpha1 "sigs.k8s.io/kueue/client-go/applyconfiguration/kueue/v1alpha1"
	kueuev1beta1 "sigs.k8s.io/kueue/client-go/applyconfiguration/kueue/v1beta1"
	applyconfigurationvisibilityv1alpha1 "sigs.k8s.io/kueue/client-go/applyconfiguration/visibility/v1alpha1"
	applyconfigurationvisibilityv1beta1 "sigs.k8s.io/kueue/client-go/applyconfiguration/visibility/v1beta1"
)

// ForKind returns an apply configuration type for the given GroupVersionKind, or nil if no
// apply configuration type exists for the given GroupVersionKind.
func ForKind(kind schema.GroupVersionKind) interface{} {
	switch kind {
	// Group=kueue.x-k8s.io, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithKind("Topology"):
		return &kueuev1alpha1.TopologyApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("TopologyLevel"):
		return &kueuev1alpha1.TopologyLevelApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("TopologySpec"):
		return &kueuev1alpha1.TopologySpecApplyConfiguration{}

		// Group=kueue.x-k8s.io, Version=v1beta1
	case v1beta1.SchemeGroupVersion.WithKind("Admission"):
		return &kueuev1beta1.AdmissionApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("AdmissionCheck"):
		return &kueuev1beta1.AdmissionCheckApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("AdmissionCheckParametersReference"):
		return &kueuev1beta1.AdmissionCheckParametersReferenceApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("AdmissionCheckSpec"):
		return &kueuev1beta1.AdmissionCheckSpecApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("AdmissionChecksStrategy"):
		return &kueuev1beta1.AdmissionChecksStrategyApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("AdmissionCheckState"):
		return &kueuev1beta1.AdmissionCheckStateApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("AdmissionCheckStatus"):
		return &kueuev1beta1.AdmissionCheckStatusApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("AdmissionCheckStrategyRule"):
		return &kueuev1beta1.AdmissionCheckStrategyRuleApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("BorrowWithinCohort"):
		return &kueuev1beta1.BorrowWithinCohortApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("ClusterQueue"):
		return &kueuev1beta1.ClusterQueueApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("ClusterQueuePendingWorkload"):
		return &kueuev1beta1.ClusterQueuePendingWorkloadApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("ClusterQueuePendingWorkloadsStatus"):
		return &kueuev1beta1.ClusterQueuePendingWorkloadsStatusApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("ClusterQueuePreemption"):
		return &kueuev1beta1.ClusterQueuePreemptionApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("ClusterQueueSpec"):
		return &kueuev1beta1.ClusterQueueSpecApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("ClusterQueueStatus"):
		return &kueuev1beta1.ClusterQueueStatusApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("FairSharing"):
		return &kueuev1beta1.FairSharingApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("FairSharingStatus"):
		return &kueuev1beta1.FairSharingStatusApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("FlavorFungibility"):
		return &kueuev1beta1.FlavorFungibilityApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("FlavorQuotas"):
		return &kueuev1beta1.FlavorQuotasApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("FlavorUsage"):
		return &kueuev1beta1.FlavorUsageApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("KubeConfig"):
		return &kueuev1beta1.KubeConfigApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("LocalQueue"):
		return &kueuev1beta1.LocalQueueApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("LocalQueueFlavorUsage"):
		return &kueuev1beta1.LocalQueueFlavorUsageApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("LocalQueueResourceUsage"):
		return &kueuev1beta1.LocalQueueResourceUsageApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("LocalQueueSpec"):
		return &kueuev1beta1.LocalQueueSpecApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("LocalQueueStatus"):
		return &kueuev1beta1.LocalQueueStatusApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("MultiKueueCluster"):
		return &kueuev1beta1.MultiKueueClusterApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("MultiKueueClusterSpec"):
		return &kueuev1beta1.MultiKueueClusterSpecApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("MultiKueueClusterStatus"):
		return &kueuev1beta1.MultiKueueClusterStatusApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("MultiKueueConfig"):
		return &kueuev1beta1.MultiKueueConfigApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("MultiKueueConfigSpec"):
		return &kueuev1beta1.MultiKueueConfigSpecApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("PodSet"):
		return &kueuev1beta1.PodSetApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("PodSetAssignment"):
		return &kueuev1beta1.PodSetAssignmentApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("PodSetTopologyRequest"):
		return &kueuev1beta1.PodSetTopologyRequestApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("PodSetUpdate"):
		return &kueuev1beta1.PodSetUpdateApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("ProvisioningRequestConfig"):
		return &kueuev1beta1.ProvisioningRequestConfigApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("ProvisioningRequestConfigSpec"):
		return &kueuev1beta1.ProvisioningRequestConfigSpecApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("ReclaimablePod"):
		return &kueuev1beta1.ReclaimablePodApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("RequeueState"):
		return &kueuev1beta1.RequeueStateApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("ResourceFlavor"):
		return &kueuev1beta1.ResourceFlavorApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("ResourceFlavorSpec"):
		return &kueuev1beta1.ResourceFlavorSpecApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("ResourceGroup"):
		return &kueuev1beta1.ResourceGroupApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("ResourceQuota"):
		return &kueuev1beta1.ResourceQuotaApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("ResourceUsage"):
		return &kueuev1beta1.ResourceUsageApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("TopologyAssignment"):
		return &kueuev1beta1.TopologyAssignmentApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("TopologyDomainAssignment"):
		return &kueuev1beta1.TopologyDomainAssignmentApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("Workload"):
		return &kueuev1beta1.WorkloadApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("WorkloadPriorityClass"):
		return &kueuev1beta1.WorkloadPriorityClassApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("WorkloadSpec"):
		return &kueuev1beta1.WorkloadSpecApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("WorkloadStatus"):
		return &kueuev1beta1.WorkloadStatusApplyConfiguration{}

		// Group=visibility.kueue.x-k8s.io, Version=v1alpha1
	case visibilityv1alpha1.SchemeGroupVersion.WithKind("ClusterQueue"):
		return &applyconfigurationvisibilityv1alpha1.ClusterQueueApplyConfiguration{}
	case visibilityv1alpha1.SchemeGroupVersion.WithKind("LocalQueue"):
		return &applyconfigurationvisibilityv1alpha1.LocalQueueApplyConfiguration{}
	case visibilityv1alpha1.SchemeGroupVersion.WithKind("PendingWorkload"):
		return &applyconfigurationvisibilityv1alpha1.PendingWorkloadApplyConfiguration{}
	case visibilityv1alpha1.SchemeGroupVersion.WithKind("PendingWorkloadsSummary"):
		return &applyconfigurationvisibilityv1alpha1.PendingWorkloadsSummaryApplyConfiguration{}

		// Group=visibility.kueue.x-k8s.io, Version=v1beta1
	case visibilityv1beta1.SchemeGroupVersion.WithKind("ClusterQueue"):
		return &applyconfigurationvisibilityv1beta1.ClusterQueueApplyConfiguration{}
	case visibilityv1beta1.SchemeGroupVersion.WithKind("LocalQueue"):
		return &applyconfigurationvisibilityv1beta1.LocalQueueApplyConfiguration{}
	case visibilityv1beta1.SchemeGroupVersion.WithKind("PendingWorkload"):
		return &applyconfigurationvisibilityv1beta1.PendingWorkloadApplyConfiguration{}
	case visibilityv1beta1.SchemeGroupVersion.WithKind("PendingWorkloadsSummary"):
		return &applyconfigurationvisibilityv1beta1.PendingWorkloadsSummaryApplyConfiguration{}

	}
	return nil
}

func NewTypeConverter(scheme *runtime.Scheme) *testing.TypeConverter {
	return &testing.TypeConverter{Scheme: scheme, TypeResolver: internal.Parser()}
}

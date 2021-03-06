package kuber

import (
	"strings"

	"github.com/reconquest/karma-go"
	"k8s.io/apimachinery/pkg/runtime/schema"

	appsv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	batchv1beta1 "k8s.io/api/batch/v1beta1"
	networkingv1beta1 "k8s.io/api/networking/v1beta1"

	corev1 "k8s.io/api/core/v1"
)

type GroupVersionResourceKind struct {
	schema.GroupVersionResource
	Kind string
}

func (gvrk GroupVersionResourceKind) String() string {
	return strings.Join([]string{gvrk.Group, "/", gvrk.Version, ", Resource=", gvrk.Resource, ", Kind=", gvrk.Kind}, "")
}

var (
	Nodes = GroupVersionResourceKind{
		GroupVersionResource: corev1.SchemeGroupVersion.WithResource("nodes"),
		Kind:                 "Node",
	}
	Namespaces = GroupVersionResourceKind{
		GroupVersionResource: corev1.SchemeGroupVersion.WithResource("namespaces"),
		Kind:                 "Namespace",
	}
	LimitRanges = GroupVersionResourceKind{
		GroupVersionResource: corev1.SchemeGroupVersion.WithResource("limitranges"),
		Kind:                 "LimitRange",
	}
	Pods = GroupVersionResourceKind{
		GroupVersionResource: corev1.SchemeGroupVersion.WithResource("pods"),
		Kind:                 "Pod",
	}
	ReplicationControllers = GroupVersionResourceKind{
		GroupVersionResource: corev1.SchemeGroupVersion.WithResource("replicationcontrollers"),
		Kind:                 "ReplicationController",
	}

	Deployments = GroupVersionResourceKind{
		GroupVersionResource: appsv1.SchemeGroupVersion.WithResource("deployments"),
		Kind:                 "Deployment",
	}
	StatefulSets = GroupVersionResourceKind{
		GroupVersionResource: appsv1.SchemeGroupVersion.WithResource("statefulsets"),
		Kind:                 "StatefulSet",
	}
	DaemonSets = GroupVersionResourceKind{
		GroupVersionResource: appsv1.SchemeGroupVersion.WithResource("daemonsets"),
		Kind:                 "DaemonSet",
	}
	ReplicaSets = GroupVersionResourceKind{
		GroupVersionResource: appsv1.SchemeGroupVersion.WithResource("replicasets"),
		Kind:                 "ReplicaSet",
	}

	Jobs = GroupVersionResourceKind{
		GroupVersionResource: batchv1.SchemeGroupVersion.WithResource("jobs"),
		Kind:                 "Job",
	}
	CronJobs = GroupVersionResourceKind{
		GroupVersionResource: batchv1beta1.SchemeGroupVersion.WithResource("cronjobs"),
		Kind:                 "CronJob",
	}
	Ingresses = GroupVersionResourceKind{
		GroupVersionResource: networkingv1beta1.SchemeGroupVersion.WithResource("ingresses"),
		Kind:                 "Ingress",
	}
)

// TODO: Refactor to a map[kind]GVRK
func KindToGvrk(kind string) (*GroupVersionResourceKind, error) {
	switch kind {
	case Nodes.Kind:
		return &Nodes, nil
	case Namespaces.Kind:
		return &Namespaces, nil
	case Pods.Kind:
		return &Pods, nil
	case ReplicationControllers.Kind:
		return &ReplicationControllers, nil
	case Deployments.Kind:
		return &Deployments, nil
	case StatefulSets.Kind:
		return &StatefulSets, nil
	case DaemonSets.Kind:
		return &DaemonSets, nil
	case ReplicaSets.Kind:
		return &ReplicaSets, nil
	case Jobs.Kind:
		return &Jobs, nil
	case CronJobs.Kind:
		return &CronJobs, nil
	case Ingresses.Kind:
		return &Ingresses, nil
	default:
		return nil, karma.Format(
			nil,
			"unknown kind: %s",
			kind,
		)
	}
}

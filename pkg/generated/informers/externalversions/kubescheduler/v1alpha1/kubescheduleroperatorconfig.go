// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	kubescheduler_v1alpha1 "github.com/openshift/cluster-kube-scheduler-operator/pkg/apis/kubescheduler/v1alpha1"
	versioned "github.com/openshift/cluster-kube-scheduler-operator/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/openshift/cluster-kube-scheduler-operator/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/openshift/cluster-kube-scheduler-operator/pkg/generated/listers/kubescheduler/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// KubeSchedulerOperatorConfigInformer provides access to a shared informer and lister for
// KubeSchedulerOperatorConfigs.
type KubeSchedulerOperatorConfigInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.KubeSchedulerOperatorConfigLister
}

type kubeSchedulerOperatorConfigInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewKubeSchedulerOperatorConfigInformer constructs a new informer for KubeSchedulerOperatorConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewKubeSchedulerOperatorConfigInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredKubeSchedulerOperatorConfigInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredKubeSchedulerOperatorConfigInformer constructs a new informer for KubeSchedulerOperatorConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredKubeSchedulerOperatorConfigInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubeschedulerV1alpha1().KubeSchedulerOperatorConfigs().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubeschedulerV1alpha1().KubeSchedulerOperatorConfigs().Watch(options)
			},
		},
		&kubescheduler_v1alpha1.KubeSchedulerOperatorConfig{},
		resyncPeriod,
		indexers,
	)
}

func (f *kubeSchedulerOperatorConfigInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredKubeSchedulerOperatorConfigInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *kubeSchedulerOperatorConfigInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&kubescheduler_v1alpha1.KubeSchedulerOperatorConfig{}, f.defaultInformer)
}

func (f *kubeSchedulerOperatorConfigInformer) Lister() v1alpha1.KubeSchedulerOperatorConfigLister {
	return v1alpha1.NewKubeSchedulerOperatorConfigLister(f.Informer().GetIndexer())
}

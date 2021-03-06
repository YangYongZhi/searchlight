package framework

import (
	"time"

	"github.com/appscode/go/crypto/rand"
	"github.com/appscode/go/types"
	. "github.com/onsi/gomega"
	apps "k8s.io/api/apps/v1beta1"
	core "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (f *Invocation) StatefulSet() *apps.StatefulSet {
	ss := &apps.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      rand.WithUniqSuffix("statefulset"),
			Namespace: f.namespace,
			Labels: map[string]string{
				"app": f.app,
			},
		},
		Spec: apps.StatefulSetSpec{
			Replicas: types.Int32P(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": f.app,
				},
			},
			Template:    f.PodTemplate(),
			ServiceName: TEST_HEADLESS_SERVICE,
		},
	}

	ss.Spec.Template.Spec.Volumes = []core.Volume{}
	ss.Spec.VolumeClaimTemplates = []core.PersistentVolumeClaim{
		{
			ObjectMeta: metav1.ObjectMeta{
				Name: TestSourceDataVolumeName,
				Annotations: map[string]string{
					"volume.beta.kubernetes.io/storage-class": f.storageClass,
				},
			},
			Spec: core.PersistentVolumeClaimSpec{
				StorageClassName: types.StringP(f.storageClass),
				AccessModes: []core.PersistentVolumeAccessMode{
					core.ReadWriteOnce,
				},
				Resources: core.ResourceRequirements{
					Requests: core.ResourceList{
						core.ResourceStorage: resource.MustParse("5Gi"),
					},
				},
			},
		},
	}
	return ss
}

func (f *Framework) GetStatefulSet(meta metav1.ObjectMeta) (*apps.StatefulSet, error) {
	return f.kubeClient.AppsV1beta1().StatefulSets(meta.Namespace).Get(meta.Name, metav1.GetOptions{})
}

func (f *Framework) CreateStatefulSet(obj *apps.StatefulSet) (*apps.StatefulSet, error) {
	return f.kubeClient.AppsV1beta1().StatefulSets(obj.Namespace).Create(obj)
}

func (f *Framework) DeleteStatefulSet(obj *apps.StatefulSet) error {
	return f.kubeClient.AppsV1beta1().StatefulSets(obj.Namespace).Delete(obj.Name, deleteInForeground())
}

func (f *Framework) EventuallyStatefulSet(meta metav1.ObjectMeta) GomegaAsyncAssertion {
	return Eventually(
		func() *core.PodList {
			obj, err := f.kubeClient.AppsV1beta1().StatefulSets(meta.Namespace).Get(meta.Name, metav1.GetOptions{})
			Expect(err).NotTo(HaveOccurred())
			podList, err := f.GetPodList(obj)
			Expect(err).NotTo(HaveOccurred())
			return podList
		},
		time.Minute*5,
		time.Second*5,
	)
}

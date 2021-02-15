package annotate

import (
	"context"
	"errors"
	"net/http"
        "os"

	eirinix "github.com/SUSE/eirinix"
	"go.uber.org/zap"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// Extension changes pod definitions
type Extension struct{ Logger *zap.SugaredLogger }

// New returns the persi extension
func New() eirinix.Extension {
	return &Extension{}
}

// Handle manages volume claims for ExtendedStatefulSet pods
func (ext *Extension) Handle(ctx context.Context, eiriniManager eirinix.Manager, pod *corev1.Pod, req admission.Request) admission.Response {
        log := eiriniManager.GetLogger().Named("Annotation Starts!!")
	ext.Logger = log
	log.infof(".......................Eirini extension called successfully........................")
	if pod == nil {
		return admission.Errored(http.StatusBadRequest, errors.New("No pod could be decoded from the request"))
	}

	
	podCopy := pod.DeepCopy()
	log.Infof("POD Details: %s (%s)", podCopy.Name, podCopy.Namespace)
        log.Infof("PRODUCT_ID and PRODUCT_NAME Env Variables values are: %s (%s)", os.Getenv("PRODUCT_ID"), os.Getenv("PRODUCT_NAME"))
	log.Infof("PRODUCT_METRIC and PRODUCT_VERSION Env Variables values are: %s (%s)", os.Getenv("PRODUCT_METRIC"), os.Getenv("PRODUCT_VERSION"))

	podCopy.Annotations["productID"]=os.Getenv("PRODUCT_ID")
	podCopy.Annotations["productMetric"]=os.Getenv("PRODUCT_METRIC")
	podCopy.Annotations["productName"]=os.Getenv("PRODUCT_NAME")
        podCopy.Annotations["productVersion"]=os.Getenv("PRODUCT_VERSION")
	podCopy.Annotations["productChargedContainers"]=os.Getenv("PRODUCT_CHARGED_CONTAINERS")

	return eiriniManager.PatchFromPod(req, podCopy)
}

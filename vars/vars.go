package vars

import (
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/kubernetes/pkg/printers"
)

var CfgFile, Namespace, MustGatherRootPath, OutputStringVar, Id, Container, OMCVersionHash, OMCVersionTag, DiffCmd, DefaultProject, ForResource string
var AllNamespaceBoolVar, ShowLabelsBoolVar, UseLocalCRDs, Wide, ShowKind, ShowNamespace, ShowManagedFields bool

var EventTypes []string
var AliasToCrd map[string]apiextensionsv1.CustomResourceDefinition
var ArgPresent map[string]bool
var KnownResources map[string]map[string]interface{}
var TableGenerator *printers.HumanReadableGenerator
var CRD *apiextensionsv1.CustomResourceDefinition

var Schema *runtime.Scheme

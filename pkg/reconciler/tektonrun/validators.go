package tektonrun

import (
	"encoding/json"
	// "fmt"

	buildv1alpha1 "github.com/shipwright-io/build/pkg/apis/build/v1alpha1"
	tektonv1alpha1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1alpha1"
	tektonv1beta1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

// ValidateTektonRun validates that the provided Tekton Run object correctly specifies the
// Shipwright Build custom task for Tekton.
func ValidateTektonRun(tektonRun *tektonv1alpha1.Run) error {
	var allErrs field.ErrorList

	path := field.NewPath("spec")

	if tektonRun.Spec.Spec == nil && tektonRun.Spec.Ref == nil {
		allErrs = append(allErrs, field.Required(path, "one of spec or ref must be provided"))
	}

	if tektonRun.Spec.Spec != nil && tektonRun.Spec.Ref != nil {
		allErrs = append(allErrs, field.Invalid(path, "<object>", "only one of spec or ref can be provided"))
	}

	if errs := validateRunEmbeddedSpec(tektonRun.Spec.Spec, path.Child("spec")); errs != nil {
		allErrs = append(allErrs, errs...)
	}

	if errs := validateRunEmbeddedRef(tektonRun.Spec.Ref, path.Child("ref")); len(errs) > 0 {
		allErrs = append(allErrs, errs...)
	}

	if err := validateRunTimeout(tektonRun.Spec.Timeout, path.Child("timeout")); err != nil {
		allErrs = append(allErrs, err)
	}

	if err := validateRunRetires(tektonRun.Spec.Retries, path.Child("retries")); err != nil {
		allErrs = append(allErrs, err)
	}

	if err := validateRunParameters(tektonRun.Spec.Params, path.Child("params")); err != nil {
		allErrs = append(allErrs, err...)
	}

	if len(allErrs) == 0 {
		return nil
	}

	return apierrors.NewInvalid(
		schema.ParseGroupKind("Run.tekton.dev"),
		tektonRun.Name,
		allErrs,
	)
}

func validateRunEmbeddedSpec(embeddedSpec *tektonv1alpha1.EmbeddedRunSpec, path *field.Path) field.ErrorList {
	var allErrs field.ErrorList

	if embeddedSpec == nil {
		return nil
	}

	if err := validateAPIVersion(embeddedSpec.APIVersion, path.Child("apiVersion")); err != nil {
		allErrs = append(allErrs, err)
	}
	if err := validateKind(embeddedSpec.Kind, path.Child("kind")); err != nil {
		allErrs = append(allErrs, err)
	}
	if err := validateBuildSpec(embeddedSpec.Spec, path.Child("spec")); err != nil {
		allErrs = append(allErrs, err)
	}

	if len(allErrs) > 0 {
		return allErrs
	}

	return nil
}

func validateRunEmbeddedRef(embeddedRef *tektonv1beta1.TaskRef, path *field.Path) field.ErrorList {
	var allErrs field.ErrorList

	if embeddedRef == nil {
		return nil
	}

	if err := validateAPIVersion(embeddedRef.APIVersion, path.Child("apiVersion")); err != nil {
		allErrs = append(allErrs, err)
	}
	if err := validateKind(string(embeddedRef.Kind), path.Child("kind")); err != nil {
		allErrs = append(allErrs, err)
	}
	if err := validateName(embeddedRef.Name, path.Child("name")); err != nil {
		allErrs = append(allErrs, err)
	}

	if len(allErrs) > 0 {
		return allErrs
	}
	return nil
}

func validateAPIVersion(apiVersion string, path *field.Path) *field.Error {
	if apiVersion != "shipwright.io/v1alpha1" {
		return field.Invalid(path, apiVersion, "apiVersion must be shipwright.io/v1alpha1")
	}
	return nil
}

func validateKind(kind string, path *field.Path) *field.Error {
	if kind != "Build" {
		return field.Invalid(path, kind, "kind must be Build")
	}
	return nil
}

func validateName(name string, path *field.Path) *field.Error {
	if len(name) == 0 {
		return field.Required(path, "build name is required")
	}
	return nil
}

func validateBuildSpec(spec runtime.RawExtension, path *field.Path) *field.Error {
	if len(spec.Raw) == 0 {
		return field.Required(path, "Build spec must be provided")
	}
	buildSpec := &buildv1alpha1.BuildSpec{}
	err := json.Unmarshal(spec.Raw, buildSpec)
	// TODO: err is only raised if we fail to unmarshal JSON
	// We need to validate if spec is non-empty to avoid unnecessary errors.
	if err != nil {
		return field.Invalid(path, "<object>", "spec is not a valid Build spec")
	}
	return nil
}

func validateRunTimeout(timeout *metav1.Duration, path *field.Path) *field.Error {
	// TODO: Timeouts are effectively ignored by custom task implementations, but can be populated
	// default in a pipeline. Provide a warning that the timeout is ignored?
	return nil
}

func validateRunRetires(retries int, path *field.Path) *field.Error {
	if retries != 0 {
		return field.Invalid(path, retries, "retries are not supported")
	}
	return nil
}

func validateRunParameters(params []tektonv1beta1.Param, path *field.Path) field.ErrorList {
	var allErrs field.ErrorList
	allowedNames := []string{ParamSourceURL, ParamSourceRevision, ParamOutputImage}
	for _, param := range params {
		allowed := false
		for _, name := range allowedNames {
			allowed = allowed || (name == param.Name)
		}
		// if !allowed {
		// 	allErrs = append(allErrs, field.NotSupported(path.Child(fmt.Sprintf("[%d].name", i)), param.Name, allowedNames))
		// }
	}

	if len(allErrs) > 0 {
		return allErrs
	}
	return nil
}

// Copyright The Shipwright Contributors
//
// SPDX-License-Identifier: Apache-2.0

package resources

import (
	"fmt"

	tektonv1beta1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"

	buildv1alpha1 "github.com/shipwright-io/build/pkg/apis/build/v1alpha1"
	"github.com/shipwright-io/build/pkg/config"
)

// amendTaskSpecWithImageMutate add more steps to Tekton's Task in order to
// mutate the image with annotations and labels
func amendTaskSpecWithImageMutate(
	cfg *config.Config,
	taskSpec *tektonv1beta1.TaskSpec,
	buildOutput, buildRunOutput buildv1alpha1.Image,
) {
	// initialize the step from the template
	mutateStep := tektonv1beta1.Step{
		Container: *cfg.MutateImageContainerTemplate.DeepCopy(),
	}

	mutateStep.Container.Name = imageMutateContainerName

	// if labels or annotations are specified in buildRun then merge them with build's
	labels := mergeMaps(buildOutput.Labels, buildRunOutput.Labels)
	annotations := mergeMaps(buildOutput.Annotations, buildRunOutput.Annotations)

	mutateStep.Container.Args = mutateArgs(annotations, labels)

	// append the mutate step
	taskSpec.Steps = append(taskSpec.Steps, mutateStep)
}

// mergeMaps takes 2 maps as input and merge the second into the first
// values in second would takes precedence if both maps have same keys
func mergeMaps(first map[string]string, second map[string]string) map[string]string {
	if len(first) == 0 {
		first = map[string]string{}
	}
	for k, v := range second {
		first[k] = v
	}
	return first
}

func mutateArgs(annotations, labels map[string]string) []string {
	args := []string{
		"--image",
		fmt.Sprintf("$(params.%s-%s)", prefixParamsVolumes, paramOutputImage),
		"--result-file-image-digest",
		fmt.Sprintf("$(results.%s-%s.path)", prefixResults, imageDigestResult),
		"result-file-image-size",
		fmt.Sprintf("$(results.%s-%s.path)", prefixResults, imageSizeResult),
	}

	if len(annotations) > 0 {
		args = append(args, convertMutateArgs("--annotation", annotations)...)
	}

	if len(labels) > 0 {
		args = append(args, convertMutateArgs("--label", labels)...)
	}

	return args
}

// convertMutateArgs to convert the argument map to comma seprated values
func convertMutateArgs(flag string, args map[string]string) []string {
	var result []string

	for key, value := range args {
		result = append(result, flag, fmt.Sprintf("%s=%s", key, value))
	}

	return result
}

---
title: local-registries
authors:
  - "@adambkaplan"
reviewers:
  - "@coreydaley"
  - "@otaviof"
approvers:
  - "@qu1queee"
  - "@sbose78"
creation-date: 2020-07-24
last-updated: 2020-07-24
status: provisional
see-also: []  
replaces: []
superseded-by: []
---

# Local Registries

## Release Signoff Checklist

- [ ] Enhancement is `implementable`
- [ ] Design details are appropriately documented from clear requirements
- [ ] Test plan is defined
- [ ] Graduation criteria for dev preview, tech preview, GA
- [ ] User-facing documentation is created in [docs](/docs/)

## Summary

[KEP 1755](https://github.com/kubernetes/enhancements/tree/master/keps/sig-cluster-lifecycle/generic/1755-communicating-a-local-registry)
provides an official convention for clusters to declare the presence of an internal image registry.
This proposal aims to take advantage of declared local image registries, and let application
developers simplify how they declare image pull specs. This may also help `Build` objects be
portable across clusters.

## Motivation

The `Build` API in its present form requires that users provide a full image pull spec for image
references. This can be cumbersome when referencing an image on an internal registry, such as those
present in [KIND](https://kind.sigs.k8s.io/docs/user/local-registry/),
[Rancher](https://github.com/rancher/k3d/blob/master/docs/registries.md#using-a-local-registry),
[microk8s](https://microk8s.io/docs/registry-built-in), and
[OpenShift](https://docs.openshift.com/container-platform/4.5/registry/architecture-component-imageregistry.html).
This proposal will provide a means of declaring that an image spec references a local registry,
thereby allowing host information to be removed from the image reference.

### Goals

- Simplify how local image registries can be referenced in `Build` objects.
- Provide a means of declaring that an image references a cluster-local image registry.

### Non-Goals

- Automate how pull/push secrets can be added to a `Build` object.
- Support secure image pull/push if a local registry uses self-signed certificates.

## Proposal

Application developers will be able to specify `local: true|false` if an image reference points
to a cluster-local image registry.

The build controller/operator will watch the global "LocalRegistryHosting" `ConfigMap` to check if
a local image registry is declared. If declared, it will infer the local registry host information
from what is provided in the `ConfigMap`.

When a `BuildRun` is instantiated, the image specs that declare `local: true` are mutated to inject
the local image registry host.

### User Stories [optional]

As an application developer
I would like to declare that images in builds reference a local image registry
So that I do not have to provide host information in the pull spec

### Implementation Details/Notes/Constraints [optional]

#### Obtaining local registry information

First, clusters which deploy a local image registry must provide the releavant host info by
publishing it in the `local-registry-hosting`
[ConfigMap](https://github.com/kubernetes/enhancements/tree/master/keps/sig-cluster-lifecycle/generic/1755-communicating-a-local-registry#the-local-registry-hosting-configmap).
The build operator/controller must have RBAC which allows it to read the `ConfigMap` and obtain the
registry host information. For `Build` and `BuildRun` objects, the `HostFromClusterNetwork` value
should be used as the hostname.

If the `local-registry-hosting` ConfigMap is created, updated, or deleted, the build controller
and/or operator needs to react accordingly.

#### Declaring an image as local

`Image` references in the API will be updated to have the `local` field. An example using the 
`output` spec:

```yaml
spec:
  output:
    image: mynamespace/myapp:latest
    local: true
```

When an image is declared as `local`, the build controller prepends the local image registry
hostname to the provided image reference.

### Risks and Mitigations

App developers use `local: true` on clusters that don't have a local registry defined, end up
pulling images from `docker.io` or an image registry in container runtime's unqualified search
path.

Local image registry configurations can be altered while the build controller is running.

## Design Details

### Test Plan

1. Deploy the build operator/controller on a cluster that populates the `local-registry-hosting`
   ConfigMap OR as cluster-admin, manually populate the `local-registry-hosting` ConfigMap.
2. Create a `Build` and/or `BuildRun` which references a local image.
3. Ensure that a build can push an image referencing the local image spec.

### Graduation Criteria

**Dev Preview** - we can simply define an environment variable to set the local registry hostname.

**Tech Preview/GA** - fully read the `local-registry-hosting` ConfigMap, with the ability to update
the build controller if this ConfigMap changes.

#### Examples

These are generalized examples to consider, in addition to the aforementioned [maturity
levels][maturity-levels].

### Upgrade / Downgrade Strategy

Upgrades add a new field to the API. On downgrade the `local` attribute should be ignored.

### Version Skew Strategy

N/A

## Implementation History

2020-07-24: Initial proposal

## Drawbacks

TBD

The idea is to find the best form of an argument why this enhancement should _not_ be implemented.

## Alternatives

TBD

Similar to the `Drawbacks` section the `Alternatives` section is used to highlight and record other
possible approaches to delivering the value proposed by an enhancement.

## Infrastructure Needed [optional]

- A cluster which populates the `local-registry-hosting` ConfigMap

## Open Questions [optional]

1. How can we make `Build` objects with local image pull specs portable across namespaces?

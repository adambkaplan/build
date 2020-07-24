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

#### Declaring an image as local

`Image` references in the API will be updated to have the `local` field. An example using the 
`output` spec:

```yaml
spec:
  output:
    image: mynamespace/myapp:latest
    local: true
```

### Risks and Mitigations

What are the risks of this proposal and how do we mitigate. Think broadly. For example, consider
both security and how this will impact the larger OKD ecosystem.

How will security be reviewed and by whom? How will UX be reviewed and by whom?

Consider including folks that also work outside your immediate sub-project.

## Design Details

### Test Plan

**Note:** *Section not required until targeted at a release.*

Consider the following in developing a test plan for this enhancement:

- Will there be e2e and integration tests, in addition to unit tests?
- How will it be tested in isolation vs with other components?

No need to outline all of the test cases, just the general strategy. Anything that would count as
tricky in the implementation and anything particularly challenging to test should be called out.

All code is expected to have adequate tests (eventually with coverage expectations).

### Graduation Criteria

**Note:** *Section not required until targeted at a release.*

Define graduation milestones.

These may be defined in terms of API maturity, or as something else. Initial proposal should keep
this high-level with a focus on what signals will be looked at to determine graduation.

Consider the following in developing the graduation criteria for this enhancement:

- Maturity levels - `Dev Preview`, `Tech Preview`, `GA`
- Deprecation

Clearly define what graduation means.

#### Examples

These are generalized examples to consider, in addition to the aforementioned [maturity
levels][maturity-levels].

##### Dev Preview -> Tech Preview

- Ability to utilize the enhancement end to end
- End user documentation, relative API stability
- Sufficient test coverage
- Gather feedback from users rather than just developers

##### Tech Preview -> GA

- More testing (upgrade, downgrade, scale)
- Sufficient time for feedback
- Available by default

**For non-optional features moving to GA, the graduation criteria must include end to end tests.**

##### Removing a deprecated feature

- Announce deprecation and support policy of the existing feature
- Deprecate the feature

### Upgrade / Downgrade Strategy

If applicable, how will the component be upgraded and downgraded? Make sure this is in the test
plan.

Consider the following in developing an upgrade/downgrade strategy for this enhancement:

- What changes (in invocations, configurations, API use, etc.) is an existing cluster required to
  make on upgrade in order to keep previous behavior?
- What changes (in invocations, configurations, API use, etc.) is an existing cluster required to
  make on upgrade in order to make use of the enhancement?

### Version Skew Strategy

How will the component handle version skew with other components? What are the guarantees? Make sure
this is in the test plan.

Consider the following in developing a version skew strategy for this enhancement:

- During an upgrade, we will always have skew among components, how will this impact your work?
- Does this enhancement involve coordinating behavior in the control plane and in the kubelet? How
  does an n-2 kubelet without this feature available behave when this feature is used?
- Will any other components on the node change? For example, changes to CSI, CRI or CNI may require
  updating that component before the kubelet.

## Implementation History

Major milestones in the life cycle of a proposal should be tracked in `Implementation History`.

## Drawbacks

The idea is to find the best form of an argument why this enhancement should _not_ be implemented.

## Alternatives

Similar to the `Drawbacks` section the `Alternatives` section is used to highlight and record other
possible approaches to delivering the value proposed by an enhancement.

## Infrastructure Needed [optional]

Use this section if you need things from the project. Examples include a new subproject, repos
requested, github details, and/or testing infrastructure.

Listing these here allows the community to get the process for these resources started right away.

## Open Questions [optional]

This is where to call out areas of the design that require closure before deciding to implement the
design. For instance:

> 1. This locks a build strategy to run privileged pods. Should we do this?
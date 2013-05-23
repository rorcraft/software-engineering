http://continuousdelivery.com/2012/10/theres-no-such-thing-as-a-devops-team/#more-827

__Make people responsible for the consequences of their actions.__

> build reliable software that can be continuous deployed to an unreliable platform that scales horizontally. They need to be able to self-service environments and deployments. They need to understand how to write testable, maintainable code. They need to know how to do packaging, deployment, and post-deployment support.

Here’s what the devops team does in this model:
- Builds a platform that allows developers to self-service environments for testing and production (and deployments to those environments) and provides metrics to the organization as a whole. This platform is a product, and the team that builds it is doing product development, which (in particular) means the people who use the platform are your customers.
- Provides a toolchain that developers can use to build, test, deploy and run their systems.
- Coaches teams working to move to this model and provides support and training for the platform and toolchain.

__Why we need Devops__
http://itrevolution.com/books/phoenix-project-devops-book/

"Title: Why We Need DevOps Now: A Fourteen Year Study Of High Performing IT Organizations"
http://www.youtube.com/watch?v=disjFj4ruHg
Presented by: Gene Kim

__Instrumentation & Metrics__

https://vimeo.com/52711264

> "Instrumentation is Unit-testing for Ops"

* detect regressions
* validate new hypotheses
* increases tolerance to change

__Collection -> Aggregation -> Storage -> Analysis__

Devs require:
* instrumentation == collection
* concise primitives
* minimal dependencies
* fire and forget
* minimal performance impact

Ops require:
* flexibility at all other layer
* simple introspection
* simple capture

Ideal instrumentation:
* implements primitives
* captures state
* provides access

---
description: Learn about the Message Bus component needed by Response Server.
---

# Message Bus

The Message Bus is a critical component of Response Server and is required for event processing. Response Server comes bundled with an embedded message bus that will work well for isolated clusters \(where the cluster has only one Response Server instance\), or for local development and testing.

For production and scaling of Response Server, we recommend using a more robust message bus implementation.

## Drivers

The following Message Bus implementations are supported in Response Server. For configuration details, see the [Events configuration page](../configuration/events.md).

### Embedded

Use: **Development, testing, and single-server production clusters**

The _Embedded_ Message Bus is an in-memory message bus meant to be used for local development, testing, and small-scale production workloads where only a single Response Server instance is in a cluster. Because this driver runs in-memory, any events being processed while Response Server is being shut down may not be handled.

### NATS

Use: **Single-server or multi-server production clusters**

The _NATS_ Message Bus is the recommended way to handle events in Response Server. [NATS Server](https://nats.io) is extremely simple to deploy, manage, and will work out-of-the-box to provide you with impressive performance when handling events in Response Server clusters.

NATS Server can be [installed in various ways](https://docs.nats.io/nats-server/installation), however, we recommend the [use of Docker](https://docs.nats.io/nats-server/installation#installing-via-docker) to ensure the service is automatically started and restarted on failure.

To keep deployment simple, ~~read our Docker deployment guide~~ that includes Response Server, PostgreSQL, and NATS Server.


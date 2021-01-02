---
title: Events
category: Configuration
position: 104
version: 1.0
---

Response Server requires a [Message Bus](/components/message-bus) implementation to properly route events between servers in a cluster. By default, Response Server will run with an embedded message bus that isolates all events to a particular server.

<alert type="info">

Using the embedded message bus is not supported when more than one server is in a cluster and will result in [split-brain](https://en.wikipedia.org/wiki/Split-brain_(computing)) events.

</alert>

## events

`events` `(EventsConfig: nil)`

All attributes for Response Server's Message Bus connection are defined in the `events` stanza.

```hcl
events {
  # the attributes from this page are defined here...
}
```

### type

`type` `("embedded" | "nats": "embedded")`

Response Server supports two message bus implementations, these are compared below.

| Type       | Default | Clustering | Recommended            | Performance |
| ---------- | :-----: | :--------: | :--------------------- | ----------- |
| `embedded` |    Y    |     N      | Dev & Test             | Medium      |
| `nats`     |    N    |     Y      | Dev, Test & Production | Very High   |

Each message bus implementation supports a varying number of the properties below. See some common examples for each below.

#### embedded

The `embedded` implementation, enabled when the `events` stanza is omitted or when `type` is set to `embedded`.

```hcl
events {
  type = "embedded"
}
```

#### nats

The `nats` implementation, enabled when the `type` is set to `nats`. This is the recommended message bus for production and is the same message bus used by Response Cloud.

```hcl
events {
  type = "nats"
  url = "nats://localhost:4222"
}
```

### subject

`subject` `(string: "response-events")`

<alert type="info">

Applies to the `nats` implementation only.

</alert>

Setting `subject` changes the NATS subject used for all Response Server events. Every Response Server in a cluster should use the same subject in order to prevent [split-brain](https://en.wikipedia.org/wiki/Split-brain_(computing)). The `subject` defaults to `response-events` and should only be changes if you want to customize the name or plan to run multiple Response Server clusters using the same NATS Server.

### url

`url` `(string: "nats://localhost:4222")`

<alert type="info">

Applies to the `nats` implementation only.

</alert>

The `url` property must be set when using a NATS Server that is not located on the local server or if accessing NATS Server via socket. By default Response Server will attempt to connect to NATS Server using the connection URL of `nats://localhost:4222`.

---
description: Configure Response as a distributed service.
---

# Cluster

Response Server offers an embedded clustering service that requires very little configuration to work. Running multiple Response Server instances can improve performance when handling a lot of simultaneous connections.

By default, Response Server starts in an isolated mode where cache and other cluster-related metadata is not distributed through multiple Response Server instances. This is typically not the behavior you want when running multiple Response Server instances behind a load balancer and we do not support this configuration when running multiple instances, certain services \(like the live map\) will not work properly. 

## cluster

`cluster` `(ClusterConfig: nil)`

All attributes for Response Server's clustering service are defined in the `cluster` stanza.

```text
cluster {
  # the attributes from this page are defined here...
}
```

### environment

`environment` `("local" | "lan" | "wan": <required>)`

Response Server supports three environments that tweak the cluster connectivity and performance settings. There is no default environment and you must explicitly set one when configuring the cluster. Every instance in a cluster should have the same environment name set.

The environments are explained in detail below.

#### **local**

The `local` environment configures Response Server's connectivity and syncing settings for local loopback connections where all instances are on the same physical or virtual machine where each instance will have the same IP address.

```text
cluster {
  environment = "local"
}
```

#### **lan**

The `lan` environment configures Response Server to work well across the local area network. This prioritizes higher convergence over less bandwidth usage as bandwidth is typically not an issue when dealing with a local network.

```text
cluster {
  environment = "lan"
}
```

#### **wan**

The `wan` environment configures Response Server to work well across a wide area network. This prioritizes bandwidth consumption but maintains high consistency.

```text
cluster {
  environment = "wan"
}
```

### bind\_address

`bind_address` `(string: "0.0.0.0")`

Setting `bind_address` changes the network that Response Server will bind to for append-only log activities. This is used to sync data between each instance in the cluster and all Response Server instances will need to be able to communicate on this network.

### bind\_port

`bind_port` `(int: 3320)`

Setting `bind_port` changes the port that Response Server will bind to for append-only log activities. This is used to sync data between each instance in the cluster and all Response Server instances will need to be able to communicate on this network.

### memberlist\_bind\_address

`memberlist_bind_address` `(string: "0.0.0.0")`

Setting `memberlist_bind_address`  changes the network that Response Server will bind to for member-to-member and member discovery purposes. This is used to detect members and maintain quorum for in-memory data stored throughout the cluster.

### memberlist\_bind\_port

`memberlist_bind_port` `(int: 3322)`

Setting `memberlist_bind_port`  changes the port that Response Server will bind to for member-to-member and member discovery purposes. This is used to detect members and maintain quorum for in-memory data stored throughout the cluster.

### members

`members` `([]string: nil)`

The `members` property allows configuring a list of other Response Server instances that the instance being configured should contact to join into the cluster. If you know all of the peer `memberlist_bind_address:memberlist_bind_port` configurations for other Response Server instances and they will never change, setting this is a simple way to configure a cluster.

{% hint style="info" %}
You cannot configure `members` and `autojoin` . Choose one or the other.
{% endhint %}

### autojoin

This is the recommended way to configure Response Server clustering and allows automatically discovering the other Response Server instances using a service discovery medium. Response Server supports the following autojoin mechanisms.

For an almost-zero-configuration setup, we recommend using the `nats` autojoin configuration.

#### nats

The `nats` autojoin configuration allows using an existing NATS Server for automatically discovering and joining a Response Server cluster. This is the recommended approach as it requires no additional external components [because NATS Server should be used for events](events.md#nats) allowing you to re-use the same NATS Server. You can share the same NATS Server connection URL using [variables](introduction.md#declaring-variables).

```text
cluster {
  autojoin "nats" {
    url = "nats://localhost:4222"
    
    # Uncomment and change the subject if you are sharing the same
    # NATS Server for multiple Response Server clusters.
    # subject = "response-cluster"
  }
}
```


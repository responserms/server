---
description: Configure Response Server's HTTP server.
---

# HTTP

Response Server's primary role is providing HTTP-based services. Because of this, HTTP configuration is a key part of properly tweaking Response Server. This page covers all of the available configuration options for tuning Response Server's HTTP and TLS options.

{% hint style="warning" %}
TLS encryption is strongly recommended in production. Response Server can automatically obtain and renew TLS certificates via Let's Encrypt for you, read more about this in the [http.tls.auto](http.md#http-tls-auto) section.
{% endhint %}

## http

`http` `(HTTPConfig: nil)`

All attributes for Response Server's HTTP services are defined in the `http` stanza.

```text
http {
  # the attributes from this page are defined here...
}
```

### bind\_address

`http.bind_address` `(string: "0.0.0.0")`

Response Server by default listens on `0.0.0.0` which means that HTTP services will be bound to all networks. This is done to create a frictionless entrypoint to getting started with Response Server.

We recommend that you limit the networks Response Server listens on when deploying to production, especially in a cloud environment. Set the `bind_address` to a suitable network \(such as your public network\) or use one of our [helpful networking functions](https://github.com/responserms/server/tree/424cedac7d7882a355db919c1f557c58026e62b1/config/introduction/README.md#private_ip) to automatically configure the `bind_address`.

### port

`http.port` `(int: 8080)`

The `port` attribute configure the primary non-TLS port to be used by Response Server. Response Server will by default listen for non-TLS connections on port `8080`.

{% hint style="warning" %}
If you are using a Linux or Unix-based system and want to run Response Server on ports below 1024, such as 80, follow our ~~privileged ports guide~~ to give the proper capabilities to Response Server without running it as root.
{% endhint %}

### max\_upload\_size

`http.max_upload_size` `(string: "10mb")`

Response Server limits the maximum file size that can be uploaded to protect the health of your server and control bandwidth consumption. The larger the files you allow to be uploaded the greater bandwidth usage will be used to download that same content. The bandwidth usage will increase linearly as more users download this same content. Response Server supports supplying data sizes in a human-readable syntax and defaults to `10mb`.

You are free to configure this setting to whatever maximum you see fit, but please understand the implications of large files on bandwidth and system resources.

```text
http {
  # set the maximum upload size to 100 megabytes
  max_upload_size = "100mb"
}
```

## http.tls

`http.tls` `(TLSConfig: nil)`

The `tls` stanza is defined within the `http` stanza. This stanza configures the TLS certificates and other settings for Response Server.

```text
http {
  tls {
    # the attributes from this section are defined here...
  }
}
```

When the `tls` stanza is not provided TLS is disabled. When provided, Response Server expects all required attributes to be defined and will refuse to start if they are not. If you do not intend to use TLS, such as for local development purposes, simply omit this stanza.

### port

`http.tls.port` `(int: 8443)`

Response Server will by default listen for TLS connections on port `8443`.

{% hint style="warning" %}
If you are using a Linux or Unix-based system and want to run Response Server on ports below 1024, such as 443, follow our ~~privileged ports guide~~ to give the proper capabilities to Response Server without running it as root.
{% endhint %}

### cert\_path

`http.tls.cert_path` `(string: nil)`

This is the path where Response Server can find the TLS certificate to be used for encrypting traffic. We recommend using automatic TLS if your environment supports it as it allows Response Server to obtain and renew certificates for you automatically.

### key\_path

`http.tls.key_path` `(string: nil)`

This is the path where Response Server can find the TLS certificate's private key.

## http.tls.auto

`http.tls.auto` `(AutomaticTLSConfig: nil)`

Defining this stanza will enable automatic TLS. You must define all of the required attributes for this stanza before Response Server can start. If you do not wish to use automatic TLS certificates simply omit this stanza.

{% hint style="warning" %}
By defining the `http.tls.auto` stanza you accept the Certificate Authority's subscriber agreement. The full text of this agreement [may be obtained here](https://letsencrypt.org/repository/#let-s-encrypt-subscriber-agreement).
{% endhint %}

Automatic TLS is an awesome feature and we think it's pretty simple to configure if you've used Let's Encrypt before. However, if you have not ~~please see our guide on automatic TLS certificates with Response Server~~ to get started quickly. Don't hesitate to reach out on Discord if you have questions!

### production

`http.tls.auto.production` `(boolean: false)`

Response Server defaults to using the Let's Encrypt staging endpoint to prevent rate limiting from invalid configurations when starting Response Server.

We do not recommend enabling `production` until you have successfully obtained TLS certificates through the staging endpoint. Once you have done so, setting this to true will enable the production endpoint and re-generate certificates using the production Certificate Authority.

{% hint style="info" %}
If `production` is not set to true your web browser will alert you that the certificates cannot be trusted and that your connection may not be encrypted. Don't worry, your connection is still encrypted. The browser warning is expected behavior and will be resolved once you set`http.tls.auto.production` to `true` and relaunch Response Server.
{% endhint %}

### email

`http.tls.auto.email` `(string: <required>)`

This is the email address that will be used when requesting TLS certificates from Let's Encrypt. You must set this to an address that can receive email, Let's Encrypt may email you regarding service information, changes to the Subscriber Agreement, or problems with your TLS certificates.

### domains

`http.tls.auto.domains` `([]string: <required>)`

This is an array of domains for which Response Server will automatically obtain and renew TLS certificates. You must provide at least one domain.

```text
http {
  tls {
    auto {
      domains = [
        "example.responserms.com"
      ]
    }
  }
}
```

### dns

By default Let's Encrypt will verify that you have control of the provided domains by making an HTTP request to each domain and reading a value from the `/.well-known/` directory. Response Server handles serving this and solving the automated challenges for you, you only need to ensure that the DNS records are properly configured to point to Response Server.

However, you may wish to use DNS challenges instead, such as when you have not configured a public DNS record for Response Server or are running more than one Response Server instance behind a load balancer.

You may define only one DNS solver at a time. This functionality is only available on supported DNS providers. The following DNS providers are supported:

#### cloudflare

To use the `cloudflare` DNS solver, Response Server requires an **API Token** \(not globally-scoped API Keys\) from Cloudflare with the following permissions:

* Zone / Zone / Read
* Zone / DNS / Edit

You may provide this as the `token` attribute in the `dns` stanza for `cloudflare`:

```text
http {
  tls {
    auto {
      # other automatic tls settings...
      dns "cloudflare" {
        token = "your-cloudflare-token"
      }
    }
  }
}
```

#### vultr

To use the `vultr` DNS solver, Response Server requires a Vultr API Token from [here](https://my.vultr.com/settings/#settingsapi). You may provide this as the `token` attribute in teh `dns` stanza for `vultr`:

```text
http {
  tls {
    auto {
      # other automatic tls settings...
      dns "vultr" {
        token = "your-vultr-token"
      }
    }
  }
}
```

#### digitalocean

To use the `digitalocean` DNS solver, Response Server requires a DigitalOcean Personal Access Token. See [this documentation](https://www.digitalocean.com/docs/apis-clis/api/create-personal-access-token/) for creating Personal Access Tokens in your DigitalOcean account.

You may provide this as the `token` attribute in the `dns` stanza for `digitalocean`:

```text
http {
  tls {
    auto {
      # other automatic tls settings...
      dns "digitalocean" {
        token = "your-do-token"
      }
    }
  }
}
```


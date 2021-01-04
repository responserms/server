---
description: A brief introduction to configuring Response Server with HCL or JSON files.
---

# Introduction

Response Server is configured using one or more files. The format of the files is [HCL](https://github.com/hashicorp/hcl/blob/hcl2/hclsyntax/spec.md) or [JSON](https://github.com/hashicorp/hcl/blob/hcl2/json/spec.md).

This guide shows all examples, like the one below, in HCL format and we strongly recommend the use of HCL over JSON.

```text
http {
  bind_address = "127.0.0.1"
  port = 8080
}

database {
  use = "sqlite"
  path = "./database.sqlite"
}
```

After the configuration is written, use the `--config` flag with `response-server` to specify where the configuration is located. By default, the Response Server will attempt to find a `server.hcl` file in the current working directory.

```bash
response-server --config my-config.hcl
```

## Stitching

Response Server supports stitching multiple configuration files together, combining all of the declarred entries into a single configuration. This will work for both HCL and JSON file formats.

{% hint style="warning" %}
When configuration files are merged, blocks declared more than once will be considered duplicates. Some blocks do not support duplication and will result in errors.
{% endhint %}

To combine multiple config files simply start Response Server with the `--config` flag, supplying a [glob pattern](https://pkg.go.dev/path/filepath/#Match) to match multiple files. You may match JSON or HCL files \(or both\), the parser will properly detect the format assuming they have the proper file extensions: `.hcl` or `.json`.

This example shows declaring the `http` and `database` blocks in separate files.

{% tabs %}
{% tab title="database" %}
{% code title="config/http.hcl" %}
```text
http {
  bind_address = public_ip()
  port = 80

  tls {
    port = 443

    # Enable automatic TLS certificates
    auto {
      email = "john.doe@example.org"
      domains = [
        "response.example.org"
      ]

      # Use DNS solving instead of HTTP. Optional.
      dns "cloudflare" {
        token = "example-cloudflare-token"
      }
    }
  }
}
```
{% endcode %}
{% endtab %}

{% tab title="http" %}
{% code title="config/http.hcl" %}
```text
database {
  type = "postgresql"
  name = "response_database"
  host = "127.0.0.1"
  port = 5432

  # Read the username from the DATABASE_USER environment
  # variable, use "postgres" if DATABASE_USER is not set
  username = env("DATABASE_USER", "postgres")

  # Read the password from the DATABASE_PASSWORD environment
  # variable, use an empty password if DATABASE_PASSWORD is
  # not set
  password = env("DATABASE_PASSWORD", "")

  options = {
    sslmode = env("DATABASE_SSL", "required")
  }
}
```
{% endcode %}
{% endtab %}
{% endtabs %}

With our files defined and each containing their respective block definitions, let's start Response Server with a [glob pattern](https://pkg.go.dev/path/filepath/#Match) to find all of our `.hcl` files in the relative `config/` directory.

```bash
response-server --config config/*.hcl
```

Now we know exactly where each configuration stanza is defined and we can add more files if we start to tweak Response Server further without having to change our startup command.

## Declaring Variables

The Response Server configuration file supports declaring global variables that may be reused throughout the configuration file. To declare variables, use the `variables` block as shown below.

```text
variables {
  my_encryption_key = "some-encryption-key"
}

# This is pointless, but it works!
encryption_key = my_encryption_key
```

## Functions

Response Server's configuration file supports a number of functions to ease configuration in various deployment scenarios using the same configuration file.

### env

`env(env_var_name: string, fallback_value: string = "")` `string`

The `env` function allows reading an environment variable and using its value, falling back to the `fallback_value` if the environment variable is not set. If no `fallback_value` is provided and the environment variable with the name is not set, an empty string is returned.

```text
encryption_key = env("ENCRYPTION_KEY", "my-default-key")
```

If the `ENCRYPTION_KEY` environment variable is set in the OS environment, its value will be used. If the environment variable is not set `my-default-key` \(the literal value\) will be returned instead.

### private\_ip

`private_ip()` `string`

The `private_ip` function returns a string with a single IP address that is part of [RFC 6890](https://tools.ietf.org/html/rfc6890) and has a default route. If the system can't determine its IP address or find an [RFC 6890](https://tools.ietf.org/html/rfc6890) IP address, an empty string will be returned instead.

```text
http {
  bind_address = private_ip()
}
```

### public\_ip

`public_ip()` `string`

The `public_ip` function returns a string with a single IP address that is NOT part of [RFC 6890](https://tools.ietf.org/html/rfc6890) and has a default route. If the system can't determine its IP address or find a non [RFC 6890](https://tools.ietf.org/html/rfc6890) IP address, an empty string will be returned instead.

```text
http {
  bind_address = public_ip()
}
```

### interface\_ip

`interface_ip(namedIfRE: string)` `string`

The `interface_ip` function returns a string with a single IP address on the given `namedIfRE`. If multiple IP addresses exist on the interface, they are sorted by the size of the network \(i.e. IP addresses with a smaller netmask, larger network size, are sorted first\) and the first address from the sorted list is returned.

```text
listener "http" {
  bind_address = interface_ip("eth0")
  port = 8000
}
```


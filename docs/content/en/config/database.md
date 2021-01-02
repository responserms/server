---
title: Database
category: Configuration
position: 103
version: 1.0
---

Response Server persists all data in a database. By default, Response Server will run with an embedded in-memory database. This database is great for development and testing purposes but will not persist any data, which means that as soon as Response Server is stopped all data is lost.

To persist data you'll need to configure a database engine that supports persistence.

## database

`database` `(DatabaseConfig: nil)`

All attributes for Response Server's database connection are defined in the `database` stanza.

```hcl
database {
  # the attributes from this page are defined here...
}
```

### type

`type` `("embedded" | "sqlite" | "postgresql" | "mysql": "embedded")`

Response Server supports four database engines, these are compared below.

| Type         | Default | Persistence | Clustering | Recommended | Performance |
| ------------ | :-----: | :---------: | :--------: | :---------- | ----------- |
| `embedded`   |    Y    |      N      |     N      | Dev         | Low         |
| `sqlite`     |    N    |      Y      |     N      | Dev & Test  | Low         |
| `postgresql` |    N    |      Y      |     Y      | Production  | Very High   |
| `mysql`      |    N    |      Y      |     Y      | Production  | High        |

Each database engine supports a varying number of the properties below. See some common examples for each below.

#### embedded

The `embedded` engine, enabled when the `database` stanza is omitted or when `type` is set to `embedded`.

```hcl
database {
  type = "embedded"
}
```

#### sqlite

The `sqlite` engine, enabled when `type` is set to `sqlite`.

```hcl
database {
  type = "sqlite"

  # The path where SQLite should create/read/update its database file
  path = "./database.sqlite"
}
```

#### postgresql

The `postgresql` engine, enabled when `type` is set to `postgresql`. The example below uses the default host of `127.0.0.1` and port `5432`.

```hcl
database {
  type = "postgresql"

  # The database name
  name = "response_database"

  # Database credentials
  username = "response"
  password = "my-password"

  # Response will by default use the `require` sslmode. This
  # is usually not enabled for default PostgreSQL servers.
  options = {
    sslmode = "disable"
  }
}
```

#### mysql

The `mysql` engine, enabled when `type` is set to `mysql`. The example below uses the default host of `127.0.0.1` and port `3306`.

```hcl
database {
  type = "mysql"

  # The database name
  name = "response_database"

  # Database credentials
  username = "response"
  password = "my-password"

}
```

### path

`database.path` `(string: nil)`

<alert type="info">

Applies to the `sqlite` engine only.

</alert>

The path where Response Server should create the SQLite database file. This path will be used to read the SQLite database file if it already exists. The path is relative to the execution path unless set as an absolute path.

<code-group>
  <code-block label="SQLite" active>

```hcl
database {
  use = "sqlite"
  path = "database.sqlite"
}
```

  </code-block>
</code-group>

### url

`database.url` `(string: nil)`

<alert type="info">

Applies to the `postgresql` and `mysql` engines only.

</alert>

A full database connection string URL to provide all of the connecion parameters necessary to connect to the database. This URL will be merged with all other specified options to build the final connection string, however items defined in this URL take precedence over the default values of other options.

<code-group>
  <code-block label="postgresql" active>

```hcl
database {
  use = "postgresql"
  url = "postgresql://postgres:@localhost:5432/response?sslmode=disable"
}
```

  </code-block>
  <code-block label="mysql">

```hcl
database {
  use = "mysql"
  url = "mysql://root:@localhost:3306/response"
}
```

  </code-block>
</code-group>

### name

`database.name` `(string: nil)`

<alert type="info">

Applies to the `postgresql` and `mysql` engines only.

</alert>

The name of the database Response Server should store and retrieve data from. The database will not be automatically created and must exist on the database server already. The database will automatically be migrated between schema versions when upgrading Response Server.

### host

`database.host` `(string: "localhost")`

<alert type="info">

Applies to the `postgresql` and `mysql` engines only.

</alert>

The hostname or IP address of the database server. The `host` defaults to `localhost` for both PostgreSQL and MySQL databases.

### port

`database.port` `(int: 3306 | 5432)`

<alert type="info">

Applies to the `postgresql` and `mysql` engines only.

</alert>

The port of the database server. The `port` defaults to `3306` when using MySQL and `5432` when using PostgreSQL.

### username

`database.username` `(string: "root" | "postgres")`

<alert type="info">

Applies to the `postgresql` and `mysql` engines only.

</alert>

The username used to authenticate when connecting to the database server. The `username` default to `root` when using MySQL and `postgres` when using PostgreSQL.

### password

`database.password` `(string: nil)`

<alert type="info">

Applies to the `postgresql` and `mysql` engines only.

</alert>

The password used to authenticate when connecting to the database server. The `password` does not have a default value and must be set if the database server requires a password to authenticate.

### options

`database.options` `(object[string]string: nil)`

<alert type="info">

Applies to the `postgresql` and `mysql` engines only.

</alert>

The options object allows configuring additional options when connecting to the database server. This is typically required when using `postgresql` and not using `url` to configure things like `sslmode`.

<alert type="info">

The options that can be set will vary based on which engine you are using. See the respective documentation for each database for all available options when establishing a connection:

<br />

`postgresql`

- [PostgreSQL Connection Options](https://www.postgresql.org/docs/12/libpq-connect.html#LIBPQ-PARAMKEYWORDS)

`mysql`

- [MySQL Connection Options](https://dev.mysql.com/doc/refman/8.0/en/connection-options.html#connection-establishment-options)
- [MariaDB Connection Options](https://mariadb.com/kb/en/connecting-to-mariadb/#connection-parameters)

</alert>

```hcl
database {
  type = "postgresql"
  name = "response"

  options = {
    sslmode = "disable"
  }
}
```

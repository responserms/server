# Attributes

The following attributes are set at the root-level of the configuration file and are not included in configuration blocks. These values are set with the `=` assignment operator, like the following:

```bash
key = "value"
```

## Encryption

The following attributes relate to encryption at rest within Response Server. Response Server supports AES 256-bit encryption and requires a base64-encoded 32-byte encryption key.

You can use the `response-server` executable to generate a suitable encryption key. To do so, use the `--generate-key` flag.

```bash
response-server --generate-key
```

Response Server will not start and will instead print a suitable encryption key and exit. Provide the returned encryption key in your configuration file or use the [env function](./#env) as you see fit.

### encryption\_key

`encryption_key` `(string: nil)`

The `encryption_key` property configures the encryption key to be used by Response when encrypting and decrypting values. Every Response Server instance should use the same encryption key and the encryption key cannot change or previously encrypted data will fail during decryption causing data to be inaccessible.

**Response Server will not start without `encryption_key` set to a suitable value.**


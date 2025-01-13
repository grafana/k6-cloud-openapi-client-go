# Development

## Code generation

### Requirements

> [!NOTE]
> Although the `openapi-generator` is the only tool really required, it's strongly recommended to have `goimports`
> installed as well, because the code generated isn't correctly formatted, and sometimes it even has some errors with
> non-used imports.

The Go code present on this repository has been generated with [openapi-generator](https://openapi-generator.tech/).

You can find further details about how to install it [here](https://openapi-generator.tech/docs/installation).

### Retrieve the schema

Once you have the `openapi-generator` installed, the first step is to retrieve the schema,
which can downloaded from: https://api.k6.io/cloud/v6/openapi.

To do it, run:

```sh
make update-schema
```

Or alternatively:

```sh
wget -O schema.yaml https://api.k6.io/cloud/v6/openapi
```

We always try to keep a copy of the most up-to-date schema, used to generate the client's code, in the repository's
root,
as [`schema.yaml`](./schema.yaml).

### Generate the client

To generate the client's source code run the following command:

```sh
make generate
```

Or alternatively:

```sh
openapi-generator-cli generate -i schema.yaml -g go -o ./k6 --git-user-id grafana --git-repo-id k6-cloud-openapi-client/go --package-name k6 -p isGoSubmodule=true -p disallowAdditionalPropertiesIfNotPresent=false -p withGoMod=false -t ./templates
```

Once the client is generated, you need to run `goimports`, cause the code generated isn't correctly formatted,
and sometimes it even has some errors with non-used imports. If you did run `make generate`, it already formats
the code, but if you didn't, you can run:

```sh
make format
```

Or alternatively:

```sh
find k6 -name \*.go -exec goimports -w {} \;
```

## Additional information

### Templates

In order to customize the code generated, we use
the [Mustache templates](https://openapi-generator.tech/docs/templating)
present in the [templates](./templates) directory, which contains some of the
[default templates](https://github.com/OpenAPITools/openapi-generator/tree/master/modules/openapi-generator/src/main/resources/go)
used by the `openapi-generator`, but modified for the sake of the development experience, or other concrete needs.

[The default ones](https://github.com/OpenAPITools/openapi-generator/tree/master/modules/openapi-generator/src/main/resources/go)
are used for those that are not present in the [templates](./templates) directory.

**If you want to modify the code generated, you'll likely need to modify these templates.**

#### Versioning

This project follows the [Semantic Versioning](https://semver.org/) guidelines,
using [pre-release syntax](https://semver.org/#spec-item-9) to indicate the client's version.

So, for instance, we can fix a bug in the client if needed, while keeping it targeting the same API version.

For instance, given `1.0.0-0.1.0`:
- The API version (from spec) would be: `1.0.0`
- The client version would be `0.1.0`.

#### GitHub Action

In the Grafana's [`shared-workflows` repository](https://github.com/grafana/shared-workflows), there is a GitHub Action
named [
`generate-openapi-clients`](https://github.com/grafana/shared-workflows/tree/main/actions/generate-openapi-clients),
which is almost equivalent to the command above, but meant to be executed within a GitHub workflow.

In the future, we might explore the use of that action to "automate" the generation of the client.
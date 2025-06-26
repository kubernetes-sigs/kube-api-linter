# Kube API Linter

Kube API Linter (KAL) is a Golang based linter for Kubernetes API types.
It checks for common mistakes and enforces best practices.
The rules implemented by the Kube API Linter, are based on the [Kubernetes API Conventions][api-conventions].

Kube API Linter is aimed at being an assistant to API review, by catching the mechanical elements of API review, and allowing reviewers to focus on the more complex aspects of API design.

[api-conventions]: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md

## Installation

Kube API Linter ships as a golangci-lint plugin, and a golangci-lint module.

### Golangci-lint Module

To install the `golangci-lint` module, first you must have `golangci-lint` installed.
If you do not have `golangci-lint` installed, review the `golangci-lint` [install guide][golangci-lint-install].

[golangci-lint-install]: https://golangci-lint.run/welcome/install/

### Using with golangci-lint v2 Plugin

Kube API Linter can be used as a plugin with `golangci-lint` v2. This setup allows integrating Kube API Linter alongside other static analysis tools within the same linter workflow.

More information about golangci-lint plugins can be found in the [golangci-lint plugin documentation][golangci-lint-plugin-docs].

#### Plugin Setup

First, build the plugin `.so` file:

```shell
go build -buildmode=plugin -o ./bin/kube-api-linter.so sigs.k8s.io/kube-api-linter/pkg/plugin
```

This will generate a `kube-api-linter.so` plugin in the `bin` directory.

#### Configuration

Update your `.golangci.yml` with the following snippet:

```yaml
version: "2"

linters:
  enable:
    - kubeapilinter
    ...

  settings:
    revive:
      rules:
        - name: comment-spacings
        - name: import-shadowing
    custom:
      kubeapilinter:
        path: "./bin/kube-api-linter.so"
        description: "Kube API Linter plugin"
        original-url: "sigs.k8s.io/kube-api-linter"
        settings:
          linters: {}
          lintersConfig: {}

exclusions:
  rules:
    # Only run kubeapilinter inside the 'api/' directory
    # Prevents running API-specific checks on unrelated packages like controllers or tests
    - path-except: "^api/"
      linters:
        - kubeapilinter
 ...
```

This approach allows full integration of Kube API Linter as a plugin, scoped correctly to relevant directories, while using `golangci-lint` v2 and coexisting with other linters.

**NOTE** For GolangCi Versions lower than 2x you might can do as follows:

You will need to create a `.custom-gcl.yml` file to describe the custom linters you want to run. The following is an example of a `.custom-gcl.yml` file:

```yaml
version:  v1.64.8 # GolangCi version
name: golangci-kube-api-linter
destination: ./bin
plugins:
- module: 'sigs.k8s.io/kube-api-linter'
  version: 'v0.0.0' # Replace with the latest version
```

Once you have created the custom configuration file, you can run the following command to build the custom `golangci-kal` binary:

```shell
golangci-lint custom
```

The output binary will be a combination of the initial `golangci-lint` binary and the Kube API linter plugin.
This means that you can use any of the standard `golangci-lint` configuration or flags to run the binary, but may also include the Kube API Linter rules.

If you wish to only use the Kube API Linter rules, you can configure your `.golangci.yml` file to only run the Kube API Linter:

```yaml
linters-settings:
  custom:
    kubeapilinter:
      type: "module"
      description: Kube API LInter lints Kube like APIs based on API conventions and best practices.
      settings:
        linters: {}
        lintersConfig: {}
linters:
  disable-all: true
  enable:
    - kubeapilinter

# To only run Kube API Linter on specific path
issues:
  exclude-rules:
    - path-except: "api/*"
      linters:
        - kubeapilinter
```

#### Optional Configurations

##### Enabling Specific Linters (Valid for Versions < 2.x)

> **Note**: The following configuration is only supported when using the **module integration** via `.custom-gcl.yml` and is **not valid in plugin mode** (`.so` files with `golangci-lint` v2.x+).

If you wish to only run selected linters you can do so by specifying the linters you want to enable in the `linters` section:

```yaml
linters-settings:
  custom:
    kubeapilinter:
      type: "module"
      settings:
        linters:
          disable:
            - "*"
          enable:
            - requiredfields
            - statusoptional
            - statussubresource
```

The settings for Kube API Linter are based on the [GolangCIConfig][golangci-config-struct] struct and allow for finer control over the linter rules.

If you wish to use the Kube API Linter in conjunction with other linters, you can enable the Kube API Linter in the `.golangci.yml` file by ensuring that `kubeapilinter` is in the `linters.enabled` list.
To provide further configuration, add the `custom.kubeapilinter` section to your `linter-settings` as per the example above.

[golangci-config-struct]: https://pkg.go.dev/sigs.k8s.io/kube-api-linter/pkg/config#GolangCIConfig

Where fixes are available within a rule, these can be applied automatically with the `--fix` flag.

```shell
golangci-kube-api-linter run path/to/api/types --fix
```

#### VSCode integration

Since VSCode already integrates with `golangci-lint` via the [Go][vscode-go] extension, you can use the `golangci-kal` binary as a linter in VSCode.
If your project authors are already using VSCode and have the configuration to lint their code when saving, this can be a seamless integration.

Ensure that your project setup includes building the `golangci-kube-api-linter` binary, and then configure the `go.lintTool` and `go.alternateTools` settings in your project `.vscode/settings.json` file.

[vscode-go]: https://code.visualstudio.com/docs/languages/go

```json
{
    "go.lintTool": "golangci-lint",
    "go.alternateTools": {
        "golangci-lint": "${workspaceFolder}/bin/golangci-kube-api-linter",
    }
}
```

Alternatively, you can also replace the binary with a script that runs the `golangci-kube-api-linter` binary,
allowing for customisation or automatic copmilation of the project should it not already exist.

```json
{
    "go.lintTool": "golangci-lint",
    "go.alternateTools": {
        "golangci-lint": "${workspaceFolder}/hack/golangci-lint.sh",
    }
}
```

# Contributing

New linters can be added by following the [New Linter][new-linter] guide.

[new-linter]: docs/new-linter.md

# Linters

For a complete list of available linters and their configuration options, see [docs/linters.md](docs/linters.md).

# License

KAL is licensed under the Apache License, Version 2.0. See [LICENSE](LICENSE) for the full license text.

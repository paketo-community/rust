# Rust Cloud Native Buildpack

The Rust Cloud Native Buildpack provides a set of collaborating buildpacks that enable the building of a Rust-based application. These buildpacks include:

- [Rust Dist CNB](https://github.com/dmikusa/rust-dist-cnb)
- [Rust Cargo CNB](https://github.com/dmikusa/rust-cargo-cnb)

## tl;dr

You can build your apps with either of these commands:

- `pack build <image-name> -b dmikusa/rust`
- `pack build <image-name> --builder dmikusa/rust-builder`

## What's Included

### A Builder

In this repo is a sample `builder.toml` that you can use to create your own builder. To create the builder, just run `pack create-builder <your-name>/rust-builder --builder-config builder.toml`. For example, `pack create-builder dmikusa/rust-builder --builder-config builder.toml`.

You can then build an app with it using `pack build <image-name> --builder <your-name>/rust-builder`.

This is the same builder that's published [to Docker Hub](https://hub.docker.com/repository/docker/dmikusa/rust-builder). You can build against that with `pack build <image-name> --builder dmikusa/rust-builder`.

### A Meta Buildpack

Also included in this repo is a `buildpack.toml` that defines the Rust meta buildpack. It allows you to reference all of the Rust related CNBs in one convenient way & know that the order is set correctly.

To build, run the following:

```bash
export VERSION=v0.0.X
git checkout $VERSION
mkdir -p build
tar czf build/buildpack.tgz buildpack.toml
pack package-buildpack build/rust-$VERSION.cnb --package-config package.toml --format file
pack package-buildpack dmikusa/rust --package-config package.toml --format image
```

That will leave you with three things:
- The file `build/buildpack.tgz` which is the raw buildpack. Not totally useful, but required for `pack package-buildpack`.
- The file `build/rust-$VERSION.cnb` which can be consumed by `pack build`. Ex: `pack build -b build/rust-$VERSION.cnb`.
- A docker image called `dmikusa/rust`. You can then `pack build -b <your-name>/rust@<version>`. You can `docker tag` and `docker push` to your registry of choice as well.

If you don't want to build:

1. A pre-built build package is attached to every release under the [releases section of this repo](https://github.com/dmikusa/rust-cnb/releases). This is the same as the `.cnb` file you'd get if you build.

2. A pre-built docker image is available from [Docker Hub](https://hub.docker.com/repository/docker/dmikusa/rust). This is the same as what you'd get if you built a docker image with the previous instructions.

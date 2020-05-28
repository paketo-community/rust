# Rust Cloud Native Buildpack

The Rust Cloud Native Buildpack provides a set of collaborating buildpacks that enable the building of a Rust-based application. These buildpacks include:

    [Rust Dist CNB](https://github.com/dmikusa/rust-dist-cnb)
    [Rust Cargo CNB](https://github.com/dmikusa/rust-cargo-cnb)

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

To build your own just run: `./scripts/package.sh --version <version>`, for example `./scripts/package.sh --version v0.0.1` (there must be a git tag for the version specified). After this runs, look in the `./build/` directory. It'll have a tgz and cnb file which can be used with `pack build <image-name> -b ./build/buildpackage.cnb`.

If you don't want to build, a pre-built build package is available under the releases section. This is the same as building the meta buildpack as stated in the previous paragraph.

You can also build this meta buildpack and publish it as an image by running `pack package-buildpack <your-name>/rust --package-config build/package.toml --format image` (`./build/` must exist, if not run the `package.sh` command from the previous paragraph). You can then `pack build -b <your-name>/rust@<version>`. You can `docker tag` and `docker push` to your registry of choice as well.

If you don't want to do that, you can use the image from [Docker Hub](https://hub.docker.com/repository/docker/dmikusa/rust).

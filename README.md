# `docker.io/paketocommunity/rust`

The Rust Cloud Native Buildpack provides a set of collaborating buildpacks that enable the building of a Rust-based application.

## Included Buildpacks

- [`paketo-community/rustup`](https://github.com/paketo-community/rustup)
- [`paketo-community/rust-dist`](https://github.com/paketo-community/rust-dist)
- [`paketo-community/cargo`](https://github.com/paketo-community/cargo)

## tl;dr

You can build your apps with either of these commands:

- `pack build <image-name> -b docker.io/paketocommunity/rust`

## What's Included

### A Meta Buildpack

Included in this repo is a `buildpack.toml` that defines the Rust meta buildpack. It allows you to reference all of the Rust related CNBs in one convenient way & know that the order is set correctly.

This is available throuh a pre-built docker image on [Docker Hub](https://hub.docker.com/repository/docker/paketocommunity/rust).

To use: `pack build <image-name> -b docker.io/paketocommunity/rust`

### A Builder

In this repo is a sample `example-builder.toml` that you can use to create your own builder. We do not publish a builder, so if you'd like to use this route you must first create the builder. 

To create the builder, just run `pack builder create <published-to>/rust-builder --config example-builder.toml`. For example, `pack builder create paketocommunity/rust-builder --config example-builder.toml`.

You can then build an app with it using `pack build <image-name> --builder <published-to>/rust-builder`.

## License

This buildpack is released under version 2.0 of the [Apache License][a].

[a]: http://www.apache.org/licenses/LICENSE-2.0

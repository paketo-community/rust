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

The builder is configure to use the base build and run images, which is a reasonable mix of size and functionality. You may change the build and run images to use the full stack, which has a lot more libraries and tools installed but is quite a bit larger, or you can use the tiny build and run images which presents a very small container but doesn't even contain a shell, which can make debugging difficult.

To switch, swap the stack that you'd like to use in the builder file.

Example full stack:

```
[stack]
  id = "io.buildpacks.stacks.bionic"
  build-image = "gcr.io/paketo-buildpacks/build:full-cnb-cf"
  run-image = "gcr.io/paketo-buildpacks/run:full-cnb-cf"
```

Example tiny stack:

```
[stack]
  id = "io.paketo.stacks.tiny"
  build-image = "gcr.io/paketo-buildpacks/build:tiny-cnb"
  run-image = "gcr.io/paketo-buildpacks/run:tiny-cnb"
```

## License

This buildpack is released under version 2.0 of the [Apache License][a].

[a]: http://www.apache.org/licenses/LICENSE-2.0

# `paketocommunity/rust`

The Rust Cloud Native Buildpack provides a set of collaborating buildpacks that enable the building of a Rust-based application.

## Included Buildpacks

- [`paketo-community/rustup`](https://github.com/paketo-community/rustup)
- [`paketo-community/rust-dist`](https://github.com/paketo-community/rust-dist)
- [`paketo-community/cargo`](https://github.com/paketo-community/cargo)
- [`paketo-buildpacks/procfile`](https://github.com/paketo-buildpacks/procfile)

## tl;dr

You can build your app with two steps:

1. Run `pack build <image-name> -b docker.io/paketocommunity/rust`. An image will be generated that you can run.
2. Run the image with `docker run -it <image-name>`.

By default, the Rust buildpack will add process types for all of the binary projects in your workspace. You may optionally add a [`Procfile`](https://paketo.io/docs/howto/configuration/#procfiles) though if you need to customize the start command for your container.

## What's Included

### A Meta Buildpack

Included in this repo is a `buildpack.toml` that defines the Rust meta buildpack. It allows you to reference all of the Rust related CNBs in one convenient way & know that the order is set correctly.

This is available through a pre-built docker image on [Docker Hub](https://hub.docker.com/repository/docker/paketocommunity/rust).

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
  id = "io.buildpacks.stacks.jammy"
  build-image = "docker.io/paketobuildpacks/build-jammy-base:latest"
  run-image = "docker.io/paketobuildpacks/run-jammy-base:latest"
```

Example tiny stack:

```
[stack]
  id = "io.buildpacks.stacks.jammy.tiny"
  build-image = "docker.io/paketobuildpacks/build-jammy-tiny"
  run-image = "docker.io/paketobuildpacks/run-jammy-tiny"
```

## Tips

- When building for the Paketo Tiny stack, you must build with the target set to `x86_64-unknown-linux-musl`. This is because if you build with the glibc target, you will end up with a dependency for the `libgcc_s.so.1` library that does not exist on the Paketo Tiny run image. If we target musl instead, you'll get a binary that is statically compiled against musl and does not have this requirement, thus it will run on the Tiny stack.

    When you run with the Tiny stack, the buildpacks will detect this and automatically adjust. This includes installing the additional target to cross-compile and modifying `cargo` commands cross-compile using the musl target.

- By default, the Rust buildpack will install `tini`, which is a very small process manager that handles the PID1 responsibilities (it's also the same running `docker run --init`), into your application image. Because your application will run as PID1, this ensures that your application will properly handle signals. If you are running your image with Docker, this is what makes CTRL+C function correctly. If you're running on Kubernetes, this is what makes graceful termination work correctly.

    It is generally recommended that you keep `tini` installed. It's extremely low-overhead. That said, if you really, really know what you're doing and your application is properly handling signals and reaping child processes then you may set `$BP_CARGO_TINI_DISABLED=true` and the buildpacks will skip installing `tini.`.

- What's the difference between using a builder and using the composite buildpack (i.e. `docker.io/paketocommunity/rust`)? There's not a lot of difference, so feel free to use which one is more convenient for your use cases. The composite buildpack is published to Docker Hub so it's often easier to get started using it. We are not publishing a builder image at the moment, so you need to build that on your own. The builder does offer some additional customization options though. Plus, you can set a default builder with `pack`, so if you're building many different applications it can reduce the amount of flags you need to send to `pack build`. You do not need to use both the builder and composite buildpack.

- Can I specify multiple processes in my `Procfile`? Yes! To make this work, just add more lines to `Procfile`. One of the lines needs to be named `web`, which is the expected name for the default process. Alternatively, you may use different names but pass the `--default-process` argument to `pack build` to indicate the default process name in your `Procfile`. It is highly recommended that you have a default process. See [here for more details on multiple process containers](https://buildpacks.io/docs/app-developer-guide/run-an-app/).

- Can I pass arguments to my application? Yes, you just append them to the end of your `docker run` command. See [here for an example](https://buildpacks.io/docs/app-developer-guide/run-an-app/#default-process-type-with-additional-arguments).

- How do I run a non-default process from my `Procfile`? Pass `--entrypoint <process name>` to `docker run`. See [here for an example](https://buildpacks.io/docs/app-developer-guide/run-an-app/#non-default-process-type).

- How do I launch a shell in my container to debug? There are two ways to do this. [First](https://buildpacks.io/docs/app-developer-guide/run-an-app/#user-provided-shell-process), `docker run --rm --entrypoint launcher -it <image-name> bash`. The first will set up the shell with the standard environment, the same as what your application will get when it launches. [Second](https://buildpacks.io/docs/app-developer-guide/run-an-app/#no-launcher), `docker run --rm --entrypoint bash -it <image-name>`. The second will give you a blank environment without any buildpack specific environment variables set.

## License

This buildpack is released under version 2.0 of the [Apache License][a].

[a]: http://www.apache.org/licenses/LICENSE-2.0

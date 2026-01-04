## About

#### Introduction

This repository modifies a raw virtual machine image for a given
hardenedBSD release. This allows GitHub Action runners to execute
tests and builds on hardenedBSD virtual machines. If you just want
to use the GitHub Action, then see the
[hardenedbsd-vm](https://github.com/0x1eef/hardenedbsd-vm)
repository.

#### How does it work?

1. A raw virtual machine image is built and uploaded as a GitHub release artifact
2. The workflow downloads this image and mounts it so it can be modified
3. The image is mounted with the help of mdconfig(8) and mount(8)
4. Configuration files are copied into the image (see [config/](config/))
5. The image is then unmounted and released as a new GitHub release artifact

See [.github/workflows/build.yml](.github/workflows/build.yml) for details.

#### Notes

The [hardenedbsd-builder](https://github.com/0x1eef/hardenedbsd-builder)
repository expects to be run on either FreeBSD or hardenedBSD hosts. You
can see how this is acheived via a GitHub workflow in
[.github/workflows/build.yml](.github/workflows/build.yml). The original virtual
machine images are built manually from `/usr/src/release` and this may change
once hardenedBSD provides official images.

## About

This repository modifies a raw virtual machine image for a given
hardenedBSD release. This allows GitHub action runners to execute
tests and builds on hardenedBSD virtual machines. If you just want
to use the GitHub action, then see the
[hardenedbsd-vm](https://github.com/0x1eef/hardenedbsd-vm)
repository.

## Workflow

**Overview**

1. A raw virtual machine image is built and uploaded as a GitHub release artifact
2. The workflow downloads this image and mounts it so it can be modified
3. The image is mounted with the help of mdconfig(8) and mount(8)
4. Configuration files are copied into the image (see [config/](config/))
5. The image is then unmounted and released as a new GitHub release artifact

See [.github/workflows/build.yml](.github/workflows/build.yml) for details.

**Documentation**

* 1. [How to build a hardenedBSD VM image](https://0x1eef.github.io/posts/how-to-build-a-hardenedbsd-vm-image/)

## Sources

* [github.com/@0x1eef](https://github.com/0x1eef/hardenedbsd-vm)
* [git.hardenedBSD.org/@0x1eef](https://git.hardenedBSD.org/0x1eef/hardenedbsd-vm)
* [bsd.cafe/@0x1eef](https://brew.bsd.cafe/0x1eef/hardenedbsd-vm)

## License

[BSD Zero Clause](https://choosealicense.com/licenses/0bsd/) <br>
See [LICENSE](./LICENSE)

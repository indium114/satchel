# satchel

**satchel** is a command-line file clipboard, designed to make file management easier.

## Installation

### with Nix

Add the repo to your flake inputs...

```nix
inputs = {
  satchel.url = "github:indium114/satchel";
};
```

And pass it into your `environment.systemPackages`...

```nix
environment.systemPackages = [
  inputs.satchel.packages.${pkgs.stdenv.hostPlatform.system}.satchel
]
```

### with Go

To install, simply run:

```shell
go install github.com/indium114/satchel@latest
```

## Usage

### Add a file to the satchel

```shell
# satchel add <filepath>
satchel add foo.txt
```

### List the contents of the satchel

```shell
satchel ls
```

### Paste a file from the satchel

```shell
# satchel put <id>
satchel put 1
```

### Remove a file from the satchel

```shell
# satchel rm <id>
satchel rm 1
```

### Drop the contents of the satchel altogether

```shell
satchel drop
```

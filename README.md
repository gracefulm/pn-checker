# pn-checker (Prime Number Checker)

pn-checker is a slightly silly and useless command line tool.
It checks whether the number specified is a prime number or not.
And sometimes, pn-checker playfully returns you "Think for yourself, stupid!".

## Installation

```shell
> go install github.com/gracefulm/pn-checker@latest
```

## Usage

```shell
> pn-checker -h
pn-checker checks whether the number specified is a prime number or not.
And sometimes, it says "Think for yourself, stupid!" to you.

Usage:
  pn-checker [flags] [arg(positive integer)]

Examples:
  pn-checker 997

Flags:
  -d, --diligent   seriously check if it's a prime number
  -h, --help       help for pn-checker
  -v, --version    version for pn-checker
```

## Example

```text
> pn-checker 13
Congrats!! The number 13 is a prime number.
```

OR

```text
> pn-checker 13

 _____ _     _       _       __
|_   _| |__ (_)_ __ | | __  / _| ___  _ __
  | | | '_ \| | '_ \| |/ / | |_ / _ \| '__|
  | | | | | | | | | |   <  |  _| (_) | |
  |_| |_| |_|_|_| |_|_|\_\ |_|  \___/|_|
                                 _  __        _               _     _   _
 _   _  ___  _   _ _ __ ___  ___| |/ _|   ___| |_ _   _ _ __ (_) __| | | |
| | | |/ _ \| | | | '__/ __|/ _ \ | |_   / __| __| | | | '_ \| |/ _  | | |
| |_| | (_) | |_| | |  \__ \  __/ |  _|  \__ \ |_| |_| | |_) | | (_| | |_|
 \__, |\___/ \__,_|_|  |___/\___|_|_|( ) |___/\__|\__,_| .__/|_|\__,_| (_)
 |___/                               |/                |_|

```

`toy-tax-calculator`
====================

Calculates fictitious taxes given gross income and a tax year.

Usage
-----

```
$ toy-tax-calculator --help
Calculates (fictitious) taxes based on gross income for a given year

Usage:
  toy-tax-calculator [flags]

Flags:
      --gross-income string   Total amount of gross income (default "0")
  -h, --help                  help for toy-tax-calculator
      --tax-year int          Year where taxes apply (default 2019)
```

Developing
----------

### Requirements

* Go ≥ 1.13
* _(optionally)_ Podman or Docker
* GNU Make

### Building

```
make
```

Optionally if using Podman or Docker:

```
podman build -t toy-tax-calculator:develop .
```

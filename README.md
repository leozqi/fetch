# fetch

fetch is the soon-to-be new universal package manager.
This is intended to be a learner's package manager to introduce me to many new concepts:

- DAG's and the handling of dependencies
- deterministic builds, etc.

I will be using the following as inspiration:

- Homebrew's `brew` structure
- Debian's `dpkg`/`apt`
- Alpine's `apk`
- Python's `pip`
- Arch's `pacman  `

 as inspiration.

# Planned features

- Written in `golang`
- Package target/repo as a online git repository
- Packages as folders in that repository with manifest, checksum, and code
- install commands per platform

The code aims to be:

- complex but not complicated
- comprehensively tested and documented
- to golang best practices
- not one person's project, once I get this going


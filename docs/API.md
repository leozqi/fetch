This is the user interface I'm thinking of:

- `fetch refresh`: Updates internal package listing (IF this can be done each command such that the app doesn't need an internal storage file, that's even better)
- `fetch install <PACKAGE> --OPTS`: install a package with OPTions.

Fetch will NOT be configured with user/machine level presets.
There will be "default" presets, and there will be options specified when running the install command.
That way you'll always know what you are getting as the user.

- `fetch update <PACKAGE> --OPTS`: update a package.
- `fetch uninstall <PACKAGE> --OPTS`: delete a package.

- `fetch list --installed | --other-opts` : list packages installed, all packages, etc.
- `fetch view <PACKAGE>`: prints out ALL locations where fetch has dumped stuff on your computer regarding this package
- `fetch depends <PACKAGE> --recurse`: print out all the dependencies of this package.
If recurse, include the nth level dependencies (actually every package that must be downloaded)

- `fetch from <GITHUB SRC>`: add a github repo for packages
- `fetch sources`: list all the sources currently being used.


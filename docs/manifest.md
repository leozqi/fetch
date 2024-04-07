# Information on the manifest.toml used

The `manifest.toml` contains the build files

# fetch on disk

Fetch will store the following in `os.UserCacheDir()/fetch`:

```
package1/
    manifest.toml
package2/
    manifest.toml
```
https://github.com/src-d/go-git/issues/430

versioning will be done with `git log --follow -- filename`
Architecture is one of GOOS/GOARCH https://pkg.go.dev/runtime#GOOS

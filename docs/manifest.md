# Information on the manifest.toml used

The `manifest.toml` contains the build files

# fetch on disk

Fetch will store the following in `os.UserConfigDir()/fetch`:

```
sources/
- repo1/
- repo2/

state.json
```
https://github.com/src-d/go-git/issues/430
tree ~/Library/Application\ Support/fetch    
versioning will be done with `git log --follow -- filename`
Architecture is one of GOOS/GOARCH https://pkg.go.dev/runtime#GOOS

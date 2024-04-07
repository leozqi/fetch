package gitpipe

import (
	"fmt"
	//	"github.com/BurntSushi/toml"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
	"io"
	"os"
	"path/filepath"
)

func getCacheDir() string {
	writeTo, _ := os.UserCacheDir()
	writeTo += "/fetch"
	fmt.Println(writeTo)
	return writeTo
}

// Comment
func GetOrigin(url string) error {
	origin := memfs.New()

	_, err := git.Clone(memory.NewStorage(), origin, &git.CloneOptions{
		URL:   url,
		Depth: 1,
	})
	if err != nil {
		return err
	}

	files, err := origin.ReadDir("/")
	if err != nil {
		return err
	}

	for _, pkg := range files {
		if !pkg.IsDir() {
			continue
		}

		src, err := origin.Open(origin.Join(pkg.Name(), "/manifest.toml"))
		if err != nil {
			return err
		}


        err = os.MkdirAll(filepath.Join(getCacheDir(), pkg.Name()), 0750)
        if err != nil {
            return err
        }

        out, err := os.Create(filepath.Join(getCacheDir(), pkg.Name(), filepath.Base("manifest.toml")))
		if err != nil {
			return err
		}

		_, err = io.Copy(out, src)
		if err != nil {
			return err
		}

	}
	return nil
}

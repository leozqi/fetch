package manager

import (
	"encoding/json"
	"errors"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	// github.com/BurntSushi/toml
)

func getStateFilePath() string {
	path, err := os.UserConfigDir()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Join(path, "fetch", filepath.Base("state.json"))
}

type State struct {
	Sources map[string]string `json:"sources"`
}

func LoadState() *State {
	if _, err := os.Stat(getStateFilePath()); errors.Is(err, os.ErrNotExist) {
		return &State{
			Sources: map[string]string{},
		}
	}

	file, err := os.Open(getStateFilePath())
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	state := State{}
	content, err := ioutil.ReadFile(getStateFilePath())

	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(content, &state)
	if err != nil {
		log.Fatal(err)
	}

	return &state
}

func getSourceDir() string {
	writeTo, err := os.UserConfigDir()

	if err != nil {
		log.Fatal("Unable to access os.UserConfigDir")
	}
	return filepath.Join(writeTo, "fetch")
}

func AddSource(state *State, name string, url string) error {
	_, err := git.PlainClone(
		filepath.Join(getSourceDir(), "sources", name),
		false,
		&git.CloneOptions{
			URL:   url,
			Depth: 1,
		},
	)
	if err != nil {
		return err
	}

	state.Sources[name] = url

	return nil
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

		err = os.MkdirAll(filepath.Join(getStateFilePath(), pkg.Name()), 0750)
		if err != nil {
			return err
		}

		out, err := os.Create(filepath.Join(getStateFilePath(), pkg.Name(), filepath.Base("manifest.toml")))
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

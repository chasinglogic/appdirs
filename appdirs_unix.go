//go:build linux || freebsd || netbsd || openbsd
// +build linux freebsd netbsd openbsd

package appdirs

import (
	"os"
	"path/filepath"
)

func naiveTildeExpand(path string) string {
	if path[0] == '~' {
		return filepath.Join(homeDir(), path[1:])
	}

	return path
}

func userDataDir(name, author, version string, roaming bool) (path string) {
	if path = os.Getenv("XDG_DATA_HOME"); path == "" {
		path = filepath.Join(homeDir(), ".local", "share")
	}

	if name != "" {
		path = filepath.Join(path, name, version)
	}

	return path
}

func SiteDataDirs(name, author, version string) (paths []string) {
	var path string

	if path = os.Getenv("XDG_DATA_DIRS"); path == "" {
		paths = []string{"/usr/local/share", "/usr/share"}
	} else {
		paths = filepath.SplitList(path)
	}

	for i, path := range paths {
		path = naiveTildeExpand(path)

		if name != "" {
			path = filepath.Join(path, name, version)
		}

		paths[i] = path
	}

	return paths
}

func siteDataDir(name, author, version string) (path string) {
	return SiteDataDirs(name, author, version)[0]
}

func userConfigDir(name, author, version string, roaming bool) (path string) {
	if path = os.Getenv("XDG_CONFIG_HOME"); path == "" {
		path = filepath.Join(homeDir(), ".config")
	}

	if name != "" {
		path = filepath.Join(path, name, version)
	}

	return path
}

func SiteConfigDirs(name, author, version string) (paths []string) {
	var path string

	if path = os.Getenv("XDG_CONFIG_DIRS"); path == "" {
		paths = []string{"/etc/xdg"}
	} else {
		paths = filepath.SplitList(path)
	}

	for i, path := range paths {
		path = naiveTildeExpand(path)

		if name != "" {
			path = filepath.Join(path, name, version)
		}

		paths[i] = path
	}

	return paths
}

func siteConfigDir(name, author, version string) (path string) {
	return SiteConfigDirs(name, author, version)[0]
}

func userCacheDir(name, author, version string, opinion bool) (path string) {
	if path = os.Getenv("XDG_CACHE_HOME"); path == "" {
		path = filepath.Join(homeDir(), ".cache")
	}

	if name != "" {
		path = filepath.Join(path, name, version)
	}

	return path
}

func userLogDir(name, author, version string, opinion bool) (path string) {
	path = UserCacheDir(name, author, version, opinion)

	return filepath.Join(path, "log")
}

//go:build linux || freebsd || netbsd || openbsd
// +build linux freebsd netbsd openbsd

package appdirs

import (
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func naiveTildeExpand(path string) string {
	if path[0] == '~' {
		return filepath.Join(homeDir(), path[1:])
	}

	return path
}

// User path functions
func userDataDir(name, author, version string, roaming bool) (path string) {
	if path = os.Getenv("XDG_DATA_HOME"); path == "" {
		path = filepath.Join(homeDir(), ".local", "share")
	}
	if name != "" {
		path = filepath.Join(path, name, version)
	}
	return path
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

func userStateDir(name, author, version string, roaming bool) (path string) {
	if path = os.Getenv("XDG_STATE_HOME"); path == "" {
		path = filepath.Join(homeDir(), ".local", "state")
	}
	if name != "" {
		path = filepath.Join(path, name, version)
	}
	return path
}

func userRuntimeDir(name, author, version string) (path string) {
	if path = os.Getenv("XDG_RUNTIME_DIR"); path == "" {
		uid := os.Getuid()
		uidStr := strconv.Itoa(uid)
		switch {
		case strings.HasPrefix(runtimeOS(), "openbsd"):
			path = filepath.Join("/tmp/run/user", uidStr)
		case strings.HasPrefix(runtimeOS(), "freebsd"), strings.HasPrefix(runtimeOS(), "netbsd"):
			path = filepath.Join("/var/run/user", uidStr)
		default:
			path = filepath.Join("/run/user", uidStr)
		}
	}
	if name != "" {
		path = filepath.Join(path, name, version)
	}
	return path
}

func userDocumentsDir() string    { return userMediaDir("XDG_DOCUMENTS_DIR", "~/Documents") }
func userDownloadsDir() string    { return userMediaDir("XDG_DOWNLOAD_DIR", "~/Downloads") }
func userPicturesDir() string     { return userMediaDir("XDG_PICTURES_DIR", "~/Pictures") }
func userVideosDir() string       { return userMediaDir("XDG_VIDEOS_DIR", "~/Videos") }
func userMusicDir() string        { return userMediaDir("XDG_MUSIC_DIR", "~/Music") }
func userDesktopDir() string      { return userMediaDir("XDG_DESKTOP_DIR", "~/Desktop") }
func userBinDir() string          { return filepath.Join(homeDir(), ".local", "bin") }
func userApplicationsDir() string { return filepath.Join(homeDir(), ".local", "share", "applications") }

// Site path functions
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

func siteCacheDir(name, author, version string, opinion bool) (path string) {
	path = "/var/cache"
	if name != "" {
		path = filepath.Join(path, name, version)
	}
	return path
}

func siteStateDir(name, author, version string) (path string) {
	path = "/var/lib"
	if name != "" {
		path = filepath.Join(path, name, version)
	}
	return path
}

func siteLogDir(name, author, version string, opinion bool) (path string) {
	path = "/var/log"
	if name != "" {
		path = filepath.Join(path, name, version)
	}
	return path
}

func siteRuntimeDir(name, author, version string) (path string) {
	switch {
	case strings.HasPrefix(runtimeOS(), "freebsd"), strings.HasPrefix(runtimeOS(), "openbsd"), strings.HasPrefix(runtimeOS(), "netbsd"):
		path = "/var/run"
	default:
		path = "/run"
	}
	if name != "" {
		path = filepath.Join(path, name, version)
	}
	return path
}

func siteBinDir() string          { return "/usr/local/bin" }
func siteApplicationsDir() string { return filepath.Join("/usr/local/share", "applications") }

func userMediaDir(envKey, fallback string) string {
	if val := strings.TrimSpace(os.Getenv(envKey)); val != "" {
		return naiveTildeExpand(val)
	}
	configPath := filepath.Join(homeDir(), ".config", "user-dirs.dirs")
	if data, err := os.ReadFile(configPath); err == nil {
		lines := strings.Split(string(data), "\n")
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if strings.HasPrefix(line, envKey+"=") {
				val := strings.TrimPrefix(line, envKey+"=")
				val = strings.Trim(val, "\"")
				val = strings.ReplaceAll(val, "$HOME", homeDir())
				return naiveTildeExpand(val)
			}
		}
	}
	return naiveTildeExpand(fallback)
}

// runtimeOS is split out for tests and to avoid importing runtime multiple times.
func runtimeOS() string { return runtime.GOOS }

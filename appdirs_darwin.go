package appdirs

import (
	"path/filepath"
)

// User path functions
func userDataDir(name, author, version string, roaming bool) (path string) {
	path = filepath.Join(homeDir(), "Library", "Application Support")
	if name != "" {
		path = filepath.Join(path, name)
	}
	if name != "" && version != "" {
		path = filepath.Join(path, version)
	}
	return path
}

func userConfigDir(name, author, version string, roaming bool) (path string) {
	return UserDataDir(name, author, version, roaming)
}

func userCacheDir(name, author, version string, opinion bool) (path string) {
	path = filepath.Join(homeDir(), "Library", "Caches")
	if name != "" {
		path = filepath.Join(path, name)
	}
	if name != "" && version != "" {
		path = filepath.Join(path, version)
	}
	return path
}

func userLogDir(name, author, version string, opinion bool) (path string) {
	path = filepath.Join(homeDir(), "Library", "Logs", name)
	if name != "" && version != "" {
		path = filepath.Join(path, version)
	}
	return path
}

func userStateDir(name, author, version string, roaming bool) string {
	return UserDataDir(name, author, version, roaming)
}

func userRuntimeDir(name, author, version string) (path string) {
	path = filepath.Join(homeDir(), "Library", "Caches", "TemporaryItems")
	if name != "" {
		path = filepath.Join(path, name)
	}
	if name != "" && version != "" {
		path = filepath.Join(path, version)
	}
	return path
}

func userDocumentsDir() string    { return filepath.Join(homeDir(), "Documents") }
func userDownloadsDir() string    { return filepath.Join(homeDir(), "Downloads") }
func userPicturesDir() string     { return filepath.Join(homeDir(), "Pictures") }
func userVideosDir() string       { return filepath.Join(homeDir(), "Movies") }
func userMusicDir() string        { return filepath.Join(homeDir(), "Music") }
func userDesktopDir() string      { return filepath.Join(homeDir(), "Desktop") }
func userBinDir() string          { return filepath.Join(homeDir(), ".local", "bin") }
func userApplicationsDir() string { return filepath.Join(homeDir(), "Applications") }

// Site path functions
func siteDataDir(name, author, version string) (path string) {
	path = "/Library/Application Support"
	if name != "" {
		path = filepath.Join(path, name)
	}
	if name != "" && version != "" {
		path = filepath.Join(path, version)
	}
	return path
}

func siteConfigDir(name, author, version string) (path string) {
	return SiteDataDir(name, author, version)
}

func siteCacheDir(name, author, version string, opinion bool) (path string) {
	path = filepath.Join("/Library", "Caches")
	if name != "" {
		path = filepath.Join(path, name)
	}
	if name != "" && version != "" {
		path = filepath.Join(path, version)
	}
	return path
}

func siteStateDir(name, author, version string) string {
	return SiteDataDir(name, author, version)
}

func siteLogDir(name, author, version string, opinion bool) (path string) {
	path = filepath.Join("/Library", "Logs")
	if name != "" {
		path = filepath.Join(path, name)
	}
	if name != "" && version != "" {
		path = filepath.Join(path, version)
	}
	return path
}

func siteRuntimeDir(name, author, version string) string {
	return userRuntimeDir(name, author, version)
}

func siteBinDir() string          { return "/usr/local/bin" }
func siteApplicationsDir() string { return "/Applications" }

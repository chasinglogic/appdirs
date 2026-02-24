// A port of the excellent python module `platformdirs`.
// See https://github.com/platformdirs/platformdirs for the python version.
package appdirs

import "os/user"

// App is a helper type to create easy access across your program to the appdirs
// functions.
//
// The *App type has 6 methods that map to the 6 functions exported by `appdirs`.
// All methods take no arguments, and supply the function it wraps with arguments
// pre-set in the struct on creation.
type App struct {
	Name    string
	Author  string
	Version string
	Roaming bool
	Opinion bool
}

// New returns a new App helper that has various methods for receiving
// relevant directories for your application.
//
// The following defaults are used for the fields not settable by New:
// Author: "", Version: "", Roaming: false, Opinion: true
//
// If you want to set these, create your own App struct by the usual means or
// modify the returned value accordingly.
func New(name string) *App {
	return &App{
		Name:    name,
		Author:  "",
		Version: "",
		Roaming: false,
		Opinion: true,
	}
}

// UserData returns the full path to the user-specific data directory
func (app *App) UserData() string {
	return UserDataDir(app.Name, app.Author, app.Version, app.Roaming)
}

// UserConfig returns the full path to the user-specific configuration directory
func (app *App) UserConfig() string {
	return UserConfigDir(app.Name, app.Author, app.Version, app.Roaming)
}

// UserCache returns the full path to the user-specific cache directory
func (app *App) UserCache() string {
	return UserCacheDir(app.Name, app.Author, app.Version, app.Opinion)
}

// UserLog returns the full path to the user-specific log directory
func (app *App) UserLog() string {
	return UserLogDir(app.Name, app.Author, app.Version, app.Opinion)
}

// UserState returns the full path to the user-specific state directory
func (app *App) UserState() string {
	return UserStateDir(app.Name, app.Author, app.Version, app.Roaming)
}

// UserRuntime returns the full path to the user-specific runtime directory
func (app *App) UserRuntime() string {
	return UserRuntimeDir(app.Name, app.Author, app.Version)
}

// UserDocuments returns the user's documents directory
func (app *App) UserDocuments() string { return UserDocumentsDir() }

// UserDownloads returns the user's downloads directory
func (app *App) UserDownloads() string { return UserDownloadsDir() }

// UserPictures returns the user's pictures directory
func (app *App) UserPictures() string { return UserPicturesDir() }

// UserVideos returns the user's videos directory
func (app *App) UserVideos() string { return UserVideosDir() }

// UserMusic returns the user's music directory
func (app *App) UserMusic() string { return UserMusicDir() }

// UserDesktop returns the user's desktop directory
func (app *App) UserDesktop() string { return UserDesktopDir() }

// UserBin returns the user's bin directory
func (app *App) UserBin() string { return UserBinDir() }

// UserApplications returns the user's applications directory
func (app *App) UserApplications() string { return UserApplicationsDir() }

// SiteData returns the full path to the user-shared data directory
func (app *App) SiteData() string {
	return SiteDataDir(app.Name, app.Author, app.Version)
}

// SiteConfig returns the full path to the user-shared configuration directory
func (app *App) SiteConfig() string {
	return SiteConfigDir(app.Name, app.Author, app.Version)
}

// SiteCache returns the full path to the shared cache directory
func (app *App) SiteCache() string {
	return SiteCacheDir(app.Name, app.Author, app.Version, app.Opinion)
}

// SiteState returns the full path to the shared state directory
func (app *App) SiteState() string {
	return SiteStateDir(app.Name, app.Author, app.Version)
}

// SiteLog returns the full path to the shared log directory
func (app *App) SiteLog() string {
	return SiteLogDir(app.Name, app.Author, app.Version, app.Opinion)
}

// SiteRuntime returns the full path to the shared runtime directory
func (app *App) SiteRuntime() string {
	return SiteRuntimeDir(app.Name, app.Author, app.Version)
}

// SiteBin returns the shared bin directory
func (app *App) SiteBin() string { return SiteBinDir() }

// SiteApplications returns the shared applications directory
func (app *App) SiteApplications() string { return SiteApplicationsDir() }

// homeDir for the current user
func homeDir() string {
	user, _ := user.Current()
	return user.HomeDir
}

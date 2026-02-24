// appdirs project doc.go

/*
This is a port of a python module used for finding what directory you 'should'
be using for saving your application data such as configuration, cache files or
other files.

The location of these directories is often hard to get right. The original python
module set out to change this into a simple API that returns you the exact
directory you need. This is a port of it to Go.

Depending on platform, this package exports a broad set of functions that
return standard system directories for application data, config, cache, logs,
state, runtime files, media folders, binaries, and application shortcuts.
It also exposes one helper struct type that combines these functions into
methods for fewer arguments in your code.

Each function defined accepts a number of arguments, each argument is optional
and can be left to the types default value if omitted. Often the function will
ignore arguments if the name given is empty.

Passing in all default values into any of the functions will return you the base
directory without any of the arguments appended to it.
*/
package appdirs

/*
UserDataDir returns the full path to the user-specific data directory.

This function uses XDG_DATA_HOME as defined by the XDG spec on *nix like systems.

Examples of return values:

	Mac OS X: ~/Library/Application Support/<AppName>
	Unix: ~/.local/share/<AppName> # or in $XDG_DATA_HOME, if defined
	Win XP (not roaming): C:\Documents and Settings\<username>\Application Data\<AppAuthor>\<AppName>
	Win XP (roaming): C:\Documents and Settings\<username>\Local Settings\Application Data\<AppAuthor>\<AppName>
	Win 7 (not roaming): C:\Users\<username>\AppData\Local\<AppAuthor>\<AppName>
	Win 7 (roaming): C:\Users\<username>\AppData\Roaming\<AppAuthor>\<AppName>
*/
func UserDataDir(name, author, version string, roaming bool) string {
	return userDataDir(name, author, version, roaming)
}

/*
UserConfigDir returns the full path to the user-specific configuration directory

This function uses XDG_CONFIG_HOME as by the XDG spec on *nix like systems.

Examples of return values:

	Mac OS X: same as UserDataDir
	Unix: ~/.config/<AppName> # or in $XDG_CONFIG_HOME, if defined
	Win *: same as UserDataDir
*/
func UserConfigDir(name, author, version string, roaming bool) string {
	return userConfigDir(name, author, version, roaming)
}

/*
UserCacheDir returns the full path to the user-specific cache directory.

The opinion argument will append 'Cache' to the base directory if set to true.

Examples of return values:

	Mac OS X: ~/Library/Caches/<AppName>
	Unix: ~/.cache/<AppName> (XDG default)
	Win XP: C:\Documents and Settings\<username>\Local Settings\Application Data\<AppAuthor>\<AppName>\Cache
	Vista: C:\Users\<username>\AppData\Local\<AppAuthor>\<AppName>\Cache
*/
func UserCacheDir(name, author, version string, opinion bool) string {
	return userCacheDir(name, author, version, opinion)
}

/*
UserLogDir returns the full path to the user-specific log directory.

The opinion argument will append either 'Logs' (windows) or 'log' (unix) to
the base directory when set to true.

Examples of return values:

	Mac OS X: ~/Library/Logs/<AppName>
	Unix: ~/.cache/<AppName>/log # or under $XDG_CACHE_HOME if defined
	Win XP: C:\Documents and Settings\<username>\Local Settings\Application Data\<AppAuthor>\<AppName>\Logs
	Vista: C:\Users\<username>\AppData\Local\<AppAuthor>\<AppName>\Logs
*/
func UserLogDir(name, author, version string, opinion bool) string {
	return userLogDir(name, author, version, opinion)
}

/*
UserStateDir returns the full path to the user-specific state directory.

Examples of return values:

	Mac OS X: same as UserDataDir
	Unix: ~/.local/state/<AppName> (or under $XDG_STATE_HOME if defined)
	Windows: same as UserDataDir
*/
func UserStateDir(name, author, version string, roaming bool) string {
	return userStateDir(name, author, version, roaming)
}

/*
UserRuntimeDir returns the full path to the user-specific runtime directory.

Examples of return values:

	Mac OS X: ~/Library/Caches/TemporaryItems/<AppName>
	Unix (Linux): /run/user/<uid>/<AppName> # or under $XDG_RUNTIME_DIR if defined
	Unix (BSD): /var/run/user/<uid>/<AppName> or /tmp/run/user/<uid>/<AppName>
	Windows: C:\Users\<username>\AppData\Local\Temp\<AppAuthor>\<AppName>
*/
func UserRuntimeDir(name, author, version string) string {
	return userRuntimeDir(name, author, version)
}

/*
UserDocumentsDir returns the path to the user's documents directory.

Examples of return values:

	Mac OS X: ~/Documents
	Unix: ~/Documents (or from XDG user-dirs configuration)
	Windows: C:\Users\<username>\Documents
*/
func UserDocumentsDir() string { return userDocumentsDir() }

/*
UserDownloadsDir returns the path to the user's downloads directory.

Examples of return values:

	Mac OS X: ~/Downloads
	Unix: ~/Downloads (or from XDG user-dirs configuration)
	Windows: C:\Users\<username>\Downloads
*/
func UserDownloadsDir() string { return userDownloadsDir() }

/*
UserPicturesDir returns the path to the user's pictures directory.

Examples of return values:

	Mac OS X: ~/Pictures
	Unix: ~/Pictures (or from XDG user-dirs configuration)
	Windows: C:\Users\<username>\Pictures
*/
func UserPicturesDir() string { return userPicturesDir() }

/*
UserVideosDir returns the path to the user's videos directory.

Examples of return values:

	Mac OS X: ~/Movies
	Unix: ~/Videos (or from XDG user-dirs configuration)
	Windows: C:\Users\<username>\Videos
*/
func UserVideosDir() string { return userVideosDir() }

/*
UserMusicDir returns the path to the user's music directory.

Examples of return values:

	Mac OS X: ~/Music
	Unix: ~/Music (or from XDG user-dirs configuration)
	Windows: C:\Users\<username>\Music
*/
func UserMusicDir() string { return userMusicDir() }

/*
UserDesktopDir returns the path to the user's desktop directory.

Examples of return values:

	Mac OS X: ~/Desktop
	Unix: ~/Desktop (or from XDG user-dirs configuration)
	Windows: C:\Users\<username>\Desktop
*/
func UserDesktopDir() string { return userDesktopDir() }

/*
UserBinDir returns the path to the user's binary directory.

Examples of return values:

	Mac OS X: ~/.local/bin
	Unix: ~/.local/bin
	Windows: C:\Users\<username>\AppData\Local\Programs
*/
func UserBinDir() string { return userBinDir() }

/*
UserApplicationsDir returns the path to the user's applications directory.

Examples of return values:

	Mac OS X: ~/Applications
	Unix: ~/.local/share/applications
	Windows: C:\Users\<username>\AppData\Roaming\Microsoft\Windows\Start Menu\Programs
*/
func UserApplicationsDir() string { return userApplicationsDir() }

/*
SiteDataDir returns the full path to the user-shared data directory.

This function uses XDG_DATA_DIRS[0] as by the XDG spec on *nix like systems.

Examples of return values:

	Mac OS X: /Library/Application Support/<AppName>
	Unix: /usr/local/share/<AppName> or /usr/share/<AppName>
	Win XP: C:\Documents and Settings\All Users\Application Data\<AppAuthor>\<AppName>
	Vista: (Fail! "C:\ProgramData" is a hidden *system* directory on Vista.)
	Win 7: C:\ProgramData\<AppAuthor>\<AppName> # Hidden, but writeable on Win 7.

WARNING: Do not use this on Windows Vista, See the note above.
*/
func SiteDataDir(name, author, version string) string {
	return siteDataDir(name, author, version)
}

/*
SiteConfigDir returns the full path to the user-shared configuration directory.

This function uses XDG_CONFIG_DIRS[0] as by the XDG spec on *nix like systems.

Examples of return values:

	Mac OS X: same as SiteDataDir
	Unix: /etc/xdg/<AppName> or $XDG_CONFIG_DIRS[i]/<AppName> for each value in $XDG_CONFIG_DIRS
	Win *: same as SiteDataDir
	Vista: (Fail! "C:\ProgramData" is a hidden *system* directory on Vista.)

WARNING: Do not use this on Windows Vista, see the note above.
*/
func SiteConfigDir(name, author, version string) string {
	return siteConfigDir(name, author, version)
}

/*
SiteCacheDir returns the full path to the user-shared cache directory.

The opinion argument will append "Cache" to the base directory if set to true on
platforms that follow that convention (e.g. Windows).

Examples of return values:

	Mac OS X: /Library/Caches/<AppName>
	Unix: /var/cache/<AppName>
	Windows: C:\ProgramData\<AppAuthor>\<AppName>\Cache
*/
func SiteCacheDir(name, author, version string, opinion bool) string {
	return siteCacheDir(name, author, version, opinion)
}

/*
SiteStateDir returns the full path to the user-shared state directory.

Examples of return values:

	Mac OS X: same as SiteDataDir
	Unix: /var/lib/<AppName>
	Windows: same as SiteDataDir
*/
func SiteStateDir(name, author, version string) string {
	return siteStateDir(name, author, version)
}

/*
SiteLogDir returns the full path to the user-shared log directory.

The opinion argument will append either 'Logs' (Windows and macOS) or 'log' (Unix)
to the base directory when set to true on platforms that follow that convention.

Examples of return values:

	Mac OS X: /Library/Logs/<AppName>
	Unix: /var/log/<AppName>
	Windows: C:\ProgramData\<AppAuthor>\<AppName>\Logs
*/
func SiteLogDir(name, author, version string, opinion bool) string {
	return siteLogDir(name, author, version, opinion)
}

/*
SiteRuntimeDir returns the full path to the shared runtime directory.

Examples of return values:

	Mac OS X: same as UserRuntimeDir
	Unix (Linux): /run/<AppName>
	Unix (BSD): /var/run/<AppName>
	Windows: same as UserRuntimeDir
*/
func SiteRuntimeDir(name, author, version string) string {
	return siteRuntimeDir(name, author, version)
}

/*
SiteBinDir returns the path to the shared binary directory.

Examples of return values:

	Mac OS X: /usr/local/bin
	Unix: /usr/local/bin
	Windows: C:\ProgramData\bin
*/
func SiteBinDir() string { return siteBinDir() }

/*
SiteApplicationsDir returns the path to the shared applications directory.

Examples of return values:

	Mac OS X: /Applications
	Unix: /usr/local/share/applications
	Windows: C:\ProgramData\Microsoft\Windows\Start Menu\Programs
*/
func SiteApplicationsDir() string { return siteApplicationsDir() }

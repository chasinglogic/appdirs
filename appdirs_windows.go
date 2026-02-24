package appdirs

import (
	"path/filepath"
	"syscall"
	"unsafe"
)

var (
	shell32, _            = syscall.LoadLibrary("shell32.dll")
	getKnownFolderPath, _ = syscall.GetProcAddress(shell32, "SHGetKnownFolderPath")

	ole32, _         = syscall.LoadLibrary("Ole32.dll")
	coTaskMemFree, _ = syscall.GetProcAddress(ole32, "CoTaskMemFree")
)

// These are KNOWNFOLDERID constants that are passed to GetKnownFolderPath
var (
	rfidLocalAppData = syscall.GUID{
		0xf1b32785,
		0x6fba,
		0x4fcf,
		[8]byte{0x9d, 0x55, 0x7b, 0x8e, 0x7f, 0x15, 0x70, 0x91},
	}
	rfidRoamingAppData = syscall.GUID{
		0x3eb685db,
		0x65f9,
		0x4cf6,
		[8]byte{0xa0, 0x3a, 0xe3, 0xef, 0x65, 0x72, 0x9f, 0x3d},
	}
	rfidProgramData = syscall.GUID{
		0x62ab5d82,
		0xfdc1,
		0x4dc3,
		[8]byte{0xa9, 0xdd, 0x07, 0x0d, 0x1d, 0x49, 0x5d, 0x97},
	}
	rfidDocuments      = syscall.GUID{0xfdd39ad0, 0x238f, 0x46af, [8]byte{0xad, 0xb4, 0x6c, 0x85, 0x48, 0x03, 0x69, 0xc7}}
	rfidDownloads      = syscall.GUID{0x374de290, 0x123f, 0x4565, [8]byte{0x91, 0x64, 0x39, 0xc4, 0x92, 0x5e, 0x46, 0x7b}}
	rfidPictures       = syscall.GUID{0x33e28130, 0x4e1e, 0x4676, [8]byte{0x83, 0x5a, 0x98, 0x39, 0x5c, 0x3b, 0xc3, 0xbb}}
	rfidVideos         = syscall.GUID{0x18989b1d, 0x99b5, 0x455b, [8]byte{0x84, 0x1c, 0xab, 0x7c, 0x74, 0xe4, 0xdd, 0xfc}}
	rfidMusic          = syscall.GUID{0x4bd8d571, 0x6d19, 0x48d3, [8]byte{0xbe, 0x97, 0x42, 0x22, 0x20, 0x08, 0x0e, 0x43}}
	rfidDesktop        = syscall.GUID{0xb4bfcc3a, 0xdb2c, 0x424c, [8]byte{0xb0, 0x29, 0x7f, 0xe9, 0x9a, 0x87, 0xc6, 0x41}}
	rfidPrograms       = syscall.GUID{0xa77f5d77, 0x2e2b, 0x44c3, [8]byte{0xa6, 0xa2, 0xab, 0xa6, 0x01, 0x05, 0x4a, 0x51}}
	rfidCommonPrograms = syscall.GUID{0x0139d44e, 0x6afe, 0x49f2, [8]byte{0x86, 0x90, 0x3d, 0xaf, 0xca, 0xe6, 0xff, 0xb8}}
)

// User path functions
func userDataDir(name, author, version string, roaming bool) (path string) {
	if author == "" {
		author = name
	}
	var rfid syscall.GUID
	if roaming {
		rfid = rfidRoamingAppData
	} else {
		rfid = rfidLocalAppData
	}
	path, err := getFolderPath(rfid)
	if err != nil {
		return ""
	}
	if path, err = filepath.Abs(path); err != nil {
		return ""
	}
	if name != "" {
		path = filepath.Join(path, author, name)
	}
	if name != "" && version != "" {
		path = filepath.Join(path, version)
	}
	return path
}

func userConfigDir(name, author, version string, roaming bool) string {
	return UserDataDir(name, author, version, roaming)
}

func userCacheDir(name, author, version string, opinion bool) (path string) {
	if author == "" {
		author = name
	}
	path, err := getFolderPath(rfidLocalAppData)
	if err != nil {
		return ""
	}
	if path, err = filepath.Abs(path); err != nil {
		return ""
	}
	if name != "" {
		path = filepath.Join(path, author, name)
		if opinion {
			path = filepath.Join(path, "Cache")
		}
	}
	if name != "" && version != "" {
		path = filepath.Join(path, version)
	}
	return path
}

func userLogDir(name, author, version string, opinion bool) (path string) {
	path = UserDataDir(name, author, version, false)
	if opinion {
		path = filepath.Join(path, "Logs")
	}
	return path
}

func userStateDir(name, author, version string, roaming bool) string {
	return UserDataDir(name, author, version, roaming)
}

func userRuntimeDir(name, author, version string) (path string) {
	base, err := getFolderPath(rfidLocalAppData)
	if err != nil {
		return ""
	}
	if base, err = filepath.Abs(base); err != nil {
		return ""
	}
	path = filepath.Join(base, "Temp")
	if author == "" {
		author = name
	}
	if name != "" {
		path = filepath.Join(path, author, name)
	}
	if name != "" && version != "" {
		path = filepath.Join(path, version)
	}
	return path
}

func userDocumentsDir() string { return knownFolderPath(rfidDocuments) }
func userDownloadsDir() string { return knownFolderPath(rfidDownloads) }
func userPicturesDir() string  { return knownFolderPath(rfidPictures) }
func userVideosDir() string    { return knownFolderPath(rfidVideos) }
func userMusicDir() string     { return knownFolderPath(rfidMusic) }
func userDesktopDir() string   { return knownFolderPath(rfidDesktop) }
func userBinDir() string {
	path, err := getFolderPath(rfidLocalAppData)
	if err != nil {
		return ""
	}
	if path, err = filepath.Abs(path); err != nil {
		return ""
	}
	return filepath.Join(path, "Programs")
}
func userApplicationsDir() string { return knownFolderPath(rfidPrograms) }

// Site path functions
func siteDataDir(name, author, version string) (path string) {
	path, err := getFolderPath(rfidProgramData)
	if err != nil {
		return ""
	}
	if path, err = filepath.Abs(path); err != nil {
		return ""
	}
	if author == "" {
		author = name
	}
	if name != "" {
		path = filepath.Join(path, author, name)
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
	path, err := getFolderPath(rfidProgramData)
	if err != nil {
		return ""
	}
	if path, err = filepath.Abs(path); err != nil {
		return ""
	}
	if author == "" {
		author = name
	}
	if name != "" {
		path = filepath.Join(path, author, name)
		if opinion {
			path = filepath.Join(path, "Cache")
		}
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
	path, err := getFolderPath(rfidProgramData)
	if err != nil {
		return ""
	}
	if path, err = filepath.Abs(path); err != nil {
		return ""
	}
	if author == "" {
		author = name
	}
	if name != "" {
		path = filepath.Join(path, author, name)
	}
	if name != "" && version != "" {
		path = filepath.Join(path, version)
	}
	if name != "" && opinion {
		path = filepath.Join(path, "Logs")
	}
	return path
}

func siteRuntimeDir(name, author, version string) string {
	return userRuntimeDir(name, author, version)
}

func siteBinDir() string {
	path, err := getFolderPath(rfidProgramData)
	if err != nil {
		return ""
	}
	if path, err = filepath.Abs(path); err != nil {
		return ""
	}
	return filepath.Join(path, "bin")
}

func siteApplicationsDir() string { return knownFolderPath(rfidCommonPrograms) }

// Helper and internal functions remain unchanged
func knownFolderPath(rfid syscall.GUID) string {
	path, err := getFolderPath(rfid)
	if err != nil {
		return ""
	}
	if abs, err := filepath.Abs(path); err == nil {
		return abs
	}
	return path
}

func getFolderPath(rfid syscall.GUID) (string, error) {
	var res uintptr
	ret, _, callErr := syscall.Syscall6(
		uintptr(getKnownFolderPath),
		4,
		uintptr(unsafe.Pointer(&rfid)),
		0,
		0,
		uintptr(unsafe.Pointer(&res)),
		0,
		0,
	)
	if callErr != 0 && ret != 0 {
		return "", callErr
	}
	defer syscall.Syscall(uintptr(coTaskMemFree), 1, res, 0, 0)
	return ucs2PtrToString(res), nil
}

func ucs2PtrToString(p uintptr) string {
	ptr := (*[4096]uint16)(unsafe.Pointer(p))
	return syscall.UTF16ToString((*ptr)[:])
}

package appdirs

import (
	"fmt"
	"os/user"
	"runtime"
	"strings"
	"testing"
)

func getExpectations() map[string]string {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}
	var userName string
	if runtime.GOOS == "windows" {
		userName = extractDomain(u.Username)
	} else {
		userName = u.Username
	}
	uid := u.Uid
	expectations := map[string]map[string]string{
		"darwin": {
			"UserData":         fmt.Sprintf("/Users/%s/Library/Application Support/Testing/1.0.0", userName),
			"UserConfig":       fmt.Sprintf("/Users/%s/Library/Application Support/Testing/1.0.0", userName),
			"UserCache":        fmt.Sprintf("/Users/%s/Library/Caches/Testing/1.0.0", userName),
			"UserLog":          fmt.Sprintf("/Users/%s/Library/Logs/Testing/1.0.0", userName),
			"UserState":        fmt.Sprintf("/Users/%s/Library/Application Support/Testing/1.0.0", userName),
			"UserRuntime":      fmt.Sprintf("/Users/%s/Library/Caches/TemporaryItems/Testing/1.0.0", userName),
			"UserDocuments":    fmt.Sprintf("/Users/%s/Documents", userName),
			"UserDownloads":    fmt.Sprintf("/Users/%s/Downloads", userName),
			"UserPictures":     fmt.Sprintf("/Users/%s/Pictures", userName),
			"UserVideos":       fmt.Sprintf("/Users/%s/Movies", userName),
			"UserMusic":        fmt.Sprintf("/Users/%s/Music", userName),
			"UserDesktop":      fmt.Sprintf("/Users/%s/Desktop", userName),
			"UserBin":          fmt.Sprintf("/Users/%s/.local/bin", userName),
			"UserApplications": fmt.Sprintf("/Users/%s/Applications", userName),
			"SiteData":         "/Library/Application Support/Testing/1.0.0",
			"SiteConfig":       "/Library/Application Support/Testing/1.0.0",
			"SiteCache":        "/Library/Caches/Testing/1.0.0",
			"SiteState":        "/Library/Application Support/Testing/1.0.0",
			"SiteLog":          "/Library/Logs/Testing/1.0.0",
			"SiteRuntime":      fmt.Sprintf("/Users/%s/Library/Caches/TemporaryItems/Testing/1.0.0", userName),
			"SiteBin":          "/usr/local/bin",
			"SiteApplications": "/Applications",
		},
		"linux": {
			"UserData":         fmt.Sprintf("/home/%s/.local/share/Testing/1.0.0", userName),
			"UserConfig":       fmt.Sprintf("/home/%s/.config/Testing/1.0.0", userName),
			"UserCache":        fmt.Sprintf("/home/%s/.cache/Testing/1.0.0", userName),
			"UserLog":          fmt.Sprintf("/home/%s/.cache/Testing/1.0.0/log", userName),
			"UserState":        fmt.Sprintf("/home/%s/.local/state/Testing/1.0.0", userName),
			"UserRuntime":      fmt.Sprintf("/run/user/%s/Testing/1.0.0", uid),
			"UserDocuments":    fmt.Sprintf("/home/%s/Documents", userName),
			"UserDownloads":    fmt.Sprintf("/home/%s/Downloads", userName),
			"UserPictures":     fmt.Sprintf("/home/%s/Pictures", userName),
			"UserVideos":       fmt.Sprintf("/home/%s/Videos", userName),
			"UserMusic":        fmt.Sprintf("/home/%s/Music", userName),
			"UserDesktop":      fmt.Sprintf("/home/%s/Desktop", userName),
			"UserBin":          fmt.Sprintf("/home/%s/.local/bin", userName),
			"UserApplications": fmt.Sprintf("/home/%s/.local/share/applications", userName),
			"SiteData":         "/usr/local/share/Testing/1.0.0",
			"SiteConfig":       "/etc/xdg/Testing/1.0.0",
			"SiteCache":        "/var/cache/Testing/1.0.0",
			"SiteState":        "/var/lib/Testing/1.0.0",
			"SiteLog":          "/var/log/Testing/1.0.0",
			"SiteRuntime":      "/run/Testing/1.0.0",
			"SiteBin":          "/usr/local/bin",
			"SiteApplications": "/usr/local/share/applications",
		},
		"windows": {
			"UserData":         fmt.Sprintf("C:\\Users\\%s\\AppData\\Local\\Tester\\Testing\\1.0.0", userName),
			"UserConfig":       fmt.Sprintf("C:\\Users\\%s\\AppData\\Local\\Tester\\Testing\\1.0.0", userName),
			"UserCache":        fmt.Sprintf("C:\\Users\\%s\\AppData\\Local\\Tester\\Testing\\Cache\\1.0.0", userName),
			"UserLog":          fmt.Sprintf("C:\\Users\\%s\\AppData\\Local\\Tester\\Testing\\1.0.0\\Logs", userName),
			"UserState":        fmt.Sprintf("C:\\Users\\%s\\AppData\\Local\\Tester\\Testing\\1.0.0", userName),
			"UserRuntime":      fmt.Sprintf("C:\\Users\\%s\\AppData\\Local\\Temp\\Tester\\Testing\\1.0.0", userName),
			"UserDocuments":    fmt.Sprintf("C:\\Users\\%s\\Documents", userName),
			"UserDownloads":    fmt.Sprintf("C:\\Users\\%s\\Downloads", userName),
			"UserPictures":     fmt.Sprintf("C:\\Users\\%s\\Pictures", userName),
			"UserVideos":       fmt.Sprintf("C:\\Users\\%s\\Videos", userName),
			"UserMusic":        fmt.Sprintf("C:\\Users\\%s\\Music", userName),
			"UserDesktop":      fmt.Sprintf("C:\\Users\\%s\\Desktop", userName),
			"UserBin":          fmt.Sprintf("C:\\Users\\%s\\AppData\\Local\\Programs", userName),
			"UserApplications": fmt.Sprintf("C:\\Users\\%s\\AppData\\Roaming\\Microsoft\\Windows\\Start Menu\\Programs", userName),
			"SiteData":         "C:\\ProgramData\\Tester\\Testing\\1.0.0",
			"SiteConfig":       "C:\\ProgramData\\Tester\\Testing\\1.0.0",
			"SiteCache":        "C:\\ProgramData\\Tester\\Testing\\Cache\\1.0.0",
			"SiteState":        "C:\\ProgramData\\Tester\\Testing\\1.0.0",
			"SiteLog":          "C:\\ProgramData\\Tester\\Testing\\1.0.0\\Logs",
			"SiteRuntime":      fmt.Sprintf("C:\\Users\\%s\\AppData\\Local\\Temp\\Tester\\Testing\\1.0.0", userName),
			"SiteBin":          "C:\\ProgramData\\bin",
			"SiteApplications": "C:\\ProgramData\\Microsoft\\Windows\\Start Menu\\Programs",
		},
	}
	return expectations[runtime.GOOS]
}

// extractDomain extract an optional domain from a Windows username
// e.g. "domain\\user" => "user"
func extractDomain(s string) string {
	p := strings.Split(s, "\\")
	if len(p) > 1 {
		return p[1]
	}
	return p[0]
}

func testApp() *App {
	app := New("Testing")
	app.Author = "Tester"
	app.Version = "1.0.0"
	return app
}

func TestUserData(t *testing.T) {
	app := testApp()
	expected := getExpectations()

	if r := app.UserData(); r != expected["UserData"] {
		t.Fatalf("Expected %s for UserData got: %s", expected["UserData"], r)
	}
}

func TestSiteData(t *testing.T) {
	app := testApp()
	expected := getExpectations()

	if r := app.SiteData(); r != expected["SiteData"] {
		t.Fatalf("Expected %s for SiteData got: %s", expected["SiteData"], r)
	}
}

func TestSiteConfig(t *testing.T) {
	app := testApp()
	expected := getExpectations()

	if r := app.SiteConfig(); r != expected["SiteConfig"] {
		t.Fatalf("Expected %s for SiteConfig got: %s", expected["SiteConfig"], r)
	}
}

func TestUserCache(t *testing.T) {
	app := testApp()
	expected := getExpectations()

	if r := app.UserCache(); r != expected["UserCache"] {
		t.Fatalf("Expected %s for UserCache got: %s", expected["UserCache"], r)
	}
}

func TestUserConfig(t *testing.T) {
	app := testApp()
	expected := getExpectations()

	if r := app.UserConfig(); r != expected["UserConfig"] {
		t.Fatalf("Expected %s for UserConfig got: %s", expected["UserConfig"], r)
	}
}

func TestUserLog(t *testing.T) {
	app := testApp()
	expected := getExpectations()

	if r := app.UserLog(); r != expected["UserLog"] {
		t.Fatalf("Expected %s for UserLog got: %s", expected["UserLog"], r)
	}
}

func TestSiteCache(t *testing.T) {
	app := testApp()
	expected := getExpectations()

	if r := app.SiteCache(); r != expected["SiteCache"] {
		t.Fatalf("Expected %s for SiteCache got: %s", expected["SiteCache"], r)
	}
}

func TestUserState(t *testing.T) {
	app := testApp()
	expected := getExpectations()

	if r := app.UserState(); r != expected["UserState"] {
		t.Fatalf("Expected %s for UserState got: %s", expected["UserState"], r)
	}
}

func TestSiteState(t *testing.T) {
	app := testApp()
	expected := getExpectations()

	if r := app.SiteState(); r != expected["SiteState"] {
		t.Fatalf("Expected %s for SiteState got: %s", expected["SiteState"], r)
	}
}

func TestSiteLog(t *testing.T) {
	app := testApp()
	expected := getExpectations()

	if r := app.SiteLog(); r != expected["SiteLog"] {
		t.Fatalf("Expected %s for SiteLog got: %s", expected["SiteLog"], r)
	}
}

func TestUserRuntime(t *testing.T) {
	app := testApp()
	expected := getExpectations()

	if r := app.UserRuntime(); r != expected["UserRuntime"] {
		t.Fatalf("Expected %s for UserRuntime got: %s", expected["UserRuntime"], r)
	}
}

func TestSiteRuntime(t *testing.T) {
	app := testApp()
	expected := getExpectations()

	if r := app.SiteRuntime(); r != expected["SiteRuntime"] {
		t.Fatalf("Expected %s for SiteRuntime got: %s", expected["SiteRuntime"], r)
	}
}

func TestUserDocuments(t *testing.T) {
	app := testApp()
	expected := getExpectations()

	if r := app.UserDocuments(); r != expected["UserDocuments"] {
		t.Fatalf("Expected %s for UserDocuments got: %s", expected["UserDocuments"], r)
	}
}

func TestUserDownloads(t *testing.T) {
	app := testApp()
	expected := getExpectations()

	if r := app.UserDownloads(); r != expected["UserDownloads"] {
		t.Fatalf("Expected %s for UserDownloads got: %s", expected["UserDownloads"], r)
	}
}

func TestUserPictures(t *testing.T) {
	app := testApp()
	expected := getExpectations()

	if r := app.UserPictures(); r != expected["UserPictures"] {
		t.Fatalf("Expected %s for UserPictures got: %s", expected["UserPictures"], r)
	}
}

func TestUserVideos(t *testing.T) {
	app := testApp()
	expected := getExpectations()

	if r := app.UserVideos(); r != expected["UserVideos"] {
		t.Fatalf("Expected %s for UserVideos got: %s", expected["UserVideos"], r)
	}
}

func TestUserMusic(t *testing.T) {
	app := testApp()
	expected := getExpectations()

	if r := app.UserMusic(); r != expected["UserMusic"] {
		t.Fatalf("Expected %s for UserMusic got: %s", expected["UserMusic"], r)
	}
}

func TestUserDesktop(t *testing.T) {
	app := testApp()
	expected := getExpectations()

	if r := app.UserDesktop(); r != expected["UserDesktop"] {
		t.Fatalf("Expected %s for UserDesktop got: %s", expected["UserDesktop"], r)
	}
}

func TestUserBin(t *testing.T) {
	app := testApp()
	expected := getExpectations()

	if r := app.UserBin(); r != expected["UserBin"] {
		t.Fatalf("Expected %s for UserBin got: %s", expected["UserBin"], r)
	}
}

func TestSiteBin(t *testing.T) {
	app := testApp()
	expected := getExpectations()

	if r := app.SiteBin(); r != expected["SiteBin"] {
		t.Fatalf("Expected %s for SiteBin got: %s", expected["SiteBin"], r)
	}
}

func TestUserApplications(t *testing.T) {
	app := testApp()
	expected := getExpectations()

	if r := app.UserApplications(); r != expected["UserApplications"] {
		t.Fatalf("Expected %s for UserApplications got: %s", expected["UserApplications"], r)
	}
}

func TestSiteApplications(t *testing.T) {
	app := testApp()
	expected := getExpectations()

	if r := app.SiteApplications(); r != expected["SiteApplications"] {
		t.Fatalf("Expected %s for SiteApplications got: %s", expected["SiteApplications"], r)
	}
}

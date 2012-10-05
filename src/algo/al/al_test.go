package al

import "testing"
import "runtime"

// import "fmt"

// some parameters
const expected_version = 83887873

func TestGetAllegroVersion(t *testing.T) {
	version := GetAllegroVersion()
	if version != expected_version {
		t.Errorf("unexpected version of Allegro: %d in stead of %d!",
			version, expected_version)
	}
}

// Test system installation and deinstallation
func TestSystemInstall(t *testing.T) {
	if IsSystemInstalled() {
		t.Errorf("System should not be installed before install\n")
		return
	}
	InstallSystem()
	if !IsSystemInstalled() {
		t.Errorf("System should be installed after install\n")
		return
	}
	UninstallSystem()
	if IsSystemInstalled() {
		t.Errorf("System should not be installed after uninstall\n")
		return
	}
}

// Test USTR
func TestUSTR(t *testing.T) {
	s1 := "Hello no unicode!"
	s2 := "Hello µ unicode!"
	u1 := USTRV(s1)
	u2 := USTRV(s2)
	r1 := u1.String()
	r2 := u2.String()
	if s1 != r1 {
		t.Errorf("USTR roundtrip failed: %s->%s", s1, r1)
	}
	if s2 != r2 {
		t.Errorf("USTR roundtrip failed: %s->%s", s2, r2)
	}
	u1.Free()
	u1.Free()
	u1.Free()
	if u1.String() != "<destroyed>" {
		t.Error("USTR.String() should return <destroyed> after Free()")
	}

}

// Test timer functions 
func TestGetTimeRest(t *testing.T) {
	InstallSystem()
	defer UninstallSystem()
	rest := 0.123
	t1 := GetTime()
	Rest(rest)
	t2 := GetTime()
	del := t2 - t1 - rest
	if (del > 0.001) || (del < -0.001) {
		t.Errorf("Rest/GetTime for %f not precise %f %f %f", rest, t1, t2, del)
	}
}

// Test path functions
func TestPath(t *testing.T) {
	InstallSystem()
	defer UninstallSystem()
	path := GetStandardPath(TEMP_PATH)
	str := path.String()
	tmp := "/tmp/"
	// special case for windows...
	if runtime.GOOS == "windows" {
		tmp = `C:\TMP\`
	}
	if str != tmp {
		t.Errorf("GetStandardPath(TEMP_PATH) is not %s but %s", tmp, str)
	}
}

// Test screen saver inhibition.
func TestInhibitScreensaver(t *testing.T) {
	InstallSystem()
	defer UninstallSystem()
	ok := InhibitScreensaver(true)
	if !ok {
		t.Errorf("InhibitScreensaver failed: %v", ok)
	}
}

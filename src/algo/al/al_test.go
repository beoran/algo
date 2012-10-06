package al

import "testing"
import "runtime"

// import "fmt"

// some parameters
const expected_version = 83887873
const SCREEN_W         = 640 
const SCREEN_H         = 480
const TEST_FULLSCREEN  = false

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



// Test joystick functions, works better with a joystick plugged in ;)
func TestJoystick(t *testing.T) {
  InstallSystem()
  defer UninstallSystem()
  InstallJoystick()
  defer UninstallJoystick()
  num := GetNumJoysticks()
  t.Logf("Found %d joysticks\n", num)
  for index :=0 ; index < num ; index ++ {
    js      := GetJoystick(index)
    jsname  := js.GetName()
    sticks  := js.GetNumSticks()
    buttons := js.GetNumButtons()
    t.Logf("Joystick %s (nr %d) has %d sticks and %d buttons:\n", 
            jsname, index, sticks, buttons)
    for sdex  := 0 ; sdex < sticks ; sdex++ {
      axes    := js.GetNumAxes(sdex)
      sname   := js.GetStickName(sdex)
      sfname  := js.GetStickFlagsName(sdex)
      t.Logf("Stick %s (nr %d, %s) has %d axes: ", sname, sdex, sfname, axes)
      for adex :=0 ; adex < axes ; adex++ {
        aname  := js.GetAxisName(sdex, adex)
        t.Logf("%s (nr %d) ",aname, adex)
      }
    }
    t.Logf("\nButtons :")
    for bdex := 0 ; bdex < buttons; bdex++ {
      bname := js.GetButtonName(bdex);
        t.Logf("%s (nr %d) ",bname, bdex)
    }
    t.Logf("\n")
  }
}

// Makesa display for testing, using the test's setting above 
func makeDisplay() (*Display) {
  flags := 0
  // Use full screen mode if needed.
  if TEST_FULLSCREEN  { 
    flags = FULLSCREEN | GENERATE_EXPOSE_EVENTS;
  } 
  SetNewDisplayFlags(flags)
  // Create a window to display things on: 640x480 pixels.
  display := CreateDisplay(SCREEN_W, SCREEN_H)  
  return display
}

// Test basic display functions 
func TestBasicDisplay(t *testing.T) {
  InstallSystem()
  defer UninstallSystem()  
  display := makeDisplay()
  if display == nil {
    t.Error("Error creating display.")
  }
  /*
  if !display.Resize(SCREEN_W, SCREEN_H) {
    t.Error("Resize of display failed.")
  }
  */
  blue   := CreateColor(0.0, 0.0, 1.0, 1.0)
  yellow := CreateColor(1.0, 1.0, 0.0, 1.0)
  ClearToColor(blue)
  DrawPixel(20.0, 10.0, yellow)  
  FlipDisplay()
  ClearToColor(yellow)
  DrawPixel(20.0, 10.0, blue)
  FlipDisplay()
  display.Destroy() 
}

// Benchmark basic display function ClearToColor
func BenchmarkClearToColor(b *testing.B) {
    b.StopTimer()
    InstallSystem()
    defer UninstallSystem()  
    display := makeDisplay()
    blue    := CreateColor(0.0, 0.0, 1.0, 1.0)
    if display == nil {
      b.Fatal("Error creating display. Cannot benchmark it.")
    }
    b.StartTimer()
    for i := 0; i < b.N; i++ {
        ClearToColor(blue)
        FlipDisplay()
    }    
    display.Destroy()
}

// Benchmark basic display function FlipDisplay
func BenchmarkFlipDisplay(b *testing.B) {
    b.StopTimer()
    InstallSystem()
    defer UninstallSystem()  
    display := makeDisplay()
    if display == nil {
      b.Fatal("Error creating display. Cannot benchmark it.")
    }
    b.StartTimer()
    for i := 0; i < b.N; i++ {
        FlipDisplay()
    }
    display.Destroy()
}




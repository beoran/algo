package al

import "testing"
import "runtime"
import "flag"
import "path/filepath"
import "math/rand"

// some parameters
const expected_version = 84017665 // 83952897
const SCREEN_W = 640
const SCREEN_H = 480
const TEST_FULLSCREEN = true

var fullscreen = flag.Bool("fullscreen", false, "Run fullscreen or not")

func TestGetAllegroVersion(t *testing.T) {
    version := GetAllegroVersion()
    if version != expected_version {
        t.Errorf("unexpected version of Allegro: %d in stead of %d!",
            version, expected_version)
    }
}

const CALLBACK_RESULT = 77
const BMP_W = 11
const BMP_H = 23


// Helper that loads bitmaps from the testdata folder

func loadBitmap(t *testing.T, name string) * Bitmap {
    path := filepath.Join("testdata", name) // relative path
    bmp := LoadBitmap(path)
    if bmp == nil {
        t.Fatalf("Could not load bitmap: %s: %d\n", path, Errno())
    }
    return bmp
}

// Test C callbacks, for example create bitmap  

// Test system installation and deinstallation
func TestSystemInstall(t *testing.T) {
    if IsSystemInstalled() {
        t.Errorf("System should not be installed before install\n")
        return
    }
    if !InstallSystem() {
        t.Errorf("System should be installed now\n")
        return
    }
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

// Test display info 
func TestGetInfo(t *testing.T) {
    InstallSystem()
    defer UninstallSystem()
    nv := NumVideoAdapters()
    if nv < 1 {
        t.Error("No video adapters found!")
    }
    for index := 0; index < nv; index++ {
        info := GetMonitorInfo(index)
        if info == nil {
            t.Errorf("Video adapter %d not found!", index)
            continue
        }
        t.Logf("MonitorInfo for %d: %d %d %d %d\n", index,
            info.X1(), info.Y1(), info.X2(), info.Y2())
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
    for index := 0; index < num; index++ {
        js := GetJoystick(index)
        jsname := js.GetName()
        sticks := js.GetNumSticks()
        buttons := js.GetNumButtons()
        t.Logf("Joystick %s (nr %d) has %d sticks and %d buttons:\n",
            jsname, index, sticks, buttons)
        for sdex := 0; sdex < sticks; sdex++ {
            axes := js.GetNumAxes(sdex)
            sname := js.GetStickName(sdex)
            sfname := js.GetStickFlagsName(sdex)
            t.Logf("Stick %s (nr %d, %s) has %d axes: ", sname, sdex, sfname, axes)
            for adex := 0; adex < axes; adex++ {
                aname := js.GetAxisName(sdex, adex)
                t.Logf("%s (nr %d) ", aname, adex)
            }
        }
        t.Logf("\nButtons :")
        for bdex := 0; bdex < buttons; bdex++ {
            bname := js.GetButtonName(bdex)
            t.Logf("%s (nr %d) ", bname, bdex)
        }
        t.Logf("\n")
    }
}

// Makesa display for testing, using the test's setting above 
func makeDisplay() *Display {
    flags := 0
    // Use full screen mode if needed.
    if *fullscreen {
        flags = FULLSCREEN // | GENERATE_EXPOSE_EVENTS
    } else {
        SetNewDisplayFlags(flags)
    }
    // Create a window to display things on: 640x480 pixels.
    display := CreateDisplay(SCREEN_W, SCREEN_H)
    display.Resize(SCREEN_W, SCREEN_H)
    if !(*fullscreen) {
        display.SetTitle("Algo test window")
    }
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
    HoldBitmapDrawing(true)
    if !IsBitmapDrawingHeld() {
        t.Error("Bitmap drawing hold failed")
    }
    HoldBitmapDrawing(false)
    if IsBitmapDrawingHeld() {
        t.Error("Bitmap drawing hold release failed")
    }

    /*
       if ! {
         t.Error("Resize of display failed.")
       }
    */
    blue := CreateColor(0.0, 0.0, 1.0, 1.0)
    yellow := CreateColor(1.0, 1.0, 0.0, 1.0)
    ClearToColor(blue)
    DrawPixel(20.0, 10.0, yellow)
    FlipDisplay()
    Rest(1.0)
    display.SetWindowPosition(50, 100)
    ClearToColor(yellow)
    DrawPixel(20.0, 10.0, blue)
    FlipDisplay()
    display.Destroy()
    Rest(1.0)
}

// Test some font functions 
func TestFonts(t *testing.T) {
    InstallSystem()
    defer UninstallSystem()
    InitFontAddon()
    defer ShutdownFontAddon()
    
    display := makeDisplay()
    if display == nil {
        t.Error("Error creating display.")
    }
  
    
    font := CreateBuiltinFont()
    if font == nil {
        t.Error("Cannot create built in font.")
    }
    ranges , count := font.Ranges()
    t.Logf("Built in font has ranges: %v, %d\n", ranges, count);
    blue := CreateColor(0.0, 0.0, 1.0, 1.0)
    yellow := CreateColor(1.0, 1.0, 0.0, 1.0)
    ClearToColor(blue)
    font.DrawMultilineTextf(yellow, 20, 30, 100, 10, 0, "This is a rather long text that should flow over multiple lines, it also has placeholders like this one: %d, and this should all work fine.", 7)
    FlipDisplay()
    Rest(1.0)
    font.Destroy()
    display.Destroy()
}


func randomColor() Color {
    return CreateColor(rand.Float32(), rand.Float32(), rand.Float32(), 1.0)
}

// Test some bitmap functions 
func TestBitmaps(t *testing.T) {
    InstallSystem()
    defer UninstallSystem()
    InitImageAddon()
    defer ShutdownImageAddon()
    
    display := makeDisplay()
    if display == nil {
        t.Error("Error creating display.")
    }
    
    bmp := loadBitmap(t, "gin_feather.png")
    blue := CreateColor(0.0, 0.0, 1.0, 1.0)
    yellow := CreateColor(1.0, 1.0, 0.0, 1.0)
    ClearToColor(blue)
    bmp.DrawTinted(yellow, 20, 30, 0)
    FlipDisplay()
    tb := TargetBitmap()
    if (!tb.Save(filepath.Join("testdata", "TestBitmaps.out.png"))) {
        t.Errorf("Could not save output file.")
    }
    Rest(1.0)
    display.Destroy()
}

// Test some primitive functions
func TestPrimitives(t *testing.T) {
    InstallSystem()
    defer UninstallSystem()
    InitPrimitivesAddon()
    defer ShutdownPrimitivesAddon()
    
    display := makeDisplay()
    if display == nil {
        t.Error("Error creating display.")
    }

    blue := CreateColor(0.0, 0.0, 1.0, 1.0)
    yellow := CreateColor(1.0, 1.0, 0.0, 1.0)
    ClearToColor(blue)
    DrawLine(0,0,640,480, yellow, 1)   
    DrawCircle(50,50,10, randomColor(), 3.0)
    DrawEllipse(70,70,20,30, randomColor(), 5.0)
    DrawFilledRoundedRectangle(120, 120, 300, 300, 7, 5, randomColor())
    FlipDisplay()
    Rest(1.0)
    display.Destroy()
}

// Test some bitmap functions 
func TestPrimitives2(t *testing.T) {
    InstallSystem()
    defer UninstallSystem()
    InitImageAddon()
    defer ShutdownImageAddon()
    InitPrimitivesAddon()
    defer ShutdownPrimitivesAddon()

    
    display := makeDisplay()
    if display == nil {
        t.Error("Error creating display.")
    }
    
    bmp := loadBitmap(t, "gin_feather.png")
    blue := CreateColor(0.0, 0.0, 1.0, 1.0)
    yellow := CreateColor(1.0, 1.0, 0.0, 1.0)
    ClearToColor(blue)
    v := make([]Vertex, 3)
    v[0].x = 10
    v[0].y = 20
    v[0].color = yellow.toC()
    v[0].u = 0.0
    v[0].v = 0.0
    
    v[1].x = 110
    v[1].y = 120
    v[0].u = 1.0
    v[0].v = 0.0

    v[2].x = 120
    v[2].y = 210
    v[0].u = 0.0
    v[0].v = 1.0
    
    DrawPrim(v, bmp, 0, 3, PRIM_TRIANGLE_LIST)
    FlipDisplay()
    Rest(1.0)
    display.Destroy()
}


// Benchmark basic display function ClearToColor
func BenchmarkClearToColor(b *testing.B) {
    b.StopTimer()
    InstallSystem()
    defer UninstallSystem()
    display := makeDisplay()
    blue := CreateColor(0.0, 0.0, 1.0, 1.0)
    if display == nil {
        b.Fatal("Error creating display. Cannot benchmark it.")
    }
    b.StartTimer()
    for i := 0; i < b.N; i++ {
        ClearToColor(blue)
    }
    // FlipDisplay()
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

// Benchmarking of C call overhead
func BenchmarkDoNothing(b *testing.B) {
    for i := 0; i < b.N; i++ {
        DoNothing()
    }
}

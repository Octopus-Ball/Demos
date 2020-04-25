import (
    "log"
    "os"
    "path/filepath"
    "runtime/pprof"
)
// 进行CPU监控
func CreateProfileFile() {
    dir, err := os.Getwd()
    if err != nil {
        log.Fatalln("get current directory failed.", err)
    }
    
    fileName := filepath.Join(dir, "pprof", "profile_file", "profile_file")
    f, _ := os.Create(fileName)
    // start to record CPU profile and write to file `f`
    _ = pprof.StartCPUProfile(f)
    // stop to record CPU profile
    defer pprof.StopCPUProfile()
    // TODO do something
}
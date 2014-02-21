package main

import (
    "os"
    "io/ioutil"
    "strings"
    "bytes"
    "fmt"
    "path"
    "path/filepath"
)

const (
    dirSep = "/"
    configFile = ".hgenconfig"
)

func fileExists(dpath string) bool {
    info, err := os.Stat(dpath)
    if err != nil {
        return false
    }
    return !info.IsDir()
}

func buildName(dstPath string) (string, error) {
    fname := path.Base(dstPath)
    curPath := path.Dir(dstPath)
    parts := []string{}

    for !fileExists(path.Join(curPath, configFile)) {
        parts = append([]string{path.Base(curPath)}, parts...)
        newCurPath := path.Dir(curPath)

        if newCurPath == curPath {
            fail("Couldn't find %s. Did you forget it?", configFile)
        }

        curPath = newCurPath
    }

    parts = append(parts, fname)
    return path.Join(parts...), nil
}

func guardFromPath(fpath string) string {
    uscores := strings.Replace(fpath, dirSep, "_", -1)
    nodots := strings.Replace(uscores, ".", "_", -1)
    caps := strings.ToUpper(nodots)
    return "__" + caps + "_"
}

func genGuard(fpath string) string {
    b := &bytes.Buffer{}
    fname := path.Base(fpath)
    guardName := guardFromPath(fpath)

    fmt.Fprintln(b)
    fmt.Fprintf(b, "/** @file %s\n", fname)
    fmt.Fprintf(b, " *  @brief Function prototypes for foo.\n")
    fmt.Fprintf(b, " */\n")
    fmt.Fprintln(b)

    fmt.Fprintf(b, "#ifndef %s\n", guardName)
    fmt.Fprintf(b, "#define %s\n", guardName)

    fmt.Fprintln(b)
    fmt.Fprintln(b)

    fmt.Fprintf(b, "#endif\n");

    return string(b.Bytes())
}

func fail(format string, params ...interface{}) {
    fmt.Printf("FAIL: " + format + "\n", params...)
    os.Exit(-1)
}

func usage(format string, params ...interface{}) {
    fmt.Printf(format + "\n", params...)
    fmt.Printf("usage: %s [headerfile]\n", os.Args[0])
    os.Exit(0)
}

func main() {
    if len(os.Args) != 2 {
        usage("You must supply a filename to create")
    }

    dstPath := os.Args[1]

    if _, err := os.Stat(dstPath); err == nil {
        fail("file %s already exists", dstPath)
    }

    absPath, err := filepath.Abs(dstPath)
    if err != nil {
        fail("Couldn't resolve absolute path of %s", absPath)
    }

    relPath, err := buildName(absPath)
    if err != nil {
        fail("Error buidling the relative path: %v", err)
    }

    sGuard := genGuard(relPath)

    if err := ioutil.WriteFile(dstPath, []byte(sGuard), 0644); err != nil {
        fail("Error writing the file: %v", err)
    }
}

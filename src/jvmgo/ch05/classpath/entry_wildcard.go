package classpath

import "os"
import "path/filepath"
import (
	"strings"
	"fmt"
)

func newWildcardEntry(path string) CompositeEntry {
	fmt.Printf("[newWildcardEntry] "+path+"\n");

	baseDir := path[:len(path)-1]
	compositeEntry := []Entry{}

	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}
		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
			jarEntry := newZipEntry(path)
			compositeEntry = append(compositeEntry, jarEntry)
		}
		return nil
	}

	filepath.Walk(baseDir, walkFn)

	for i := range compositeEntry{
		fmt.Printf("To search from: " + compositeEntry[i].String())
		fmt.Printf("\n")
	}

	return compositeEntry
}

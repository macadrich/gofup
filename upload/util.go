package upload

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// CHashFile hash files in directory and return list of hash file (in slice)
func CHashFile(dir string) ([]File, error) {
	var out []File
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("directory walk failed: %s", err)
		}

		if info.IsDir() {
			return nil
		}

		// got file in directory
		fi, err := os.Open(path)
		if err != nil {
			return err
		}
		defer fi.Close()

		h := sha256.New()
		nb, err := io.Copy(h, fi) // copy file to sha256 allocated ()
		if err != nil {
			return err
		}

		sum := h.Sum(nil)                     // hash the file
		hashstr := hex.EncodeToString(sum[:]) // encode to string

		// add file information including hash in array (as collection)
		out = append(out, File{
			Name: filepath.Base(path),
			Size: nb,
			Hash: hashstr,
		})
		return nil
	})

	if err != nil {
		return nil, err
	}

	return out, nil
}

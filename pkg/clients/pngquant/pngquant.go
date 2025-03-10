package pngquant

import (
	"log"
	"os"
	"os/exec"
)

func Compress(source string, target string) error {
	log.Println("[PNGQUANT] Compressing file", source, "to", target)
	cmd := exec.Command("pngquant", "--quality=10", "--output", target, "--force", source)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

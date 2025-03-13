package guetzli

import "log"

func Compress(source string, target string) error {
	log.Println("[GUETZLI] Compressing file", source, "to", target)
	return nil
}

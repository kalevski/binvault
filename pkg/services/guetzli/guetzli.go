package guetzli

import "log"

func Compress(source string, target string) {
	log.Println("[GUETZLI] Compressing file", source, "to", target)
}

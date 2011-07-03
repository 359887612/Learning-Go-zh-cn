package main
import "os"

func main() {
	buf := make([]byte, 1024)
	f, _ := os.Open("/etc/passwd", os.O_RDONLY, 0666) |\longremark{Open the file;}|
	defer f.Close() |\longremark{Make sure we close it again;}|
	for {
		n, _ := f.Read(buf) |\longremark{Read up to 1023 bytes at the time;}|
		if n == 0 { break } |\longremark{We have reached the end of the file;}|
		os.Stdout.Write(buf[0:n]) |\longremark{Write the contents to \var{Stdout};}|
	}
}

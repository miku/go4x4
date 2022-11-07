// A short tool to fetch a single file from FTP.
//
// Usage of /tmp/go-build3945962022/b001/exe/main:
//   -H string
//         ftp hostname (default "ftp.ncbi.nlm.nih.gov:21")
//   -L    use lftp
//   -P string
//         password (default "anonymous")
//   -U string
//         username (default "anonymous")
//   -f string
//         filepath to retrieve (default "/pub/pmc/readme.txt")

package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/jlaffaye/ftp"
)

var (
	hostPort = flag.String("H", "ftp.ncbi.nlm.nih.gov:21", "ftp hostname")
	ftpPath  = flag.String("f", "/pub/pmc/readme.txt", "filepath to retrieve")
	username = flag.String("U", "anonymous", "username")
	password = flag.String("P", "anonymous", "password")
	useLftp  = flag.Bool("L", false, "use lftp")
)

func RetrLftp(hostPort, path, username, password string) error {
	var (
		tmpPath     = filepath.Join(os.TempDir(), fmt.Sprintf("filefetch-%d", rand.Intn(999999999)))
		lftpCommand = fmt.Sprintf(`pget %s -o %s; bye`, path, tmpPath)
		cmd         = exec.Command(
			"lftp", "-u", fmt.Sprintf("%s,%s", username, password), "-e", lftpCommand, hostPort)
	)
	log.Println(cmd)
	if b, err := cmd.CombinedOutput(); err != nil {
		log.Println(string(b))
		return err
	}
	f, err := os.Open(tmpPath)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := io.Copy(os.Stdout, f); err != nil {
		return err
	}
	return nil
}

func Retr(hostPort, path, username, password string) error {
	c, err := ftp.Dial(hostPort, ftp.DialWithTimeout(10*time.Second))
	if err != nil {
		return err
	}
	err = c.Login(username, password)
	if err != nil {
		return err
	}
	r, err := c.Retr(path)
	if err != nil {
		return err
	}
	// b, err := ioutil.ReadAll(r)
	// if err != nil {
	// 	return err
	// }
	// fmt.Println(string(b))
	_, err = io.Copy(os.Stdout, r)
	if err != nil {
		return err
	}
	// if _, err := io.Copy(os.Stdout, r); err != nil {
	// 	return err
	// }
	return c.Quit()
}

func main() {
	rand.Seed(time.Now().UnixNano())
	flag.Parse()
	var err error
	if *useLftp {
		err = RetrLftp(*hostPort, *ftpPath, *username, *password)
	} else {
		err = Retr(*hostPort, *ftpPath, *username, *password)
	}
	if err != nil {
		log.Fatal(err)
	}
}

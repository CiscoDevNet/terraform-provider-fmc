// Package ftd handles the functionality specific to Secure Firewall Threat Defense boxes.
package ftd

import (
	"crypto/rand"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"net/url"
	"os"
	"time"

	"golang.org/x/crypto/ssh"
)

// MustInitFromEnv initializes an FTD so that it would be ready for
// a registration attempt made by the given Firewall Management Center (FMC) if
// it uses identical TF_VAR_registration_key. It panics on any error.
//
// Reads from env FTD_ADDR, FTD_USERNAME, FTD_PASSWORD, FMC_URL, TF_VAR_registration_key.
// Assumes FTD_ADDR and FMC_URL are non-empty, checks others.
//
// It writes results of SSH session to the stdout.
//
// Func is not secure enough for production, use only for lab/testing.
func MustInitFromEnv() {
	// Determine address from FMC_URL
	fmcUrl, err := url.Parse(os.Getenv("FMC_URL"))
	if err != nil {
		panic(err)
	}
	fmcAddr, _, err := net.SplitHostPort(fmcUrl.Host)
	if err != nil {
		fmcAddr = fmcUrl.Host
	}

	if os.Getenv("FTD_USERNAME") == "" {
		panic("FTD_USERNAME env variable must be set")
	}
	if os.Getenv("FTD_PASSWORD") == "" {
		panic("FTD_PASSWORD env variable must be set")
	}
	if os.Getenv("TF_VAR_registration_key") == "" {
		panic("TF_VAR_registration_key env variable must be set")
	}

	err = initialize(
		os.Getenv("FTD_ADDR"),
		os.Getenv("FTD_USERNAME"),
		os.Getenv("FTD_PASSWORD"),
		fmcAddr,
		os.Getenv("TF_VAR_registration_key"),
		os.Getenv("TF_VAR_nat_id" /* can be unset */),
	)
	if err != nil {
		panic(err)
	}
}

func initialize(ftdAddr, ftdUser, ftdPass, fmcAddr, registrationKey, natID string) error {
	client, err := ssh.Dial("tcp",
		net.JoinHostPort(ftdAddr, "22"),
		&ssh.ClientConfig{
			Timeout: 5 * time.Second,
			User:    ftdUser,
			Auth: []ssh.AuthMethod{
				ssh.Password(ftdPass),
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		})
	if err != nil {
		return err
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	stdin, err := session.StdinPipe()
	if err != nil {
		return err
	}

	stderr, err := session.StderrPipe()
	if err != nil {
		return err
	}

	stdout, err := session.StdoutPipe()
	if err != nil {
		return err
	}

	err = session.Shell()
	if err != nil {
		return err
	}

	fmt.Fprintf(stdin, "configure manager delete %s\n", natID)
	fmt.Fprintf(stdin, "configure manager delete %s\n", fmcAddr)
	fmt.Fprintln(stdin, "configure manager delete")
	fmt.Fprintf(stdin, "configure manager add  %s %s %s\n", fmcAddr, registrationKey, natID)
	fmt.Fprint(stdin, "show managers verbose\n")
	stdin.Close()

	go func() {
		_, _ = io.Copy(os.Stdout, stdout)
	}()

	err = session.Wait()
	if err != nil {
		return err
	}
	written, err := io.Copy(os.Stdout, stderr) // blocks
	if err != nil {
		return err
	}
	if written != 0 {
		return errors.New("stderr is non-empty")
	}

	return nil
}

// MustRandomizeKey returns a random registration key, for example "0x21ae0cdc4518cc2d", or panics.
func MustRandomizeKey() string {
	ret, err := RandomizeKey()
	if err != nil {
		panic(err)
	}
	return ret
}

// RandomizeKey returns a random registration key, for example "0x21ae0cdc4518cc2d".
func RandomizeKey() (string, error) {
	randb := make([]byte, 8)
	_, err := rand.Read(randb)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("0x%016x", binary.BigEndian.Uint64(randb)), nil
}

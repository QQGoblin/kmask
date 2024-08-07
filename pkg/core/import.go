package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/QQGoblin/kmask/pkg/background"
	"github.com/QQGoblin/kmask/pkg/core/encrypt"
	"io"
	"os"
)

func Import(config, passphrase, picture string) error {

	inByte, err := os.ReadFile(config)
	if err != nil {
		return err
	}

	in := &KMaskData{}
	if err = json.Unmarshal(inByte, in); err != nil {
		return err
	}

	for _, secret := range in.Secrets {
		for _, secretFile := range secret.Files {
			if err = secretFile.Load(); err != nil {
				return err
			}
		}
	}

	var (
		scByte  []byte
		crypted []byte
	)

	scByte, err = json.Marshal(in.Secrets)
	if err != nil {
		return err
	}

	crypted, err = encrypt.AESEncrypt(scByte, []byte(passphrase))
	if err != nil {
		return err
	}

	return writeWithPicture(picture, in.DBFileName(), crypted)
}

func writeWithPicture(picture string, outfile string, data []byte) error {

	var (
		in  io.Reader
		f   *os.File
		err error
	)

	if picture != "" {
		f, err = os.Open(picture)
		if err != nil {
			return err
		}
		in = f
		defer f.Close()
	} else {
		var bgdata []byte
		bgdata, err = background.Asset(background.Default)
		if err != nil {
			return err
		}
		in = bytes.NewReader(bgdata)
	}

	if _, err = os.Stat(outfile); err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	} else {
		return fmt.Errorf("%s is exist, please backup", outfile)
	}
	out, _ := os.Create(outfile)
	defer out.Close()

	return encrypt.Encode(in, bytes.NewReader(data), out)
}

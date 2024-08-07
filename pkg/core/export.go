package core

import (
	"encoding/json"
	"fmt"
	"github.com/QQGoblin/kmask/pkg/core/encrypt"
	"os"
	"text/tabwriter"
)

func Export(data, passphrase, output string, all bool) error {

	secrets, err := loadSecret(data, passphrase)
	if err != nil {
		return err
	}

	if all {
		k := KMaskData{
			Secrets: secrets,
		}
		return k.Export(output)
	}

	printSecret(secrets)

	return nil
}

func loadSecret(filepath, passphrase string) ([]*Secret, error) {

	in, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer in.Close()

	cryptBytes, err := encrypt.Decode(in)
	if err != nil {
		return nil, err
	}

	data, err := encrypt.AESDecrypt(cryptBytes, []byte(passphrase))
	if err != nil {
		return nil, err
	}

	secrets := make([]*Secret, 0)

	if err = json.Unmarshal(data, &secrets); err != nil {
		return nil, err
	}

	return secrets, nil

}

func printSecret(secrets []*Secret) {
	w := tabwriter.NewWriter(os.Stdout, 10, 0, 5, ' ', tabwriter.TabIndent)
	fmt.Fprintln(w, "UID\tKey\tValue")
	for _, secret := range secrets {
		firstLine := true

		f := func(key, value string) {
			if firstLine {
				fmt.Fprintf(w, "%s\t%s\t%s\n", secret.UID, key, value)
				firstLine = false
				return
			}
			fmt.Fprintf(w, "%s\t%s\t%s\n", " ", key, value)
		}

		for k, v := range secret.Keys {
			f(k, v)
		}

		for _, file := range secret.Files {
			f(file.Name, file.Description)
		}
	}
	w.Flush()
}

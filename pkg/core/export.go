package core

import (
	"encoding/json"
	"fmt"
	"github.com/QQGoblin/kmask/pkg/core/encrypt"
	"github.com/WeiZhang555/tabwriter"
	"os"
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
	w := tabwriter.NewWriter(os.Stdout, 20, 0, 5, ' ', tabwriter.TabIndent)
	fmt.Fprintln(w, "UID\tKey\tValue")
	for _, secret := range secrets {
		firstLine := true
		f := func(key, value string) {
			uid := " "
			if firstLine {
				uid = secret.UID
				firstLine = false
			}
			// TODO: 这里golang原生tabwriter，不能正确计算中文字符的宽度，因此使用一个第三方的库
			// 具体参考：https://github.com/golang/go/issues/12073，https://github.com/golang/go/issues/13989
			fmt.Fprintf(w, "%s\t%s\t%s\n", uid, key, value)
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

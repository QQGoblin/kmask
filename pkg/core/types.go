package core

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
)

const (
	DefaultExportConfigName = "secret.json"
	DefaultDBName           = "result.png"
	defaultSecretFilepath   = "files"
)

type Secret struct {
	UID   string            `json:"uid"`
	Keys  map[string]string `json:"keys"`
	Files []*SecretFile     `json:"files,omitempty"` // filename and filepath
}

type SecretFile struct {
	Name        string `json:"name"`
	Path        string `json:"path,omitempty"`
	Description string `json:"description,omitempty"`
	Data        string `json:"data,omitempty"`
}

type KMaskData struct {
	Secrets []*Secret `json:"secrets"`
}

func (f *SecretFile) Load() error {

	if f.Data != "" {
		return nil
	}

	b, err := os.ReadFile(f.Path)
	if err != nil {
		return errors.Wrapf(err, "load %s", f.Path)
	}
	f.Data = base64.StdEncoding.EncodeToString(b)
	return nil
}

func (f *KMaskData) Export(output string) error {

	if _, err := os.Stat(output); err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	} else {
		if err = os.RemoveAll(output); err != nil {
			return err
		}
	}

	secretFilepath := filepath.Join(output, defaultSecretFilepath)
	if err := os.MkdirAll(secretFilepath, 0755); err != nil {
		return err
	}

	copySecrets := make([]*Secret, 0)
	for _, secret := range f.Secrets {
		copySecret := Secret{
			UID:   secret.UID,
			Keys:  secret.Keys,
			Files: make([]*SecretFile, len(secret.Files)),
		}
		for i, sfile := range secret.Files {

			outputSFilePath := filepath.Join(secretFilepath, sfile.Name)

			data, err := base64.StdEncoding.DecodeString(sfile.Data)
			if err != nil {
				return err
			}

			if err = os.WriteFile(outputSFilePath, data, 0600); err != nil {
				return err
			}
			copySecret.Files[i] = &SecretFile{
				Name:        sfile.Name,
				Path:        filepath.Join(defaultSecretFilepath, sfile.Name),
				Description: sfile.Description,
			}
		}

		copySecrets = append(copySecrets, &copySecret)

	}
	outConfig := filepath.Join(output, DefaultExportConfigName)
	outKMask := KMaskData{Secrets: copySecrets}
	b, err := json.Marshal(outKMask)
	if err != nil {
		return err
	}
	var fmtJson bytes.Buffer
	if err = json.Indent(&fmtJson, b, "", "    "); err != nil {
		return err
	}

	return os.WriteFile(outConfig, fmtJson.Bytes(), 0600)
}

func (f *KMaskData) DBFileName() string {

	return DefaultDBName
}

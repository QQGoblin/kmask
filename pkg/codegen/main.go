package main

import (
	"github.com/go-bindata/go-bindata"
	"k8s.io/klog/v2"
)

func main() {

	bc := &bindata.Config{
		Input: []bindata.InputConfig{
			{
				Path: "pkg/codegen/background",
			},
		},
		Package:    "background",
		NoCompress: false,
		NoMemCopy:  true,
		NoMetadata: true,
		Output:     "pkg/background/zz_generated_bindata.go",
	}
	if err := bindata.Translate(bc); err != nil {
		klog.Fatal(err)
	}

}

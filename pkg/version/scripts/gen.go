package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/elliotxx/go-cli-prototype/pkg/version"
)

func main() {
	versionInfo, err := version.NewInfo()
	if err != nil {
		log.Fatal(err)
	}

	data := makeUpdateVersionGoFile(versionInfo)

	err = ioutil.WriteFile("../z_update_version.go", []byte(data), 0o666)
	if err != nil {
		log.Fatalf("ioutil.WriteFile: err = %v", err)
	}

	fmt.Println(versionInfo.String())
}

func makeUpdateVersionGoFile(v *version.Info) string {
	return fmt.Sprintf(`// Auto generated by 'go run gen.go', DO NOT EDIT.

package version

func init() {
	info = &Info{
		ReleaseVersion: %q,
		GitInfo: &GitInfo{
			LatestTag:   %q,
			Commit:      %q,
			TreeState:   %q,
		},
		BuildInfo: &BuildInfo{
			GoVersion: %q,
			GOOS:      %q,
			GOARCH:    %q,
			NumCPU:    %d,
			Compiler:  %q,
			BuildTime: %q,
		},
	}
}
`,
		v.ReleaseVersion,
		v.GitInfo.LatestTag,
		v.GitInfo.Commit,
		v.GitInfo.TreeState,
		v.BuildInfo.GoVersion,
		v.BuildInfo.GOOS,
		v.BuildInfo.GOARCH,
		v.BuildInfo.NumCPU,
		v.BuildInfo.Compiler,
		v.BuildInfo.BuildTime,
	)
}

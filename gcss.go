package gcss

// github.com/yosssi/gcss binding for gonzo.
// No Configuration required.

import (
	"bytes"
	"io/ioutil"
	"strings"

	"github.com/omeid/gonzo/context"

	"github.com/omeid/gonzo"
	"github.com/yosssi/gcss"
)

func Compile() gonzo.Stage {
	return func(ctx context.Context, in <-chan gonzo.File, out chan<- gonzo.File) error {

		for {
			select {
			case file, ok := <-in:
				if !ok {
					return nil
				}

				buff := new(bytes.Buffer)
				name := strings.TrimSuffix(file.FileInfo().Name(), ".gcss") + ".css"
				ctx.Infof("Compiling %s to %s", file.FileInfo().Name(), name)
				n, err := gcss.Compile(buff, file)
				if err != nil {
					return err
				}

				file = gonzo.NewFile(ioutil.NopCloser(buff), file.FileInfo())
				file.FileInfo().SetSize(int64(n))
				file.FileInfo().SetName(name)

				out <- file
			case <-ctx.Done():
				return nil
			}
		}
	}
}

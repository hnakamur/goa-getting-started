package cli

import (
	"encoding/json"
	"fmt"
	"github.com/goadesign/goa"
	goaclient "github.com/goadesign/goa/client"
	uuid "github.com/goadesign/goa/uuid"
	"github.com/hnakamur/goa-getting-started/client"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

type (
	// ShowBottleCommand is the command line data structure for the show action of bottle
	ShowBottleCommand struct {
		// ボトル ID
		BottleID    int
		PrettyPrint bool
	}

	// DownloadCommand is the command line data structure for the download command.
	DownloadCommand struct {
		// OutFile is the path to the download output file.
		OutFile string
	}
)

// RegisterCommands registers the resource action CLI commands.
func RegisterCommands(app *cobra.Command, c *client.Client) {
	var command, sub *cobra.Command
	command = &cobra.Command{
		Use:   "show",
		Short: `指定されたIDのボトルを取得する`,
	}
	tmp1 := new(ShowBottleCommand)
	sub = &cobra.Command{
		Use:   `bottle ["/bottles/BOTTLEID"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp1.Run(c, args) },
	}
	tmp1.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp1.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)

	dl := new(DownloadCommand)
	dlc := &cobra.Command{
		Use:   "download [PATH]",
		Short: "Download file with given path",
		RunE: func(cmd *cobra.Command, args []string) error {
			return dl.Run(c, args)
		},
	}
	dlc.Flags().StringVar(&dl.OutFile, "out", "", "Output file")
	app.AddCommand(dlc)
}

func intFlagVal(name string, parsed int) *int {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func float64FlagVal(name string, parsed float64) *float64 {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func boolFlagVal(name string, parsed bool) *bool {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func stringFlagVal(name string, parsed string) *string {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func hasFlag(name string) bool {
	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "--"+name) {
			return true
		}
	}
	return false
}

func jsonVal(val string) (*interface{}, error) {
	var t interface{}
	err := json.Unmarshal([]byte(val), &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func jsonArray(ins []string) ([]interface{}, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []interface{}
	for _, id := range ins {
		val, err := jsonVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, val)
	}
	return vals, nil
}

func timeVal(val string) (*time.Time, error) {
	t, err := time.Parse("RFC3339", val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func timeArray(ins []string) ([]time.Time, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []time.Time
	for _, id := range ins {
		val, err := timeVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func uuidVal(val string) (*uuid.UUID, error) {
	t, err := uuid.FromString(val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func uuidArray(ins []string) ([]uuid.UUID, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []uuid.UUID
	for _, id := range ins {
		val, err := uuidVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func float64Val(val string) (*float64, error) {
	t, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func float64Array(ins []string) ([]float64, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []float64
	for _, id := range ins {
		val, err := float64Val(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func boolVal(val string) (*bool, error) {
	t, err := strconv.ParseBool(val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func boolArray(ins []string) ([]bool, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []bool
	for _, id := range ins {
		val, err := boolVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

// Run downloads files with given paths.
func (cmd *DownloadCommand) Run(c *client.Client, args []string) error {
	var (
		fnf func(context.Context, string) (int64, error)
		fnd func(context.Context, string, string) (int64, error)

		rpath   = args[0]
		outfile = cmd.OutFile
		logger  = goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
		ctx     = goa.WithLogger(context.Background(), logger)
		err     error
	)

	if rpath[0] != '/' {
		rpath = "/" + rpath
	}
	if rpath == "/swagger.json" {
		fnf = c.DownloadSwaggerJSON
		if outfile == "" {
			outfile = "swagger.json"
		}
		goto found
	}
	if strings.HasPrefix(rpath, "/swaggerui/") {
		fnd = c.DownloadSwaggerui
		rpath = rpath[11:]
		if outfile == "" {
			_, outfile = path.Split(rpath)
		}
		goto found
	}
	return fmt.Errorf("don't know how to download %s", rpath)
found:
	ctx = goa.WithLogContext(ctx, "file", outfile)
	if fnf != nil {
		_, err = fnf(ctx, outfile)
	} else {
		_, err = fnd(ctx, rpath, outfile)
	}
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	return nil
}

// Run makes the HTTP request corresponding to the ShowBottleCommand command.
func (cmd *ShowBottleCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/bottles/%v", cmd.BottleID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ShowBottle(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowBottleCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var bottleID int
	cc.Flags().IntVar(&cmd.BottleID, "bottleID", bottleID, `ボトル ID`)
}

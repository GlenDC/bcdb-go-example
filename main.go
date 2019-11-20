package main

import (
	mrand "math/rand"
	"crypto/rand"
	"encoding/json"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/glendc/bcdb-go-example/pkg/tftexplorer"
	tftexplt "github.com/glendc/bcdb-go-example/pkg/types/tftexplorer"
	"github.com/threefoldtech/zos/pkg/schema"
)

var (
	flagConsensusChangeID = flag.String("ccid", "", "consensus change ID to be set during SET")
	flagHeight            = flag.Int64("height", 0, "height to be set during SET")
	flagTimestamp         = flag.Int64("timestamp", 0, "timestamp to be set during SET")
	flagBlockID           = flag.String("block", "", "block ID to be set during SET")
	flagRandom            = flag.Bool("random", false, "use random data during SET")
)

func main() {
	// define method and remaining args
	var (
		args   = flag.Args()
		method string
	)
	if argn := len(args); argn == 0 {
		method = "get"
	} else {
		method = strings.ToLower(args[0])
		args = args[1:]
	}
	if len(args) > 0 {
		fmt.Fprintf(os.Stderr, "unexpected remaining arguments: %v\n", args)
		os.Exit(1)
	}

	// get client
	tftexpl := tftexplorer.NewClient("https://172.17.0.2")

	// set/get chain context
	switch method {
	case "get":
		get(tftexpl)

	case "set", "post":
		set(tftexpl)
		get(tftexpl)

	default:
		fmt.Fprintf(os.Stderr, "unexpected method %s, valid methods: get, set\n", method)
		os.Exit(1)
	}
}

func get(client *tftexplorer.Client) {
	var chainCtx tftexplt.TftExplorerChainContext1
	err := client.Get("get_chain_context", nil, &chainCtx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "get error: %v\n", err)
		os.Exit(1)
	}
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "\t ")
	err = enc.Encode(chainCtx)
	if err != nil {
		panic(err)
	}
}

func set(client *tftexplorer.Client) {
	var chainCtx tftexplt.TftExplorerChainContext1
	if *flagRandom {
		chainCtx.ConsensusChangeId = randomHexString(32)
		chainCtx.Height = randomInt64(1, 500000)
		chainCtx.Timestamp = schema.Date{Time: time.Unix(randomInt64(1522453800, 1590392100), 0)}
		chainCtx.BlockId = randomHexString(32)
	} else {
		chainCtx.ConsensusChangeId = *flagConsensusChangeID
		chainCtx.Height = *flagHeight
		chainCtx.Timestamp = schema.Date{Time: time.Unix(*flagTimestamp, 0)}
		chainCtx.BlockId = *flagBlockID
	}
	err := client.Set("set_chain_context", chainCtx, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "set error: %v\n", err)
		os.Exit(1)
	}
}

func randomHexString(n int) string {
	if n <= 0 {
		return ""
	}
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		panic(fmt.Sprintln("randomHexString", n, err))
	}
	return hex.EncodeToString(b)
}

func randomInt64(min, max int64) int64 {
	if min == max {
		return min
	}
	if min > max {
		panic(fmt.Sprintf("max %d cannot be smaller then min %d", max, min))
	}
	return mrand.Int63n(max-min) + min
}

func init() {
	flag.Parse()
}

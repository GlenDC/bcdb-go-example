# bcdb-go-example

A small example how to I/O with BCDB using Go.

It interacts with BCDB over https web gedis,
using the `tft_explorer` 3Bot package found at
<https://github.com/threefoldtech/jumpscaleX_threebot/tree/8fe7611df1a80e6bc36a92714eaf33a50f971449/ThreeBotPackages/threefold/tft_explorer>.

JumpscaleX_Core was at commit `007522983f93bdd44cf69e3a25e1e1fc5dba7e40` at the time of developing this small example.

## Install

```bash
make install
```

Or

```
go install .
```

Or

```
go get -u github.com/glendc/bcdb-go-example
```

## Usage

> Prerequisites:
> - Have a threebot server running somewhere with the `tft_explorer` running;
> - Make sure the used commit version of your `jumpscaleX_threebot` matches the one defined at the top of this document;
> - Make sure the used commit version of your `jumpscaleX_core` matches the one defined at the top of this document;
> - Make sure your threebot server is reachable over HTTPS Gedis.

Get the chain context at this point:

```bash
$ bcdb-go-example get
{
         "consensus_change_id": "",
         "height": 0,
         "timestamp": 0,
         "block_id": ""
}
```

Set the chain context with random data:

```bash
$ bcdb-go-example --random set
{
         "consensus_change_id": "60807b432f03ed7a991a6f47ee6105db0f43acd2147dafe52aecdbd2f536cff1",
         "height": 171378,
         "timestamp": 1529971200,
         "block_id": "ceddf2658437e8dec207d9c2579098dc7fa607cdaf044b41d876f0b4293db674"
}
```

> Note: The updated chain context is returned in each set call.

Set the chain context with all properties equal to their default null value:

```bash
bcdb-go-example set
{
         "consensus_change_id": "",
         "height": 0,
         "timestamp": 0,
         "block_id": ""
}
```

Set the chain context with custom properties (all properties not given as flags use the default null value):

```bash
bcdb-go-example --ccid "foo" --block "bar" --height 42 --timestamp 1574208000 set
{
         "consensus_change_id": "foo",
         "height": 42,
         "timestamp": 1574208000,
         "block_id": "bar"
}
```

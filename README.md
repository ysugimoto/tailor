# Tailor

`Tailor` is a simple log tailing, summarize tool.

## Installation

This product supplied by single binary for `MacOS` and `Linux`. Download latest binary from [releases](https://github.com/ysugimoto/tailor/releases).

## Basic Usage

Put a binary on your `PATH` directory.

```
$ tailor [options] [,file]
```

### Simple file tailing

```
$ tailor [file]
```

It works like `tail -f` command.

### Tailing from stdin

```
$ [some command] | tailor --stdin
```

It works continous tailing command output.

## Transporting

`tailor` can tarnsport data over the network like `fluendtd` . It will useful for correct same logs from multiple machine's, e.g. docker containers, Web servers under the load balancing.
Transport some data from CLI, and realtime watch these data on Web UI.

### Central server

First, execute central-server process at external server that enable to access by browser, (sample is `example.com`)

```
$ tailor -C [-d]
```

If you want to run with daemon, add `-d` option.

`tailor` supply the Web Interface at http://example.com:9000 (enable to change listen host/port number with `-h`/`-p` option), access `http://examole.com:9000` on browser.
Central server accepts data on HTTP interface with CORS. In this sample case, endpoint is `http://example.com:9000/remote`. then, send data from any client support HTTP interface, `curl`, `XMLHttpRequest`, etc....

### Tailing server(s)

Second, Execute tail data and transporting process at web servers you want to watch:

```
$ tailor [-d] -P http://example.com:9000 /var/log/httpd/access_log
```

If you want to run with daemon, add `-d` option.

`-P` option's value is a central server's host:port. Then, tailing data will transport to central server, and can watch Web Interface.

### Stopping server

If you stop the daemon, execute with `--kill` option:

```
$ tailor --kill
```

## Web UI

Central server supply GUI on browser, updating realtime logs by WebSocket.
You can see logs that classified by host, one, two, four splitted panes.

## Note

- Currently we support transport interface only **HTTP** , not support **HTTPS**. So that, central/transporting server recommend to be an appropriate settings, with Firewall.
- Transported data does not save anywhere, watiching only.

## Command Options

| option          | description                                                  | default   |
|-----------------|--------------------------------------------------------------|-----------|
| -p, --port      | Change listen port.                                          | 9000      |
| -h, --host      | Change listen host.                                          | 0.0.0.0   |
| -C, --central   | Run with centran server mode.                                | -         |
| -t, --transport | Run with tansport mode. Determine central server's full URI. | -         |
| -d, --daemon    | Run with daemon.                                             | -         |
| -c, --client    | Run with client mode.                                        | -         |
| -k, --kill      | Kill the daemon process.                                     | -         |
| --help          | Show command help.                                           | -         |
| --stdin         | Tailing from stdin data.                                     | -         |

## Contributing

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request

## License

MIT License

## Author

Yoshiaki Sugimoto


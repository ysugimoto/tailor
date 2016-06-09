## Tailor

Tailor is a simple log tailing, summarize tool, supplied by single binary for `MoxOS` and `Linux`. Download latest binary from [releases](https://github.com/ysugimoto/tailor/releases).

## Basic

Put a binary on your `PATH` directory.

```
$ tailor [options] [,file]
```

### Simple file tailing

```
$ tailor [file]
```

It works like `tail` command.

### Tailing from stdin

```
$ [some command] | tailor --stdin
```

It works continous tailing command output.

## Transporting

`tailor` can tarnsport tailing data over the network. It will useful for correct same logs from multiple machine's, e.g. docker containers.

### Central server

First, Execute central-server process at external server that enable to access by browser, (for example, example.com)

```
$ tailor -C
```

If you want to run with daemon, add `-d` option:

```
$ tailor -C -d
```

And, tailor supply the Web Interface at http://example.com:9000 (enable to change listen port number with `-p` option), access `http://examole.com:9000` on browser.
Central server accepts data on HTTP interface, in this sample case, endpoint is `http://example.com:9000/remote`. then, send data from any client support HTTP interface, `curl`, `XMLHttpRequest`, and so on.

### Tailing server

Second, Execute tail data and transporting process at web servers you want to watch:

```
$ tailor -P http://example.com:9000 /var/log/httpd/access_log
```

If you want to run with daemon, add `-d` option:

```
$ tailor -d -P http://example.com:9000 /var/log/httpd/access_log
```

`-P` option's value is a central server's host:port. Then, tailing data will transport to central server, and can watch Web Interface.


### Stopping server

If you stop the daemon, execute with `--kill` option:

```
$ tailor --kill
```

## Options

| option          | description                                                  | default   |
|-----------------|--------------------------------------------------------------|-----------|
| -p, --port      | Change listen port.                                          | 9000      |
| -h, --host      | Change listen host.                                          | 0.0.0.0   |
| -C, --central   | Run with centran server mode.                                | -         |
| -t, --transport | Run with tansport mode. Determine central server's full URI. | -         |
| -d, --daemon    | Run with daemon.                                             | -         |
| -c, --client    | Run with daemon.                                             | -         |
| -h, --help      | Show the help.                                               | -         |
| --kill          | Kill the daemon process.                                     | -         |
| --stdin         | Tailing from stdin data.                                     | -         |


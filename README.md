**scv** is a simple and modular task runner. It supports loading triggers and actions as plugins.

# Scripts
`start.sh` runs scv in the background using `config.json` as the config file.

`stop.sh` kills a running scv and cleans up the pid file.

# Example
Save following snippet as `config.json` and run `./scv -conf=config.json` to monitor the current directory recursively and run `make` when a file is written:
```
[
  {
    "trigger": {"type": "watchdir", "param": "."},
    "action": {"type": "build", "param": "make"}
  }
]
```

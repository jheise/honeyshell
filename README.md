# HoneyShell
---
Simple shell for use in honeypots.

HoneyShell is a limited shell designed for use in ssh honeypots, all it does is read from standard in and writes
json encoded log lines to a file.

```
honeyshell </path/to/logfil> <ipAddr:port>
```

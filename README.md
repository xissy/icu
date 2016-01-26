# icu
> Detect whatsoever charset.

```
usage: icu [<flags>] [<file>]

Detect whatsoever charset

Flags:
      --help           Show context-sensitive help (also try --help-long and --help-man).
  -b, --buffer=100000  Buffer size of the file.
  -t, --text           Enable text more, default is HTML mode.
      --version        Show application version.

Args:
  [<file>]  Input file name, or you can use pipe or redirection.
```

## with iconv
```
iconv -f $(icu a.smi) -t utf8 a.smi > a_utf8.smi
```

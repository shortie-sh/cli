# `shortie` - The CLI companion to shortie.sh

## Usage: 
```
shortie <url>

create a new redirect

Commands:
  shortie new <url>     create a new redirect                          [default]
  shortie get <ending>  get existing redirect

Positionals:
  url  target URL to be redirected                                      [string]

Options:
      --version   Show version number                                  [boolean]
  -i, --instance  instance of shortie.sh to use
                                        [string] [default: "https://shortie.sh"]
      --help      Show help                                            [boolean]
  -e, --ending    custom ending for redirect                            [string]
```
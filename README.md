# apg.go
_apg.go_ is a simple APG-like password generator script written in Go. It tries to replicate the
functionality of the
"[Automated Password Generator](https://web.archive.org/web/20130313042424/http://www.adel.nursat.kz:80/apg)",
which hasn't been maintained since 2003. Since more and more Unix distributions are abondoning the tool, I was
looking for an alternative. FreeBSD for example recommends "security/makepasswd", which is written in Perl
but requires a lot of dependency packages and doesn't offer the feature-set/flexibility of APG. Therefore, as a
first attempt, I decided to write [my own implementation in Perl](https://github.com/wneessen/passwordGen), but since
I just started learning Go, I gave it another try and reproduced apg.pl in Go as apg.go. Again, as I never used
the "pronouncable password" functionality, I left this out in this version.

## Usage
Either use the binary releases, unzip them and simply execute them:
```sh
$ unzip apg_v0.2.0_linux_amd64.zip
$ chmod +x apg
$ ./apg
```
Or download the sources and build the binary yourselves:
```sh
$ go build apg.go
$ ./apg
```

## Systemwide installation
To be a proper APG replacement, i suggest to install it into a directory in your PATH and symlink it to "apg":
```sh
$ sudo cp apg /usr/local/bin/apg
```

## CLI options
_apg.go_ replicates some of the parameters of the original APG. Some parameters are different though:

- ```-m, --minpasslen <length>```: The minimum length of the password to be generated
- ```-x, --maxpasslen <length>```: The maximum length of the password to be generated
- ```-n, --numofpass <number of passwords>```: The amount of passwords to be generated
- ```-E, --exclude <list of characters>```: Do not use the specified characters in generated passwords
- ```-U, --uppercase```: Use uppercase characters in passwords
- ```-N, --numbers```: Use numeric characters in passwords
- ```-S, --special```: Use special characters in passwords
- ```-H, --human```: Avoid ambiguous characters in passwords (i. e.: 1, l, I, o, O, 0)
- ```-c, --complex```: Generate complex passwords (implies -U -N -S and disables -H)
- ```-h, --help```: Show a CLI help text
- ```-v, --version```: Show the version number
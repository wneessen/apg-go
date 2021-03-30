# Advanced Password Generator (APG) Clone
_apg-go_ is a simple APG-like password generator script written in Go. It tries to replicate the
functionality of the
"[Automated Password Generator](https://web.archive.org/web/20130313042424/http://www.adel.nursat.kz:80/apg)",
which hasn't been maintained since 2003. Since more and more Unix distributions are abondoning the tool, I was
looking for an alternative. FreeBSD for example recommends "security/makepasswd", which is written in Perl
but requires a lot of dependency packages and doesn't offer the feature-set/flexibility of APG.

Therefore, as a first attempt, I decided to write 
[my own implementation in Perl](https://github.com/wneessen/passwordGen), but since I just started learning Go, 
I gave it another try and reproduced apg.pl in Go as apg.go.

Since FIPS-181 (pronouncable passwords) has been withdrawn in 2015, I didn't see any use in replicating that
feature. Therfore apg.go does not support pronouncable passwords.

## Installation
### Binary releases
#### Linux/BSD/MacOS
* Download release
  ```sh
  $ curl -LO https://github.com/wneessen/apg.go/releases/download/v<version>/apg-v<version>-<os>-<architecture>.tar.gz
  $ curl -LO https://github.com/wneessen/apg.go/releases/download/v<version>/apg-v<version>-<os>-<architecture>.tar.gz.sha256
  ```
* Verify the checksum
  ```sh
  $ sha256 apg-v<version>-<os>-<architecture>.tar.gz 
  $ cat apg-v<version>-<os>-<architecture>.tar.gz.sha256
  ```
  **Make sure the checksum of the downloaded file and the checksum in the .sha256 match**
* Extract archive
  ```sh
  $ tar xzf apg-v<version>-<os>-<architecture>.tar.gz
  ```
* Execute
  ```sh
  $ ./apg
  ```
#### Windows
* Download release
  ```PowerShell
  PS> Invoke-RestMethod -Uri https://github.com/wneessen/apg.go/releases/download/v<version>/apg-v<version>-windows-<architecture>.zip -OutFile apg-v<version>-windows-<architecure>.zip
  PS> Invoke-RestMethod -Uri https://github.com/wneessen/apg.go/releases/download/v<version>/apg-v<version>-windows-<architecture>.zip.sha256 -OutFile apg-v<version>-windows-<architecure>.zip.sha256
  ```
* Verify the checksum
  ```PowerShell
  PS> Get-FileHash apg-v<version>-windows-<architecture>.zip | Format-List
  PS> type apg-v<version>-windows-<architecture>.zip.sha256
  ```
  **Make sure the checksum of the downloaded file and the checksum in the .sha256 match**
* Extract archive
  ```PowerShell
  PS> Expand-Archive -LiteralPath apg-v<version>-windows-<architecture>
  ```
* Execute
  ```PowerShell
  PS> cd apg-v<version>-windows-<architecture> 
  PS> apg.exe
  ```

### Sources
* Download sources
  ```sh
  $ curl -LO https://github.com/wneessen/apg.go/archive/refs/tags/v<version>.tar.gz
  ```
* Extract source
  ```sh
  $ tar xzf v<version>.tar.gz
  ```
* Build binary
  ```sh
  $ cd apg.go-<version>
  $ go build -o apg ./...
  ```
* Execute the brand new binary
  ```sh
  $ ./apg
  ```

### Systemwide installation
It is recommed to install apg in a directory of your ```$PATH``` environment. To do so run:
(In this example we use ```/usr/local/bin``` as system-wide binary path. YMMV)
```sh
$ sudo cp apg /usr/local/bin/apg
```

## CLI parameters
_apg.go_ replicates some of the parameters of the original APG. Some parameters are different though:

- ```-m <length>```: The minimum length of the password to be generated (Default: 20)
- ```-x <length>```: The maximum length of the password to be generated (Default: 20)
- ```-n <number of passwords>```: The amount of passwords to be generated (Default: 1)
- ```-E <list of characters>```: Do not use the specified characters in generated passwords
- ```-M <[LUNSHClunshc]>```: New style password parameters (upper-case enables, lower-case disables)
- ```-L```: Use lower-case characters in passwords (Default: on)
- ```-U```: Use upper-case characters in passwords (Default: on)
- ```-N```: Use numeric characters in passwords (Default: on)
- ```-S```: Use special characters in passwords (Default: off)
- ```-H```: Avoid ambiguous characters in passwords (i. e.: 1, l, I, o, O, 0) (Default: off)
- ```-C```: Generate complex passwords (implies -L -U -N -S and disables -H) (Default: off)
- ```-l```: Spell generated passwords (Default: off)
- ```-h```: Show a CLI help text
- ```-v```: Show the version number

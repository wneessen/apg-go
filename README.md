# A "Automated Password Generator"-clone
[![Go Reference](https://pkg.go.dev/badge/github.com/wneessen/apg-go.svg)](https://pkg.go.dev/github.com/wneessen/apg-go) [![Go Report Card](https://goreportcard.com/badge/github.com/wneessen/apg-go)](https://goreportcard.com/report/github.com/wneessen/apg-go) [![Build Status](https://api.cirrus-ci.com/github/wneessen/apg-go.svg)](https://cirrus-ci.com/github/wneessen/apg-go) ![CodeQL workflow](https://github.com/wneessen/apg-go/actions/workflows/codeql-analysis.yml/badge.svg) <a href="https://ko-fi.com/D1D24V9IX"><img src="https://uploads-ssl.webflow.com/5c14e387dab576fe667689cf/5cbed8a4ae2b88347c06c923_BuyMeACoffee_blue.png" height="20" alt="buy ma a coffee"></a>

_apg-go_ is a simple APG-like password generator written in Go. It tries to replicate the
functionality of the
"[Automated Password Generator](https://web.archive.org/web/20130313042424/http://www.adel.nursat.kz:80/apg)",
which hasn't been maintained since 2003. Since more and more Unix distributions are abondoning the tool, I was
looking for an alternative. FreeBSD for example recommends "security/makepasswd", which is written in Perl
but requires a lot of dependency packages and doesn't offer the feature-set/flexibility of APG.

Since FIPS-181 (pronouncable passwords) has been withdrawn in 2015, apg-go does not follow this standard. Instead
it implements the [Koremutake Syllables System](https://shorl.com/koremutake.php) in its pronouncable password mode.

## Installation
### Ports/Packages
#### FreeBSD
apg-go can be found as `/security/apg` in the [FreeBSD ports](https://cgit.freebsd.org/ports/tree/security/apg)
tree.
#### Arch Linux
Find apg-go in [Arch Linux AUR](https://aur.archlinux.org/packages/apg-go/). \
Alternatively use the [PKGBUILD](https://github.com/wneessen/apg-go/tree/main/buildfiles/arch-linux) file 
in this git repository
### Binary releases
#### Linux/BSD/MacOS
* Download release
  ```sh
  $ curl -LO https://github.com/wneessen/apg-go/releases/download/v<version>/apg-v<version>-<os>-<architecture>.tar.gz
  $ curl -LO https://github.com/wneessen/apg-go/releases/download/v<version>/apg-v<version>-<os>-<architecture>.tar.gz.sha256
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
  PS> Invoke-RestMethod -Uri https://github.com/wneessen/apg-go/releases/download/v<version>/apg-v<version>-windows-<architecture>.zip -OutFile apg-v<version>-windows-<architecure>.zip
  PS> Invoke-RestMethod -Uri https://github.com/wneessen/apg-go/releases/download/v<version>/apg-v<version>-windows-<architecture>.zip.sha256 -OutFile apg-v<version>-windows-<architecure>.zip.sha256
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
  $ curl -LO https://github.com/wneessen/apg-go/archive/refs/tags/v<version>.tar.gz
  ```
* Extract source
  ```sh
  $ tar xzf v<version>.tar.gz
  ```
* Build binary
  ```sh
  $ cd apg-go-<version>
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

## Programmatic interface
Since v0.4.0 the CLI and the main package functionality have been separated from each other, which makes
it easier to use the `apg-go` package in other Go code as well. This way you can make of the password
generation in your own code without having to rely on the actual apg-go binary.

Code examples on how to use the package can be found in the [example-code](example-code) directory.

## Usage examples
### Default behaviour
By default apg-go will generate 6 passwords, with a minimum length of 12 characters and a 
maxiumum length of 20 characters. The generated password will use a character set constructed 
from lower case, upper case and numeric characters.
```shell
$ ./apg-go
R8rCC8bw5NvJmTUK2g
cHB9qogTbfdzFgnH
hoHfpWAHHSNa4Q
QyjscIsZkQGh
904YqsU5SnoqLo2w
utdFKXdeiXFzM
```
### Modifying the character sets
#### Old style
Let's assume you want to generate a single password, constructed out of upper case, numeric
and special characters. Since lower case is part of the default set, you would need to disable them
by setting the `-L` parameter. In addition you would set the `-S` parameter to enable special 
characters. Finally the parameter `-n 1` is needed to keep apg-go from generating more than one
password:
```shell
$ ./apg-go -n 1 -L -S
XY7>}H@5U40&_A1*9I$
```

#### New/modern style
Since the old style switches can be kind of confusing, it is recommended to use the "new style" 
parameters instead. The new style is all combined in the `-M` parameter. Using the upper case
version of a parameter argument enables a feature, while the lower case version disabled it. The
previous example could be represented like this in new style:
```shell
$ ./apg-go -n 1 -M lUSN
$</K?*|M)%8\U$5JA5~
```

#### Human readability
Generated passwords can sometimes be a bit hard to read for humans, especially when ambiguous 
characters are part of the password. Some characters in the ASCII character set look similar to 
each other. In example it can be hard to differentiate an upper case I from a lower case l. 
Same applies to the number zero (0) and the upper case O. To not run into issues with human 
readability, you can set the `-H` parameter to toggle on the "human readable" feature. When the
option is set, apg-go will avoid using any of the typical ambiguous characters in the generated
passwords.
```shell
$ ./apg-go -n 1 -M LUSN -H
YpranThY3b6b5%\6ARx
```

#### Character exclusion
Let's assume, that for whatever reason, your generated password can never include a colon (:) sign. For
this specific case, you can use the `-E` parameter to specify a list of characters that are to be excluded 
from the password generation character set:
```shell
$ ./apg-go -n 1 -M lUSN -H -E :
~B2\%E_|\VV|/5C7EF=
```

#### Complex passwords
If you want to generate complex passwords, there is a shortcut for this as well. By setting the `-C`
parameter, apg-go will automatically default to the most secure settings. The complex parameter 
basically implies that the password will use all available characters (lower case, upper case, 
numeric and special) and will make sure that human readability is disabled.
```shell
$ ./apg-go -n 1 -C
{q6cvz9le5_fo"X7
```

### Password length
By default, apg-go will generate a password with a random length between 12 and 20 characters. If you
want to be more specific, you can use the `-m` and `-x` parameters to override the defaults. Let's 
assume you want a single complex password with a length of exactly 32 characters, you can do so by
running:
```shell
$ ./apg-go -n 1 -C -m 32 -x 32
5lc&HBvx=!EUY*;'/t&>B|~sudhtyDBu
```

### Password spelling
If you need to read out a password, it can be helpful to know the corresponding word for that character in
the phonetic alphabet. By setting the `-l` parameter, agp-go will provide you with the phonetic spelling 
(english language) of your newly created password:
```shell
$ ./apg-go -n 1 -M LUSN -H -E : -l
fUTDKeFsU+zn3r= (foxtrot/Uniform/Tango/Delta/Kilo/echo/Foxtrot/sierra/Uniform/PLUS_SIGN/zulu/november/THREE/romeo/EQUAL_SIGN)
```

### Pronouncable passwords
Since v0.4.0 apg-go supports pronouncable passwords, anologous to the original c-apg using the `-a 0`
flag. The original c-apg implemented FIPS-181, which was withdrawn in 2015 for generating pronouncable
passwords. Since the standard is not recommended anymore, `apg-go` instead make use of the
[Koremutake Syllables System](https://shorl.com/koremutake.php). Similar to the original apg, `agp-go`
will automatically randomly add special characters and number (from the human-readable pool) to each
generated pronouncable password. Additionally it will perform a "coinflip" for each Koremutake syllable
and decided if it should switch the case of one of the characters to an upper-case character.

Using the `-t` parameter, `apg-go` will display a spelled out version of the pronouncable password, where
each syllable or number/special character is seperated with a "-" (dash) and if the syllable is not a
Koremutake syllable the character will be spelled out the same was as with activated `-l` in the
non-pronouncable password mode (`-a 1`).

**Note on password length**: The `-m` and `-x` parameters will work in prouncable password mode, but
please keep in mind, that due to the nature how syllables work, your generated password might exceed 
the desired length by one complete syllable (which can be up to 3 characters long).

**Security consideration:** Please keep in mind, that pronouncable passwords are less secure then truly
randomly created passwords, due to the nature how syllables work. As a rule of thumb, it is recommended
to multiply the length of your generated pronouncable passwords by at least 1.5 times, compared to truly
randomly generated passwords. It might also be helpful to run the pronoucable password mode with enabled
"[HIBP](#have-i-been-pwned)" flag, so that each generated password is automatically checked against "Have I Been Pwned" 
database.
```shell
$ ./apg-go -a 0 -n 1
KebrutinernMy

$ ./apg-go -a 0 -n 1 -m 15 -x 15 -t
pEnbocydrageT*En (pEn-bo-cy-dra-geT-ASTERISK-En)
```

### Have I Been Pwned
Even though, the passwords that apg-go generated for you, are secure, there is a minimal chance, that 
someone on the planet used exactly the same password before and that this person was part of an 
internet leak or hack, which exposed the password to the public. Such passwords are not considered 
secure anymore as they usually land on public available password lists, that are used by crackers.

To be on the safe side, you can use the `-p` parameter, to enable a HIBP check. When the feature is 
enabled, apg-go will check the HIBP database at https://haveibeenpwned.com if that password has been
leaked before and provide you with a warning if that is the case.

Please be aware, that this is a live check against the HIBP API, which not only requires internet
connectivity, but also might take between 500ms to 1s to complete. When you generating a bigger list
of password `-n 100`, the process could take much longer than without the `-p` feature enabled.

## CLI parameters
_apg-go_ replicates most of the parameters of the original c-apg. Some parameters are different though:

- `-a <algorithm>`: Choose password generation algorithm (Default: 1)
  - `0`: Pronouncable password generation (Koremutake syllables)
  - `1`: Random password generation according to password modes/flags
- `-m <length>`: The minimum length of the password to be generated (Default: 12)
- `-x <length>`: The maximum length of the password to be generated (Default: 20)
- `-n <number of passwords>`: The amount of passwords to be generated (Default: 6)
- `-E <list of characters>`: Do not use the specified characters in generated passwords
- `-M <[LUNSHClunshc]>`: New style password parameters (upper-case enables, lower-case disables)
- `-L`: Use lower-case characters in passwords (Default: on)
- `-U`: Use upper-case characters in passwords (Default: on)
- `-N`: Use numeric characters in passwords (Default: on)
- `-S`: Use special characters in passwords (Default: off)
- `-H`: Avoid ambiguous characters in passwords (i. e.: 1, l, I, o, O, 0) (Default: off)
- `-C`: Generate complex passwords (implies -L -U -N -S and disables -H) (Default: off)
- `-l`: Spell generated passwords in random password mode (Default: off)
- `-t`: Spell generated passwords in pronouncable password mode (Default: off)
- `-p`: Check the HIBP database if the generated passwords was found in a leak before (Default: off) // *this feature requires internet connectivity*
- `-h`: Show a CLI help text
- `-v`: Show the version number

## Contributors
Thanks to the following people for contributing to the apg-go codebase:
* [Romain Tartière](https://github.com/smortex)
* [Abraham Ingersoll](https://github.com/aberoham)

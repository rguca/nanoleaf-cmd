# nanoleaf-cmd
nanoleaf command line utility

## Usage

```
nanoleaf-cmd [effect|color temp] [brightness]
```
Call without arguments toggles on/off.
Call with unknown effect lists effects.

Color temp range 1200k-6500k.

Brightness range 0-100.

```
nanoleaf-cmd -p [<profile-index>]
```
If the profile-index is omitted, profiles are listed with their index.

### Config

Edit config file under config/nanoleaf-cmd.yml.

## Dependencies

-	github.com/adnanbrq/nanoleaf
-	github.com/jinzhu/configor

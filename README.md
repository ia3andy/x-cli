# XD CLI

This tiny tool wrap all the complex commands you always use in a project and make it easy to find/use them:
```
$ xd help                               // Show me the commands I always use in this project
...
COMMANDS:
   start     Execute the command: java -jar target/getting-started-1.0-SNAPSHOT-runner.jar
   dev       Execute the command: ./mvnw compile quarkus:dev
   ext:add   Execute the command: mvn quarkus:add-extension -Dextensions="${name}"
   ext:list  Execute the command: mvn quarkus:list-extensions
   package   Execute the command: ./mvnw package
   help, h   Shows a list of commands or help for one command
...
```

It's close to what's provided by **npm scripts** or **makefile**. The key differences are:
- lightness and simplicity
- straight forward (no learning needed)
- params with default values
- cross-language
- cross-platform

## Install

Download the latest release and add it to your path, or add it to your project as part of your sources:
https://github.com/ia3andy/xd/releases

## Configuration file

The configuration file can be in `json` or `yaml`, it should be named `.xd.yaml`, `.xd.yml` or `.xd.json`.

You just need a `script` field with a map of name and command.

You can specify available params in the form of `${param1}` and you can specify a default value `${param1|Some default value}`

### Example:

```yaml
--- 
scripts:
  dev: ./mvnw compile quarkus:dev
  ext:add: mvn quarkus:add-extension -Dextensions="${name}"
  ext:list: mvn quarkus:list-extensions
  package: ./mvnw package
  start: java -jar target/getting-started-1.0-SNAPSHOT-runner.jar
```

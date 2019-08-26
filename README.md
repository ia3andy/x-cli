# XD CLI

This tiny tool wrap all the complex commands you always use in a project:

```
$ xd help                               // Show me the commands I always use in this project
...
COMMANDS:
   deploy     Execute the command: i-hate-this-command --deploy --bin ./target/my-bin --target-host ${to|192.168.1.1}
   [other commands ...]
...
$ xd deploy --to 192.168.1.4
Running command: i-love-this-tool --deploy --bin ./target/my-bin --target-host 192.168.1.4
[command output ...]
$ xd deploy --to 192.168.1.4
Running command: i-love-this-tool --deploy --bin ./target/my-bin --target-host 192.168.1.1
[command output ...]
```

It really looks like **npm scripts**, but this is a **cross-language** tool.

## Install

Download the latest release and add it to your path, or publish it as part of your sources:
https://github.com/ia3andy/xd/releases

## Configuration file

The configuration file can be in `json` or `yaml`, it should be named `.xd.yaml`, `.xd.yml` or `.xd.json`.

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

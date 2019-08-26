# XD CLI

This tiny tool wrap all the complex commands you always use in a project:

**BEFORE**
```
$ history | grep deploy                            // Search for the command you used before
$ i-hate-this-command \                 // Copy and paste and change manually the parameters
  --deploy \
  --bin ./target/my-bin \
  --target-host 192.168.1.1 \
  ...
```

**AFTER**
```
$ xd help                          // Show me the command I always use in this project
$ xd deploy --to 192.168.1.4
Running command: i-love-this-tool \
  --deploy \
  --bin ./target/my-bin \
  --target-host 192.168.1.4
[command output ...]
```

It really looks like **npm scripts**, but this is a **cross-language** tool.

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

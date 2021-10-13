# path-part: print the interesting part from stdin to stdout

This program was written because the widely distributed classic unix `basename` and `dirname` programs expect data input as command line arguments, which composes poorly with unix pipelines and is generally an eyesore in a suite of tools whith mostly a consistent design. After years of working around this with `xargs` I finally decided to just write them the proper way.

- input comes on standard input. It's newline-separated filesystem paths.
- output goes on standard output. Again newline-separated.
- options are provided as command line arguments.

## Usage

`path-part {base|extension|extensions|path|name} < my-file-list.txt`

`path-part` expects a single command line argument, which describes the path component you want to keep, discarding implicitly the other parts of the path. It is an enumeration of valid values as follows:

* `base` Print the final component without any extensions
* `extension` Print the final period-delimited extension
  * This command is also aliased as `ext`
* `extensions` Print all period-delimited filename extensions
  * This command is also aliased as `exts`
* `name` Similar to `basename`. Print only the final path component.
  * This command is also aliased as `last` and `basename`
* `path` Similar to `dirname`. Strip the final component and print the remainder.
  * This command is also aliased as `directory`, `dir`, or `dirname'

## Examples

```
find examples -type f | path-part name
file1.txt
file3.markdown
file2.tar.bz2
file11.yml
file 20.png
file10.yaml
file12.json.crdownload
```


```
find examples -type f | ./path-part path
examples
examples
examples
examples/subdir1
examples/subdir1/subdir2
examples/subdir1
examples/subdir1
```

```
find examples -type f | ./path-part extensions
txt
markdown
tar.bz2
yml
png
yaml
json.crdownload
```

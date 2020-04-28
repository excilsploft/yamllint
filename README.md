# yamllint

A [GO](https://golang.org/) based simple cli app that attempts to verify a YAML file is valid.


This does not verify the "Schema" of a YAML file that would difficult without a defined Schema, what this does is simply take a yaml file in StdIn or traverses a path for YAML files and attempt to open
parse them, additionally, it will show duplicate keys if the top level type is a JSON Object/YAML Object


## usage

verbose from stdin
	`yamllint -v < some.yaml`
or
	`cat some.yaml | yamllint -v`

not verbose
	`yamllint ./some.yaml`


some directory or directories filled with yaml (this recusively descends into it)
	`yamllint ./some_directory/`

verbosity flag will show each file processed

## install

get the latest release from the [GitHub Release Page](https://github.com/excilsploft/yamllint/releases)

Download the zip for your platform:
OS                   | Link
---------------------|-------------------------------------------------------------------------------------------------------------------
 Linux (any distro)  | [linux-amd64-yamllint.zip](https://github.com/excilsploft/yamllint/releases/download/v1.1/linux-amd64-yamllint.zip)
 MacOsx              | [darwin-amd64-yamllint.zip](https://github.com/excilsploft/yamllint/releases/download/v1.1/darwin-amd64-yamllint.zip)
 Windows             | [windows-amd64-yamllint.zip](https://github.com/excilsploft/yamllint/releases/download/v1.1/windows-amd64-yamllint.zip)


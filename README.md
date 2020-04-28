# yamllint

A (GO)[https://golang.org/] based simple cli app that attempts to verify a YAML file is valid.


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

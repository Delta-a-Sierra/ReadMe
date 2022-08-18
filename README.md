# ReadMe

A simple command-line readme generator.

## Description

A command line tool that's used to generate readmes from templates as well as create custom templates.

## Agruments

### template

- use
  - accepts one positonal agrument the template name
- add
  - accepts postional agruments the desired name of the template and the file location.
  - tag flag can be used to supply list of tags to associate with the template
- remove
  - accepts one agrument the template name
- list
  - no agruments lists all templates sorted a-z.
  - tag flag can be used to list only templates of certain tag.
  - sort tag can be used to sort results in alternate ways, most-used.

### template tags

- list
- add
- remove
  - one postional agrument the tag name, displays number of templates attached to it. before confirming delete.
- rename
  - two postional arguments current tag name and desired name, no duplicate tags

<!--
### custom

- create

### custom sections

- list
  - preview flag to show the markdown for each section.
- add
  - default to interactive mode
- remove
  - takes one argument a positional argument. -->

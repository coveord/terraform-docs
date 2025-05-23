## terraform-docs asciidoc

Generate AsciiDoc of inputs and outputs

### Synopsis

Generate AsciiDoc of inputs and outputs

```
terraform-docs asciidoc [PATH] [flags]
```

### Options

```
  -h, --help         help for asciidoc
      --indent int   indention level of AsciiDoc sections [1, 2, 3, 4, 5] (default 2)
      --required     show Required column or section (default true)
      --sensitive    show Sensitive column or section (default true)
```

### Options inherited from parent commands

```
  -c, --config string               config file name (default ".terraform-docs.yml")
      --header-from string          relative path of a file to read header from (default "main.tf")
      --hide strings                hide section [header, inputs, outputs, providers, requirements]
      --hide-all                    hide all sections (default false)
      --output-values               inject output values into outputs (default false)
      --output-values-from string   inject output values from file into outputs (default "")
      --show strings                show section [header, inputs, outputs, providers, requirements]
      --show-all                    show all sections (default true)
      --sort                        sort items (default true)
      --sort-by-required            sort items by name and print required ones first (default false)
      --sort-by-type                sort items by type of them (default false)
```

### SEE ALSO

* [terraform-docs asciidoc document](/docs/formats/asciidoc-document.md)	 - Generate AsciiDoc document of inputs and outputs
* [terraform-docs asciidoc table](/docs/formats/asciidoc-table.md)	 - Generate AsciiDoc tables of inputs and outputs

###### Auto generated by spf13/cobra on 25-Nov-2020

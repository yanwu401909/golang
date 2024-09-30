module test1

go 1.22.1

require nazul.top/projects/test1/utils v0.0.1

replace nazul.top/projects/test1/utils v0.0.1 => ./utils

require nazul.top/projects/test1/model v0.0.1

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/testify v1.9.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace nazul.top/projects/test1/model v0.0.1 => ./model

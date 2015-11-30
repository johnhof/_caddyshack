## Gorilla Example

Basic [Gorilla](http://www.gorillatoolkit.org/) driven server

## Directory structure

```bash
.
├── app.go # main file
├── configuration
│   └── index.json # catch all configuration
├── router # router package
│   └── router.go # route-controller mapping functionality
└── controllers
    ├── core # core controller package
    │   └── read.go # GET core handlers
    └── note # note controller package
        └── read.go # GET note handlers
```

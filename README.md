# Rack
A simple command-line utility to keep your files and directories organized. It cleanly copies the files based on their types into the directories you want.
## Get Started
1. Clone the repo

2. Build it by running ```go build```

## Usage
```bash
rack source-dir target-dir ( --type <filetype> | --allfiles )
``` 
`--type` flag specifies the type of file to be copied from `source-dir` to `target-dir`

`--allfiles` specifying it moves all the files to its own directory based on the file type.

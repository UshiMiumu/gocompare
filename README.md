# GoCompare

Simple program to compare the files & directories within two different directories.

It takes a source directory and a destination directory and then prints out a list of files & directories that are in the source directory but not in the destination.

It does NOT does any kind of comparison such as hashing of the files themselves, only checks the names. It was made purely so that I could confirm all the files had been copied over to a new location after a large copy/paste

## Usage

./gocompare -src "source path" -dst "destination path"
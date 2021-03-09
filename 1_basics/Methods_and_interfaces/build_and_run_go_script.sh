#! /bin/bash

main () {
	echo "use 'gohr main'"
	exit 0

  # built_file=$1
  built_file="main"
  go build -o $built_file
  ./$built_file
}

main "$1"
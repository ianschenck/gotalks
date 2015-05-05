#!/bin/sh
readonly container=$1
readonly element=$2
readonly input=$3
gofmt -r "TYPE -> $container" $input \
	 | gofmt -r "ELEMENT_TYPE -> $element" \
	 | tail -n +3 \
			  > ${container}_${element}_${input}

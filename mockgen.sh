#!/bin/sh

file=$1

echo ${file}

dir=`dirname ${file}`/mock
base=`basename ${file}`
mkdir -p ${dir}
mockfile=${dir}/${base}

rm ${mockfile}
mockgen -source ${file} > ${mockfile}
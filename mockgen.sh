#!/bin/sh

pkg=$1
file=$2

mockfile=`echo ${file} | sed s/.go$/_mock.go/g`
mockgen -source ${file} -package `basename ${pkg}` -self_package ${pkg} > ${mockfile}
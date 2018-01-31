#!/bin/sh
set -e
mkdir -p classes
export CLASSPATH=$(pwd)/$(ls lib/*.jar | tr '\n' ':')
set -x
/usr/lib/jvm/java-1.8.0/bin/javac -cp "$CLASSPATH" -sourcepath src/main/java/ -d classes src/main/java/hu/unosoft/pst2eml/Dump.java
set -x
java -cp "$CLASSPATH":classes pst2eml.Dump "$@"

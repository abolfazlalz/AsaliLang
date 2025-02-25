#!/bin/sh

FILE=antlr.jar

if [ ! -f "$FILE" ]; then
    curl https://www.antlr.org/download/antlr-4.13.0-complete.jar -o $FILE
fi
alias antlr4='java -Xmx500M -cp "$FILE:$CLASSPATH" org.antlr.v4.Tool'
antlr4 -Dlanguage=Go -visitor -package parsing *.g4 -o ../go/parsing
antlr4 -Dlanguage=Python2 *.g4 -o ../python/parsing -visitor
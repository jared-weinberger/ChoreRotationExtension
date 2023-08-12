#!/bin/sh
alias antlr4='java -cp "./antlr-4.13.0-complete.jar:$CLASSPATH" org.antlr.v4.Tool'
antlr4 -Dlanguage=Go -no-visitor -package parsing -o ../parsing *.g4
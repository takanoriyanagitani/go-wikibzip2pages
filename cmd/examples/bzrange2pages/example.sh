#!/bin/sh

export WIKIPAGE_FILE=~/Downloads/enwiki-20250801-pages-articles-multistream.xml.bz2
export WIKIPAGE_OFFSET=14034442273
export WIKIPAGE_SIZE=114781

./bzrange2pages | grep Hello

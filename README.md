# ad - add delimiter

simple command line tool which adds a delimiter to output after it has been 'silent' for a couple of seconds

## usage:

    Usage of 'ad':
      -d string
          the delimiter to be printed out (default "#########################################")
      -n int
          number of times to print delimiter (default 3)
      -t int
          number of seconds to wait before printing the delimiter (default 3)

## example:

    dmesg --follow | bin/ad

## screencast:

https://youtu.be/98iLeZiRxwI

## changelog

* 2026-09-26 first working version
----


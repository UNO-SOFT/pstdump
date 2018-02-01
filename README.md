# pstdump
As I have some .pst files what [readpst](http://www.five-ten-sg.com/libpst/rn01re01.html) cannot dump,
I had to use [java-libpst](https://github.com/rjohnsondev/java-libpst) to parse the .pst
and dump it as JSON elements (with [gson](https://github.com/google/gson)).

The JSON is parsed with Go using the `github.com/UNO-SOFT/pstdump/parse` package.

	go get github.com/UNO-SOFT/pstdump

will create a `pstdump` executable, which can parse the output of the `dump` shell script,
which dumps the given .pst:

	./dump pesky.pst | pstdump dest_dir

This will dump the contents of `pesky.pst` under `dest_dir`, at least the .eml (message/rfc822) files.

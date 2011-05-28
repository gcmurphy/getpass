include $(GOROOT)/src/Make.inc

TARG=github.com/gcmurphy/getpass
CGOFILES=\
	getpass.go\

include $(GOROOT)/src/Make.pkg
%: install %.go
	$(GC) $*.go
	$(LD) -o $@ $*.$O

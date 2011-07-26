include $(GOROOT)/src/Make.inc

TARG=ttime
GOFILES=\
	*.go\

include $(GOROOT)/src/Make.pkg

fmt:
	gofmt -w $GOFILES

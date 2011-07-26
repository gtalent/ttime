include $(GOROOT)/src/Make.inc

TARG=gtime
GOFILES=\
	*.go\

include $(GOROOT)/src/Make.pkg

fmt:
	gofmt -w $GOFILES

TOPTARGETS := all clean images build test

SUBDIRS := $(wildcard *-aware-scheduling/.)

$(TOPTARGETS): $(SUBDIRS)
$(SUBDIRS):
	$(MAKE) -C $@ $(MAKECMDGOALS)

.PHONY: $(TOPTARGETS) $(SUBDIRS)

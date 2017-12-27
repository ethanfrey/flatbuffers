.PHONY: tools flatc install demo benchmark flatb-Windows flatb-Linux flatb-Darwin

flatc:
	flatc -g myschema.fbs

install:
	@ go install ./cmd/fbdemo

demo: install
	@ fbdemo

deps: glide
	@ glide install

glide:
	@ go get github.com/Masterminds/glide

# benchmark:
# 	cd benchmarks && go test -bench=.


### cross-platform check for installing flatb ###

OS := $(shell uname -s)
VERSION := 1.8.0
ZIP := v$(VERSION).tar.gz
FBDIR := flatbuffers-$(VERSION)
FBBIN := $(HOME)/bin/flatc

/usr/local/bin/flatc:
	which cmake || sudo apt-get install cmake
	@ curl -L https://github.com/google/flatbuffers/archive/$(ZIP) > $(ZIP)
	@ tar xzf $(ZIP)
	@ rm $(ZIP)
	@ cd $(FBDIR) && cmake -G "Unix Makefiles" && make flatc
	sudo mv $(FBDIR)/flatc /usr/local/bin/
	@ rm -rf $(FBDIR)

install-fb-Linux: /usr/local/bin/flatc

install-fb-Darwin:
	echo "Only works if homebrew is installed..."
	brew update
	brew install flatbuffers

tools: install-fb-$(OS) glide
	@ which flatc > /dev/null || echo please add flatc to $PATH


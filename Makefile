GB := gb
BINDATA := go-bindata-assetfs


all: src/webui-aria2/compiled.go
	@$(GB) build


src/webui-aria2/compiled.go: webui-aria2/*
	@$(BINDATA) webui-aria2/...
	@mv bindata_assetfs.go $@

.PHONY: all app

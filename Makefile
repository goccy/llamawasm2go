LLAMA_WASM_REPO     ?= goccy/llama-wasm
LLAMA_WASM_VERSION  ?= v0.4.0
# llama-wasm emits its release attestations from release.yml (the v* tag
# workflow), NOT build.yml — releasing lives only in release.yml there.
LLAMA_WASM_WORKFLOW ?= goccy/llama-wasm/.github/workflows/release.yml

TARBALL          := llama_wasm2go.tar.gz
SHA256SUMS       := llama_wasm2go.sha256
RELEASE_URL       = https://github.com/$(LLAMA_WASM_REPO)/releases/download/$(LLAMA_WASM_VERSION)
ATTESTATION_API   = https://api.github.com/repos/$(LLAMA_WASM_REPO)/attestations

.PHONY: bundle download verify verify-release verify-attestation verify-tree link-check build

## bundle: download release artifacts, verify their attestations, check
## the tree against the manifest, and link a consumer per asm target.
## Run this whenever LLAMA_WASM_VERSION bumps.
bundle: download verify verify-tree link-check

## download: fetch the wasm2go tarball + sha256 manifest and extract
## them into the repo root. The tarball is discarded once unpacked.
download:
	curl -fSL --proto '=https' --tlsv1.2 -o $(TARBALL)    $(RELEASE_URL)/$(TARBALL)
	curl -fSL --proto '=https' --tlsv1.2 -o $(SHA256SUMS) $(RELEASE_URL)/$(SHA256SUMS)
	tar xzf $(TARBALL)
	rm -f $(TARBALL)

## verify: byte-check each in-tree file against the sha256 manifest AND
## confirm each carries a valid GitHub artifact attestation signed by
## the upstream release.yml workflow. Either check failing aborts.
verify: verify-release verify-attestation

## verify-release: confirm every in-tree file matches the entries in
## $(SHA256SUMS). Fast sanity check; not a trust anchor on its own.
verify-release:
	@echo "==> verifying in-tree files against $(SHA256SUMS)"
	@shasum -a 256 -c $(SHA256SUMS)

## verify-attestation: confirm every in-tree artifact is a signed
## subject of the upstream SLSA build attestation. The release emits one
## attestation whose subject list (via subject-checksums) covers every
## file in the manifest, so we fetch the bundle once anonymously from the
## public attestation API and then offline-verify each file via
## `gh attestation verify --bundle`. No GH access token is required.
verify-attestation:
	@set -eu; \
	tmpdir=$$(mktemp -d); \
	bundle=$$tmpdir/bundle.jsonl; \
	trap 'rm -rf $$tmpdir' EXIT; \
	probe=$$(awk 'NR==1 {print $$2}' $(SHA256SUMS) | sed 's|^\./||'); \
	digest=$$(shasum -a 256 $$probe | awk '{print $$1}'); \
	echo "==> fetching attestation bundle via $$probe (sha256:$$digest)"; \
	curl -fsSL --proto '=https' --tlsv1.2 \
	  "$(ATTESTATION_API)/sha256:$$digest" \
	  | jq -c '.attestations[].bundle' > $$bundle; \
	files=$$(awk '{print $$2}' $(SHA256SUMS) | sed 's|^\./||'); \
	for f in $$files; do \
	  echo "==> verifying $$f"; \
	  GH_TOKEN= GITHUB_TOKEN= gh attestation verify "$$f" \
	    -R $(LLAMA_WASM_REPO) \
	    --bundle $$bundle \
	    --signer-workflow $(LLAMA_WASM_WORKFLOW); \
	done

## verify-tree: the manifest and the tracked tree must describe the
## same file set. `tar xzf` never deletes: a file the previous bundle
## shipped and the new one dropped would silently survive extraction,
## and `shasum -c` only checks files the manifest lists. Every tracked
## file outside the repo's own scaffolding must be in the manifest, and
## every manifest entry must exist.
TREE_EXEMPT := .gitignore .github/% Makefile README.md LICENSE llama_wasm2go.sha256 scripts/%
verify-tree:
	@set -eu; \
	man=$$(mktemp); trk=$$(mktemp); trap 'rm -f $$man $$trk' EXIT; \
	awk '{print $$2}' $(SHA256SUMS) | sed 's|^\./||' | sort > $$man; \
	git ls-files | sort > $$trk; \
	stale=$$(comm -13 $$man $$trk | grep -v -E '^($(subst %,.*,$(subst $(eval) ,|,$(TREE_EXEMPT))))$$' || true); \
	missing=$$(comm -23 $$man $$trk); \
	if [ -n "$$stale" ]; then echo "tracked files not in $(SHA256SUMS) (stale from a previous bundle?):"; echo "$$stale"; exit 1; fi; \
	if [ -n "$$missing" ]; then echo "manifest entries missing from the tree:"; echo "$$missing"; exit 1; fi; \
	echo "==> tree matches $(SHA256SUMS) ($$(wc -l < $$man | tr -d ' ') files)"

## link-check: link a consumer binary against this bundle for every asm
## target (linux/arm64, linux/amd64 v2, linux/amd64 v1). `go build ./...`
## does not run the linker's nosplit / ABI-wrapper resolution; a bundle
## that compiles everywhere can still fail to link (v0.3.1 did, on the
## cross-chunk direct-call layout). Cross-compiles only; nothing runs.
link-check:
	@set -eu; dir=$$(mktemp -d); trap 'rm -rf $$dir' EXIT; \
	scripts/bundle-link-consumer.sh "$(CURDIR)" "$$(awk '/^module /{print $$2}' go.mod)" "$$dir"

## build: sanity-build every package (compile only — see link-check for
## the link-time checks).
build:
	go build ./...

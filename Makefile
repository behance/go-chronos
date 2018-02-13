build:
	go build ./...

test:
	golint $(shell glide nv)
	go vet $(shell glide nv)
	ginkgo -r -trace -failFast -v --cover --randomizeAllSpecs --randomizeSuites -p
	echo "" && for i in $$(ls **/*.coverprofile); do echo "$${i}" && go tool cover -func=$${i} && echo ""; done

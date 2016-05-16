.PHONY: test compile package clean

test:
	go test -v -cover github.com/ihcsim/go-soundex/...

compile: test clean
	go build -v -o bin/soundex github.com/ihcsim/go-soundex/cmd/cli/...
	go build -v -o bin/soundex-web github.com/ihcsim/go-soundex/cmd/web/...	

package: compile
	docker run --rm \
		-v `PWD`/bin:/workspace \
		-w /workspace \
		isim/fpm \
		-t deb \
		-s dir \
		-n soundex \
		-v 1.0.0 \
		--maintainer "Ivan Sim" \
		--description "A Go Soundex algorithm program" \
		--prefix=/usr/bin \
		--deb-no-default-config-files \
		--license="Apache License Version 2.0 http://www.apache.org/licenses/" \
		soundex

	docker run --rm \
		-v `PWD`/bin:/workspace \
		-w /workspace \
		isim/fpm \
		-t deb \
		-s dir \
		-n soundex-web \
		-v 1.0.0 \
		--maintainer "Ivan Sim" \
		--description "A Go Soundex algorithm web application" \
		--prefix=/usr/bin \
		--deb-no-default-config-files \
		--license="Apache License Version 2.0 http://www.apache.org/licenses/" \
		soundex-web

clean:
	rm -rf bin

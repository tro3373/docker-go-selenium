CONTAINERNAME=golang
all_container=$$(docker ps -a -q)
active_container=$$(docker ps -q)
images=$$(docker images | awk '/^<none>/ { print $$3 }')
dld=src/downloads/.seluser

default: build
all: build
build:
	docker-compose build

up: start
start:
	[[ ! -e ${dld} ]] && mkdir -p ${dld}; \
	sudo chown 1000.1000 -R ${dld}; \
	docker-compose up -d && docker-compose logs
stop: down
down:
	docker-compose down
restart: stop start

console: attach
attach:
	docker exec -it $(CONTAINERNAME) /bin/bash
do:
	# docker exec -it $(CONTAINERNAME) /works/sample
	docker exec $(CONTAINERNAME) go run /works/main.go

logs:
	docker-compose logs
logsf:
	docker-compose logs -f

clean: clean_container clean_images
clean_images:
	@if [ "$(images)" != "" ] ; then \
		docker rmi $(images); \
	fi
clean_container:
	@for a in $(all_container) ; do \
		for b in $(active_container) ; do \
			if [ "$${a}" = "$${b}" ] ; then \
				continue 2; \
			fi; \
		done; \
		docker rm $${a}; \
	done

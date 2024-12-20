install:
	go install -v
	
.PHONY: install

validate:
		swagger validate ./swagger/swagger.yml

gen: validate
	swagger generate server \
		--target=./swagger/server \
		--spec=./swagger/swagger.yml \
		--exclude-main \
		--name=pap
		
.PHONY: install gen validate

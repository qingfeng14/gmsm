PACKAGE=gmsm
UNITPACKAGE=gmsm/sm2 \
			gmsm/sm3 \
			gmsm/sm4 \
			gmsm/tests
COUNT=1
# Go parameters
GOCMD=go
GOINSTALL=$(GOCMD) install
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get
GOTEST=$(GOCMD) test -count=$(COUNT) -timeout=20m \
		$(SYMBOL)




# all: install
all: test

# test: unit test
test: unit

# unit: unit test and coverage test
unit:
	@echo "unit test"
	@for pkg in $(UNITPACKAGE); do \
		$(GOTEST) $$pkg -cover; \
	done

bench:
	@$(GOTEST)  -bench=. -benchmem -cpu=1,2,4 gmsm/tests
	# @$(GOTEST)  -bench=. -benchmem gmsm/tests

cpu:
	@go test -benchmem -run=^$ gmsm/tests -bench ^BenchmarkTJVerify$ -memprofile=mem.out -cpuprofile=cpu.out

cpuprof:
	@go tool pprof tests.test cpu.out

memprof:
	@go tool pprof -alloc_space tests.test mem.out

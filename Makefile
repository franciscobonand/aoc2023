day1-%:
	@echo "Running day 1, file ${@:day1-%=%}"
	@cd day1 && go run . ${@:day1-%=%}
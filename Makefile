day1-%:
	@echo "Running day 1, file ${@:day1-%=%}"
	@cd day1 && go run . ${@:day1-%=%}

day2-%:
	@echo "Running day 2, file ${@:day2-%=%}"
	@cd day2 && go run . ${@:day2-%=%}
	
day3-%:
	@echo "Running day 3, file ${@:day3-%=%}"
	@cd day3 && go run . ${@:day3-%=%}
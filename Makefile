go.day1-%:
	@echo "Running day 1, file ${@:go.day1-%=%}"
	@cd go/day1 && go run . ${@:go.day1-%=%}

go.day2-%:
	@echo "Running day 2, file ${@:go.day2-%=%}"
	@cd go/day2 && go run . ${@:go.day2-%=%}
	
go.day3-%:
	@echo "Running day 3, file ${@:go.day3-%=%}"
	@cd go/day3 && go run . ${@:go.day3-%=%}

go.day4-%:
	@echo "Running day 4, file ${@:go.day4-%=%}"
	@cd go/day4 && go run . ${@:go.day4-%=%}

go.day5-%:
	@echo "Running day 5, file ${@:go.day5-%=%}"
	@cd go/day5 && go run . ${@:go.day5-%=%}
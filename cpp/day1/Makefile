.PHONY: day1 clear

# TO-DO: Support other compilers than clang
day1: main.o CalibrationInterpreter.o
	clang++ --std=c++17 main.o CalibrationInterpreter.o -o day1

# TO-DO: Before the linking stage maybe list all object files with *
# TO-DO: Support other compilers than clang
main.o: main.cpp
	clang++ --std=c++17 -c main.cpp

CalibrationInterpreter.o: ./AuxFunctions/CalibrationInterpreter.cpp
	clang++ --std=c++17 -c ./AuxFunctions/CalibrationInterpreter.cpp


clear:
	rm -rf *.o
	rm -rf day1
.PHONY: day2 clear

# Same TO-DOs from previous day are valid here
# Change the possibilities for a compiler
# Use pattern matching to compile cpp into object files better

day2: main.o ElfSackInterpreter.o
	clang++ main.o ElfSackInterpreter.o -o day2

main.o: main.cpp 
	clang++ --std=c++17 -c main.cpp

ElfSackInterpreter.o: ./AuxFunctions/ElfSackInterpreter.cpp
	clang++ --std=c++17 -c ./AuxFunctions/ElfSackInterpreter.cpp

clear:
	rm -rf *.o
	rm -rf day2
#include <iostream>
#include "AuxFunctions/CalibrationInterpreter.hpp"

int main(int argc, char* argv[])
{
    // We expect to receive two values in the terminal, program and file name
    if (argc < 2)
        return 1;
    CalibrationInterpreter interp(argv[1]);
    interp.computeSumOfCalibrationValues();
    return 0;
}
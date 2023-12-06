import readline from 'readline';

function extractCalibrationValues(input) {
    const lines = input.split('\n');
    let sum = 0;

    lines.forEach(line => {
        const digits = line.match(/\d/g);
        if (digits && digits.length >= 2) {
            const calibrationValue = parseInt(digits[0] + digits[digits.length - 1]);
            sum += calibrationValue;
        }
        if (digits && digits.length == 1) {
            const calibrationValueUnique = parseInt(digits + digits);
            sum += calibrationValueUnique;
        }
    });

    return sum;
}

function readInput() {
    const rl = readline.createInterface({
        input: process.stdin,
        output: process.stdout
    });

    let userInput = '';

    return new Promise(resolve => {
        console.log("Enter multiple lines of text. Press Enter on an empty line to finish.");

        rl.on('line', line => {
            if (line.trim() === '') {
                rl.close();
                resolve(userInput);
            } else {
                userInput += line + '\n';
            }
        });
    });
}

async function main() {
    const userInput = await readInput();
    const totalCalibration = extractCalibrationValues(userInput.trim());
    console.log('Total Calibration:', totalCalibration);
}

main();

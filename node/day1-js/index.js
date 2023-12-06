import fs from 'fs';

function extractCalibrationValues(input) {
    const lines = input.split('\n');
    let sum = 0;

    lines.forEach(line => {
        const digits = line.match(/\d/g);

        if (digits && digits.length >= 2) {
            const calibrationValue = parseInt(digits[0] + digits[digits.length - 1]);
            sum += calibrationValue;
        } else if (digits && digits.length >= 1) {
            const calibrationValue = parseInt(digits.join('') + digits.join(''));
            sum += calibrationValue;
        }
    });

    return sum;
}

function readFile(filePath) {
    try {
        const fileContent = fs.readFileSync(filePath, 'utf8');
        return fileContent;
    } catch (error) {
        console.error('Error reading the file:', error.message);
        process.exit(1);
    }
}

const filePath = './input.txt'; 
const fileContent = readFile(filePath);
const totalCalibration = extractCalibrationValues(fileContent.trim());
console.log('Total Calibration:', totalCalibration);

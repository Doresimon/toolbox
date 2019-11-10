const sharp = require('sharp');
const fs = require("fs")

main()

function main() {
    const PATH_TO_FILE = "./tmp/mei.jpg"
    const buffer = fs.readFileSync(PATH_TO_FILE);


    sharp(buffer)
        .resize({
            width: 200,
            // height: 200,
            fit: sharp.fit.inside,
            position: sharp.strategy.entropy
        })
        .toFile('output.jpg', (err, info) => {
            console.log(err)
            console.log(info)
        });
}


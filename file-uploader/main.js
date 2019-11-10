const FormData = require('form-data');
const fs = require("fs")
const axios = require("axios")

main()

function main() {
    const PATH_TO_FILE = "./tmp/fake.file"
    const form = new FormData();
    const stream = fs.createReadStream(PATH_TO_FILE);

    form.append('image', stream);
    form.append('xxxxxxxxxxxxxx', value = "value");

    // In Node.js environment you need to set boundary in the header field 'Content-Type' by calling method `getHeaders`
    const formHeaders = form.getHeaders();

    console.log("formheaders", formHeaders)
    console.log("form", form)
    const url = "http://localhost:80/api/test/file_upload"
    axios.post(url, form, { headers: { ...formHeaders, }, })
        .then(response => response)
        .catch(error => error)
}


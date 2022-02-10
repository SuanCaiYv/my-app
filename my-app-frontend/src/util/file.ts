const uploadLocalFile = function (filePath: string): File {
    let rawFile = new XMLHttpRequest();
    rawFile.open("GET", filePath, false);
    rawFile.onreadystatechange = function () {
        if (rawFile.readyState === 4) {
            if (rawFile.status === 200 || rawFile.status == 0) {
                console.log(rawFile.responseText.length)
                return new File(rawFile.responseText, filePath.substring(filePath.lastIndexOf("/") + 1, filePath.length))
            }
        }
        return null
    }
    rawFile.send(null);
}

export {uploadLocalFile}
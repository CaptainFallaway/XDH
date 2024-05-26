const dropArea = document.getElementById('drop-area');

function handleFile(files) {
    console.log(files);
}

dropArea.addEventListener('change', (e) => {
    handleFile(e.target.files);
});

dropArea.addEventListener('dragover', (e) => { e.preventDefault() });

dropArea.addEventListener('drop', (e) => {
    e.preventDefault();
    handleFile(e.dataTransfer.files);
});
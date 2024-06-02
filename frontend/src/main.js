import { OpenFileDialog } from "../wailsjs/go/app/App.js";

const dropArea = document.getElementById('drop-area');

dropArea.addEventListener('click', async (e) => {
    e.preventDefault();
    await OpenFileDialog();
});

// THIS WILL BE IMPLEMENTED IN WAILS VERSION 3
// dropArea.addEventListener('dragover', (e) => { e.preventDefault() });

// dropArea.addEventListener('drop', (e) => {
//     e.preventDefault();
//     handleFile(e.dataTransfer.files);
// });
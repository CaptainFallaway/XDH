import { GetModels, OpenFileDialog, GetDropArea } from "../wailsjs/go/app/App.js";
import { EventsOn } from "../wailsjs/runtime/runtime.js";

const content = document.getElementById('content');
var sortingMetal = "Sn";

EventsOn('modelLoaded', async () => loadModels())
loadDropArea()

document.querySelectorAll('input[name="sortingMetal"]').forEach((radio) => {
    radio.addEventListener('change', async (e) => {
        sortingMetal = e.target.id
        await loadModels()
    })
});

async function loadDropArea() {
    content.innerHTML = await GetDropArea()
    document.getElementById('drop-area').addEventListener('click', async (e) => {
        e.preventDefault();
        await OpenFileDialog();
    });
}

async function loadModels() {
    content.innerHTML = await GetModels(sortingMetal)
}

// THIS WILL BE IMPLEMENTED IN WAILS VERSION 3
// dropArea.addEventListener('dragover', (e) => { e.preventDefault() });

// dropArea.addEventListener('drop', (e) => {
//     e.preventDefault();
//     handleFile(e.dataTransfer.files);
// });
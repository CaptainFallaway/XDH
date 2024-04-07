/// This is supposed to be like a lightweight ViewModel

import {GetModels} from '../wailsjs/go/app/ModelInterface';

// Helper function to bind events to elements
function bindEvent(id, event, callback) {
    let element = document.querySelector(id);
    if (!element) {
        console.error("Element not found: " + id);
        return;
    }
    element.addEventListener(event, callback);
}

function swapInner(id, content) {
    let element = document.getElementById(id);
    if (!element) {
        console.error("Element not found: " + id);
        return;
    }
    element.innerHTML = content;
}

export class ViewModel {
    constructor() {
        // Bindings here

        bindEvent("#test", "click", this.render);
        bindEvent("#model", "click", this.render);
    }

    // Just Testing
    render() {
        GetModels()
        .then((result) => {
            swapInner("models", result)
        })
        .catch((err) => {
            console.log(err);
        });
    }
}
 
new ViewModel()
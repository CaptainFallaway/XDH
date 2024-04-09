// Common shared fucntions for the presenter class

export async function updateModelsList(html) {
    let element = document.getElementById("models");
    if (!element) {
        console.error("Element not found: #models");
        return;
    }
    element.innerHTML = html;
}
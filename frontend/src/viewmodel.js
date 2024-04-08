/// This is supposed to be like a lightweight ViewModel

import { GetModels } from '../wailsjs/go/app/ModelInterface';

export class ViewModel {
    async updateModelsList(html) {
        let element = document.querySelector("#models");
        if (!element) {
            console.error("Element not found: #models");
            return;
        }
        element.innerHTML = html;
    }

    // Returns html string of models, since templ is nice
    async fetchModels() {
        return GetModels()
            .then((result) => result)
            .catch((err) => console.error(err));
    }
}
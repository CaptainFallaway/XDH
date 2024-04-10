/// This is supposed to be like a lightweight Presenter

import { GetModel, GetModels } from '../../wailsjs/go/app/ModelInterface';
import { updateModelsList  } from './util';

export class Presenter {
    async updateModels() {
        return GetModels()
        .then((result) => {
            updateModelsList(result)
        })
        .catch((err) => console.error(err));
    }

    async updateModel(id) {
        return GetModel(id, true)
        .then((result) => {
            let element = document.getElementById(id);
            element.innerHTML = result;
            element.classList.add("active");
        })
        .catch((err) => console.error(err));
    }

    async closeActiveModel() {
        let elements = document.getElementsByClassName("active");
        console.log(elements)
        if (elements.length == 0) {
            return;
        }

        let element = elements[0];

        return GetModel(element.id, false)
        .then((result) => {
            element.innerHTML = result;
            element.classList.remove("active");
        })
        .catch((err) => console.error(err));
    }
}

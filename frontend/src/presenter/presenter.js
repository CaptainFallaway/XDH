/// This is supposed to be like a lightweight Presenter

import { GetModels } from '../../wailsjs/go/app/ModelInterface';
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
        console.log("Updating model: " + id);
    }
}

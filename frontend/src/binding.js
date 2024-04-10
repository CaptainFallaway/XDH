import { Presenter } from "./presenter/presenter";

const presenter = new Presenter()

// This is where the bindings are going to be made (not so sure about this architecture i have yet...)

document.getElementById("test").addEventListener("click", async () => {
    await presenter.updateModels();
    let children = document.getElementById("models").children;
    [...children].forEach((child) => {
        child.addEventListener("click", async () => {
            console.log(child.id);
            await presenter.closeActiveModel();
            await presenter.updateModel(child.id);
        });
    });
});

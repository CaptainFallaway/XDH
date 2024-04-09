import { Presenter } from "./presenter/presenter";

const presenter = new Presenter()

// This is where most bindings are made

document.getElementById("test").addEventListener("click", async () => {
    await presenter.updateModels()
    let children = document.getElementById("models").children
    
});

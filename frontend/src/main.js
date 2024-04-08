import { ViewModel } from "./viewmodel";

// Create a new ViewModel
let viewModel = new ViewModel();

document.getElementById("test").addEventListener("click", () => {
    viewModel.fetchModels()
        .then((html) => {
            viewModel.updateModelsList(html);
        }).catch((err) => {
            console.error(err);
        });
});
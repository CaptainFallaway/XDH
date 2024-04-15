import { getParentElementWithRef } from "./helpers.js";
import { handleWailsBindClick } from "./Services/click_handler.js";

// Entry point for the Wails bindings
// The point of this is to create a similar experience / philosophy to htmx
// Where backend returns html that gets swapped or put into the dom

// the attributes i have defined are:
// wails-ref: object:method
// wails-state: comma separated list of parameters (default is empty)
// wails-swap: inner, outer or none (default is none)
// wails-trgt: either class or id of target element (default is self)
// wails-ignore: if present, the click event will not be handled

document.addEventListener("click", async (e) => {
    let element = e.target;

    if (element.hasAttribute("wails-ref")) {
        e.preventDefault();
        e.stopPropagation();
        
        await handleWailsBindClick(element);
    } 

    let seekedWailsParent = await getParentElementWithRef(element);

    if (seekedWailsParent !== null) {
        e.preventDefault();
        e.stopPropagation();

        await handleWailsBindClick(seekedWailsParent);
    }
})
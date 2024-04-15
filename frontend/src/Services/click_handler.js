
async function getObjAndMethod(element) {
    // this attr does not need to be checked for existence, as it is checked before handler
    let reference = element.getAttribute("wails-ref");

    if (!reference.includes(":")) {
        throw new Error("Invalid reference: " + reference + "\nExpected format: object:method");
    }
    
    return reference.split(":");
}

async function getStates(element) {
    let params = element.getAttribute("wails-state");
    
    if (!params) {
        return [];
    } else {
        return params.split(",").map((param) => param.trim());
    }
}

async function getSwap(element) {
    let temp = element.getAttribute("wails-swap");

    if (temp) {
        return temp;
    } else {
        return "none"
    }
}

async function getTarget(element) {
    let target = element.getAttribute("wails-trgt");

    if (target) {
        let temp = document.getElementById(target);

        if (!temp) {
            throw new Error("Invalid target: " + target);
        }

        return temp;
    } else {
        return element;
    }

}

export async function handleWailsBindClick(element) {    
    let [object, method] = await getObjAndMethod(element);
    let states = await getStates(element);
    let swap = await getSwap(element);
    let target = await getTarget(element);

    console.log(object, method, states, swap, target)

    if (!window['go']['app'][object]) {
        throw new Error("Invalid object: " + object);
    }

    if (!window['go']['app'][object][method]) {
        throw new Error("Invalid method: " + method);
    }
        
    let result = await window['go']['app'][object][method](...states);

    switch (swap) {
        case "inner":
            target.innerHTML = result;
            break;
        case "outer":
            target.outerHTML = result;
            break;
        case "none":
            break;
        default:
            throw new Error("Invalid swap: " + swap);
    }
}

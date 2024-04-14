
export async function getParentElementWithRef(childElement) {
    if (childElement.hasAttribute("wails-ignore")) {
        return null;
    }

    let parent = childElement.parentElement;

    while (parent != document.body) {
        if (parent.hasAttribute("wails-ignore")) {
            return parent;
        }

        if (parent.hasAttribute("wails-ref")) {
            return parent;
        }

        parent = parent.parentElement;
    }

    return null;
}
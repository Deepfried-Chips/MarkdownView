import Options from './dropdown.svelte';

export function OpenMenuOptions(node, {
    elements = [],
}) {
    let componentDOM
    function OpenDropdown(event) {
        if (componentDOM) {
            componentDOM.$destroy();
            componentDOM = undefined;
        }else {
            componentDOM = new Options({
                target: node,
                props: {
                    parent: event.target,
                    elements: elements
                }
            });
        }
    }

    node.addEventListener("click",OpenDropdown)
}

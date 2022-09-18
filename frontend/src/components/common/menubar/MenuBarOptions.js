import Options from './dropdown.svelte';

export function OpenMenuOptions(node, {
    elements = [],
}) {
    console.log(node);
    function OpenDropdown(event) {
        console.log(event);
        console.log('open');
        const options = new Options({
            target: node,
            props: {
                parent: event.target
            }
        });
    };

    node.addEventListener("click",OpenDropdown)
}

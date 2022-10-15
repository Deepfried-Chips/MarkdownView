<script>
    export let elements = [];
    export let parent = undefined;
</script>

<svelte:options accessors/>

<svelte:element class="dropdown" id="dropdown" this=parent>
    <div class="dropdown-content" bind:this={parent}>
        {#each elements as element}
            <button on:click={element.action} class="dropbtn">{element.text}</button>
        {/each}
    </div>
</svelte:element>

<style>

    /* Dropdown button */
    .dropdown .dropbtn {
        font-size: 16px;
        display: flex;
        border: none;
        outline: none;
        color: white;
        padding: 14px 16px;
        background-color: inherit;
        font-family: inherit; /* Important for vertical align on mobile phones */
        margin: 0; /* Important for vertical align on mobile phones */
    }

    /* Add a red background color to navbar links on hover */
    .dropdown .dropbtn:hover {
        background-color: #6d6d6d;
    }

    /* Dropdown content (hidden by default) */
    .dropdown-content {
        display: block;
        position: absolute;
        background-color: rgba(55,55,55,0.90);
        min-width: 60px;
        box-shadow: 0 8px 16px 0 rgba(0,0,0,0.2);
        z-index: 1;
    }

    /* Show the dropdown menu on hover */
    .dropdown:hover .dropdown-content {
        display: block;
    }
</style>

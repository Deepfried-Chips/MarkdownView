<script>
    import { Quit, WindowMinimise, WindowToggleMaximise } from "../../wailsjs/runtime/runtime";

    export let macButtons;

    function minimize() {
        WindowMinimise();
    }
    function maximize() {
        WindowToggleMaximise();
    }

</script>

<header style="--wails-draggable:drag" class="titlebar {macButtons === true ? 'type-mac' : 'type-standard'}">
    <p class="wordmark">Markdown Viewer</p>

    <div style="--wails-draggable:no-drag" class="window-controls">
        {#if macButtons === true}
            <button tabindex="-1" on:click={Quit} id="close">
                <svg width="12" height="12" viewBox="0 0 12 12">
                    <path stroke="#4c0000" fill="none" d="M8.5,3.5 L6,6 L3.5,3.5 L6,6 L3.5,8.5 L6,6 L8.5,8.5 L6,6 L8.5,3.5 Z"></path>
                </svg>
            </button>
            <button tabindex="-1" on:click={maximize} id="maximize">
                <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 20 20" fill="none">
                    <path d="M2 9.75C2 9.33579 2.33579 9 2.75 9H17.25C17.6642 9 18 9.33579 18 9.75C18 10.1642 17.6642 10.5 17.25 10.5H2.75C2.33579 10.5 2 10.1642 2 9.75Z"/>
                </svg>
            </button>
            <button tabindex="-1" on:click={minimize} id="minimize">
                <svg width="12" height="12" viewBox="0 0 12 12">
                    <rect fill="#975500" width="6" height="1" x="3" y="5.5" fill-rule="evenodd"></rect>
                </svg>
            </button>
        {:else}
            <button tabindex="-1" on:click={minimize} id="minimize">
                <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 20 20" fill="none">
                    <path d="M2 9.75C2 9.33579 2.33579 9 2.75 9H17.25C17.6642 9 18 9.33579 18 9.75C18 10.1642 17.6642 10.5 17.25 10.5H2.75C2.33579 10.5 2 10.1642 2 9.75Z"/>
                </svg>
            </button>
            <button tabindex="-1" on:click={maximize} id="maximize">
                <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 20 20" fill="none">
                    <rect stroke="#ffffff" stroke-opacity="1" stroke-width="2" fill-opacity="0" fill-rule="nonzero" width="20" height="20"></rect>
                </svg>
            </button>
            <button tabindex="-1" on:click={Quit} id="close">
                <svg xmlns="http://www.w3.org/2000/svg" width="28" height="28" viewBox="0 0 28 28" fill="none">
                    <path d="M3.52499 3.71761L3.61612 3.61612C4.07173 3.1605 4.79155 3.13013 5.28239 3.52499L5.38388 3.61612L14 12.233L22.6161 3.61612C23.1043 3.12796 23.8957 3.12796 24.3839 3.61612C24.872 4.10427 24.872 4.89573 24.3839 5.38388L15.767 14L24.3839 22.6161C24.8395 23.0717 24.8699 23.7915 24.475 24.2824L24.3839 24.3839C23.9283 24.8395 23.2085 24.8699 22.7176 24.475L22.6161 24.3839L14 15.767L5.38388 24.3839C4.89573 24.872 4.10427 24.872 3.61612 24.3839C3.12796 23.8957 3.12796 23.1043 3.61612 22.6161L12.233 14L3.61612 5.38388C3.1605 4.92827 3.13013 4.20845 3.52499 3.71761L3.61612 3.61612L3.52499 3.71761Z"/>
                </svg>
            </button>
        {/if}
    </div>
</header>

<style>
    .titlebar {
        background-color: rgb(59, 59, 59);
        color: white;
        height: 28px;
        display: flex;
        align-items: center;
        user-select: none;
    }

    .window-controls {
        display: flex;
        align-items: center;
        margin-left: auto;
    }

    .wordmark {
        width: auto;
        height: auto;
        alignment: left;
        margin: 0 8px;
    }

    .window-controls button {
        display: flex;
        color: white;
        align-items: center;
        justify-content: center;
        padding: 0;
        border: none;
    }

    /* Standard Titlebar */

    .type-standard button {
        height: 28px;
        width: 40px;
        transition: 50ms ease;
        background-color: transparent;
    }

    .type-standard button svg {
        width: 12px;
        height: 12px;
        fill: currentColor;
    }

    .type-standard button#close:hover {
        background-color: #d13d3d;
        color: #fff;
    }

    .type-standard button#close:active {
        background-color: #b12a2a;
        color: #fff;
    }

    .type-standard button#minimize:hover {
        background-color: rgba(175, 175, 175, 0.76);
        color: #fff;
    }

    .type-standard button#minimize:active {
        background-color: rgba(175, 175, 175, 0.76);
        color: rgba(255, 255, 255, 0.78);
    }

    .type-standard button#maximize:hover {
        background-color: rgba(175, 175, 175, 0.76);
        color: #fff;
    }

    .type-standard button#maximize:active {
        background-color: rgba(175, 175, 175, 0.76);
        color: rgba(255, 255, 255, 0.78);
    }

    /* Mac Titlebar */

    .type-mac {
        justify-content: space-between;
    }

    .type-mac .window-controls {
        order: -1;
        margin: 0 3px;
    }

    .type-mac .window-controls button {
        margin: 0 4px;
        width: 12px;
        height: 12px;
        border-radius: 50%;
        background-size: auto 12px;
        box-shadow: inset 0 0 0 1px rgba(0,0,0,0.12);
    }

    .type-mac .window-controls svg {
        visibility: hidden;
        width: 12px;
        height: 12px;
    }

    .type-mac .window-controls:hover svg {
        visibility: visible;
    }

    .type-mac .window-controls button:not([disabled]):active {
        box-shadow: inset 0 0 0 1px rgba(0,0,0,0.12), inset 0 0 0 12px rgba(0,0,0,0.25);
    }

    .type-mac .window-controls #close {
        margin-left: 6px;
        background-color: #ff5e57;
    }

    .type-mac .window-controls #minimize {
        background-color: #ffbb2e;
    }

    .type-mac .window-controls #maximize {
        background-color: #82ff2e;
    }
</style>

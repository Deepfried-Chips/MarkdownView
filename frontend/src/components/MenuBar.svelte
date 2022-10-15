<script>
    import { Quit, WindowMinimise, WindowToggleMaximise } from "../../wailsjs/runtime/runtime";
    import {OpenMenuOptions} from "./common/menubar/MenuBarOptions.js";

    export let mac;

    function minimize() {
        WindowMinimise();
    }
    function maximize() {
        WindowToggleMaximise();
    }

</script>

<header style="--wails-draggable:drag" class="menubar {mac === true ? 'type-mac' : 'type-standard'}">

    <div class="menu-controls">
            <button tabindex="-1" on:click={minimize} id="minimize">
                File
            </button>
            <button tabindex="-1" use:OpenMenuOptions={{elements:[{text:"About",action:maximize}]}} id="maximize">
                View
            </button>
            <button tabindex="-1" use:OpenMenuOptions={{elements:[{text:"About",action:Quit}]}} id="close">
                Help
            </button>
    </div>
</header>

<style>
    .menubar {
        background-color: rgb(59, 59, 59);
        color: white;
        height: 28px;
        align-items: center;
        user-select: none;
    }

    .menu-controls {
        width: auto;
        height: auto;
        display: flex;
    }

    .menu-controls button {
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

    .type-standard button:hover {
        background-color: rgba(175, 175, 175, 0.76);
        color: #fff;
    }

    .type-standard button:active {
        background-color: rgba(175, 175, 175, 0.76);
        color: rgba(255, 255, 255, 0.78);
    }
</style>

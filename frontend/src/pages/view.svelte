<script>
    import FileBrowser from "../components/FileBrowser.svelte";
</script>

<div class="left">
    <FileBrowser/>
</div>
<div class="right">
    <p>Hello!</p>
</div>

<style>
    .left {
        background-color: rgba(55,55,55,0.24);
        color: #fff;
        float: left;
        height: 100%;
        overflow: auto;
        padding: 2px;
        width: 40%;
        user-select: none;
    }
    .right {
        background-color: rgba(55,55,55,0.24);
        float: right;
        height: 100%;
        overflow: auto;
        padding: 15px;
        width: 60%;
    }
</style>

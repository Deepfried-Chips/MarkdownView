<script>
  import { register, init, getLocaleFromNavigator, isLoading } from 'svelte-i18n';
  import {GetPlatform} from "../wailsjs/go/main/App.js";
  import View from './pages/view.svelte'
  import Titlebar from "./components/Titlebar.svelte";
  import MenuBar from "./components/MenuBar.svelte";

  import "./app.css";


  let platform = "win32";

  register('en', () => import('./i18n/en.json'));

  init({
    fallbackLocale: 'en',
    initialLocale: getLocaleFromNavigator(),
  });

  function PlatformProcess() {
    GetPlatform().then(result => platform = result);
  }

  PlatformProcess();
</script>

<div class="main-window-{platform}">
  <Titlebar macButtons={platform === "darwin"}/>
  {#if $isLoading}
    <div>Loading...</div>
  {:else}
    <MenuBar mac={platform === "darwin"}/>
    <View/>
  {/if}
</div>

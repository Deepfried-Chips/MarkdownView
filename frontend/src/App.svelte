<script>
  import { register, init, getLocaleFromNavigator, _, isLoading } from 'svelte-i18n';
  import {GetPlatform} from "../wailsjs/go/main/App.js";
  import Home from './pages/home.svelte'
  import Titlebar from "./components/Titlebar.svelte";

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

<div class="main-window-{GetPlatform()}">
  <Titlebar macButtons={platform === "darwin"}/>
  {#if $isLoading}
    <div>Loading...</div>
  {:else}
    <Home/>
  {/if}
</div>

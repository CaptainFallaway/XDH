<script>
    import "./app.css";

    import { ModeWatcher, mode } from "mode-watcher";
    import { toggleValue } from "$lib/stores.js";

    import Model from "$lib/Model.svelte";
    import MenuBar from "$lib/MenuBar.svelte";

    import * as app from "$lib/wailsjs/go/app/App.js";
</script>

<ModeWatcher />

<MenuBar />

{#await app.GetModels($toggleValue)}
    <div style="--main: var(--{$mode}-main); --secon: var(--{$mode}-secon)" class="w-full h-screen flex justify-center">
        <div  class="loader m-auto bottom-[80px]"></div>
    </div>
{:then modelList}
    {#each modelList as model}
        <Model id={model.index.toString()} {model} />
    {/each}
{:catch error}
    <p>{error.message}</p>
{/await}

<style>
:root {
    --dark-main: #fff;
    --dark-secon: #fff2;
    --light-main: #000;
    --light-secon: #0002;
}
.loader {
  
  width: 16px;
  height: 16px;
  border-radius: 50%;
  background-color: var(--main);
  box-shadow: 32px 0 var(--main), -32px 0 var(--main);
  position: relative;
  animation: flash 0.5s ease-out infinite alternate;
}

@keyframes flash {
  0% {
    background-color: var(--secon);
    box-shadow: 32px 0 var(--secon), -32px 0 var(--main);
  }
  50% {
    background-color: var(--main);
    box-shadow: 32px 0 var(--secon), -32px 0 var(--secon);
  }
  100% {
    background-color: var(--secon);
    box-shadow: 32px 0 var(--main), -32px 0 var(--secon);
  }
}
</style>
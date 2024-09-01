<script>
    import "./app.css";

    import { ModeWatcher } from "mode-watcher";
    import { toggleValue } from "$lib/stores.js";

    import Model from "$lib/Model.svelte";
    import MenuBar from "$lib/MenuBar.svelte";

    import * as app from "$lib/wailsjs/go/app/App.js";
</script>

<ModeWatcher />

<MenuBar />

{#await app.GetModels($toggleValue)}
    <div class="w-full h-screen flex justify-center">
        <div class="loader m-auto bottom-[80px]"></div>
    </div>
{:then modelList}
    {#each modelList as model}
        <Model id={model.index.toString()} {model} />
    {/each}
{:catch error}
    <p>{error.message}</p>
{/await}

<style>
.loader {
  width: 16px;
  height: 16px;
  border-radius: 50%;
  background-color: #fff;
  box-shadow: 32px 0 #fff, -32px 0 #fff;
  position: relative;
  animation: flash 0.5s ease-out infinite alternate;
}

@keyframes flash {
  0% {
    background-color: #FFF2;
    box-shadow: 32px 0 #FFF2, -32px 0 #FFF;
  }
  50% {
    background-color: #FFF;
    box-shadow: 32px 0 #FFF2, -32px 0 #FFF2;
  }
  100% {
    background-color: #FFF2;
    box-shadow: 32px 0 #FFF, -32px 0 #FFF2;
  }
}
</style>
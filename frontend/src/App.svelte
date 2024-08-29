<script>
    import "./app.css";

    import { ModeWatcher } from "mode-watcher";
    import { toggleValue } from "$lib/stores.js";

    import Model from "$lib/Model.svelte";
    import MenuBar from "$lib/MenuBar.svelte";

    import * as wails from "$lib/wailsjs/go/app/App.js";

    let modelListPromise;
    $: modelListPromise = wails.GetModels($toggleValue);
</script>

<ModeWatcher />

<MenuBar />

<!-- svelte-ignore empty-block -->
<!-- in the future i want to add a loading icon -->
{#await modelListPromise then modelList}
    {#each modelList as model}
        <Model id={model.index.toString()} {model} />
    {/each}
{:catch error}
    <p>{error.message}</p>
{/await}

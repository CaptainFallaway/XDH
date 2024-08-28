<script>
    import "./app.css";

    import { ModeWatcher } from "mode-watcher";

    import MenuBar from "$lib/MenuBar.svelte";
    import { toggleValue } from "$lib/stores.js";

    import * as wails from "$lib/wailsjs/go/app/App.js";

    $: {
        wails.GetModels($toggleValue).then((models) => {
            console.log(models);
        });
    }
</script>

<ModeWatcher />

<MenuBar />

<!-- <p>{modelsList}</p> -->

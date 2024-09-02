<script lang="ts">
    import { Button } from "$lib/components/ui/button";
    import * as Tabs from "$lib/components/ui/tabs/index.ts";

    import Dropdown from "$lib/Dropdown.svelte";

    import * as app from "$lib/wailsjs/go/app/App.js";

    import { toggleValue } from "$lib/stores.ts";

    function openFileDialog() {
        app.OpenFileDialog()
        const temp = $toggleValue
        $toggleValue = ""
        $toggleValue = temp
    }
</script>

<div class="flex flex-row right-0 left-0 m-5">
    <div class="basis-1/2 text-left place-items-center">
        <Tabs.Root bind:value={$toggleValue} class="justify-start">
            <Tabs.List class="shadow-xl">
                <Tabs.Trigger value="Sn">Sn</Tabs.Trigger>
                <Tabs.Trigger value="Pb">Pb</Tabs.Trigger>
                <!-- <Tabs.Trigger value="Cu">Cu</Tabs.Trigger>
                <Tabs.Trigger value="Zn">Zn</Tabs.Trigger> -->
            </Tabs.List>
        </Tabs.Root>
    </div>
    <div class="basis-1/2 flex gap-1 justify-end">
        <Button on:click={openFileDialog} class="shadow-xl">Open File</Button>
        <Dropdown />
    </div>
</div>

<script>
    import { toggleMode } from "mode-watcher";
    import { Button } from "$lib/components/ui/button";
    import * as Tabs from "$lib/components/ui/tabs/index.js";

    import Sun from "lucide-svelte/icons/sun";
    import Moon from "lucide-svelte/icons/moon";

    import * as wails from "$lib/wailsjs/go/app/App.js";

    import { toggleValue } from "$lib/stores.js";
</script>

<div class="flex flex-row right-0 left-0 m-5">
    <div class="basis-1/2 text-left place-items-center">
        <Tabs.Root bind:value={$toggleValue} class="justify-start">
            <Tabs.List>
                <Tabs.Trigger value="Pb">Pb</Tabs.Trigger>
                <Tabs.Trigger value="Sn">Sn</Tabs.Trigger>
            </Tabs.List>
        </Tabs.Root>
    </div>
    <div class="basis-1/2 flex gap-1 justify-end">
        <Button on:click={wails.OpenFileDialog}>Open File</Button>
        <Button on:click={toggleMode}>
            <Sun
                class="h-[1.2rem] w-[1.2rem] rotate-0 scale-100 transition-all dark:-rotate-90 dark:scale-0"
            />
            <Moon
                class="absolute h-[1.2rem] w-[1.2rem] rotate-90 scale-0 transition-all dark:rotate-0 dark:scale-100"
            />
            <span class="sr-only">Toggle theme</span>
        </Button>
    </div>
</div>

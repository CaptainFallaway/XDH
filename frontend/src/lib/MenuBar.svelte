<script>
    import { toggleMode } from "mode-watcher";
    import { Button } from "$lib/components/ui/button";
    import * as ToggleGroup from "$lib/components/ui/toggle-group/index.js";

    import Sun from "lucide-svelte/icons/sun";
    import Moon from "lucide-svelte/icons/moon";
    import FolderArchive from "lucide-svelte/icons/folder-archive";

    import * as wails from "$lib/wailsjs/go/app/App.js";

    import { toggleValue } from "$lib/stores.js";
</script>

<div class="flex flex-row border-2 rounded-lg right-0 left-0 m-3 p-2">
    <div class="basis-1/2 text-left place-items-center">
        <ToggleGroup.Root
            bind:value={$toggleValue}
            id="toggle"
            type="single"
            variant="outline"
            class="justify-start"
        >
            <ToggleGroup.Item value="Pb" aria-label="Tenn">
                <p>Pb</p>
            </ToggleGroup.Item>
            <ToggleGroup.Item value="Sn" aria-label="Fuck">
                <p>Sn</p>
            </ToggleGroup.Item>
        </ToggleGroup.Root>
    </div>
    <div class="basis-1/2 text-right place-items-center">
        <Button on:click={wails.OpenFileDialog}>Allah</Button>
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

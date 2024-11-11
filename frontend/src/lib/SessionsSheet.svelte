<script lang="ts">
    import * as Sheet from "$lib/components/ui/sheet";
    import ScrollArea from "./components/ui/scroll-area/scroll-area.svelte";
    import { Button } from "$lib/components/ui/button";
    import { PanelLeft } from "lucide-svelte";
    import * as Tabs from "$lib/components/ui/tabs";

    import { setMode, mode, resetMode } from "mode-watcher";

    import * as app from "$lib/wailsjs/go/app/App.js";
    import { Sun, Moon } from "lucide-svelte";

    let _mode: "dark" | "light" | undefined;
    mode.subscribe((value) => {
        _mode = value;
    });

    import { toggleValue } from "$lib/globalstores";

    function openFileDialog() {
      app.OpenFileDialog()
      const temp = $toggleValue
      $toggleValue = ""
      $toggleValue = temp
    }
</script>

<Sheet.Root>
    <Sheet.Trigger asChild let:builder>
      <Button builders={[builder]} variant="outline"><PanelLeft/></Button>
    </Sheet.Trigger>
    <Sheet.Content side="left" class="flex flex-col">
      <Sheet.Header>
        <Sheet.Title>
          <div class="flex items-center">
            Mätnings Sessioner
            <div class="flex justify-end">
              <Button on:click={openFileDialog}>Ny Session</Button>
            </div>
          </div>
        </Sheet.Title>
      </Sheet.Header>

      <ScrollArea class="gap-1 mt-4 flex-1">
        <div class="w-full p-2 hover:bg-secondary rounded-md">
            <p>Hello World</p>
        </div>
        <div class="w-full p-2 hover:bg-secondary rounded-md">
            <p>Hello World</p>
        </div>
      </ScrollArea>
      <div class="flex w-full">
        <Tabs.Root value="dark" class="w-[800px]">
          <Tabs.List>
            <Tabs.Trigger value="dark"><Moon /></Tabs.Trigger>
            <Tabs.Trigger value="light"><Sun /></Tabs.Trigger>
          </Tabs.List>
        </Tabs.Root>
      </div>
    </Sheet.Content>
</Sheet.Root>
<script>
    import * as DropdownMenu from "$lib/components/ui/dropdown-menu/index.js";
    import { Button } from "$lib/components/ui/button";
    import { Menu } from "lucide-svelte";

    import { setMode, mode, resetMode } from "mode-watcher";
    import Separator from "./components/ui/separator/separator.svelte";

    import { Sun, Moon } from "lucide-svelte";

    let _mode;
    mode.subscribe((value) => {
        _mode = value;
    });
</script>

<DropdownMenu.Root closeOnItemClick={false}>
    <DropdownMenu.Trigger>
        <Button class="shadow-xl" variant="outline">
            <Menu size={24} />
        </Button>
    </DropdownMenu.Trigger>
    <DropdownMenu.Content>
        <DropdownMenu.Sub>
            <DropdownMenu.SubTrigger>Theme</DropdownMenu.SubTrigger>
            <DropdownMenu.SubContent>
                <DropdownMenu.RadioGroup bind:value={_mode}>
                    <DropdownMenu.RadioItem
                        value="dark"
                        on:click={() => setMode("dark")}
                        >Dark <div
                            class="flex content-center justify-end w-full"
                        >
                            <Moon size={18} />
                        </div>
                    </DropdownMenu.RadioItem>
                    <DropdownMenu.RadioItem
                        value="light"
                        on:click={() => setMode("light")}
                        >Light <div
                            class="flex content-center justify-end w-full"
                        >
                            <Sun size={18} />
                        </div>
                    </DropdownMenu.RadioItem>
                </DropdownMenu.RadioGroup>
                <DropdownMenu.Separator />
                <DropdownMenu.Item on:click={resetMode}
                    >System</DropdownMenu.Item
                >
            </DropdownMenu.SubContent>
        </DropdownMenu.Sub>
        <DropdownMenu.Group>
            <DropdownMenu.Item>Billing</DropdownMenu.Item>
        </DropdownMenu.Group>
        <DropdownMenu.Group>
            <DropdownMenu.Item>Team</DropdownMenu.Item>
        </DropdownMenu.Group>
        <DropdownMenu.Group>
            <DropdownMenu.Item>Subscription</DropdownMenu.Item>
        </DropdownMenu.Group>
    </DropdownMenu.Content>
</DropdownMenu.Root>

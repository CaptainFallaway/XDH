<script>
    import { cubicOut } from "svelte/easing";
    import { fly, slide } from "svelte/transition";

    import { toggleValue } from "$lib/stores.js";
    // import { Badge } from "$lib/components/ui/badge/";
    import { Separator } from "$lib/components/ui/separator";
    import Info from "$lib/Info.svelte";
    import * as Tabs from "$lib/components/ui/tabs/index.js";
    import Table from "$lib/Table.svelte";
    import { ChevronUp, ChevronDown } from "lucide-svelte";

    export let id = "0";
    export let model = {};

    const expandMotion = {
        duration: 500,
        easing: cubicOut,
    };

    // let status;
    let statusColor;

    switch (model.violations[$toggleValue]) {
        case 0:
            // status = "safe";
            statusColor = "green";
            break;
        case 1:
            // status = "warning";
            statusColor = "yellow";
            break;
        default:
            // status = "danger";
            statusColor = "red";
            break;
    }

    let show = false;

    function toggleShow() {
        show = !show;
    }

    function handleKeyPress(event) {
        if (event.key === "Enter" || event.key === " ") {
            toggleShow();
        }
    }
</script>

<!-- shadow-red-500 shadow-yellow-500 shadow-green-500 (Not used right now) -->
<!-- border-l-red-500 border-l-yellow-500 border-l-green-500 -->

<div
    {id}
    role="button"
    on:click={toggleShow}
    on:keydown={handleKeyPress}
    tabindex="0"
    transition:fly={{ y: 200, duration: 500 }}
    class="border-2 rounded-lg right-0 left-0 m-5 content-center shadow-lg border-l-8 border-l-{statusColor}-500"
>
    <div aria-label="info" class="flex gap-10 m-8 content-center">
        <Info top="Boat ID" bottom={model.boatID} />
        <Info top="Operators" bottom={model.operators} />
        <Info top="First Date" bottom={model.firstDate} />
        <Info top="Last Date" bottom={model.lastDate} />
        <Info top="Units" bottom={model.unit} />
        <div class="w-full flex items-center justify-end">
            {#if show}
                <ChevronDown size={48} />
            {:else}
                <ChevronUp size={48} />
            {/if}
        </div>
    </div>
    {#if show}
        <Separator />
        <div
            class="m-2 flex"
            in:slide={{ ...expandMotion }}
            out:slide={{ ...expandMotion }}
        >
            <div
                class="w-full mx-5 mb-5 flex flex-col"
                role="button"
                tabindex="0"
                on:click={(e) => e.stopPropagation()}
                on:keydown={(e) => e.stopPropagation()}
            >
                <div class="w-full flex content-center justify-center gap-8">
                    <Info top="Total Scans" bottom={model.totalScans} />
                    <Info top="Valid Scans" bottom={model.validScans} />
                    <Info top="Invalid Scans" bottom={model.invalidScans} />
                </div>
                <Tabs.Root class="w-full">
                    <Tabs.List>
                        <Tabs.Trigger value="Valid">Valid</Tabs.Trigger>
                        <Tabs.Trigger value="Invalid">Invalid</Tabs.Trigger>
                    </Tabs.List>
                    <Tabs.Content value="Valid">
                        <Table scans={model.scans} />
                    </Tabs.Content>
                    <Tabs.Content value="Invalid">
                        <Table scans={model.invalidScans} />
                    </Tabs.Content>
                </Tabs.Root>
            </div>
        </div>
    {/if}
</div>
